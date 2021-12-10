// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/theme/download.proto

package theme

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

// 批量下载主题信息请求
type BatchDownloadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主题文件信息列表
	Files []*File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *BatchDownloadReq) Reset() {
	*x = BatchDownloadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_theme_download_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDownloadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDownloadReq) ProtoMessage() {}

func (x *BatchDownloadReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_theme_download_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDownloadReq.ProtoReflect.Descriptor instead.
func (*BatchDownloadReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_theme_download_proto_rawDescGZIP(), []int{0}
}

func (x *BatchDownloadReq) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

// 主题文件信息
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件Id
	// @gotags: json:"fileId" validate:"required,len=20"
	FileId string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"fileId" validate:"required,len=20"`
	// 文件另存名字
	// @gotags: json:"name" validate:"omitempty,min=1,without_special_symbol,filename"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" validate:"omitempty,min=1,without_special_symbol,filename"`
	// 操作类型
	// @gotags: json:"type" default:"1" validate:"required,oneof=1 2 3"
	Type cloudfile.OperationType `protobuf:"varint,4,opt,name=type,proto3,enum=relo.core.cloudfile.OperationType" json:"type" default:"1" validate:"required,oneof=1 2 3"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_theme_download_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_theme_download_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_relo_protocol_theme_download_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetType() cloudfile.OperationType {
	if x != nil {
		return x.Type
	}
	return cloudfile.OperationType_RESERVE
}

var File_relo_protocol_theme_download_proto protoreflect.FileDescriptor

var file_relo_protocol_theme_download_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x74, 0x68, 0x65, 0x6d, 0x65, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x1a, 0x28, 0x72, 0x65, 0x6c, 0x6f, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x6b, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x72, 0x65,
	0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x3b, 0x74, 0x68, 0x65,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_theme_download_proto_rawDescOnce sync.Once
	file_relo_protocol_theme_download_proto_rawDescData = file_relo_protocol_theme_download_proto_rawDesc
)

func file_relo_protocol_theme_download_proto_rawDescGZIP() []byte {
	file_relo_protocol_theme_download_proto_rawDescOnce.Do(func() {
		file_relo_protocol_theme_download_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_theme_download_proto_rawDescData)
	})
	return file_relo_protocol_theme_download_proto_rawDescData
}

var file_relo_protocol_theme_download_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_theme_download_proto_goTypes = []interface{}{
	(*BatchDownloadReq)(nil),     // 0: relo.protocol.theme.BatchDownloadReq
	(*File)(nil),                 // 1: relo.protocol.theme.File
	(cloudfile.OperationType)(0), // 2: relo.core.cloudfile.OperationType
}
var file_relo_protocol_theme_download_proto_depIdxs = []int32{
	1, // 0: relo.protocol.theme.BatchDownloadReq.files:type_name -> relo.protocol.theme.File
	2, // 1: relo.protocol.theme.File.type:type_name -> relo.core.cloudfile.OperationType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_relo_protocol_theme_download_proto_init() }
func file_relo_protocol_theme_download_proto_init() {
	if File_relo_protocol_theme_download_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_theme_download_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDownloadReq); i {
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
		file_relo_protocol_theme_download_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_relo_protocol_theme_download_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_theme_download_proto_goTypes,
		DependencyIndexes: file_relo_protocol_theme_download_proto_depIdxs,
		MessageInfos:      file_relo_protocol_theme_download_proto_msgTypes,
	}.Build()
	File_relo_protocol_theme_download_proto = out.File
	file_relo_protocol_theme_download_proto_rawDesc = nil
	file_relo_protocol_theme_download_proto_goTypes = nil
	file_relo_protocol_theme_download_proto_depIdxs = nil
}
