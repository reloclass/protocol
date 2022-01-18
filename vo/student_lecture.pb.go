// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/student_lecture.proto

package vo

import (
	course "github.com/reloclass/core/course"
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

// 学生参与的录播课信息
type StudentLecture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string"`
	// 创建时间
	// @gotags: json:"createdAt"
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"createdAt"`
	// 修改时间
	// @gotags: json:"updatedAt"
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updatedAt"`
	// 学校编号
	// @gotags: json:"schoolId,string"
	SchoolId int64 `protobuf:"varint,5,opt,name=school_id,json=schoolId,proto3" json:"schoolId,string"`
	// 课程名称
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// 课程类型
	Type course.Type `protobuf:"varint,7,opt,name=type,proto3,enum=relo.core.course.Type" json:"type,omitempty"`
	// 封面
	Cover string `protobuf:"bytes,8,opt,name=cover,proto3" json:"cover,omitempty"`
	// 创建人
	// @gotags: json:"creatorId,string"
	CreatorId int64 `protobuf:"varint,9,opt,name=creator_id,json=creatorId,proto3" json:"creatorId,string"`
	// 介绍
	Info string `protobuf:"bytes,10,opt,name=info,proto3" json:"info,omitempty"`
	// 上课模式
	// @gotags: json:"classMode"
	ClassMode course.ClassMode `protobuf:"varint,11,opt,name=class_mode,json=classMode,proto3,enum=relo.core.course.ClassMode" json:"classMode"`
	// 封面
	// @gotags: json:"coverUrl"
	CoverUrl string `protobuf:"bytes,12,opt,name=cover_url,json=coverUrl,proto3" json:"coverUrl"`
	// 讲次数
	// @gotags: json:"lectureCount"
	SectionCount int64 `protobuf:"varint,13,opt,name=section_count,json=sectionCount,proto3" json:"lectureCount"`
	// 作业数
	// @gotags: json:"homeworkCount"
	HomeworkCount int64 `protobuf:"varint,14,opt,name=homework_count,json=homeworkCount,proto3" json:"homeworkCount"`
	// 测验数
	// @gotags: json:"testCount"
	TestCount int64 `protobuf:"varint,15,opt,name=test_count,json=testCount,proto3" json:"testCount"`
	// 完成讲次数
	// @gotags: json:"finishLectureCount"
	FinishedSectionCount int64 `protobuf:"varint,16,opt,name=finished_section_count,json=finishedSectionCount,proto3" json:"finishLectureCount"`
	// 完成作业数
	// @gotags: json:"finishHomeworkCount"
	FinishedHomeworkCount int64 `protobuf:"varint,17,opt,name=finished_homework_count,json=finishedHomeworkCount,proto3" json:"finishHomeworkCount"`
	// 完成测验数
	// @gotags: json:"finishTestCount"
	FinishedTestCount int64 `protobuf:"varint,18,opt,name=finished_test_count,json=finishedTestCount,proto3" json:"finishTestCount"`
}

func (x *StudentLecture) Reset() {
	*x = StudentLecture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_student_lecture_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentLecture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentLecture) ProtoMessage() {}

func (x *StudentLecture) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_student_lecture_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentLecture.ProtoReflect.Descriptor instead.
func (*StudentLecture) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_student_lecture_proto_rawDescGZIP(), []int{0}
}

func (x *StudentLecture) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StudentLecture) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *StudentLecture) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *StudentLecture) GetSchoolId() int64 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *StudentLecture) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudentLecture) GetType() course.Type {
	if x != nil {
		return x.Type
	}
	return course.Type_COURSE_TYPE_UNSPECIFIED
}

func (x *StudentLecture) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *StudentLecture) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *StudentLecture) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *StudentLecture) GetClassMode() course.ClassMode {
	if x != nil {
		return x.ClassMode
	}
	return course.ClassMode_CLASS_MODE_UNSPECIFIED
}

func (x *StudentLecture) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *StudentLecture) GetSectionCount() int64 {
	if x != nil {
		return x.SectionCount
	}
	return 0
}

func (x *StudentLecture) GetHomeworkCount() int64 {
	if x != nil {
		return x.HomeworkCount
	}
	return 0
}

func (x *StudentLecture) GetTestCount() int64 {
	if x != nil {
		return x.TestCount
	}
	return 0
}

func (x *StudentLecture) GetFinishedSectionCount() int64 {
	if x != nil {
		return x.FinishedSectionCount
	}
	return 0
}

func (x *StudentLecture) GetFinishedHomeworkCount() int64 {
	if x != nil {
		return x.FinishedHomeworkCount
	}
	return 0
}

func (x *StudentLecture) GetFinishedTestCount() int64 {
	if x != nil {
		return x.FinishedTestCount
	}
	return 0
}

var File_relo_protocol_vo_student_lecture_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_student_lecture_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x04, 0x0a, 0x0e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_relo_protocol_vo_student_lecture_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_student_lecture_proto_rawDescData = file_relo_protocol_vo_student_lecture_proto_rawDesc
)

func file_relo_protocol_vo_student_lecture_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_student_lecture_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_student_lecture_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_student_lecture_proto_rawDescData)
	})
	return file_relo_protocol_vo_student_lecture_proto_rawDescData
}

var file_relo_protocol_vo_student_lecture_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_student_lecture_proto_goTypes = []interface{}{
	(*StudentLecture)(nil), // 0: relo.protocol.vo.StudentLecture
	(course.Type)(0),       // 1: relo.core.course.Type
	(course.ClassMode)(0),  // 2: relo.core.course.ClassMode
}
var file_relo_protocol_vo_student_lecture_proto_depIdxs = []int32{
	1, // 0: relo.protocol.vo.StudentLecture.type:type_name -> relo.core.course.Type
	2, // 1: relo.protocol.vo.StudentLecture.class_mode:type_name -> relo.core.course.ClassMode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_student_lecture_proto_init() }
func file_relo_protocol_vo_student_lecture_proto_init() {
	if File_relo_protocol_vo_student_lecture_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_student_lecture_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentLecture); i {
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
			RawDescriptor: file_relo_protocol_vo_student_lecture_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_student_lecture_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_student_lecture_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_student_lecture_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_student_lecture_proto = out.File
	file_relo_protocol_vo_student_lecture_proto_rawDesc = nil
	file_relo_protocol_vo_student_lecture_proto_goTypes = nil
	file_relo_protocol_vo_student_lecture_proto_depIdxs = nil
}
