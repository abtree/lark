// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v4.22.1
// source: api_msg.proto

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

type APIMsgId int32

const (
	APIMsgId_None      APIMsgId = 0
	APIMsgId_BackToWeb APIMsgId = 1
	APIMsgId_GetAllCfg APIMsgId = 2
	APIMsgId_EGiftPack APIMsgId = 3
	APIMsgId_EGetid    APIMsgId = 4
	APIMsgId_EAuth     APIMsgId = 5
)

// Enum value maps for APIMsgId.
var (
	APIMsgId_name = map[int32]string{
		0: "None",
		1: "BackToWeb",
		2: "GetAllCfg",
		3: "EGiftPack",
		4: "EGetid",
		5: "EAuth",
	}
	APIMsgId_value = map[string]int32{
		"None":      0,
		"BackToWeb": 1,
		"GetAllCfg": 2,
		"EGiftPack": 3,
		"EGetid":    4,
		"EAuth":     5,
	}
)

func (x APIMsgId) Enum() *APIMsgId {
	p := new(APIMsgId)
	*p = x
	return p
}

func (x APIMsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (APIMsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_api_msg_proto_enumTypes[0].Descriptor()
}

func (APIMsgId) Type() protoreflect.EnumType {
	return &file_api_msg_proto_enumTypes[0]
}

func (x APIMsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use APIMsgId.Descriptor instead.
func (APIMsgId) EnumDescriptor() ([]byte, []int) {
	return file_api_msg_proto_rawDescGZIP(), []int{0}
}

var File_api_msg_proto protoreflect.FileDescriptor

var file_api_msg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x2a, 0x58, 0x0a, 0x08, 0x41, 0x50, 0x49, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x63,
	0x6b, 0x54, 0x6f, 0x57, 0x65, 0x62, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x66, 0x67, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x47, 0x69, 0x66, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x47, 0x65, 0x74, 0x69, 0x64,
	0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x41, 0x75, 0x74, 0x68, 0x10, 0x05, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_msg_proto_rawDescOnce sync.Once
	file_api_msg_proto_rawDescData = file_api_msg_proto_rawDesc
)

func file_api_msg_proto_rawDescGZIP() []byte {
	file_api_msg_proto_rawDescOnce.Do(func() {
		file_api_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_proto_rawDescData)
	})
	return file_api_msg_proto_rawDescData
}

var file_api_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_msg_proto_goTypes = []interface{}{
	(APIMsgId)(0), // 0: pb.APIMsgId
}
var file_api_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_msg_proto_init() }
func file_api_msg_proto_init() {
	if File_api_msg_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_msg_proto_goTypes,
		DependencyIndexes: file_api_msg_proto_depIdxs,
		EnumInfos:         file_api_msg_proto_enumTypes,
	}.Build()
	File_api_msg_proto = out.File
	file_api_msg_proto_rawDesc = nil
	file_api_msg_proto_goTypes = nil
	file_api_msg_proto_depIdxs = nil
}