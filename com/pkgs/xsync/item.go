package xsync

import (
	"google.golang.org/protobuf/proto"
)

type XSyncItem struct {
	SyncId int64
}

func (t *XSyncItem) SetSyncId(id int64) {
	t.SyncId = id
}

func (t *XSyncItem) Return(msg proto.Message) {
	if t.SyncId <= 0 {
		return //没有返回消息
	}
	Sign(t.SyncId, msg)
}
