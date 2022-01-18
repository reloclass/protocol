// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/lecture/update.proto

package lecture

import (
	lecture "github.com/reloclass/core/lecture"
	vo "github.com/reloclass/protocol/vo"
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

// 更新讲次请求
type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 讲次编号
	// @gotags: param:"id"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty" param:"id"`
	// 课程编号
	// @gotags: json:"courseId,string" validate:"required"
	CourseId int64 `protobuf:"varint,3,opt,name=course_id,json=courseId,proto3" json:"courseId,string" validate:"required"`
	// 讲次类型
	// @gotags: default:"0" json:"type" validate:"omitempty,oneof=0 1"
	Type lecture.Type `protobuf:"varint,4,opt,name=type,proto3,enum=relo.core.lecture.Type" json:"type" default:"0" validate:"omitempty,oneof=0 1"`
	// 章节编号
	// @gotags: json:"parentId,string"
	ParentId int64 `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3" json:"parentId,string"`
	// 顺序
	// @gotags: default:"0" json:"sequence"
	Sequence int32 `protobuf:"varint,6,opt,name=sequence,proto3" json:"sequence" default:"0"`
	// 名称
	// @gotags: json:"name" validate:"required"
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name" validate:"required"`
	// 简介
	// @gotags: json:"profile" validate:"omitempty,max=255"
	Profile string `protobuf:"bytes,8,opt,name=profile,proto3" json:"profile" validate:"omitempty,max=255"`
	// 讲次内容
	// @gotags: json:"lectureContents" validate:"omitempty,dive"
	Contents []*vo.LectureSectionContent `protobuf:"bytes,9,rep,name=contents,proto3" json:"lectureContents" validate:"omitempty,dive"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_lecture_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_lecture_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_lecture_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReq) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *UpdateReq) GetType() lecture.Type {
	if x != nil {
		return x.Type
	}
	return lecture.Type_CHAPTER
}

func (x *UpdateReq) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *UpdateReq) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *UpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateReq) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *UpdateReq) GetContents() []*vo.LectureSectionContent {
	if x != nil {
		return x.Contents
	}
	return nil
}

// 交换讲次请求
type SwitchSequenceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"switchItems"
	SwitchItems []*vo.Lecture `protobuf:"bytes,2,rep,name=switch_items,json=switchItems,proto3" json:"switchItems"`
}

func (x *SwitchSequenceReq) Reset() {
	*x = SwitchSequenceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_lecture_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchSequenceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchSequenceReq) ProtoMessage() {}

func (x *SwitchSequenceReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_lecture_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchSequenceReq.ProtoReflect.Descriptor instead.
func (*SwitchSequenceReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_lecture_update_proto_rawDescGZIP(), []int{1}
}

func (x *SwitchSequenceReq) GetSwitchItems() []*vo.Lecture {
	if x != nil {
		return x.SwitchItems
	}
	return nil
}

var File_relo_protocol_lecture_update_proto protoreflect.FileDescriptor

var file_relo_protocol_lecture_update_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x1c, 0x72, 0x65, 0x6c,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x72, 0x65, 0x6c, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x65, 0x6c,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x4c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x51, 0x0a,
	0x11, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x4c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x0b, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x3b, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_lecture_update_proto_rawDescOnce sync.Once
	file_relo_protocol_lecture_update_proto_rawDescData = file_relo_protocol_lecture_update_proto_rawDesc
)

func file_relo_protocol_lecture_update_proto_rawDescGZIP() []byte {
	file_relo_protocol_lecture_update_proto_rawDescOnce.Do(func() {
		file_relo_protocol_lecture_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_lecture_update_proto_rawDescData)
	})
	return file_relo_protocol_lecture_update_proto_rawDescData
}

var file_relo_protocol_lecture_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_lecture_update_proto_goTypes = []interface{}{
	(*UpdateReq)(nil),                // 0: relo.protocol.lecture.UpdateReq
	(*SwitchSequenceReq)(nil),        // 1: relo.protocol.lecture.SwitchSequenceReq
	(lecture.Type)(0),                // 2: relo.core.lecture.Type
	(*vo.LectureSectionContent)(nil), // 3: relo.protocol.vo.LectureSectionContent
	(*vo.Lecture)(nil),               // 4: relo.protocol.vo.Lecture
}
var file_relo_protocol_lecture_update_proto_depIdxs = []int32{
	2, // 0: relo.protocol.lecture.UpdateReq.type:type_name -> relo.core.lecture.Type
	3, // 1: relo.protocol.lecture.UpdateReq.contents:type_name -> relo.protocol.vo.LectureSectionContent
	4, // 2: relo.protocol.lecture.SwitchSequenceReq.switch_items:type_name -> relo.protocol.vo.Lecture
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relo_protocol_lecture_update_proto_init() }
func file_relo_protocol_lecture_update_proto_init() {
	if File_relo_protocol_lecture_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_lecture_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_relo_protocol_lecture_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchSequenceReq); i {
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
			RawDescriptor: file_relo_protocol_lecture_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_lecture_update_proto_goTypes,
		DependencyIndexes: file_relo_protocol_lecture_update_proto_depIdxs,
		MessageInfos:      file_relo_protocol_lecture_update_proto_msgTypes,
	}.Build()
	File_relo_protocol_lecture_update_proto = out.File
	file_relo_protocol_lecture_update_proto_rawDesc = nil
	file_relo_protocol_lecture_update_proto_goTypes = nil
	file_relo_protocol_lecture_update_proto_depIdxs = nil
}
