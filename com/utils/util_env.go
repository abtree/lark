package utils

import (
	"os"
	"strings"
)

/*
	获取环境变量
	envs 为需要获取的环境变量列表 可以带默认值
*/
func GetEnvs(envs map[string]string) {
	//读取环境变量 填充数据
	dp := os.Environ()
	for _, e := range dp {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			if _, ok := envs[parts[0]]; ok {
				envs[parts[0]] = parts[1]
			}
		}
	}
}
