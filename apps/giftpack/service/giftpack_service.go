package service

import (
	"lark/com/pkgs/xgiftpack"
	"lark/com/pkgs/xlog"
	"lark/com/pkgs/xsync"
	"lark/com/tools"
	"lark/pb"
)

type GiftpackTask struct {
	xsync.XSyncItem
	msg *pb.GiftPackProto
}

type GiftpackService struct {
	giftpacks map[uint32]*pb.GiftPack //礼包码数据
	index     uint32                  //礼包码id使用记录
	key       uint32                  //使用掉的礼包码
	c         chan *GiftpackTask      //礼包码操作通道
	exit      chan struct{}           //该服务关闭通道
}

func NewGiftpackService() *GiftpackService {
	return &GiftpackService{
		giftpacks: map[uint32]*pb.GiftPack{},
		c:         make(chan *GiftpackTask),
		exit:      make(chan struct{}),
		key:       1,
	}
}

// 初始化服务
func (t *GiftpackService) Init() {
	cfg, err := tools.GetAllCfg()
	if err != nil {
		xlog.Warn(err.Error())
		return
	}
	xgiftpack.NewGiftPack(cfg.Configs.Giftpack)
	go t.Run()
}

// 开启消息监听
func (t *GiftpackService) Run() {
	select {
	case task := <-t.c:
		t.OnProto(task)
	case <-t.exit:
		return
	}
}

// 关闭服务
func (t *GiftpackService) Exit() {
	t.exit <- struct{}{}
}

// 向服务注入任务
func (t *GiftpackService) PostTask(syncid int64, msg *pb.GiftPackProto) {
	task := &GiftpackTask{
		msg: msg,
	}
	task.SetSyncId(syncid)
	t.c <- task
}

// 处理任务
func (t *GiftpackService) OnProto(msg *GiftpackTask) {
	switch msg.msg.Op {
	case pb.GiftPackProto_Create:
		t.OnCreate(msg)
	case pb.GiftPackProto_Update:
		t.OnUpdate(msg)
	case pb.GiftPackProto_GetAll:
		t.OnGetAll(msg)
	}
}

func (t *GiftpackService) OnCreate(msg *GiftpackTask) {
	gp := msg.msg.Giftpack
	t.index++
	gp.Id = t.index
	if !gp.IsShare {
		//不是共享码
		gp.Code = []*pb.GiftPackSlice{}
		n := &pb.GiftPackSlice{
			Start: int32(t.key),
			End:   int32(t.key + gp.Count),
		}
		t.key += gp.Count
		gp.Code = append(gp.Code, n)
	}
	t.giftpacks[gp.Id] = gp
	xlog.Infof("create giftpack %d finished \n", gp.Id)
	msg.Return(msg.msg)
	// SendToHttp(msg)
}

func (t *GiftpackService) OnUpdate(msg *GiftpackTask) {
	gp, ok := t.giftpacks[msg.msg.Giftpack.Id]
	if !ok || gp == nil {
		return
	}
	gp.Prizes = msg.msg.Giftpack.Prizes
	if msg.msg.Giftpack.Count > gp.Count {
		add := msg.msg.Giftpack.Count - gp.Count
		n := &pb.GiftPackSlice{
			Start: int32(t.key),
			End:   int32(t.key + add),
		}
		t.key += add
		gp.Code = append(gp.Code, n)
	}
	gp.Count = msg.msg.Giftpack.Count
	xlog.Infof("update giftpack %d finished \n", gp.Id)
}

func (t *GiftpackService) OnGetAll(msg *GiftpackTask) {
	gp, ok := t.giftpacks[msg.msg.Giftpack.Id]
	if !ok || gp == nil {
		return
	}
	for _, v := range gp.Code {
		for i := v.Start; i < v.End; i++ {
			src := xgiftpack.GetKey(int(i))
			xlog.Debugf("index %d key %s \n", i, gp.Prefix+src)
		}
	}
}
