package xloadcfg

import (
	"encoding/json"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var oriname string

func init() {
	pk := &ParserKandler{
		Category: PARSER_CG_EXT,
		Parser:   decodeTxt,
		Param:    ".txt",
	}
	RegisterParser(pk)
}

func decodeTxt(dat, nex string, field *reflect.Value) {
	lines := strings.FieldsFunc(dat, func(c rune) bool { return c == '\n' })
	//先查找title
	titlestr := ""
	i := 0
	size := len(lines)
	for i < size {
		str := strings.TrimSpace(lines[i])
		i++
		if strings.HasPrefix(str, "^") {
			//找到title
			str = strings.TrimPrefix(str, "^")
			titlestr = strings.TrimSpace(str)
			break
		}
		if strings.HasPrefix(str, "#") {
			//没有表头
			return
		}
	}
	if titlestr == "" {
		//没有表头
		return
	}
	title := &cfgTitle{
		columns: map[int]*cfgColumn{},
	}
	dcolumns(titlestr, title)

	//加入到map
	if field.IsNil() {
		field.Set(reflect.MakeMap(field.Type()))
	}
	//解析行
	for i < size {
		str := strings.TrimSpace(lines[i])
		i++
		if !strings.HasPrefix(str, "#") {
			continue
		}
		str = strings.TrimPrefix(str, "#")
		drow(strings.TrimSpace(str), title, *field)
	}
}

func dcolumns(source string, title *cfgTitle) {
	strs := strings.FieldsFunc(source, func(c rune) bool { return c == '\t' })
	title.pbname = "pb." + strings.TrimSpace(strs[0])
	strs = strs[1:]
	for i, v := range strs {
		v = strings.TrimSpace(v)
		v = strings.Trim(v, "\"")
		//特殊字段解析
		pos := strings.LastIndexByte(v, ':')
		if strings.HasPrefix(v, ".key") {
			//特殊key
			v = v[pos+1:]
			ss := strings.Split(v, ",")
			for _, si := range ss {
				if id, err := strconv.Atoi(si); err != nil {
					log.Fatalf("Parser Configs Key Error: %v", err)
				} else {
					title.key = append(title.key, id-1)
				}
			}
			continue
		}
		if strings.HasPrefix(v, ".multi") {
			//重复key字段
			title.multi = v[pos+1:]
			continue
		}
		//普通字段解析
		if pos == -1 {
			continue //忽略的字段
		} else if pos == 0 {
			title.columns[i] = &cfgColumn{
				typ:  "",
				name: v[1:],
			}
		} else {
			title.columns[i] = &cfgColumn{
				typ:  v[:pos],
				name: v[pos+1:],
			}
		}
	}
	if title.multi == "" {
		//需要特殊解析的拼接字段
		for _, col := range title.columns {
			//拼接字符串直接解析为对象 不能有点号
			if strings.IndexByte(col.name, '.') == -1 {
				typ := col.typ.(string)
				typ = strings.ReplaceAll(typ, "\"\"", "\"")
				if strings.HasPrefix(typ, "[{") || strings.HasPrefix(typ, "{") {
					var config interface{}
					err := json.Unmarshal([]byte(typ), &config)
					if err != nil {
						log.Panicf(err.Error())
					}
					col.typ = config
					col.oristr = typ
				}
			}
		}
	}
}

func drow(line string, title *cfgTitle, field reflect.Value) {
	rowName := ""
	defer func() {
		if err := recover(); err != nil {
			log.Panicf("Parse %s Row %s Err %s %s", title.pbname, rowName, err, line)
		}
	}()

	strs := strings.FieldsFunc(line, func(c rune) bool { return c == '\t' })
	//先取出key
	key := strs[0]
	if len(title.key) > 0 {
		key = ""
		for _, v := range title.key {
			key += strs[v] + ","
		}
		key = key[:len(key)-1]
	}
	//先生成pb对象
	fkeyT, _ := parseValue(field.Type().Key().Kind(), key)
	fkey := reflect.ValueOf(fkeyT).Convert(field.Type().Key())
	row := field.MapIndex(fkey)
	if !row.IsValid() || row.IsZero() {
		row = newPB(title.pbname)
		if row.IsNil() {
			log.Panicln("Not find pb ", title.pbname)
		}
		field.SetMapIndex(fkey, row)
	}

	//先生成multi对象
	if title.multi != "" {
		initMulti(row, title.multi)
	}

	//解析每一列
	for i, value := range strs {
		col, ok := title.columns[i]
		if !ok {
			continue //忽略该字段
		}
		//用于打印解析失败的列
		rowName = col.name
		if !fullData(row, col, value) {
			log.Panicln("Parse column err", col.name)
		}
	}
}

func initMulti(t reflect.Value, field string) {
	tv := t.Elem().FieldByName(field)
	if tv.IsNil() {
		slic := reflect.MakeSlice(tv.Type(), 0, 0)
		tv.Set(slic)
	}
	//log.Debugf("%v", tv.Type().Elem().Elem())
	vv := reflect.New(tv.Type().Elem().Elem())
	tv.Set(reflect.Append(tv, vv))
}

func fullData(t reflect.Value, col *cfgColumn, value string) bool {
	line := col.name
	if line == "" {
		return false
	}
	split := ""
	if line[0] < 'A' || line[0] > 'Z' {
		//特殊符号
		split = string(line[0])
		line = line[1:]
	}
	pars := strings.Split(line, ".")
	if len(pars) == 1 {
		field := t.Elem().FieldByName(pars[0])
		if reflect.TypeOf(col.typ).Kind() == reflect.String {
			fullField(&field, value, split)
		} else {
			fullSubObj(col, field, value)
		}
	} else {
		tv := t
		for _, str := range pars {
			if tv.Type().Kind() == reflect.Slice {
				i, err := strconv.Atoi(str)
				if err != nil {
					//.multi 多行的情况 (取最后一个,需要提前将他填进去)
					i = tv.Len() - 1
					tv = tv.Index(i)
					tv = tv.Elem().FieldByName(str)
					initValue(&tv)
				} else {
					for tv.Len() <= i {
						vvv := reflect.New(tv.Type().Elem().Elem())
						tv.Set(reflect.Append(tv, vvv))
					}
					tv = tv.Index(i)
				}
			} else {
				tv = tv.Elem().FieldByName(str)
				if !tv.IsValid() {
					return false
				}
				initValue(&tv)
			}
		}
		if reflect.TypeOf(col.typ).Kind() == reflect.String {
			fullField(&tv, value, split)
		} else {
			fullSubObj(col, tv, value)
		}
	}
	return true
}

func fullSubObj(title *cfgColumn, val reflect.Value, dat string) {
	oriname = title.oristr
	if reflect.TypeOf(title.typ).Kind() == reflect.Slice {
		x := title.typ.([]interface{})[0]
		strs := strings.FieldsFunc(dat, func(c rune) bool { return c == '|' })
		for _, v := range strs {
			f := reflect.New(val.Type().Elem().Elem())
			fullstruct(f.Elem(), v, x.(map[string]interface{}))
			val.Set(reflect.Append(val, f))
		}
	} else {
		f := reflect.New(val.Type().Elem())
		fullstruct(f.Elem(), dat, title.typ.(map[string]interface{}))
		val.Set(f)
	}
}

func sortkeys(m map[string]interface{}) (ret []string) {
	poss := []int{}
	keys := map[int]string{}
	for k := range m {
		//这里是因为key 都是带分隔符的 分隔符不会重复 所以key不会重复
		//可以用key在字符串中出现的位置 进行排序
		pos := strings.Index(oriname, k)
		poss = append(poss, pos)
		keys[pos] = k
	}
	sort.Ints(poss)
	for _, p := range poss {
		ret = append(ret, keys[p])
	}
	return
}

func fullstruct(field reflect.Value, val string, m map[string]interface{}) {
	//先获取分隔符
	var split rune = '|'
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.String {
			split = rune(k[0])
			break //找到了基础类型时 就找到了正确的分隔符
		}
		//如果没有找到基础类型 就不用切分了(下面处理时再切分)
	}
	strs := strings.FieldsFunc(val, func(c rune) bool { return c == split })
	pos := 0
	keys := sortkeys(m)
	for _, k := range keys {
		v := m[k]
		kind := reflect.TypeOf(v).Kind()
		switch kind {
		case reflect.Slice: //数组
			x := v.([]interface{})[0]
			split := rune(k[0])
			vls := strings.FieldsFunc(strs[pos], func(c rune) bool { return c == split })
			f := field.FieldByName(strings.Title(k[1:]))
			if reflect.TypeOf(x).Kind() == reflect.Map {
				for _, v := range vls {
					ff := reflect.New(f.Type().Elem().Elem())
					// log.Println(ff.Elem().Type().Name())
					// log.Println(ff.Elem().NumField())
					fullstruct(ff.Elem(), v, x.(map[string]interface{}))
					f.Set(reflect.Append(f, ff))
				}
			} else {
				//基本类型的切片 直接填值
				for _, v := range vls {
					fullField(&f, v, "")
				}
			}
		case reflect.Map: //struct
			f := field.FieldByName(strings.Title(k[1:]))
			initValue(&f)
			fullstruct(f, strs[pos], v.(map[string]interface{}))
			field.Set(f)
		default: //普通字段
			f := field.FieldByName(strings.Title(k[1:]))
			fullField(&f, strs[pos], "")
		}
		pos++
	}
}
