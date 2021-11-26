// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/client_detail.proto

package vo

import (
	core "github.com/reloclass/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 客户端信息
type ClientDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 客户端编号
	// @gotags: json:"id,string"`
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,string"`
	// 版本信息
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// 客户端类型
	Type core.ClientType `protobuf:"varint,3,opt,name=type,proto3,enum=relo.core.ClientType" json:"type,omitempty"`
	// 发布状态
	Status core.ClientStatus `protobuf:"varint,4,opt,name=status,proto3,enum=relo.core.ClientStatus" json:"status,omitempty"`
	// 版本描述
	Desc string `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	// 文件编号
	// @gotags: json:"fileId"
	FileId string `protobuf:"bytes,6,opt,name=file_id,json=fileId,proto3" json:"fileId"`
	// 文件MD5
	Md5 string `protobuf:"bytes,7,opt,name=md5,proto3" json:"md5,omitempty"`
	// 更新类型
	// @gotags: json:"updateType"
	UpdateType core.UpdateType `protobuf:"varint,8,opt,name=update_type,json=updateType,proto3,enum=relo.core.UpdateType" json:"updateType"`
	// 更新文件类型
	// @gotags：json:"updateFileType"
	UpdateFileType core.UpdateFileType `protobuf:"varint,9,opt,name=update_file_type,json=updateFileType,proto3,enum=relo.core.UpdateFileType" json:"update_file_type,omitempty"`
	// 发布时间
	// @gotags: json:"publishedAt"
	PublishedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=published_at,json=publishedAt,proto3" json:"publishedAt"`
	// 创建时间
	Created *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created,proto3" json:"created,omitempty"`
	// 更新时间
	Updated *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *ClientDetail) Reset() {
	*x = ClientDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_client_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDetail) ProtoMessage() {}

func (x *ClientDetail) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_client_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDetail.ProtoReflect.Descriptor instead.
func (*ClientDetail) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_client_detail_proto_rawDescGZIP(), []int{0}
}

func (x *ClientDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClientDetail) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ClientDetail) GetType() core.ClientType {
	if x != nil {
		return x.Type
	}
	return core.ClientType_CLIENT_TYPE_UNSPECIFIED
}

func (x *ClientDetail) GetStatus() core.ClientStatus {
	if x != nil {
		return x.Status
	}
	return core.ClientStatus_CLIENT_STATUS_UNSPECIFIED
}

func (x *ClientDetail) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ClientDetail) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ClientDetail) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *ClientDetail) GetUpdateType() core.UpdateType {
	if x != nil {
		return x.UpdateType
	}
	return core.UpdateType_UPDATE_TYPE_UNSPECIFIED
}

func (x *ClientDetail) GetUpdateFileType() core.UpdateFileType {
	if x != nil {
		return x.UpdateFileType
	}
	return core.UpdateFileType_UPDATE_FILE_TYPE_UNSPECIFIED
}

func (x *ClientDetail) GetPublishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

func (x *ClientDetail) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ClientDetail) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

var File_relo_protocol_vo_client_detail_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_client_detail_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x03, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x6c,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12,
	0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x36, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_relo_protocol_vo_client_detail_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_client_detail_proto_rawDescData = file_relo_protocol_vo_client_detail_proto_rawDesc
)

func file_relo_protocol_vo_client_detail_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_client_detail_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_client_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_client_detail_proto_rawDescData)
	})
	return file_relo_protocol_vo_client_detail_proto_rawDescData
}

var file_relo_protocol_vo_client_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_client_detail_proto_goTypes = []interface{}{
	(*ClientDetail)(nil),          // 0: relo.protocol.vo.ClientDetail
	(core.ClientType)(0),          // 1: relo.core.ClientType
	(core.ClientStatus)(0),        // 2: relo.core.ClientStatus
	(core.UpdateType)(0),          // 3: relo.core.UpdateType
	(core.UpdateFileType)(0),      // 4: relo.core.UpdateFileType
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_relo_protocol_vo_client_detail_proto_depIdxs = []int32{
	1, // 0: relo.protocol.vo.ClientDetail.type:type_name -> relo.core.ClientType
	2, // 1: relo.protocol.vo.ClientDetail.status:type_name -> relo.core.ClientStatus
	3, // 2: relo.protocol.vo.ClientDetail.update_type:type_name -> relo.core.UpdateType
	4, // 3: relo.protocol.vo.ClientDetail.update_file_type:type_name -> relo.core.UpdateFileType
	5, // 4: relo.protocol.vo.ClientDetail.published_at:type_name -> google.protobuf.Timestamp
	5, // 5: relo.protocol.vo.ClientDetail.created:type_name -> google.protobuf.Timestamp
	5, // 6: relo.protocol.vo.ClientDetail.updated:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_client_detail_proto_init() }
func file_relo_protocol_vo_client_detail_proto_init() {
	if File_relo_protocol_vo_client_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_client_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientDetail); i {
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
			RawDescriptor: file_relo_protocol_vo_client_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_client_detail_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_client_detail_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_client_detail_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_client_detail_proto = out.File
	file_relo_protocol_vo_client_detail_proto_rawDesc = nil
	file_relo_protocol_vo_client_detail_proto_goTypes = nil
	file_relo_protocol_vo_client_detail_proto_depIdxs = nil
}
