package xloadcfg

import (
	"encoding/json"
	"log"
	"reflect"
	"regexp"
	"strings"
)

var (
	rInit *regexp.Regexp //匹配到第一个大写字母
	rSub  *regexp.Regexp //从第一个大写字母切子串
)

func init() {
	var err error
	rInit, err = regexp.Compile(`^(\[\])?[a-z]`)
	if err != nil {
		log.Fatalf("regexp init with alpha error with : %v", err)
	}
	rSub, err = regexp.Compile(`^(\[\])?[a-z][a-z0-9_]*[A-Z]`)
	if err != nil {
		log.Fatalf("regexp sub string error with : %v", err)
	}

	pk := &ParserKandler{
		Category: PARSER_CG_EXT,
		Parser:   decodeIni,
		Param:    ".ini.txt",
	}
	RegisterParser(pk)
	pk.Param = ".ini"
	RegisterParser(pk)
}

func regInit(str string) bool {
	if rInit.MatchString(str) {
		return true
	}
	return false
}

func findSubStr(str string) string {
	s := rSub.FindString(str)
	if s == "" {
		return s
	}
	return s[:len(s)-1]
}

func decodeIni(dat, nex string, cfg *reflect.Value) {
	initValue(cfg)
	field := cfg.Elem()
	lines := strings.FieldsFunc(dat, func(c rune) bool { return c == '\n' })
	i := 0
	size := len(lines)
	//先找到pb名称
	for i < size {
		str := strings.TrimSpace(lines[i])
		i++
		if strings.HasPrefix(str, "[") && !strings.HasPrefix(str, "[]") {
			//找到表头
			break
		}
	}
	//解析字段
	for i < size {
		str := strings.TrimSpace(lines[i])
		i++
		if str[0] == ';' {
			continue //注释行
		}
		pos := strings.IndexByte(str, '=')
		if pos < 1 {
			continue //结构不完整
		}
		typ := strings.TrimSpace(str[:pos])
		val := strings.TrimSpace(str[pos+1:])
		if strings.HasPrefix(str, "[{") || strings.HasPrefix(str, "{") {
			p := -1
			if typ[0] == '[' {
				p = strings.LastIndexByte(typ, ']')
			} else {
				p = strings.LastIndexByte(typ, '}')
			}
			js := typ[:p+1]
			par := typ[p+1:]
			var config interface{}
			err := json.Unmarshal([]byte(js), &config)
			if err != nil {
				log.Panicf(err.Error())
			}
			title := &cfgColumn{
				typ:    config,
				name:   par,
				oristr: js,
			}
			fullSubObj(title, field.FieldByName(par), val)
		} else if strings.HasPrefix(val, "[{") || strings.HasPrefix(val, "{") {
			ff := field.FieldByName(typ)
			if ff.Kind() == reflect.Slice {
				log.Panicln("reflect can not Unmarshal slice json, change slice to object")
				// slic := reflect.MakeSlice(ff.Type(), 0, 0)
				// vv := slic.Interface()
				// decodeJson(val, &vv)
				// ff.Set(slic)
			} else {
				initValue(&ff)
				vv := ff.Interface()
				dojson(val, vv)
			}
		} else {
			//取出正在的名称
			if regInit(typ) {
				str := findSubStr(typ)
				typ = strings.TrimPrefix(typ, str)
			}
			ff := field.FieldByName(typ)
			split := ""
			if ff.Kind() == reflect.Slice {
				split = "_"
			}
			fullField(&ff, val, split)
		}
	}
}
