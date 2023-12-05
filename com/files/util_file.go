package files

import (
	"io"
	"log"
	"os"
	"strings"
)

// path 是否是文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 根据path 创建文件夹
func MKDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return os.Chmod(path, os.ModePerm)
}

// 读取整个文件
func ReadAll(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Open file [%s] error %s", path, err.Error())
		return ""
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	str := string(byteValue)
	if strings.HasPrefix(str, "\uFEFF") {
		str = strings.TrimPrefix(str, "\uFEFF") //去除utf8 with bom中的bom
	}
	return str
}
