package xcfgstructpb

import (
	"fmt"
	"regexp"
)

/*
	解析配置文件需要用到的正则表达式
*/
var (
	rFloat *regexp.Regexp //匹配float数字
	rInt   *regexp.Regexp //匹配int数字
	rInit  *regexp.Regexp //匹配到第一个大写字母
	rSub   *regexp.Regexp //从第一个大写字母切子串
	rArray *regexp.Regexp //匹配数组(以下划线'_'分割)
)

func init() {
	var err error
	rFloat, err = regexp.Compile(`^[-+]?[0-9]+\.[0-9]+$`)
	if err != nil {
		fmt.Printf("regexp float error with : %v", err)
	}
	rInt, err = regexp.Compile(`^[-+]?\d+$`)
	if err != nil {
		fmt.Printf("regexp int error with : %v", err)
	}
	rInit, err = regexp.Compile(`^(\[\])?[a-z]`)
	if err != nil {
		fmt.Printf("regexp init with alpha error with : %v", err)
	}
	rSub, err = regexp.Compile(`^(\[\])?[a-z][a-z0-9_]*[A-Z]`)
	if err != nil {
		fmt.Printf("regexp sub string error with : %v", err)
	}
	rArray, err = regexp.Compile(`^(.+_)+`)
	if err != nil {
		fmt.Printf("regexp sub string error with : %v", err)
	}
}

func regFloat(str string) bool {
	if rFloat.MatchString(str) {
		return true
	}
	return false
}

func regInt(str string) bool {
	if rInt.MatchString(str) {
		return true
	}
	return false
}

func regInit(str string) bool {
	if rInit.MatchString(str) {
		return true
	}
	return false
}

func regArray(str string) bool {
	if rArray.MatchString(str) {
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
