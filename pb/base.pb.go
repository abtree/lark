// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v4.22.1
// source: base.proto

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

type EItemType int32

const (
	EItemType_EIT_UnUse EItemType = 0
	EItemType_EIT_Good  EItemType = 1 //道具
)

// Enum value maps for EItemType.
var (
	EItemType_name = map[int32]string{
		0: "EIT_UnUse",
		1: "EIT_Good",
	}
	EItemType_value = map[string]int32{
		"EIT_UnUse": 0,
		"EIT_Good":  1,
	}
)

func (x EItemType) Enum() *EItemType {
	p := new(EItemType)
	*p = x
	return p
}

func (x EItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[0].Descriptor()
}

func (EItemType) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[0]
}

func (x EItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EItemType.Descriptor instead.
func (EItemType) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type ErrorMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ErrorMsg) Reset() {
	*x = ErrorMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMsg) ProtoMessage() {}

func (x *ErrorMsg) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMsg.ProtoReflect.Descriptor instead.
func (*ErrorMsg) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorMsg) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type MsgStr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val string `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *MsgStr) Reset() {
	*x = MsgStr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgStr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgStr) ProtoMessage() {}

func (x *MsgStr) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgStr.ProtoReflect.Descriptor instead.
func (*MsgStr) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *MsgStr) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type CPrizeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MType   EItemType `protobuf:"varint,1,opt,name=MType,proto3,enum=pb.EItemType" json:"MType,omitempty"` //类型
	Oriname string    `protobuf:"bytes,2,opt,name=Oriname,proto3" json:"Oriname,omitempty"`                //原始名
	Count   int32     `protobuf:"zigzag32,3,opt,name=Count,proto3" json:"Count,omitempty"`                 //数量
	Exinfo  int32     `protobuf:"zigzag32,4,opt,name=Exinfo,proto3" json:"Exinfo,omitempty"`               //额外参数
	Inner   int32     `protobuf:"zigzag32,5,opt,name=inner,proto3" json:"inner,omitempty"`                 //内部参数
}

func (x *CPrizeItem) Reset() {
	*x = CPrizeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPrizeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPrizeItem) ProtoMessage() {}

func (x *CPrizeItem) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPrizeItem.ProtoReflect.Descriptor instead.
func (*CPrizeItem) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *CPrizeItem) GetMType() EItemType {
	if x != nil {
		return x.MType
	}
	return EItemType_EIT_UnUse
}

func (x *CPrizeItem) GetOriname() string {
	if x != nil {
		return x.Oriname
	}
	return ""
}

func (x *CPrizeItem) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CPrizeItem) GetExinfo() int32 {
	if x != nil {
		return x.Exinfo
	}
	return 0
}

func (x *CPrizeItem) GetInner() int32 {
	if x != nil {
		return x.Inner
	}
	return 0
}

type WebProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid int64  `protobuf:"zigzag64,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WebProto) Reset() {
	*x = WebProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebProto) ProtoMessage() {}

func (x *WebProto) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebProto.ProtoReflect.Descriptor instead.
func (*WebProto) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *WebProto) GetGuid() int64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

func (x *WebProto) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x30, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x1a, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x8f,
	0x01, 0x0a, 0x0a, 0x43, 0x50, 0x72, 0x69, 0x7a, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x23, 0x0a,
	0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x4d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x06, 0x45, 0x78, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x22, 0x32, 0x0a, 0x08, 0x57, 0x65, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x2a, 0x28, 0x0a, 0x09, 0x45, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x49, 0x54, 0x5f, 0x55, 0x6e, 0x55, 0x73, 0x65, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x45, 0x49, 0x54, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_base_proto_goTypes = []interface{}{
	(EItemType)(0),     // 0: pb.EItemType
	(*ErrorMsg)(nil),   // 1: pb.ErrorMsg
	(*MsgStr)(nil),     // 2: pb.MsgStr
	(*CPrizeItem)(nil), // 3: pb.CPrizeItem
	(*WebProto)(nil),   // 4: pb.WebProto
}
var file_base_proto_depIdxs = []int32{
	0, // 0: pb.CPrizeItem.MType:type_name -> pb.EItemType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMsg); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgStr); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPrizeItem); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebProto); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		EnumInfos:         file_base_proto_enumTypes,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}