// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/filechunk/complete.proto

package filechunk

import (
	cloudfile "github.com/reloclass/core/cloudfile"
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

// 完成分块上传请求
type CompleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件编号
	// @gotags: json:"fileId" validate:"required,len=20"
	FileId string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"fileId" validate:"required,len=20"`
	// 腾讯Cos对象存储分块上传时唯一标识Id
	// @gotags: json:"uploadId" validate:"required"
	UploadId string `protobuf:"bytes,3,opt,name=upload_id,json=uploadId,proto3" json:"uploadId" validate:"required"`
	// 每一部分上传返回信息
	// @gotags: json:"parts" validate:"required,dive"
	Parts []*cloudfile.ChunkInfo `protobuf:"bytes,4,rep,name=parts,proto3" json:"parts" validate:"required,dive"`
}

func (x *CompleteReq) Reset() {
	*x = CompleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_filechunk_complete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteReq) ProtoMessage() {}

func (x *CompleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_filechunk_complete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteReq.ProtoReflect.Descriptor instead.
func (*CompleteReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_filechunk_complete_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteReq) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *CompleteReq) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *CompleteReq) GetParts() []*cloudfile.ChunkInfo {
	if x != nil {
		return x.Parts
	}
	return nil
}

var File_relo_protocol_filechunk_complete_proto protoreflect.FileDescriptor

var file_relo_protocol_filechunk_complete_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x69, 0x6c,
	0x65, 0x1a, 0x24, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65,
	0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x70, 0x61, 0x72,
	0x74, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x3b, 0x66, 0x69,
	0x6c, 0x65, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_filechunk_complete_proto_rawDescOnce sync.Once
	file_relo_protocol_filechunk_complete_proto_rawDescData = file_relo_protocol_filechunk_complete_proto_rawDesc
)

func file_relo_protocol_filechunk_complete_proto_rawDescGZIP() []byte {
	file_relo_protocol_filechunk_complete_proto_rawDescOnce.Do(func() {
		file_relo_protocol_filechunk_complete_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_filechunk_complete_proto_rawDescData)
	})
	return file_relo_protocol_filechunk_complete_proto_rawDescData
}

var file_relo_protocol_filechunk_complete_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_filechunk_complete_proto_goTypes = []interface{}{
	(*CompleteReq)(nil),         // 0: relo.protocol.cloudfile.CompleteReq
	(*cloudfile.ChunkInfo)(nil), // 1: relo.core.cloudfile.ChunkInfo
}
var file_relo_protocol_filechunk_complete_proto_depIdxs = []int32{
	1, // 0: relo.protocol.cloudfile.CompleteReq.parts:type_name -> relo.core.cloudfile.ChunkInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relo_protocol_filechunk_complete_proto_init() }
func file_relo_protocol_filechunk_complete_proto_init() {
	if File_relo_protocol_filechunk_complete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_filechunk_complete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteReq); i {
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
			RawDescriptor: file_relo_protocol_filechunk_complete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_filechunk_complete_proto_goTypes,
		DependencyIndexes: file_relo_protocol_filechunk_complete_proto_depIdxs,
		MessageInfos:      file_relo_protocol_filechunk_complete_proto_msgTypes,
	}.Build()
	File_relo_protocol_filechunk_complete_proto = out.File
	file_relo_protocol_filechunk_complete_proto_rawDesc = nil
	file_relo_protocol_filechunk_complete_proto_goTypes = nil
	file_relo_protocol_filechunk_complete_proto_depIdxs = nil
}