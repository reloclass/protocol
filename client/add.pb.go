// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/client/add.proto

package client

import (
	core "github.com/reloclass/core"
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

// 添加客户端请求
type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件编号
	// @gotags: json:"appInfo" validate:"required"
	AppInfo *AppInfo `protobuf:"bytes,1,opt,name=app_info,json=appInfo,proto3" json:"appInfo" validate:"required"`
	// 发布类型
	// @gotags: json:"type" validate:"required_with=0 1"
	Type core.PublishType `protobuf:"varint,2,opt,name=type,proto3,enum=relo.core.PublishType" json:"type" validate:"required_with=0 1"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_client_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_client_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_client_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetAppInfo() *AppInfo {
	if x != nil {
		return x.AppInfo
	}
	return nil
}

func (x *AddReq) GetType() core.PublishType {
	if x != nil {
		return x.Type
	}
	return core.PublishType_PUBLISH_RIGHT_NOW
}

type AppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Windows 文件编号
	// @gotags: json:"fileId" validate:"required,len=20"
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"fileId" validate:"required,len=20"`
	// 文件md5
	// @gotags: json:"md5" validate:"omitempty"
	Md5 string `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5" validate:"omitempty"`
	// 版本号
	// @gotags: json:"version" validate:"required"
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version" validate:"required"`
	// 版本说明
	// @gotags: json:"versionDesc" validate:"required,min=2,max=4095"
	VersionDesc string `protobuf:"bytes,4,opt,name=version_desc,json=versionDesc,proto3" json:"versionDesc" validate:"required,min=2,max=4095"`
	// 更新类型
	// @gotags: json:"updateType" validate:"required_with=1 2 3"
	UpdateType core.UpdateType `protobuf:"varint,5,opt,name=update_type,json=updateType,proto3,enum=relo.core.UpdateType" json:"updateType" validate:"required_with=1 2 3"`
	// 客户端类型
	// @gotags: json:"clientType" validate:"required_with=0 1"
	ClientType core.ClientType `protobuf:"varint,6,opt,name=client_type,json=clientType,proto3,enum=relo.core.ClientType" json:"clientType" validate:"required_with=0 1"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_client_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_client_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfo.ProtoReflect.Descriptor instead.
func (*AppInfo) Descriptor() ([]byte, []int) {
	return file_relo_protocol_client_add_proto_rawDescGZIP(), []int{1}
}

func (x *AppInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *AppInfo) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *AppInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AppInfo) GetVersionDesc() string {
	if x != nil {
		return x.VersionDesc
	}
	return ""
}

func (x *AppInfo) GetUpdateType() core.UpdateType {
	if x != nil {
		return x.UpdateType
	}
	return core.UpdateType_UPDATE_TYPE_UNSPECIFIED
}

func (x *AppInfo) GetClientType() core.ClientType {
	if x != nil {
		return x.ClientType
	}
	return core.ClientType_CLIENT_TYPE_UNSPECIFIED
}

// 添加响应
type AddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 客户端编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,string"`
}

func (x *AddRsp) Reset() {
	*x = AddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_client_add_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRsp) ProtoMessage() {}

func (x *AddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_client_add_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRsp.ProtoReflect.Descriptor instead.
func (*AddRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_client_add_proto_rawDescGZIP(), []int{2}
}

func (x *AddRsp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_relo_protocol_client_add_proto protoreflect.FileDescriptor

var file_relo_protocol_client_add_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e,
	0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x6c,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x70, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe1,
	0x01, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x36, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_client_add_proto_rawDescOnce sync.Once
	file_relo_protocol_client_add_proto_rawDescData = file_relo_protocol_client_add_proto_rawDesc
)

func file_relo_protocol_client_add_proto_rawDescGZIP() []byte {
	file_relo_protocol_client_add_proto_rawDescOnce.Do(func() {
		file_relo_protocol_client_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_client_add_proto_rawDescData)
	})
	return file_relo_protocol_client_add_proto_rawDescData
}

var file_relo_protocol_client_add_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_relo_protocol_client_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),        // 0: relo.protocol.client.AddReq
	(*AppInfo)(nil),       // 1: relo.protocol.client.AppInfo
	(*AddRsp)(nil),        // 2: relo.protocol.client.AddRsp
	(core.PublishType)(0), // 3: relo.core.PublishType
	(core.UpdateType)(0),  // 4: relo.core.UpdateType
	(core.ClientType)(0),  // 5: relo.core.ClientType
}
var file_relo_protocol_client_add_proto_depIdxs = []int32{
	1, // 0: relo.protocol.client.AddReq.app_info:type_name -> relo.protocol.client.AppInfo
	3, // 1: relo.protocol.client.AddReq.type:type_name -> relo.core.PublishType
	4, // 2: relo.protocol.client.AppInfo.update_type:type_name -> relo.core.UpdateType
	5, // 3: relo.protocol.client.AppInfo.client_type:type_name -> relo.core.ClientType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_relo_protocol_client_add_proto_init() }
func file_relo_protocol_client_add_proto_init() {
	if File_relo_protocol_client_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_client_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_relo_protocol_client_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfo); i {
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
		file_relo_protocol_client_add_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRsp); i {
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
			RawDescriptor: file_relo_protocol_client_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_client_add_proto_goTypes,
		DependencyIndexes: file_relo_protocol_client_add_proto_depIdxs,
		MessageInfos:      file_relo_protocol_client_add_proto_msgTypes,
	}.Build()
	File_relo_protocol_client_add_proto = out.File
	file_relo_protocol_client_add_proto_rawDesc = nil
	file_relo_protocol_client_add_proto_goTypes = nil
	file_relo_protocol_client_add_proto_depIdxs = nil
}
