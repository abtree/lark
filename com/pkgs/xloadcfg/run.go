package xloadcfg

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

type cfgColumn struct {
	typ    interface{}
	name   string
	oristr string
}

type cfgTitle struct {
	key     []int
	multi   string
	columns map[int]*cfgColumn
	pbname  string
}

func Run(path string, dat, yyact interface{}) {
	ReloadPath(path, dat, yyact)
}

func initValue(tv *reflect.Value) {
	//这里初始化
	if tv.Type().Kind() == reflect.Ptr {
		if tv.Type().Elem().Kind() == reflect.Struct {
			if tv.IsNil() {
				vvv := reflect.New(tv.Type().Elem())
				tv.Set(vvv)
			}
		} else if tv.Type().Elem().Kind() == reflect.Slice {
			if tv.IsNil() {
				slic := reflect.MakeSlice(tv.Type(), 0, 0)
				tv.Set(slic)
			}
		}
	}
}

func newPB(name string) reflect.Value {
	return reflect.New(proto.MessageType(name).Elem())
}

func fullField(field *reflect.Value, value string, split string) bool {
	switch field.Type().Kind() {
	case reflect.String,
		reflect.Int32,
		reflect.Int,
		reflect.Int64,
		reflect.Uint32,
		reflect.Uint,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Bool:
		{
			v, ok := parseValue(field.Type().Kind(), value)
			if !ok {
				return false
			}
			field.Set(reflect.ValueOf(v).Convert(field.Type()))
		}
	case reflect.Slice:
		{
			typ := field.Type().Elem().Kind()
			vals := []string{}
			if split != "" {
				vals = strings.Split(value, split)
			} else {
				vals = append(vals, value)
			}
			for _, vs := range vals {
				v, ok := parseValue(typ, vs)
				if !ok {
					return false
				}
				field.Set(reflect.Append(*field, reflect.ValueOf(v).Convert(field.Type().Elem())))
			}
		}
	case reflect.Ptr:
		{
			vals := []string{}
			if split != "" {
				vals = strings.Split(value, split)
			} else {
				vals = append(vals, value)
			}
			initValue(field)
			// log.Println(field.Elem().NumField())
			for i, vs := range vals {
				fi := field.Elem().Field(i)
				fullField(&fi, vs, "")
			}
		}
	}
	return true
}

func parseValue(typ reflect.Kind, value string) (interface{}, bool) {
	switch typ {
	case reflect.String:
		{
			return value, true
		}
	case reflect.Int32:
		{
			f, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalln(err.Error())
				return nil, false
			}
			return int32(f), true
		}
	case reflect.Int,
		reflect.Int64:
		{
			f, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				log.Fatalln(err.Error())
				return nil, false
			}
			return f, true
		}
	case reflect.Uint32:
		{
			f, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				log.Fatalln(err.Error())
				return nil, false
			}
			return f, true
		}
	case reflect.Uint,
		reflect.Uint64:
		{
			f, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.Fatalln(err.Error())
				return nil, false
			}
			return f, true
		}
	case reflect.Float32:
		{
			f, err := strconv.ParseFloat(value, 10)
			if err != nil {
				log.Fatalln(err.Error())
				return nil, false
			}
			return float32(f), true
		}
	case reflect.Float64:
		{
			f, err := strconv.ParseFloat(value, 10)
			if err != nil {
				log.Fatalln(err.Error())
				return nil, false
			}
			return f, true
		}
	case reflect.Bool:
		{
			f, err := strconv.ParseBool(value)
			if err != nil {
				log.Fatalln(err.Error())
				return nil, false
			}
			return f, true
		}
	}
	return nil, false
}
