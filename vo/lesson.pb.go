// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/lesson.proto

package vo

import (
	course "github.com/reloclass/core/course"
	lesson "github.com/reloclass/core/lesson"
	schedule "github.com/reloclass/core/schedule"
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

// 课节信息
type Lesson struct {
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
	Status lesson.Status `protobuf:"varint,10,opt,name=status,proto3,enum=relo.core.lesson.Status" json:"status,omitempty"`
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
	AllowedType schedule.AllowedType `protobuf:"varint,15,opt,name=allowed_type,json=allowedType,proto3,enum=relo.core.schedule.AllowedType" json:"allowedType"`
	// 验证码
	Captcha string `protobuf:"bytes,16,opt,name=captcha,proto3" json:"captcha,omitempty"`
	// 录制操作类型
	// @gotags: json:"recordType"
	RecordType schedule.RecordType `protobuf:"varint,17,opt,name=record_type,json=recordType,proto3,enum=relo.core.schedule.RecordType" json:"recordType"`
	// 是否允许直播
	// @gotags: json:"liveAllowedType"
	LiveAllowedType schedule.LiveAllowedType `protobuf:"varint,18,opt,name=live_allowed_type,json=liveAllowedType,proto3,enum=relo.core.schedule.LiveAllowedType" json:"liveAllowedType"`
	// 视频质量
	// @gotags: json:"videoDimension"
	VideoDimension schedule.VideoDimensionType `protobuf:"varint,19,opt,name=video_dimension,json=videoDimension,proto3,enum=relo.core.schedule.VideoDimensionType" json:"videoDimension"`
	// 上台人数
	// @gotags: json:"stageNumber"
	StageNumber int32 `protobuf:"varint,20,opt,name=stage_number,json=stageNumber,proto3" json:"stageNumber"`
}

func (x *Lesson) Reset() {
	*x = Lesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_lesson_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lesson) ProtoMessage() {}

func (x *Lesson) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_lesson_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lesson.ProtoReflect.Descriptor instead.
func (*Lesson) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_lesson_proto_rawDescGZIP(), []int{0}
}

func (x *Lesson) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lesson) GetSchoolId() int64 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *Lesson) GetCourseScheduleId() int64 {
	if x != nil {
		return x.CourseScheduleId
	}
	return 0
}

func (x *Lesson) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *Lesson) GetTeacherId() int64 {
	if x != nil {
		return x.TeacherId
	}
	return 0
}

func (x *Lesson) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lesson) GetClassTime() string {
	if x != nil {
		return x.ClassTime
	}
	return ""
}

func (x *Lesson) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Lesson) GetStatus() lesson.Status {
	if x != nil {
		return x.Status
	}
	return lesson.Status_CREATED
}

func (x *Lesson) GetMeetingId() string {
	if x != nil {
		return x.MeetingId
	}
	return ""
}

func (x *Lesson) GetMeetingNo() uint64 {
	if x != nil {
		return x.MeetingNo
	}
	return 0
}

func (x *Lesson) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Lesson) GetClassMode() course.ClassMode {
	if x != nil {
		return x.ClassMode
	}
	return course.ClassMode_CLASS_MODE_UNSPECIFIED
}

func (x *Lesson) GetAllowedType() schedule.AllowedType {
	if x != nil {
		return x.AllowedType
	}
	return schedule.AllowedType_ALLOWED_TYPE_UNSPECIFIED
}

func (x *Lesson) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

func (x *Lesson) GetRecordType() schedule.RecordType {
	if x != nil {
		return x.RecordType
	}
	return schedule.RecordType_RECORD_TYPE_UNSPECIFIED
}

func (x *Lesson) GetLiveAllowedType() schedule.LiveAllowedType {
	if x != nil {
		return x.LiveAllowedType
	}
	return schedule.LiveAllowedType_LIVE_ALLOWED_TYPE_UNSPECIFIED
}

func (x *Lesson) GetVideoDimension() schedule.VideoDimensionType {
	if x != nil {
		return x.VideoDimension
	}
	return schedule.VideoDimensionType_VIDEO_DIMENSION_TYPE_UNSPECIFIED
}

func (x *Lesson) GetStageNumber() int32 {
	if x != nil {
		return x.StageNumber
	}
	return 0
}

var File_relo_protocol_vo_lesson_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_lesson_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76,
	0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x72, 0x65, 0x6c, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x72, 0x65,
	0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x06, 0x0a, 0x06,
	0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x65,
	0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x12, 0x3f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x4f, 0x0a, 0x11, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0f, 0x6c, 0x69, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x64, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_lesson_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_lesson_proto_rawDescData = file_relo_protocol_vo_lesson_proto_rawDesc
)

func file_relo_protocol_vo_lesson_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_lesson_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_lesson_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_lesson_proto_rawDescData)
	})
	return file_relo_protocol_vo_lesson_proto_rawDescData
}

var file_relo_protocol_vo_lesson_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_lesson_proto_goTypes = []interface{}{
	(*Lesson)(nil),                   // 0: relo.protocol.vo.Lesson
	(lesson.Status)(0),               // 1: relo.core.lesson.Status
	(course.ClassMode)(0),            // 2: relo.core.course.ClassMode
	(schedule.AllowedType)(0),        // 3: relo.core.schedule.AllowedType
	(schedule.RecordType)(0),         // 4: relo.core.schedule.RecordType
	(schedule.LiveAllowedType)(0),    // 5: relo.core.schedule.LiveAllowedType
	(schedule.VideoDimensionType)(0), // 6: relo.core.schedule.VideoDimensionType
}
var file_relo_protocol_vo_lesson_proto_depIdxs = []int32{
	1, // 0: relo.protocol.vo.Lesson.status:type_name -> relo.core.lesson.Status
	2, // 1: relo.protocol.vo.Lesson.class_mode:type_name -> relo.core.course.ClassMode
	3, // 2: relo.protocol.vo.Lesson.allowed_type:type_name -> relo.core.schedule.AllowedType
	4, // 3: relo.protocol.vo.Lesson.record_type:type_name -> relo.core.schedule.RecordType
	5, // 4: relo.protocol.vo.Lesson.live_allowed_type:type_name -> relo.core.schedule.LiveAllowedType
	6, // 5: relo.protocol.vo.Lesson.video_dimension:type_name -> relo.core.schedule.VideoDimensionType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_lesson_proto_init() }
func file_relo_protocol_vo_lesson_proto_init() {
	if File_relo_protocol_vo_lesson_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_lesson_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lesson); i {
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
			RawDescriptor: file_relo_protocol_vo_lesson_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_lesson_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_lesson_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_lesson_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_lesson_proto = out.File
	file_relo_protocol_vo_lesson_proto_rawDesc = nil
	file_relo_protocol_vo_lesson_proto_goTypes = nil
	file_relo_protocol_vo_lesson_proto_depIdxs = nil
}
