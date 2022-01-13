// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/user/register.proto

package user

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

// 检查用户是否注册请求
type CheckRegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	// @gotags: json:"phone" validate:"required,printascii,mobile"
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone" validate:"required,printascii,mobile"`
}

func (x *CheckRegisterReq) Reset() {
	*x = CheckRegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRegisterReq) ProtoMessage() {}

func (x *CheckRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRegisterReq.ProtoReflect.Descriptor instead.
func (*CheckRegisterReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_register_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRegisterReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

// 检查用户是否注册响应
type CheckRegisterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	// 是否注册
	Register bool `protobuf:"varint,3,opt,name=register,proto3" json:"register,omitempty"`
}

func (x *CheckRegisterRsp) Reset() {
	*x = CheckRegisterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRegisterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRegisterRsp) ProtoMessage() {}

func (x *CheckRegisterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRegisterRsp.ProtoReflect.Descriptor instead.
func (*CheckRegisterRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_register_proto_rawDescGZIP(), []int{1}
}

func (x *CheckRegisterRsp) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CheckRegisterRsp) GetRegister() bool {
	if x != nil {
		return x.Register
	}
	return false
}

var File_relo_protocol_user_register_proto protoreflect.FileDescriptor

var file_relo_protocol_user_register_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x22, 0x44, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_user_register_proto_rawDescOnce sync.Once
	file_relo_protocol_user_register_proto_rawDescData = file_relo_protocol_user_register_proto_rawDesc
)

func file_relo_protocol_user_register_proto_rawDescGZIP() []byte {
	file_relo_protocol_user_register_proto_rawDescOnce.Do(func() {
		file_relo_protocol_user_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_user_register_proto_rawDescData)
	})
	return file_relo_protocol_user_register_proto_rawDescData
}

var file_relo_protocol_user_register_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_user_register_proto_goTypes = []interface{}{
	(*CheckRegisterReq)(nil), // 0: relo.protocol.user.CheckRegisterReq
	(*CheckRegisterRsp)(nil), // 1: relo.protocol.user.CheckRegisterRsp
}
var file_relo_protocol_user_register_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_protocol_user_register_proto_init() }
func file_relo_protocol_user_register_proto_init() {
	if File_relo_protocol_user_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_user_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRegisterReq); i {
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
		file_relo_protocol_user_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRegisterRsp); i {
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
			RawDescriptor: file_relo_protocol_user_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_user_register_proto_goTypes,
		DependencyIndexes: file_relo_protocol_user_register_proto_depIdxs,
		MessageInfos:      file_relo_protocol_user_register_proto_msgTypes,
	}.Build()
	File_relo_protocol_user_register_proto = out.File
	file_relo_protocol_user_register_proto_rawDesc = nil
	file_relo_protocol_user_register_proto_goTypes = nil
	file_relo_protocol_user_register_proto_depIdxs = nil
}
