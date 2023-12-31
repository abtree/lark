// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v4.22.1
// source: auto_config.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Jgiftpack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtroPos []int32 `protobuf:"zigzag32,1,rep,packed,name=extroPos,proto3" json:"extroPos,omitempty"`
	Mix      []int32 `protobuf:"zigzag32,2,rep,packed,name=mix,proto3" json:"mix,omitempty"`
	Chars    string  `protobuf:"bytes,3,opt,name=chars,proto3" json:"chars,omitempty"`
	Key      string  `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	BasePos  []int32 `protobuf:"zigzag32,5,rep,packed,name=basePos,proto3" json:"basePos,omitempty"`
}

func (x *Jgiftpack) Reset() {
	*x = Jgiftpack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jgiftpack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jgiftpack) ProtoMessage() {}

func (x *Jgiftpack) ProtoReflect() protoreflect.Message {
	mi := &file_auto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jgiftpack.ProtoReflect.Descriptor instead.
func (*Jgiftpack) Descriptor() ([]byte, []int) {
	return file_auto_config_proto_rawDescGZIP(), []int{0}
}

func (x *Jgiftpack) GetExtroPos() []int32 {
	if x != nil {
		return x.ExtroPos
	}
	return nil
}

func (x *Jgiftpack) GetMix() []int32 {
	if x != nil {
		return x.Mix
	}
	return nil
}

func (x *Jgiftpack) GetChars() string {
	if x != nil {
		return x.Chars
	}
	return ""
}

func (x *Jgiftpack) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Jgiftpack) GetBasePos() []int32 {
	if x != nil {
		return x.BasePos
	}
	return nil
}

//Websocket
type WebSocketCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WriteWait          uint32 `protobuf:"varint,1,opt,name=WriteWait,proto3" json:"WriteWait,omitempty"`
	PongWait           uint32 `protobuf:"varint,2,opt,name=PongWait,proto3" json:"PongWait,omitempty"`
	PingWait           uint32 `protobuf:"varint,3,opt,name=PingWait,proto3" json:"PingWait,omitempty"`
	ReadBufferSize     uint32 `protobuf:"varint,4,opt,name=ReadBufferSize,proto3" json:"ReadBufferSize,omitempty"`
	WriteBufferSize    uint32 `protobuf:"varint,5,opt,name=WriteBufferSize,proto3" json:"WriteBufferSize,omitempty"`
	ReadMaxBufferSize  uint32 `protobuf:"varint,6,opt,name=ReadMaxBufferSize,proto3" json:"ReadMaxBufferSize,omitempty"`
	ChanRegisterSize   uint32 `protobuf:"varint,7,opt,name=ChanRegisterSize,proto3" json:"ChanRegisterSize,omitempty"`
	ChanUnregisterSize uint32 `protobuf:"varint,8,opt,name=ChanUnregisterSize,proto3" json:"ChanUnregisterSize,omitempty"`
	ChanReadSize       uint32 `protobuf:"varint,9,opt,name=ChanReadSize,proto3" json:"ChanReadSize,omitempty"`
	ChanWriteSize      uint32 `protobuf:"varint,10,opt,name=ChanWriteSize,proto3" json:"ChanWriteSize,omitempty"`
	RoutineRead        uint32 `protobuf:"varint,11,opt,name=RoutineRead,proto3" json:"RoutineRead,omitempty"`
	MaxConn            uint32 `protobuf:"varint,12,opt,name=MaxConn,proto3" json:"MaxConn,omitempty"`
	RWDeadLine         uint32 `protobuf:"varint,13,opt,name=RWDeadLine,proto3" json:"RWDeadLine,omitempty"`
}

func (x *WebSocketCfg) Reset() {
	*x = WebSocketCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebSocketCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebSocketCfg) ProtoMessage() {}

func (x *WebSocketCfg) ProtoReflect() protoreflect.Message {
	mi := &file_auto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebSocketCfg.ProtoReflect.Descriptor instead.
func (*WebSocketCfg) Descriptor() ([]byte, []int) {
	return file_auto_config_proto_rawDescGZIP(), []int{1}
}

func (x *WebSocketCfg) GetWriteWait() uint32 {
	if x != nil {
		return x.WriteWait
	}
	return 0
}

func (x *WebSocketCfg) GetPongWait() uint32 {
	if x != nil {
		return x.PongWait
	}
	return 0
}

func (x *WebSocketCfg) GetPingWait() uint32 {
	if x != nil {
		return x.PingWait
	}
	return 0
}

func (x *WebSocketCfg) GetReadBufferSize() uint32 {
	if x != nil {
		return x.ReadBufferSize
	}
	return 0
}

func (x *WebSocketCfg) GetWriteBufferSize() uint32 {
	if x != nil {
		return x.WriteBufferSize
	}
	return 0
}

func (x *WebSocketCfg) GetReadMaxBufferSize() uint32 {
	if x != nil {
		return x.ReadMaxBufferSize
	}
	return 0
}

func (x *WebSocketCfg) GetChanRegisterSize() uint32 {
	if x != nil {
		return x.ChanRegisterSize
	}
	return 0
}

func (x *WebSocketCfg) GetChanUnregisterSize() uint32 {
	if x != nil {
		return x.ChanUnregisterSize
	}
	return 0
}

func (x *WebSocketCfg) GetChanReadSize() uint32 {
	if x != nil {
		return x.ChanReadSize
	}
	return 0
}

func (x *WebSocketCfg) GetChanWriteSize() uint32 {
	if x != nil {
		return x.ChanWriteSize
	}
	return 0
}

func (x *WebSocketCfg) GetRoutineRead() uint32 {
	if x != nil {
		return x.RoutineRead
	}
	return 0
}

func (x *WebSocketCfg) GetMaxConn() uint32 {
	if x != nil {
		return x.MaxConn
	}
	return 0
}

func (x *WebSocketCfg) GetRWDeadLine() uint32 {
	if x != nil {
		return x.RWDeadLine
	}
	return 0
}

type MsgConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unhandle  map[string][]byte `protobuf:"bytes,1,rep,name=unhandle,proto3" json:"unhandle,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Giftpack  *Jgiftpack        `protobuf:"bytes,2,opt,name=Giftpack,proto3" json:"Giftpack,omitempty"`
	Websocket *WebSocketCfg     `protobuf:"bytes,3,opt,name=Websocket,proto3" json:"Websocket,omitempty"`
}

func (x *MsgConfigs) Reset() {
	*x = MsgConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auto_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgConfigs) ProtoMessage() {}

func (x *MsgConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_auto_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgConfigs.ProtoReflect.Descriptor instead.
func (*MsgConfigs) Descriptor() ([]byte, []int) {
	return file_auto_config_proto_rawDescGZIP(), []int{2}
}

func (x *MsgConfigs) GetUnhandle() map[string][]byte {
	if x != nil {
		return x.Unhandle
	}
	return nil
}

func (x *MsgConfigs) GetGiftpack() *Jgiftpack {
	if x != nil {
		return x.Giftpack
	}
	return nil
}

func (x *MsgConfigs) GetWebsocket() *WebSocketCfg {
	if x != nil {
		return x.Websocket
	}
	return nil
}

type MsgYYactConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unhandle map[string][]byte `protobuf:"bytes,1,rep,name=unhandle,proto3" json:"unhandle,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MsgYYactConfigs) Reset() {
	*x = MsgYYactConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auto_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgYYactConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgYYactConfigs) ProtoMessage() {}

func (x *MsgYYactConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_auto_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgYYactConfigs.ProtoReflect.Descriptor instead.
func (*MsgYYactConfigs) Descriptor() ([]byte, []int) {
	return file_auto_config_proto_rawDescGZIP(), []int{3}
}

func (x *MsgYYactConfigs) GetUnhandle() map[string][]byte {
	if x != nil {
		return x.Unhandle
	}
	return nil
}

type MsgAllConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs *MsgConfigs      `protobuf:"bytes,1,opt,name=Configs,proto3" json:"Configs,omitempty"`
	Yyacts  *MsgYYactConfigs `protobuf:"bytes,2,opt,name=Yyacts,proto3" json:"Yyacts,omitempty"`
}

func (x *MsgAllConfigs) Reset() {
	*x = MsgAllConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auto_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAllConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAllConfigs) ProtoMessage() {}

func (x *MsgAllConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_auto_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAllConfigs.ProtoReflect.Descriptor instead.
func (*MsgAllConfigs) Descriptor() ([]byte, []int) {
	return file_auto_config_proto_rawDescGZIP(), []int{4}
}

func (x *MsgAllConfigs) GetConfigs() *MsgConfigs {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *MsgAllConfigs) GetYyacts() *MsgYYactConfigs {
	if x != nil {
		return x.Yyacts
	}
	return nil
}

var File_auto_config_proto protoreflect.FileDescriptor

var file_auto_config_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x7b, 0x0a, 0x09, 0x4a, 0x67, 0x69, 0x66, 0x74,
	0x70, 0x61, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x72, 0x6f, 0x50, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x11, 0x52, 0x08, 0x65, 0x78, 0x74, 0x72, 0x6f, 0x50, 0x6f, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x78, 0x18, 0x02, 0x20, 0x03, 0x28, 0x11, 0x52, 0x03, 0x6d,
	0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61,
	0x73, 0x65, 0x50, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x11, 0x52, 0x07, 0x62, 0x61, 0x73,
	0x65, 0x50, 0x6f, 0x73, 0x22, 0xe6, 0x03, 0x0a, 0x0c, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x43, 0x66, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x57, 0x61,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x57,
	0x61, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f, 0x6e, 0x67, 0x57, 0x61, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6f, 0x6e, 0x67, 0x57, 0x61, 0x69, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x57, 0x61, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x57, 0x61, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x52,
	0x65, 0x61, 0x64, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x78, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61,
	0x78, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x43,
	0x68, 0x61, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x55,
	0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x52,
	0x65, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x43,
	0x68, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x43,
	0x68, 0x61, 0x6e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x57, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x52, 0x57, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x22, 0xde, 0x01,
	0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x38, 0x0a, 0x08,
	0x75, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x55,
	0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x6e,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x69, 0x66, 0x74, 0x70, 0x61,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x67,
	0x69, 0x66, 0x74, 0x70, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x47, 0x69, 0x66, 0x74, 0x70, 0x61, 0x63,
	0x6b, 0x12, 0x2e, 0x0a, 0x09, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x66, 0x67, 0x52, 0x09, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x1a, 0x3b, 0x0a, 0x0d, 0x55, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8d,
	0x01, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x59, 0x59, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x75, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x73, 0x67, 0x59, 0x59, 0x61,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x55, 0x6e, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x55, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x66,
	0x0a, 0x0d, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x28, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x59, 0x79, 0x61,
	0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x73, 0x67, 0x59, 0x59, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x06,
	0x59, 0x79, 0x61, 0x63, 0x74, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auto_config_proto_rawDescOnce sync.Once
	file_auto_config_proto_rawDescData = file_auto_config_proto_rawDesc
)

func file_auto_config_proto_rawDescGZIP() []byte {
	file_auto_config_proto_rawDescOnce.Do(func() {
		file_auto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_auto_config_proto_rawDescData)
	})
	return file_auto_config_proto_rawDescData
}

var file_auto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_auto_config_proto_goTypes = []interface{}{
	(*Jgiftpack)(nil),       // 0: pb.Jgiftpack
	(*WebSocketCfg)(nil),    // 1: pb.WebSocketCfg
	(*MsgConfigs)(nil),      // 2: pb.MsgConfigs
	(*MsgYYactConfigs)(nil), // 3: pb.MsgYYactConfigs
	(*MsgAllConfigs)(nil),   // 4: pb.MsgAllConfigs
	nil,                     // 5: pb.MsgConfigs.UnhandleEntry
	nil,                     // 6: pb.MsgYYactConfigs.UnhandleEntry
}
var file_auto_config_proto_depIdxs = []int32{
	5, // 0: pb.MsgConfigs.unhandle:type_name -> pb.MsgConfigs.UnhandleEntry
	0, // 1: pb.MsgConfigs.Giftpack:type_name -> pb.Jgiftpack
	1, // 2: pb.MsgConfigs.Websocket:type_name -> pb.WebSocketCfg
	6, // 3: pb.MsgYYactConfigs.unhandle:type_name -> pb.MsgYYactConfigs.UnhandleEntry
	2, // 4: pb.MsgAllConfigs.Configs:type_name -> pb.MsgConfigs
	3, // 5: pb.MsgAllConfigs.Yyacts:type_name -> pb.MsgYYactConfigs
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_auto_config_proto_init() }
func file_auto_config_proto_init() {
	if File_auto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jgiftpack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebSocketCfg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auto_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgConfigs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auto_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgYYactConfigs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auto_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAllConfigs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auto_config_proto_goTypes,
		DependencyIndexes: file_auto_config_proto_depIdxs,
		MessageInfos:      file_auto_config_proto_msgTypes,
	}.Build()
	File_auto_config_proto = out.File
	file_auto_config_proto_rawDesc = nil
	file_auto_config_proto_goTypes = nil
	file_auto_config_proto_depIdxs = nil
}
