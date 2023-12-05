package api

import (
	"lark/com/pkgs/xloadcfg"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Inst   MainInstance
	Signal chan os.Signal
	SysCfg *pb.MsgSysConfigs
	SrvCfg *pb.CfgServers
	Envs   map[string]string
}

func NewServer(typ pb.ServerType, inst MainInstance) *Server {
	svr := &Server{
		Inst: inst,
		Envs: make(map[string]string),
	}
	svr.Envs[ServerType] = utils.ToString(typ)
	return svr
}

func (s *Server) init() {
	//读取系统配置
	s.SysCfg = &pb.MsgSysConfigs{}
	xloadcfg.Run(*confFile, s.SysCfg, nil)
	//读取环境变量
	s.getEnv()
	//初始化日志系统
	xlog.Shared(s.SysCfg.Logger, s.SrvCfg.Name)

	err := s.Inst.Init()
	if err != nil {
		log.Panic(err)
	}
	//注册消息
	if x, ok := s.Inst.(IMsgServer); ok {
		x.RegisterMsg()
	}
}

// 设置默认值
func (s *Server) defEnv() {
	// s.Envs = map[string]string{}
	// s.Envs[ServerType] = "1"
	s.Envs[ServerHost] = "127.0.0.1"
}

func (s *Server) getEnv() {
	s.defEnv()
	utils.GetEnvs(s.Envs)
	st := utils.StrToUint32(s.Envs[ServerType])
	if svr, ok := s.SysCfg.Servers[st]; !ok {
		xlog.Warnf("server type %d not configs", st)
		return
	} else {
		s.SrvCfg = svr
	}
	s.SrvCfg.Host = s.Envs[ServerHost]
	if p, o := s.Envs[ServerId]; o {
		s.SrvCfg.ServerId = utils.StrToUint32(p)
	}
	if p, o := s.Envs[ServerPort]; o {
		s.SrvCfg.Port = utils.StrToUint32(p)
	}
}

func (s *Server) Run() {
	s.Inst.RunLoop()
}

func (s *Server) WaitSignal() {
	s.Signal = make(chan os.Signal, 1)
	signal.Notify(s.Signal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		c := <-s.Signal
		switch c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.Inst.Destory()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func (s *Server) GetSvrNameBySvrType(typ pb.ServerType) string {
	cfg, ok := s.SysCfg.Servers[uint32(typ)]
	if !ok {
		return ""
	}
	return cfg.Name
}

// func (s *Server) SetServerType(typ pb.ServerType) {
// 	s.Envs[ServerType] = utils.ToString(typ)
// }
