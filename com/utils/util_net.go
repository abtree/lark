package utils

import (
	"lark/com/pkgs/xlog"
	"net"
	"strings"
)

//通过读取网络配置 获取当前ip (当存在多网卡时 查询并不准确)
func GetServerIp() string {
	inters, err := net.Interfaces()
	if err != nil {
		xlog.Error(err.Error())
		return "127.0.0.1"
	}
	for _, inter := range inters {
		if (inter.Flags&net.FlagUp) != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				xlog.Error(err.Error())
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return "127.0.0.1"
}
