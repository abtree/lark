package xloadcfg

import (
	"encoding/json"
	"log"
	"reflect"
)

func init() {
	pk := &ParserKandler{
		Category: PARSER_CG_EXT,
		Parser:   decodeJson,
		Param:    ".json.txt",
	}
	RegisterParser(pk)
	pk.Param = ".json"
	RegisterParser(pk)
}

func decodeJson(dat, nex string, field *reflect.Value) {
	initValue(field)
	v := field.Interface()
	dojson(dat, &v)
}

func dojson(dat string, v interface{}) {
	err := json.Unmarshal([]byte(dat), v)
	if err != nil {
		log.Panicf(err.Error())
	}
}
