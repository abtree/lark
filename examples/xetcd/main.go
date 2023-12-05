package main

import (
	"fmt"
	"lark/com/pkgs/xetcd"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	go Write()
	Read()
}

func Write() {
	xetcd.RegisterEtcd("lark", []string{"192.168.0.174:30379"}, "127.0.0.1", 8081, "example", xetcd.TIME_TO_LIVE)
}

func Read() {
	r := xetcd.NewGetEtcd([]string{"192.168.0.174:30379"}, "lark", "example", "", "")
	back := r.Get(5*time.Second, Update)
	addr := []string{}
	for i := range back {
		addr = append(addr, string(back[i].Value))
	}
	fmt.Println(addr)
}

func Update(ev *clientv3.Event) {
	fmt.Println(ev.Kv.Version)
	switch ev.Type {
	case mvccpb.PUT:
		fmt.Println("Put", string(ev.Kv.Value))
	case mvccpb.DELETE:
		fmt.Println("DELETE", string(ev.Kv.Key))
	default:
	}
}
