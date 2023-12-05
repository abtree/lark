package api

import (
	"flag"
	"lark/com/utils"
	"lark/pb"
	"log"
	"math/rand"
	"runtime"
	"time"
)

/*
处理启动服务需要的初始化操作
*/

var (
	confFile = flag.String("cfg", "../../configs", "config path")
)

type MainInstance interface {
	Init() error
	RunLoop()
	Destory()
}

const (
	ServerType = "server_type"
	ServerId   = "server_id"
	ServerHost = "host"
	ServerPort = "port"
)

var App *Server

func Run(typ pb.ServerType, inst MainInstance) {
	flag.Parse()

	if inst == nil {
		log.Panicln("instance is nil, exit")
	}

	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	//新建应用
	App = NewServer(typ, inst)
	App.init()
	//启动服务
	go App.Run()
	//监听信号
	App.WaitSignal()
}

func SetServerType(typ pb.ServerType) {
	App.Envs[ServerType] = utils.ToString(typ)
}
