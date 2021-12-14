// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/client/publish.proto

package client

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

// 发布请求
type PublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 客户端信息编号
	// @gotags: param:"id" validate:"required"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" param:"id" validate:"required"`
}

func (x *PublishReq) Reset() {
	*x = PublishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_client_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReq) ProtoMessage() {}

func (x *PublishReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_client_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReq.ProtoReflect.Descriptor instead.
func (*PublishReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_client_publish_proto_rawDescGZIP(), []int{0}
}

func (x *PublishReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 发布响应
type PublishRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 客户端信息编号
	// @gotags: json:"affected,string"
	Affected int64 `protobuf:"varint,1,opt,name=affected,proto3" json:"affected,string"`
}

func (x *PublishRsp) Reset() {
	*x = PublishRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_client_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRsp) ProtoMessage() {}

func (x *PublishRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_client_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRsp.ProtoReflect.Descriptor instead.
func (*PublishRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_client_publish_proto_rawDescGZIP(), []int{1}
}

func (x *PublishRsp) GetAffected() int64 {
	if x != nil {
		return x.Affected
	}
	return 0
}

var File_relo_protocol_client_publish_proto protoreflect.FileDescriptor

var file_relo_protocol_client_publish_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x1c, 0x0a, 0x0a, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_client_publish_proto_rawDescOnce sync.Once
	file_relo_protocol_client_publish_proto_rawDescData = file_relo_protocol_client_publish_proto_rawDesc
)

func file_relo_protocol_client_publish_proto_rawDescGZIP() []byte {
	file_relo_protocol_client_publish_proto_rawDescOnce.Do(func() {
		file_relo_protocol_client_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_client_publish_proto_rawDescData)
	})
	return file_relo_protocol_client_publish_proto_rawDescData
}

var file_relo_protocol_client_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_client_publish_proto_goTypes = []interface{}{
	(*PublishReq)(nil), // 0: relo.protocol.client.PublishReq
	(*PublishRsp)(nil), // 1: relo.protocol.client.PublishRsp
}
var file_relo_protocol_client_publish_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_protocol_client_publish_proto_init() }
func file_relo_protocol_client_publish_proto_init() {
	if File_relo_protocol_client_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_client_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReq); i {
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
		file_relo_protocol_client_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRsp); i {
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
			RawDescriptor: file_relo_protocol_client_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_client_publish_proto_goTypes,
		DependencyIndexes: file_relo_protocol_client_publish_proto_depIdxs,
		MessageInfos:      file_relo_protocol_client_publish_proto_msgTypes,
	}.Build()
	File_relo_protocol_client_publish_proto = out.File
	file_relo_protocol_client_publish_proto_rawDesc = nil
	file_relo_protocol_client_publish_proto_goTypes = nil
	file_relo_protocol_client_publish_proto_depIdxs = nil
}
