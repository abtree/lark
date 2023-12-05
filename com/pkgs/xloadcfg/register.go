package xloadcfg

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

const (
	PARSER_CG_EXT    uint8 = iota //后缀名
	PARSER_CG_PREFIX              //前缀字符串
	PARSER_CG_FILE                //文件名匹配
)

//读取文件回调函数
type PaserFunc func(path, nex string, field *reflect.Value)

//注册解释函数 和 筛选方式
type ParserKandler struct {
	Category uint8
	Parser   PaserFunc
	Param    string
	Ext      string //在以文件或文件夹注册时 可指定后缀名
}

//定义解析器类型
type parsers struct {
	ext_parser    map[string]PaserFunc
	prefix_parser map[string]PaserFunc
	file_parser   map[string]PaserFunc
	lock          sync.Mutex
}

//定义解析器
var handles = parsers{
	ext_parser:    make(map[string]PaserFunc),
	prefix_parser: make(map[string]PaserFunc),
	file_parser:   make(map[string]PaserFunc),
}

func (p *parsers) isRegister(pk *ParserKandler) bool {
	switch pk.Category {
	case PARSER_CG_EXT:
		{
			_, ok := handles.ext_parser[pk.Param]
			return ok
		}
	case PARSER_CG_PREFIX:
		{
			key := pk.Param
			if pk.Ext != "" {
				key = fmt.Sprintf("[%s]%s", pk.Ext, pk.Param)
			}
			_, ok := handles.prefix_parser[key]
			return ok
		}
	case PARSER_CG_FILE:
		{
			key := pk.Param
			if pk.Ext != "" {
				key = fmt.Sprintf("[%s]%s", pk.Ext, pk.Param)
			}
			_, ok := handles.file_parser[key]
			return ok
		}
	}
	return false
}

func (p *parsers) matchFunc(name, ext string) (PaserFunc, bool) {
	//优先匹配文件名
	key := fmt.Sprintf("[%s]%s", ext, name)
	if f, ok := p.file_parser[key]; ok {
		return f, true
	}
	if f, ok := p.file_parser[name]; ok {
		return f, true
	}
	//其次匹配前缀
	for k, v := range p.prefix_parser {
		if ok := strings.HasPrefix(key, k); ok {
			return v, true
		}
		if ok := strings.HasPrefix(name, k); ok {
			return v, true
		}
	}
	//最后匹配后缀名
	if r, ok := p.ext_parser[ext]; ok {
		return r, true
	}
	return nil, false
}

//注册配置的解析函数
func RegisterParser(pk *ParserKandler) bool {
	handles.lock.Lock()
	defer handles.lock.Unlock()

	if ok := handles.isRegister(pk); ok {
		return false
	}
	switch pk.Category {
	case PARSER_CG_EXT:
		{
			if handles.ext_parser == nil {
				handles.ext_parser = make(map[string]PaserFunc)
			}
			handles.ext_parser[pk.Param] = pk.Parser
		}
	case PARSER_CG_PREFIX:
		{
			if handles.prefix_parser == nil {
				handles.prefix_parser = make(map[string]PaserFunc)
			}
			key := pk.Param
			if pk.Ext != "" {
				key = fmt.Sprintf("[%s]%s", pk.Ext, pk.Param)
			}
			handles.prefix_parser[key] = pk.Parser
		}
	case PARSER_CG_FILE:
		{
			if handles.file_parser == nil {
				handles.file_parser = make(map[string]PaserFunc)
			}
			key := pk.Param
			if pk.Ext != "" {
				key = fmt.Sprintf("[%s]%s", pk.Ext, pk.Param)
			}
			handles.file_parser[key] = pk.Parser
		}
	}
	return true
}
