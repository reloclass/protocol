// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/course.proto

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

// 课程信息
type Course struct {
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
	// 用户数
	// @gotags: json:"userCount"
	UserCount int64 `protobuf:"varint,12,opt,name=user_count,json=userCount,proto3" json:"userCount"`
	// 讲次数
	// @gotags: json:"lectureCount"
	LectureCount int64 `protobuf:"varint,13,opt,name=lecture_count,json=lectureCount,proto3" json:"lectureCount"`
	// 封面
	// @gotags: json:"coverUrl"
	CoverUrl string `protobuf:"bytes,14,opt,name=cover_url,json=coverUrl,proto3" json:"coverUrl"`
	// 老师列表
	Teachers []*SchoolUser `protobuf:"bytes,15,rep,name=teachers,proto3" json:"teachers,omitempty"`
	// 创建人
	Creator *SchoolUserDetail `protobuf:"bytes,16,opt,name=creator,proto3" json:"creator,omitempty"`
	// 学生列表
	Students []*SchoolUser `protobuf:"bytes,17,rep,name=students,proto3" json:"students,omitempty"`
	// 课程教学资料
	Materials []*TeachingMaterial `protobuf:"bytes,18,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Course) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Course) GetSchoolId() int64 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetType() course.Type {
	if x != nil {
		return x.Type
	}
	return course.Type_COURSE_TYPE_UNSPECIFIED
}

func (x *Course) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Course) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Course) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *Course) GetClassMode() course.ClassMode {
	if x != nil {
		return x.ClassMode
	}
	return course.ClassMode_CLASS_MODE_UNSPECIFIED
}

func (x *Course) GetUserCount() int64 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *Course) GetLectureCount() int64 {
	if x != nil {
		return x.LectureCount
	}
	return 0
}

func (x *Course) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Course) GetTeachers() []*SchoolUser {
	if x != nil {
		return x.Teachers
	}
	return nil
}

func (x *Course) GetCreator() *SchoolUserDetail {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Course) GetStudents() []*SchoolUser {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *Course) GetMaterials() []*TeachingMaterial {
	if x != nil {
		return x.Materials
	}
	return nil
}

var File_relo_protocol_vo_course_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_course_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76,
	0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x76, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x05, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x3a, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x08,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76,
	0x6f, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x40,
	0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_course_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_course_proto_rawDescData = file_relo_protocol_vo_course_proto_rawDesc
)

func file_relo_protocol_vo_course_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_course_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_course_proto_rawDescData)
	})
	return file_relo_protocol_vo_course_proto_rawDescData
}

var file_relo_protocol_vo_course_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_course_proto_goTypes = []interface{}{
	(*Course)(nil),           // 0: relo.protocol.vo.Course
	(course.Type)(0),         // 1: relo.core.course.Type
	(course.ClassMode)(0),    // 2: relo.core.course.ClassMode
	(*SchoolUser)(nil),       // 3: relo.protocol.vo.SchoolUser
	(*SchoolUserDetail)(nil), // 4: relo.protocol.vo.SchoolUserDetail
	(*TeachingMaterial)(nil), // 5: relo.protocol.vo.TeachingMaterial
}
var file_relo_protocol_vo_course_proto_depIdxs = []int32{
	1, // 0: relo.protocol.vo.Course.type:type_name -> relo.core.course.Type
	2, // 1: relo.protocol.vo.Course.class_mode:type_name -> relo.core.course.ClassMode
	3, // 2: relo.protocol.vo.Course.teachers:type_name -> relo.protocol.vo.SchoolUser
	4, // 3: relo.protocol.vo.Course.creator:type_name -> relo.protocol.vo.SchoolUserDetail
	3, // 4: relo.protocol.vo.Course.students:type_name -> relo.protocol.vo.SchoolUser
	5, // 5: relo.protocol.vo.Course.materials:type_name -> relo.protocol.vo.TeachingMaterial
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_course_proto_init() }
func file_relo_protocol_vo_course_proto_init() {
	if File_relo_protocol_vo_course_proto != nil {
		return
	}
	file_relo_protocol_vo_teaching_material_proto_init()
	file_relo_protocol_vo_school_user_proto_init()
	file_relo_protocol_vo_school_user_detail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
			RawDescriptor: file_relo_protocol_vo_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_course_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_course_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_course_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_course_proto = out.File
	file_relo_protocol_vo_course_proto_rawDesc = nil
	file_relo_protocol_vo_course_proto_goTypes = nil
	file_relo_protocol_vo_course_proto_depIdxs = nil
}