package xloadcfg

import (
	"io/ioutil"
	util_files "lark/com/files"
	"lark/com/utils"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func ReloadPath(path string, dat, yyact interface{}) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln("Read Configs error")
	}
	for _, file := range files {
		npath := filepath.Join(path, file.Name())
		if file.IsDir() {
			if file.Name() == "yyact" {
				//交给活动特殊处理
				walk_yyact(npath, "YY", reflect.ValueOf(yyact).Elem())
			} else {
				walk_files(npath, "", reflect.ValueOf(dat).Elem())
			}
		} else {
			handlefile(file, npath, "", reflect.ValueOf(dat).Elem())
		}
	}
}

func walk_files(path, ex string, cfg reflect.Value) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln("Read Configs error")
	}
	for _, file := range files {
		npath := filepath.Join(path, file.Name())
		if file.IsDir() {
			nex := ex + strings.Title(file.Name())
			walk_files(npath, nex, cfg)
		} else {
			handlefile(file, npath, ex, cfg)
		}
	}
}

func walk_yyact(path, ex string, cfg reflect.Value) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln("Read Configs error")
	}
	for _, file := range files {
		npath := filepath.Join(path, file.Name())
		pos := strings.LastIndexByte(file.Name(), '_')
		if pos == -1 {
			//当普通配置处理
			if file.IsDir() {
				walk_files(npath, ex, cfg)
			} else {
				handlefile(file, npath, ex, cfg)
			}
			continue
		}
		act := file.Name()[:pos]
		sub := file.Name()[pos+1:]
		pos = strings.Index(act, "_")
		fi := strings.Title(act[pos+1:])
		name := act[:pos] + fi
		nex := ex + name
		pos = strings.Index(sub, ".")
		var idx uint32
		if pos == -1 {
			x := utils.StrToInt32(sub)
			idx = uint32(x)
		} else {
			y := sub[:pos]
			x := utils.StrToUint32(y)
			idx = uint32(x)
		}
		field := cfg.FieldByName(nex)
		if field.IsNil() {
			field.Set(reflect.MakeMap(field.Type()))
		}
		fkey := reflect.ValueOf(idx)
		row := newPB("pb.Msg" + name)
		field.SetMapIndex(fkey, row)
		if file.IsDir() {
			walk_files(npath, "", row.Elem())
		} else {
			ext := sub[pos:]
			dat := util_files.ReadAll(npath)
			//获取结构对象
			curfield := row.Elem().FieldByName(fi)
			//获取结构对象
			if fn, ok := handles.matchFunc("", ext); ok {
				fn(dat, "", &curfield)
			}
		}
	}
}

func handlefile(file os.FileInfo, npath, ex string, cfg reflect.Value) {
	ext := filepath.Ext(file.Name())
	fi := strings.TrimSuffix(file.Name(), ext)

	if ext == ".txt" && strings.HasSuffix(fi, ".ini") {
		ext = ".ini.txt"
		fi = strings.TrimSuffix(fi, ".ini")
	}
	if ext == ".txt" && strings.HasSuffix(fi, ".json") {
		ext = ".json.txt"
		fi = strings.TrimSuffix(fi, ".json")
	}

	nex := ex + strings.Title(fi)
	dat := util_files.ReadAll(npath)
	field := cfg.FieldByName(nex)
	//获取结构对象
	if fn, ok := handles.matchFunc(nex, ext); ok {
		fn(dat, nex, &field)
	}
}
