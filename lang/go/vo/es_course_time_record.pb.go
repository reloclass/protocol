// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/es_course_time_record.proto

package vo

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

// 课节录课资源
type EsCourseTimeRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件编号
	// @gotags: json:"fileId"
	FileId string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"fileId"`
	// 视频名字
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 创建时间
	// @gotags: json:"createTime"
	CreateTime int64 `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"createTime"`
	// 是否删除
	// @gotags: json:"isDeleted"
	IsDeleted bool `protobuf:"varint,5,opt,name=is_deleted,json=isDeleted,proto3" json:"isDeleted"`
}

func (x *EsCourseTimeRecord) Reset() {
	*x = EsCourseTimeRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_es_course_time_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsCourseTimeRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsCourseTimeRecord) ProtoMessage() {}

func (x *EsCourseTimeRecord) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_es_course_time_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsCourseTimeRecord.ProtoReflect.Descriptor instead.
func (*EsCourseTimeRecord) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_es_course_time_record_proto_rawDescGZIP(), []int{0}
}

func (x *EsCourseTimeRecord) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *EsCourseTimeRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EsCourseTimeRecord) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *EsCourseTimeRecord) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_relo_protocol_vo_es_course_time_record_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_es_course_time_record_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f,
	0x22, 0x81, 0x01, 0x0a, 0x12, 0x45, 0x73, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_es_course_time_record_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_es_course_time_record_proto_rawDescData = file_relo_protocol_vo_es_course_time_record_proto_rawDesc
)

func file_relo_protocol_vo_es_course_time_record_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_es_course_time_record_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_es_course_time_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_es_course_time_record_proto_rawDescData)
	})
	return file_relo_protocol_vo_es_course_time_record_proto_rawDescData
}

var file_relo_protocol_vo_es_course_time_record_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_es_course_time_record_proto_goTypes = []interface{}{
	(*EsCourseTimeRecord)(nil), // 0: relo.protocol.vo.EsCourseTimeRecord
}
var file_relo_protocol_vo_es_course_time_record_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_es_course_time_record_proto_init() }
func file_relo_protocol_vo_es_course_time_record_proto_init() {
	if File_relo_protocol_vo_es_course_time_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_es_course_time_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsCourseTimeRecord); i {
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
			RawDescriptor: file_relo_protocol_vo_es_course_time_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_es_course_time_record_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_es_course_time_record_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_es_course_time_record_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_es_course_time_record_proto = out.File
	file_relo_protocol_vo_es_course_time_record_proto_rawDesc = nil
	file_relo_protocol_vo_es_course_time_record_proto_goTypes = nil
	file_relo_protocol_vo_es_course_time_record_proto_depIdxs = nil
}