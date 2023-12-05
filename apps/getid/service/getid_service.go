package service

import (
	"lark/com/api"
	"lark/com/pkgs/xsync"
	"lark/pb"
)

type GetidService struct {
	CurAtoi uint32
	c       chan *GetidTask
	exit    chan struct{}
}

type GetidTask struct {
	xsync.XSyncItem
}

func NewGetidService() *GetidService {
	return &GetidService{
		c:    make(chan *GetidTask),
		exit: make(chan struct{}),
	}
}

func (t *GetidService) Init() {
	t.CurAtoi = 1
	go t.Run()
}

// 开启消息监听
func (t *GetidService) Run() {
	for {
		select {
		case task := <-t.c:
			t.OnProto(task)
		case <-t.exit:
			return
		}
	}
}

// 关闭服务
func (t *GetidService) Exit() {
	t.exit <- struct{}{}
}

// 向服务注入任务
func (t *GetidService) PostTask(syncid int64) {
	task := &GetidTask{}
	task.SetSyncId(syncid)
	t.c <- task
}

func (t *GetidService) OnProto(task *GetidTask) {
	ret := &pb.GetidProto{}
	ret.Id = api.App.SrvCfg.Id<<24 + t.CurAtoi
	t.CurAtoi++
	task.Return(ret)
}
