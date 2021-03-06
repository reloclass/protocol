// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/teacher/check.proto

package teacher

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

// 检查是否已经是老师请求
type CheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	// @gotags: json:"phone" validate:"required,printascii,mobile"
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone" validate:"required,printascii,mobile"`
}

func (x *CheckReq) Reset() {
	*x = CheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_teacher_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckReq) ProtoMessage() {}

func (x *CheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_teacher_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckReq.ProtoReflect.Descriptor instead.
func (*CheckReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_teacher_check_proto_rawDescGZIP(), []int{0}
}

func (x *CheckReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

// 检查是否已经是老师响应
type CheckRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	// 是否为老师
	Exist bool `protobuf:"varint,3,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *CheckRsp) Reset() {
	*x = CheckRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_teacher_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRsp) ProtoMessage() {}

func (x *CheckRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_teacher_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRsp.ProtoReflect.Descriptor instead.
func (*CheckRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_teacher_check_proto_rawDescGZIP(), []int{1}
}

func (x *CheckRsp) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CheckRsp) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

var File_relo_protocol_teacher_check_proto protoreflect.FileDescriptor

var file_relo_protocol_teacher_check_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x08, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x36, 0x0a, 0x08,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x3b, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_teacher_check_proto_rawDescOnce sync.Once
	file_relo_protocol_teacher_check_proto_rawDescData = file_relo_protocol_teacher_check_proto_rawDesc
)

func file_relo_protocol_teacher_check_proto_rawDescGZIP() []byte {
	file_relo_protocol_teacher_check_proto_rawDescOnce.Do(func() {
		file_relo_protocol_teacher_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_teacher_check_proto_rawDescData)
	})
	return file_relo_protocol_teacher_check_proto_rawDescData
}

var file_relo_protocol_teacher_check_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_teacher_check_proto_goTypes = []interface{}{
	(*CheckReq)(nil), // 0: relo.protocol.teacher.CheckReq
	(*CheckRsp)(nil), // 1: relo.protocol.teacher.CheckRsp
}
var file_relo_protocol_teacher_check_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_protocol_teacher_check_proto_init() }
func file_relo_protocol_teacher_check_proto_init() {
	if File_relo_protocol_teacher_check_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_teacher_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckReq); i {
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
		file_relo_protocol_teacher_check_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRsp); i {
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
			RawDescriptor: file_relo_protocol_teacher_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_teacher_check_proto_goTypes,
		DependencyIndexes: file_relo_protocol_teacher_check_proto_depIdxs,
		MessageInfos:      file_relo_protocol_teacher_check_proto_msgTypes,
	}.Build()
	File_relo_protocol_teacher_check_proto = out.File
	file_relo_protocol_teacher_check_proto_rawDesc = nil
	file_relo_protocol_teacher_check_proto_goTypes = nil
	file_relo_protocol_teacher_check_proto_depIdxs = nil
}
