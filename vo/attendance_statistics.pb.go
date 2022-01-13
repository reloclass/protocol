// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/attendance_statistics.proto

package vo

import (
	user "github.com/reloclass/core/user"
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

// 课节出勤统计
type AttendanceStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 出勤编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string"`
	// 学校编号
	// @gotags: json:"schoolId,string"
	SchoolId int64 `protobuf:"varint,3,opt,name=school_id,json=schoolId,proto3" json:"schoolId,string"`
	// 用户在这节课中的角色类型
	// @gotags: json:"userType"
	UserType user.Type `protobuf:"varint,4,opt,name=user_type,json=userType,proto3,enum=relo.core.user.Type" json:"userType"`
	// 课节编号
	// @gotags: json:"courseTimeId,string"
	CourseTimeId int64 `protobuf:"varint,5,opt,name=course_time_id,json=courseTimeId,proto3" json:"courseTimeId,string"`
	// 出勤次数
	// @gotags: json:"attendanceTime"
	AttendanceTime int32 `protobuf:"varint,6,opt,name=attendance_time,json=attendanceTime,proto3" json:"attendanceTime"`
	// 出勤时间
	// @gotags: json:"attendanceAt"
	AttendanceAt string `protobuf:"bytes,7,opt,name=attendance_at,json=attendanceAt,proto3" json:"attendanceAt"`
	// 是否出勤
	// @gotags: json:"isAttendance"
	IsAttendance bool `protobuf:"varint,8,opt,name=is_attendance,json=isAttendance,proto3" json:"isAttendance"`
	// 是否迟到
	// @gotags: json:"isLate"
	IsLate bool `protobuf:"varint,9,opt,name=is_late,json=isLate,proto3" json:"isLate"`
	// 是否早退
	// @gotags: isEarly
	IsEarly bool `protobuf:"varint,10,opt,name=is_early,json=isEarly,proto3" json:"is_early,omitempty"`
	// 课程编号
	// @gotags: json:"courseId,string"
	CourseId int64 `protobuf:"varint,11,opt,name=course_id,json=courseId,proto3" json:"courseId,string"`
	// 课程名字
	// json:"courseName"
	CourseName string `protobuf:"bytes,12,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	// 课节名称
	Name string `protobuf:"bytes,13,opt,name=name,proto3" json:"name,omitempty"`
	// 授课时长 单位：分钟
	Duration int32 `protobuf:"varint,14,opt,name=duration,proto3" json:"duration,omitempty"`
	// 上课时间
	// @gotags: json:"classTime"
	ClassTime string `protobuf:"bytes,15,opt,name=class_time,json=classTime,proto3" json:"classTime"`
	// 上课时间排序字段
	// @gotags: json:"classTimeSortField"
	ClassTimeSortField int64 `protobuf:"varint,16,opt,name=class_time_sort_field,json=classTimeSortField,proto3" json:"classTimeSortField"`
	// 老师
	Teacher *UserDetail `protobuf:"bytes,17,opt,name=teacher,proto3" json:"teacher,omitempty"`
	// 助教
	Assistants []*UserDetail `protobuf:"bytes,18,rep,name=assistants,proto3" json:"assistants,omitempty"`
}

func (x *AttendanceStatistics) Reset() {
	*x = AttendanceStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_attendance_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttendanceStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttendanceStatistics) ProtoMessage() {}

func (x *AttendanceStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_attendance_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttendanceStatistics.ProtoReflect.Descriptor instead.
func (*AttendanceStatistics) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_attendance_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *AttendanceStatistics) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttendanceStatistics) GetSchoolId() int64 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *AttendanceStatistics) GetUserType() user.Type {
	if x != nil {
		return x.UserType
	}
	return user.Type_TYPE_UNSPECIFIED
}

func (x *AttendanceStatistics) GetCourseTimeId() int64 {
	if x != nil {
		return x.CourseTimeId
	}
	return 0
}

func (x *AttendanceStatistics) GetAttendanceTime() int32 {
	if x != nil {
		return x.AttendanceTime
	}
	return 0
}

func (x *AttendanceStatistics) GetAttendanceAt() string {
	if x != nil {
		return x.AttendanceAt
	}
	return ""
}

func (x *AttendanceStatistics) GetIsAttendance() bool {
	if x != nil {
		return x.IsAttendance
	}
	return false
}

func (x *AttendanceStatistics) GetIsLate() bool {
	if x != nil {
		return x.IsLate
	}
	return false
}

func (x *AttendanceStatistics) GetIsEarly() bool {
	if x != nil {
		return x.IsEarly
	}
	return false
}

func (x *AttendanceStatistics) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *AttendanceStatistics) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *AttendanceStatistics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttendanceStatistics) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *AttendanceStatistics) GetClassTime() string {
	if x != nil {
		return x.ClassTime
	}
	return ""
}

func (x *AttendanceStatistics) GetClassTimeSortField() int64 {
	if x != nil {
		return x.ClassTimeSortField
	}
	return 0
}

func (x *AttendanceStatistics) GetTeacher() *UserDetail {
	if x != nil {
		return x.Teacher
	}
	return nil
}

func (x *AttendanceStatistics) GetAssistants() []*UserDetail {
	if x != nil {
		return x.Assistants
	}
	return nil
}

var File_relo_protocol_vo_attendance_statistics_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_attendance_statistics_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f,
	0x1a, 0x19, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x72, 0x65, 0x6c,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf9, 0x04, 0x0a, 0x14, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61,
	0x6e, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4c, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x45, 0x61, 0x72, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6f, 0x72, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3c, 0x0a,
	0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b,
	0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_attendance_statistics_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_attendance_statistics_proto_rawDescData = file_relo_protocol_vo_attendance_statistics_proto_rawDesc
)

func file_relo_protocol_vo_attendance_statistics_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_attendance_statistics_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_attendance_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_attendance_statistics_proto_rawDescData)
	})
	return file_relo_protocol_vo_attendance_statistics_proto_rawDescData
}

var file_relo_protocol_vo_attendance_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_attendance_statistics_proto_goTypes = []interface{}{
	(*AttendanceStatistics)(nil), // 0: relo.protocol.vo.AttendanceStatistics
	(user.Type)(0),               // 1: relo.core.user.Type
	(*UserDetail)(nil),           // 2: relo.protocol.vo.UserDetail
}
var file_relo_protocol_vo_attendance_statistics_proto_depIdxs = []int32{
	1, // 0: relo.protocol.vo.AttendanceStatistics.user_type:type_name -> relo.core.user.Type
	2, // 1: relo.protocol.vo.AttendanceStatistics.teacher:type_name -> relo.protocol.vo.UserDetail
	2, // 2: relo.protocol.vo.AttendanceStatistics.assistants:type_name -> relo.protocol.vo.UserDetail
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_attendance_statistics_proto_init() }
func file_relo_protocol_vo_attendance_statistics_proto_init() {
	if File_relo_protocol_vo_attendance_statistics_proto != nil {
		return
	}
	file_relo_protocol_vo_user_detail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_attendance_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttendanceStatistics); i {
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
			RawDescriptor: file_relo_protocol_vo_attendance_statistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_attendance_statistics_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_attendance_statistics_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_attendance_statistics_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_attendance_statistics_proto = out.File
	file_relo_protocol_vo_attendance_statistics_proto_rawDesc = nil
	file_relo_protocol_vo_attendance_statistics_proto_goTypes = nil
	file_relo_protocol_vo_attendance_statistics_proto_depIdxs = nil
}
