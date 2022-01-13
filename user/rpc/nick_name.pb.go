// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/user/rpc/nick_name.proto

package rpc

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

// 设置用户头像
type SetNicknameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,string"` // 用户id
	// @gotags: json:"nickname"
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"` // 昵称
}

func (x *SetNicknameReq) Reset() {
	*x = SetNicknameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_rpc_nick_name_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNicknameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNicknameReq) ProtoMessage() {}

func (x *SetNicknameReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_rpc_nick_name_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNicknameReq.ProtoReflect.Descriptor instead.
func (*SetNicknameReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_rpc_nick_name_proto_rawDescGZIP(), []int{0}
}

func (x *SetNicknameReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetNicknameReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type SetNicknameRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetNicknameRsp) Reset() {
	*x = SetNicknameRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_rpc_nick_name_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNicknameRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNicknameRsp) ProtoMessage() {}

func (x *SetNicknameRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_rpc_nick_name_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNicknameRsp.ProtoReflect.Descriptor instead.
func (*SetNicknameRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_rpc_nick_name_proto_rawDescGZIP(), []int{1}
}

var File_relo_protocol_user_rpc_nick_name_proto protoreflect.FileDescriptor

var file_relo_protocol_user_rpc_nick_name_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x70, 0x63,
	0x22, 0x3c, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x10,
	0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x73, 0x70,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_user_rpc_nick_name_proto_rawDescOnce sync.Once
	file_relo_protocol_user_rpc_nick_name_proto_rawDescData = file_relo_protocol_user_rpc_nick_name_proto_rawDesc
)

func file_relo_protocol_user_rpc_nick_name_proto_rawDescGZIP() []byte {
	file_relo_protocol_user_rpc_nick_name_proto_rawDescOnce.Do(func() {
		file_relo_protocol_user_rpc_nick_name_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_user_rpc_nick_name_proto_rawDescData)
	})
	return file_relo_protocol_user_rpc_nick_name_proto_rawDescData
}

var file_relo_protocol_user_rpc_nick_name_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_user_rpc_nick_name_proto_goTypes = []interface{}{
	(*SetNicknameReq)(nil), // 0: relo.protocol.user.rpc.SetNicknameReq
	(*SetNicknameRsp)(nil), // 1: relo.protocol.user.rpc.SetNicknameRsp
}
var file_relo_protocol_user_rpc_nick_name_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_protocol_user_rpc_nick_name_proto_init() }
func file_relo_protocol_user_rpc_nick_name_proto_init() {
	if File_relo_protocol_user_rpc_nick_name_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_user_rpc_nick_name_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNicknameReq); i {
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
		file_relo_protocol_user_rpc_nick_name_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNicknameRsp); i {
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
			RawDescriptor: file_relo_protocol_user_rpc_nick_name_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_user_rpc_nick_name_proto_goTypes,
		DependencyIndexes: file_relo_protocol_user_rpc_nick_name_proto_depIdxs,
		MessageInfos:      file_relo_protocol_user_rpc_nick_name_proto_msgTypes,
	}.Build()
	File_relo_protocol_user_rpc_nick_name_proto = out.File
	file_relo_protocol_user_rpc_nick_name_proto_rawDesc = nil
	file_relo_protocol_user_rpc_nick_name_proto_goTypes = nil
	file_relo_protocol_user_rpc_nick_name_proto_depIdxs = nil
}
