// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/es_course_time.proto

package vo

import (
	course "github.com/reloclass/core/course"
	courseschedule "github.com/reloclass/core/courseschedule"
	coursetime "github.com/reloclass/core/coursetime"
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

// ES课节信息
type EsCourseTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课节编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string"`
	// 学校编号
	// @gotags: json:"schoolId,string"
	SchoolId int64 `protobuf:"varint,3,opt,name=school_id,json=schoolId,proto3" json:"schoolId,string"`
	// 课程计划
	// @gotags: json:"courseScheduleId,string"
	CourseScheduleId int64 `protobuf:"varint,4,opt,name=course_schedule_id,json=courseScheduleId,proto3" json:"courseScheduleId,string"`
	// 课程编号
	// @gotags: json:"courseId,string"
	CourseId int64 `protobuf:"varint,5,opt,name=course_id,json=courseId,proto3" json:"courseId,string"`
	// 主讲老师
	// @gotags: json:"teacherId,string"
	TeacherId int64 `protobuf:"varint,6,opt,name=teacher_id,json=teacherId,proto3" json:"teacherId,string"`
	// 课节名称
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	// 上课时间
	// @gotags: json:"classTime"
	ClassTime string `protobuf:"bytes,8,opt,name=class_time,json=classTime,proto3" json:"classTime"`
	// 授课时长 单位：分钟
	Duration int32 `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// 课节状态
	Status coursetime.Status `protobuf:"varint,10,opt,name=status,proto3,enum=relo.core.coursetime.Status" json:"status,omitempty"`
	// 云视讯会议Id
	// @gotags: json:"meetingId"
	MeetingId string `protobuf:"bytes,11,opt,name=meeting_id,json=meetingId,proto3" json:"meetingId"`
	// 云视讯会议No
	// @gotags: json:"meetingNo,string"
	MeetingNo uint64 `protobuf:"varint,12,opt,name=meeting_no,json=meetingNo,proto3" json:"meetingNo,string"`
	// 组ID
	// @gotags: json:"groupId,string"
	GroupId int64 `protobuf:"varint,13,opt,name=group_id,json=groupId,proto3" json:"groupId,string"`
	// 上课模式
	// @gotags: json:"classMode"
	ClassMode course.ClassMode `protobuf:"varint,14,opt,name=class_mode,json=classMode,proto3,enum=relo.core.course.ClassMode" json:"classMode"`
	// 允许类型
	// @gotags: json:"allowedType"
	AllowedType courseschedule.AllowedType `protobuf:"varint,15,opt,name=allowed_type,json=allowedType,proto3,enum=relo.core.courseschedule.AllowedType" json:"allowedType"`
	// 验证码
	Captcha string `protobuf:"bytes,16,opt,name=captcha,proto3" json:"captcha,omitempty"`
	// 录制操作类型
	// @gotags: json:"recordType"
	RecordType courseschedule.RecordType `protobuf:"varint,17,opt,name=record_type,json=recordType,proto3,enum=relo.core.courseschedule.RecordType" json:"recordType"`
	// 是否允许直播
	// @gotags: json:"liveAllowedType"
	LiveAllowedType courseschedule.LiveAllowedType `protobuf:"varint,18,opt,name=live_allowed_type,json=liveAllowedType,proto3,enum=relo.core.courseschedule.LiveAllowedType" json:"liveAllowedType"`
	// 视频质量
	// @gotags: json:"videoDimension"
	VideoDimension courseschedule.VideoDimensionType `protobuf:"varint,19,opt,name=video_dimension,json=videoDimension,proto3,enum=relo.core.courseschedule.VideoDimensionType" json:"videoDimension"`
	// 上台人数
	// @gotags: json:"stageNumber"
	StageNumber int32 `protobuf:"varint,20,opt,name=stage_number,json=stageNumber,proto3" json:"stageNumber"`
	// 排课计划
	// @gotags: json:"courseSchedule"
	CourseSchedule *CourseSchedule `protobuf:"bytes,21,opt,name=course_schedule,json=courseSchedule,proto3" json:"courseSchedule"`
	// 课程
	Course *Course `protobuf:"bytes,22,opt,name=course,proto3" json:"course,omitempty"`
	// 讲师
	Lecture *SchoolUserDetail `protobuf:"bytes,23,opt,name=lecture,proto3" json:"lecture,omitempty"`
	// 助教
	Assistants []*SchoolUserDetail `protobuf:"bytes,24,rep,name=assistants,proto3" json:"assistants,omitempty"`
	// 学生
	Students []*SchoolUserDetail `protobuf:"bytes,25,rep,name=students,proto3" json:"students,omitempty"`
	// 录课资源
	Records []*CourseTimeRecord `protobuf:"bytes,26,rep,name=records,proto3" json:"records,omitempty"`
	// 教学资料
	Materials []*TeachingMaterial `protobuf:"bytes,27,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *EsCourseTime) Reset() {
	*x = EsCourseTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_es_course_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsCourseTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsCourseTime) ProtoMessage() {}

func (x *EsCourseTime) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_es_course_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsCourseTime.ProtoReflect.Descriptor instead.
func (*EsCourseTime) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_es_course_time_proto_rawDescGZIP(), []int{0}
}

func (x *EsCourseTime) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EsCourseTime) GetSchoolId() int64 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *EsCourseTime) GetCourseScheduleId() int64 {
	if x != nil {
		return x.CourseScheduleId
	}
	return 0
}

func (x *EsCourseTime) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *EsCourseTime) GetTeacherId() int64 {
	if x != nil {
		return x.TeacherId
	}
	return 0
}

func (x *EsCourseTime) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EsCourseTime) GetClassTime() string {
	if x != nil {
		return x.ClassTime
	}
	return ""
}

func (x *EsCourseTime) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *EsCourseTime) GetStatus() coursetime.Status {
	if x != nil {
		return x.Status
	}
	return coursetime.Status_CREATED
}

func (x *EsCourseTime) GetMeetingId() string {
	if x != nil {
		return x.MeetingId
	}
	return ""
}

func (x *EsCourseTime) GetMeetingNo() uint64 {
	if x != nil {
		return x.MeetingNo
	}
	return 0
}

func (x *EsCourseTime) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *EsCourseTime) GetClassMode() course.ClassMode {
	if x != nil {
		return x.ClassMode
	}
	return course.ClassMode_CLASS_MODE_UNSPECIFIED
}

func (x *EsCourseTime) GetAllowedType() courseschedule.AllowedType {
	if x != nil {
		return x.AllowedType
	}
	return courseschedule.AllowedType_ALLOWED_TYPE_UNSPECIFIED
}

func (x *EsCourseTime) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

func (x *EsCourseTime) GetRecordType() courseschedule.RecordType {
	if x != nil {
		return x.RecordType
	}
	return courseschedule.RecordType_RECORD_TYPE_UNSPECIFIED
}

func (x *EsCourseTime) GetLiveAllowedType() courseschedule.LiveAllowedType {
	if x != nil {
		return x.LiveAllowedType
	}
	return courseschedule.LiveAllowedType_LIVE_ALLOWED_TYPE_UNSPECIFIED
}

func (x *EsCourseTime) GetVideoDimension() courseschedule.VideoDimensionType {
	if x != nil {
		return x.VideoDimension
	}
	return courseschedule.VideoDimensionType_VIDEO_DIMENSION_TYPE_UNSPECIFIED
}

func (x *EsCourseTime) GetStageNumber() int32 {
	if x != nil {
		return x.StageNumber
	}
	return 0
}

func (x *EsCourseTime) GetCourseSchedule() *CourseSchedule {
	if x != nil {
		return x.CourseSchedule
	}
	return nil
}

func (x *EsCourseTime) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

func (x *EsCourseTime) GetLecture() *SchoolUserDetail {
	if x != nil {
		return x.Lecture
	}
	return nil
}

func (x *EsCourseTime) GetAssistants() []*SchoolUserDetail {
	if x != nil {
		return x.Assistants
	}
	return nil
}

func (x *EsCourseTime) GetStudents() []*SchoolUserDetail {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *EsCourseTime) GetRecords() []*CourseTimeRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *EsCourseTime) GetMaterials() []*TeachingMaterial {
	if x != nil {
		return x.Materials
	}
	return nil
}

var File_relo_protocol_vo_es_course_time_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_es_course_time_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x65,
	0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x72, 0x65,
	0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x72, 0x65, 0x6c, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x72, 0x65,
	0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x09, 0x0a, 0x0c, 0x45, 0x73, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4e,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x45, 0x0a, 0x0b,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x11, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x6c, 0x69, 0x76, 0x65, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x0f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x0e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x30, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x42, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x18, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x1a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x40, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x1b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_relo_protocol_vo_es_course_time_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_es_course_time_proto_rawDescData = file_relo_protocol_vo_es_course_time_proto_rawDesc
)

func file_relo_protocol_vo_es_course_time_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_es_course_time_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_es_course_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_es_course_time_proto_rawDescData)
	})
	return file_relo_protocol_vo_es_course_time_proto_rawDescData
}

var file_relo_protocol_vo_es_course_time_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_es_course_time_proto_goTypes = []interface{}{
	(*EsCourseTime)(nil),                   // 0: relo.protocol.vo.EsCourseTime
	(coursetime.Status)(0),                 // 1: relo.core.coursetime.Status
	(course.ClassMode)(0),                  // 2: relo.core.course.ClassMode
	(courseschedule.AllowedType)(0),        // 3: relo.core.courseschedule.AllowedType
	(courseschedule.RecordType)(0),         // 4: relo.core.courseschedule.RecordType
	(courseschedule.LiveAllowedType)(0),    // 5: relo.core.courseschedule.LiveAllowedType
	(courseschedule.VideoDimensionType)(0), // 6: relo.core.courseschedule.VideoDimensionType
	(*CourseSchedule)(nil),                 // 7: relo.protocol.vo.CourseSchedule
	(*Course)(nil),                         // 8: relo.protocol.vo.Course
	(*SchoolUserDetail)(nil),               // 9: relo.protocol.vo.SchoolUserDetail
	(*CourseTimeRecord)(nil),               // 10: relo.protocol.vo.CourseTimeRecord
	(*TeachingMaterial)(nil),               // 11: relo.protocol.vo.TeachingMaterial
}
var file_relo_protocol_vo_es_course_time_proto_depIdxs = []int32{
	1,  // 0: relo.protocol.vo.EsCourseTime.status:type_name -> relo.core.coursetime.Status
	2,  // 1: relo.protocol.vo.EsCourseTime.class_mode:type_name -> relo.core.course.ClassMode
	3,  // 2: relo.protocol.vo.EsCourseTime.allowed_type:type_name -> relo.core.courseschedule.AllowedType
	4,  // 3: relo.protocol.vo.EsCourseTime.record_type:type_name -> relo.core.courseschedule.RecordType
	5,  // 4: relo.protocol.vo.EsCourseTime.live_allowed_type:type_name -> relo.core.courseschedule.LiveAllowedType
	6,  // 5: relo.protocol.vo.EsCourseTime.video_dimension:type_name -> relo.core.courseschedule.VideoDimensionType
	7,  // 6: relo.protocol.vo.EsCourseTime.course_schedule:type_name -> relo.protocol.vo.CourseSchedule
	8,  // 7: relo.protocol.vo.EsCourseTime.course:type_name -> relo.protocol.vo.Course
	9,  // 8: relo.protocol.vo.EsCourseTime.lecture:type_name -> relo.protocol.vo.SchoolUserDetail
	9,  // 9: relo.protocol.vo.EsCourseTime.assistants:type_name -> relo.protocol.vo.SchoolUserDetail
	9,  // 10: relo.protocol.vo.EsCourseTime.students:type_name -> relo.protocol.vo.SchoolUserDetail
	10, // 11: relo.protocol.vo.EsCourseTime.records:type_name -> relo.protocol.vo.CourseTimeRecord
	11, // 12: relo.protocol.vo.EsCourseTime.materials:type_name -> relo.protocol.vo.TeachingMaterial
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_es_course_time_proto_init() }
func file_relo_protocol_vo_es_course_time_proto_init() {
	if File_relo_protocol_vo_es_course_time_proto != nil {
		return
	}
	file_relo_protocol_vo_course_schedule_proto_init()
	file_relo_protocol_vo_course_proto_init()
	file_relo_protocol_vo_school_user_detail_proto_init()
	file_relo_protocol_vo_course_time_record_proto_init()
	file_relo_protocol_vo_teaching_material_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_es_course_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsCourseTime); i {
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
			RawDescriptor: file_relo_protocol_vo_es_course_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_es_course_time_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_es_course_time_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_es_course_time_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_es_course_time_proto = out.File
	file_relo_protocol_vo_es_course_time_proto_rawDesc = nil
	file_relo_protocol_vo_es_course_time_proto_goTypes = nil
	file_relo_protocol_vo_es_course_time_proto_depIdxs = nil
}
