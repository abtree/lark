package xcfgstructpb

import (
	"encoding/json"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var oridat string = ""

func parserTxt(path, dat string, isyyact bool) {
	pos := strings.IndexByte(dat, '^')
	if pos == -1 {
		return //没有表头
	}
	dat = dat[pos+1:]
	pos = strings.IndexByte(dat, '\n')
	dat = dat[:pos]
	spkey := strings.Contains(dat, ".key")
	strs := strings.FieldsFunc(dat, func(c rune) bool { return c == '\t' })
	key := "string"
	//解析表头
	pbname := strs[0]
	pbname = strings.TrimSpace(pbname)
	//判断key类型
	if !spkey {
		k1 := strings.TrimSpace(strs[1])
		pos := strings.IndexByte(k1, ':')
		key = k1[:pos]
	}
	//判断是否已经存在
	if _, find := had_pbs[pbname]; find {
		return //已经解析过的表头
	} else {
		had_pbs[pbname] = true
	}
	//生成pb
	strbuff := strings.Builder{}
	strbuff.WriteString("//")
	strbuff.WriteString(path)
	strbuff.WriteString(end_chat)
	strbuff.WriteString("message ")
	strbuff.WriteString(pbname)
	strbuff.WriteString("{")
	wfile.WriteString(strbuff.String())

	strbuff = strings.Builder{}
	//解析表项
	size := len(strs)
	j := 1
	for i := 1; i < size; i++ {
		str := strings.TrimSpace(strs[i])
		if str == "" || str[0] == '.' {
			continue //忽略的字段
		}
		pos := strings.LastIndexByte(str, ':')
		if pos < 1 { //-1 或 0
			continue //不需要解析的字段
		}
		typ := str[:pos]   //字段类型
		nam := str[pos+1:] //字段名称
		//解析字段名称
		idx := strings.IndexByte(nam, '.')
		if idx > 0 {
			nam = nam[:idx]
		}
		//去除分割符
		if nam[0] < 'A' || nam[0] > 'Z' {
			nam = nam[1:]
		}
		strbuff.WriteString(end_begin)
		if strings.HasPrefix(typ, pb_array) {
			//普通数组
			strbuff.WriteString("repeated ")
			strbuff.WriteString(strings.TrimPrefix(typ, pb_array))
		} else if strings.HasPrefix(typ, "{") || strings.HasPrefix(typ, "[{") {
			typ = build(nam, typ)
			strbuff.WriteString(typ)
		} else {
			strbuff.WriteString(typ)
		}

		strbuff.WriteString(" ")
		strbuff.WriteString(nam)
		//添加索引
		strbuff.WriteString(" = ")
		strbuff.WriteString(strconv.Itoa(j))
		strbuff.WriteString(";")
		j++
	}
	strbuff.WriteString(end_chat)
	strbuff.WriteString("}")
	strbuff.WriteString(end_chat)
	strbuff.WriteString(end_chat)

	//写入文件
	wfile.WriteString(strbuff.String())
	if isyyact {
		writeFile(path, pbname, key)
	} else {
		writeBig(path, pbname, key)
	}
}

func build(name, dat string) (ret string) {
	var config interface{}
	strbuilder = strings.Builder{}
	strbuilder.WriteString(end_chat)
	err := json.Unmarshal([]byte(dat), &config)
	if err != nil {
		log.Panicf(err.Error())
	}
	oridat = dat
	if reflect.TypeOf(config).Kind() == reflect.Slice {
		x := config.([]interface{})[0]
		buildObj(name, x.(map[string]interface{}))
		ret = "repeated C" + name
	} else {
		buildObj(name, config.(map[string]interface{}))
		ret = "C" + name
	}
	wfile.WriteString(strbuilder.String())
	return
}

func sortkeys(m map[string]interface{}) (ret []string) {
	poss := []int{}
	keys := map[int]string{}
	for k := range m {
		//这里是因为key 都是带分隔符的 分隔符不会重复 所以key不会重复
		//可以用key在字符串中出现的位置 进行排序
		pos := strings.Index(oridat, k)
		poss = append(poss, pos)
		keys[pos] = k
	}
	sort.Ints(poss)
	for _, p := range poss {
		ret = append(ret, keys[p])
	}
	return
}

func buildObj(nam string, m map[string]interface{}) {
	var seq = 1
	strbuilder.WriteString("message C" + nam + "{\n")
	//先排序key 不然会乱序
	keys := sortkeys(m)
	for _, k := range keys {
		v := m[k]
		nk := k[1:]
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Slice {
			x := v.([]interface{})[0]
			if reflect.TypeOf(x).Kind() == reflect.Map {
				buildObj(nk, x.(map[string]interface{}))
				strbuilder.WriteString("\t repeated C" + nk + " " + strings.Title(nk) + " = " + strconv.Itoa(seq) + ";\n")
			} else {
				strbuilder.WriteString("\t repeated " + x.(string) + " " + strings.Title(nk) + " = " + strconv.Itoa(seq) + ";\n")
			}
		} else if kind == reflect.Map {
			buildObj(nk, v.(map[string]interface{}))
			strbuilder.WriteString("\t C" + nk + " " + strings.Title(nk) + " = " + strconv.Itoa(seq) + ";\n")
		} else { //普通类型
			strbuilder.WriteString("\t" + v.(string) + " " + strings.Title(nk) + " = " + strconv.Itoa(seq) + ";\n")
		}
		seq++
	}
	strbuilder.WriteString("} \n")
}
