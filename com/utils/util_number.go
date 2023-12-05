package utils

import (
	"lark/com/pkgs/xlog"
	"reflect"
	"strconv"
)

//将 number(包括自定义number 类型) 转换为 string
func ToString(val interface{}) (ret string) {
	switch val.(type) {
	case int:
		ret = strconv.FormatInt(int64(val.(int)), 10)
	case int8:
		ret = strconv.FormatInt(int64(val.(int8)), 10)
	case int16:
		ret = strconv.FormatInt(int64(val.(int16)), 10)
	case int32:
		ret = strconv.FormatInt(int64(val.(int32)), 10)
	case int64:
		ret = strconv.FormatInt(val.(int64), 10)
	case float32:
		ret = strconv.FormatFloat(float64(val.(float32)), 'f', -1, 64)
	case float64:
		ret = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case uint8:
		ret = strconv.FormatUint(uint64(val.(uint8)), 10)
	case uint16:
		ret = strconv.FormatUint(uint64(val.(uint16)), 10)
	case uint32:
		ret = strconv.FormatUint(uint64(val.(uint32)), 10)
	case uint64:
		ret = strconv.FormatUint(val.(uint64), 10)
	case bool:
		if val.(bool) {
			ret = "true"
		} else {
			ret = "false"
		}
	case string:
		ret = val.(string)
	default:
		if reflect.TypeOf(val).Kind() == reflect.Int32 {
			ret = strconv.FormatInt(reflect.ValueOf(val).Int(), 10)
		}
	}
	return
}

func StrToInt32(str string) int32 {
	i, e := strconv.ParseInt(str, 10, 32)
	if e != nil {
		xlog.Error(e.Error())
		return 0
	}
	return int32(i)
}

func StrToInt64(str string) int64 {
	i, e := strconv.ParseInt(str, 10, 64)
	if e != nil {
		xlog.Error(e.Error())
		return 0
	}
	return i
}

func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		xlog.Error(e.Error())
		return 0
	}
	return i
}

func StrToUint32(str string) uint32 {
	i, e := strconv.ParseUint(str, 10, 32)
	if e != nil {
		xlog.Error(e.Error())
		return 0
	}
	return uint32(i)
}

func StrToUint64(str string) uint64 {
	i, e := strconv.ParseUint(str, 10, 64)
	if e != nil {
		xlog.Error(e.Error())
		return 0
	}
	return i
}

func StrToUint(str string) uint {
	i, e := strconv.ParseUint(str, 10, 64)
	if e != nil {
		xlog.Error(e.Error())
		return 0
	}
	return uint(i)
}

func StrToFloat(str string) float32 {
	f, e := strconv.ParseFloat(str, 32)
	if e != nil {
		xlog.Error(e.Error())
		return 0
	}
	return float32(f)
}

func StrToDouble(str string) float64 {
	f, e := strconv.ParseFloat(str, 64)
	if e != nil {
		xlog.Error(e.Error())
		return 0
	}
	return f
}

func StrToBool(str string) bool {
	b, e := strconv.ParseBool(str)
	if e != nil {
		xlog.Error(e.Error())
		return false
	}
	return b
}
