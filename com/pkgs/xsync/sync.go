package xsync

import (
	"errors"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"sync"
	"sync/atomic"
	"time"
)

type XSync struct {
	syncGuid int64
	waitMap  sync.Map
	TC       uint32
}

var instance *XSync

// 初始化异步结构
func getSync() *XSync {
	if instance != nil {
		return instance
	}
	instance = &XSync{
		TC: 20,
	}
	return instance
}

// 设置超时时间
func SetTimeOut(tc uint32) {
	if instance != nil {
		instance.TC = tc
	}
}

// 获取等待id
func GetGuid() int64 {
	inst := getSync()
	guid := atomic.AddInt64(&inst.syncGuid, 1)
	cc := make(chan interface{})
	inst.waitMap.Store(guid, cc)
	return guid
}

// 等待返回消息
func Wait(guid int64) (interface{}, error) {
	inst := getSync()
	c, ok := inst.waitMap.Load(guid)
	if !ok || c == nil {
		return nil, errors.New(utils.CONST_WAIT_CHANNEL_NOT_EXIST)
	}
	cc := c.(chan interface{})
	select {
	case msg := <-cc:
		close(cc)
		inst.waitMap.Delete(guid)
		return msg, nil
	case <-time.After(time.Duration(inst.TC) * time.Second):
		close(cc)
		inst.waitMap.Delete(guid)
		xlog.Warnf("Web load data time out %d", inst.TC)
		return nil, errors.New(utils.CONST_WAIT_CHANNEL_TIMEOUT)
	}
}

// 发送返回消息
func Sign(guid int64, msg interface{}) bool {
	inst := getSync()
	cc, ok := inst.waitMap.Load(guid)
	if ok {
		cc.(chan interface{}) <- msg
		return true
	}
	return false
}
