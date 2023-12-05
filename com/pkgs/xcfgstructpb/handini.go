package xcfgstructpb

import (
	"bytes"
	"strconv"
	"strings"
)

func parserIni(path, dat string, isyyact bool) {
	lines := strings.FieldsFunc(dat, func(c rune) bool { return c == '\n' })
	//先找到pb名称
	i := 0
	size := len(lines)

	pbname := ""
	for i < size {
		str := strings.TrimSpace(lines[i])
		i++
		if strings.HasPrefix(str, "[") && !strings.HasPrefix(str, "[]") {
			//找到表头
			pos := strings.IndexByte(str, ']')
			pbname = str[1:pos]
			break
		}
	}
	if pbname == "" {
		return //没有表头
	}
	if ok := had_pbs[pbname]; ok {
		return //已经创建
	} else {
		had_pbs[pbname] = true
	}
	//构建pb
	msg := bytes.NewBuffer([]byte{})
	defer func() {
		msg.WriteString(end_chat)
		msg.WriteString("}")
		msg.WriteString(end_chat)
		msg.WriteString(end_chat)
		wfile.WriteString(msg.String())
		if isyyact {
			writeFile(path, pbname, "")
		} else {
			writeBig(path, pbname, "")
		}
	}()

	msg.WriteString("//")
	msg.WriteString(path)
	msg.WriteString(end_chat)
	msg.WriteString("message ")
	msg.WriteString(pbname)
	msg.WriteString("{")
	wfile.WriteString(msg.String())
	msg = bytes.NewBuffer([]byte{})
	//解析每一列
	id := 1
	for i < size {
		str := strings.TrimSpace(lines[i])
		i++
		if str[0] == ';' || str[0] == '_' {
			continue //注释行(忽略行，仅客户端读取)
		}
		pos := strings.IndexByte(str, '=')
		if pos < 1 {
			continue //结构不完整
		}
		typ := strings.TrimSpace(str[:pos])
		val := strings.TrimSpace(str[pos+1:])
		msg.WriteString(end_begin)
		if strings.HasPrefix(typ, "[{") || strings.HasPrefix(typ, "{") {
			//解析需要生成子对象的结构
			pos := -1
			if typ[0] == '[' {
				pos = strings.LastIndexByte(typ, ']')
			} else {
				pos = strings.LastIndexByte(typ, '}')
			}
			js := typ[:pos+1]
			par := typ[pos+1:]
			typ = build(par, js)
			//添加字段名
			msg.WriteString(typ)
			msg.WriteString(" ")
			msg.WriteString(par)
		} else if strings.HasPrefix(val, "[{") || strings.HasPrefix(val, "{") {
			pbname := buildJson(typ, val)
			msg.WriteString(pbname)
			msg.WriteString(" ")
			msg.WriteString(typ)
		} else {
			if regInit(typ) { //设置了类型
				str := findSubStr(typ)
				typ = strings.TrimPrefix(typ, str)
				if strings.HasPrefix(str, "[]") {
					msg.WriteString("repeated ")
					str = strings.TrimPrefix(str, "[]")
				}
				msg.WriteString(str)
				msg.WriteString(" ")
			} else {
				//类型推导
				if strings.HasPrefix(val, `"`) { //val = "123"
					msg.WriteString("string ")
				} else if regFloat(val) {
					msg.WriteString("float ")
				} else if regInt(val) {
					msg.WriteString("sint32 ")
				} else if regArray(val) {
					typ := chooseArray(val)
					if typ == 1 {
						msg.WriteString("repeated sint32 ")
					} else if typ == 2 {
						msg.WriteString("repeated float ")
					} else {
						msg.WriteString("repeated string ")
					}
				} else {
					msg.WriteString("string ")
				}
			}
			//添加字段名
			msg.WriteString(typ)
		}
		msg.WriteString(" = ")
		msg.WriteString(strconv.Itoa(id))
		msg.WriteString(";")
		id++
	}
}

func spile(a rune) bool {
	if a == '_' {
		return true
	}
	return false
}

func chooseArray(str string) int {
	typ := 3
	strs := strings.FieldsFunc(str, spile)
	for _, s := range strs {
		if regInt(s) {
			typ = 1
			continue
		} else if regFloat(s) {
			typ = 2
			continue
		} else {
			typ = 3
			break
		}
	}

	return typ
}
