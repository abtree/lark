package xcfgstructpb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var strbuilder strings.Builder

func parserJson(path, nam, dat string, isyyact bool) {
	pbname := buildJson(nam, dat)
	if _, ok := had_pbs[pbname]; ok {
		fmt.Printf("json pb %s has repeated \n", pbname)
	}
	if isyyact {
		writeFile(path, pbname, "")
	} else {
		writeBig(path, pbname, "")
	}
}

func buildJson(nam, dat string) (pbname string) {
	var config interface{}
	strbuilder = strings.Builder{}
	err := json.Unmarshal([]byte(dat), &config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if reflect.TypeOf(config).Kind() == reflect.Slice {
		m := config.([]interface{})[0]
		parseMap(nam, m.(map[string]interface{}))
		pbname = "repeated J" + nam
	} else {
		m := config.(map[string]interface{})
		parseMap(nam, m)
		pbname = "J" + nam
	}

	wfile.WriteString(strbuilder.String())
	return
}

func parseMap(nam string, m map[string]interface{}) {
	var seq = 1
	strbuilder.WriteString("message J" + nam + "{\n")
	for k, v := range m {
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Slice {
			x := v.([]interface{})[0]
			if reflect.TypeOf(x).Kind() == reflect.Map {
				parseMap(k, x.(map[string]interface{}))
				strbuilder.WriteString("\t repeated J" + k + " " + k + " = " + strconv.Itoa(seq) + ";\n")
			} else if reflect.TypeOf(x).Kind() == reflect.Float64 {
				t := switchTyp(k, x)
				strbuilder.WriteString("\t repeated " + t + " " + k + " = " + strconv.Itoa(seq) + ";\n")
			} else {
				strbuilder.WriteString("\t repeated " + reflect.TypeOf(x).Name() + " " + k + " = " + strconv.Itoa(seq) + ";\n")
			}
			// for _, x := range v.([]interface{}) {
			// 	log.Println(reflect.TypeOf(x).Name())
			// }
		} else if kind == reflect.Map {
			parseMap(k, v.(map[string]interface{}))
			strbuilder.WriteString("\t J" + k + " " + k + " = " + strconv.Itoa(seq) + ";\n")
		} else if kind == reflect.Float64 {
			t := switchTyp(k, v)
			strbuilder.WriteString("\t " + t + " " + k + " = " + strconv.Itoa(seq) + ";\n")
		} else {
			strbuilder.WriteString("\t " + reflect.TypeOf(v).Name() + " " + k + " = " + strconv.Itoa(seq) + ";\n")
		}
		seq++
	}
	strbuilder.WriteString("} \n")
}

func keyHasPrefix(k string) string {
	arrs := []string{"double", "float", "int32", "int64", "uint32", "uint64", "bool"}
	for _, typ := range arrs {
		if strings.HasPrefix(k, typ) {
			return typ
		}
	}
	return ""
}

func switchTyp(k string, v interface{}) string {
	t := keyHasPrefix(k)
	if t != "" {
		return t
	}
	s := strconv.FormatFloat(v.(float64), 'f', -1, 64)
	if _, err := strconv.Atoi(s); err == nil {
		return "sint32"
	} else {
		return "double"
	}
}
