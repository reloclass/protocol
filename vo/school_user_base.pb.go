// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/school_user_base.proto

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

// 学校用户基本信息
type SchoolUserBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户编号
	// @gotags:
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 学校编号
	// @gotags: json:"schoolId,string"
	SchoolId int64 `protobuf:"varint,3,opt,name=school_id,json=schoolId,proto3" json:"schoolId,string"`
	// 手机号
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	// 是否老师
	// @gotags: json:"isTeacher"
	IsTeacher bool `protobuf:"varint,5,opt,name=is_teacher,json=isTeacher,proto3" json:"isTeacher"`
	// 是否学生
	// @gotags: json:"isStudent"
	IsStudent bool `protobuf:"varint,6,opt,name=is_student,json=isStudent,proto3" json:"isStudent"`
	// 是否管理员
	// @gotags: json:"isConsole"
	IsConsole bool `protobuf:"varint,7,opt,name=is_console,json=isConsole,proto3" json:"isConsole"`
	// 老师状态 0正常、1停用
	// @gotags: json:"teacherStatus"
	TeacherStatus user.Status `protobuf:"varint,8,opt,name=teacher_status,json=teacherStatus,proto3,enum=relo.core.user.Status" json:"teacherStatus"`
	// 学生状态 0正常、1停用
	// @gotags: json:"studentStatus"
	StudentStatus user.Status `protobuf:"varint,9,opt,name=student_status,json=studentStatus,proto3,enum=relo.core.user.Status" json:"studentStatus"`
	// 管理员状态 0正常、1停用
	// @gotags: json:"consoleStatus"
	ConsoleStatus user.Status `protobuf:"varint,10,opt,name=console_status,json=consoleStatus,proto3,enum=relo.core.user.Status" json:"consoleStatus"`
	// 教师名称
	// @gotags: json:"teacherName"
	TeacherName string `protobuf:"bytes,11,opt,name=teacher_name,json=teacherName,proto3" json:"teacherName"`
	// 学生名称
	// @gotags: json:"studentName"
	StudentName string `protobuf:"bytes,12,opt,name=student_name,json=studentName,proto3" json:"studentName"`
	// 后台名称
	// @gotags: json:"consoleName"
	ConsoleName string `protobuf:"bytes,13,opt,name=console_name,json=consoleName,proto3" json:"consoleName"`
	// 创建时间
	// @gotags: json:"createdAt"
	CreatedAt string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"createdAt"`
	// 更新时间
	// @gotags: json:"updatedAt"
	UpdatedAt string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updatedAt"`
}

func (x *SchoolUserBase) Reset() {
	*x = SchoolUserBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_school_user_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchoolUserBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolUserBase) ProtoMessage() {}

func (x *SchoolUserBase) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_school_user_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchoolUserBase.ProtoReflect.Descriptor instead.
func (*SchoolUserBase) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_school_user_base_proto_rawDescGZIP(), []int{0}
}

func (x *SchoolUserBase) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SchoolUserBase) GetSchoolId() int64 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *SchoolUserBase) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SchoolUserBase) GetIsTeacher() bool {
	if x != nil {
		return x.IsTeacher
	}
	return false
}

func (x *SchoolUserBase) GetIsStudent() bool {
	if x != nil {
		return x.IsStudent
	}
	return false
}

func (x *SchoolUserBase) GetIsConsole() bool {
	if x != nil {
		return x.IsConsole
	}
	return false
}

func (x *SchoolUserBase) GetTeacherStatus() user.Status {
	if x != nil {
		return x.TeacherStatus
	}
	return user.Status_NORMAL
}

func (x *SchoolUserBase) GetStudentStatus() user.Status {
	if x != nil {
		return x.StudentStatus
	}
	return user.Status_NORMAL
}

func (x *SchoolUserBase) GetConsoleStatus() user.Status {
	if x != nil {
		return x.ConsoleStatus
	}
	return user.Status_NORMAL
}

func (x *SchoolUserBase) GetTeacherName() string {
	if x != nil {
		return x.TeacherName
	}
	return ""
}

func (x *SchoolUserBase) GetStudentName() string {
	if x != nil {
		return x.StudentName
	}
	return ""
}

func (x *SchoolUserBase) GetConsoleName() string {
	if x != nil {
		return x.ConsoleName
	}
	return ""
}

func (x *SchoolUserBase) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SchoolUserBase) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_relo_protocol_vo_school_user_base_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_school_user_base_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x6c,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x04, 0x0a, 0x0e, 0x53, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x65,
	0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_school_user_base_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_school_user_base_proto_rawDescData = file_relo_protocol_vo_school_user_base_proto_rawDesc
)

func file_relo_protocol_vo_school_user_base_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_school_user_base_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_school_user_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_school_user_base_proto_rawDescData)
	})
	return file_relo_protocol_vo_school_user_base_proto_rawDescData
}

var file_relo_protocol_vo_school_user_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_school_user_base_proto_goTypes = []interface{}{
	(*SchoolUserBase)(nil), // 0: relo.protocol.vo.SchoolUserBase
	(user.Status)(0),       // 1: relo.core.user.Status
}
var file_relo_protocol_vo_school_user_base_proto_depIdxs = []int32{
	1, // 0: relo.protocol.vo.SchoolUserBase.teacher_status:type_name -> relo.core.user.Status
	1, // 1: relo.protocol.vo.SchoolUserBase.student_status:type_name -> relo.core.user.Status
	1, // 2: relo.protocol.vo.SchoolUserBase.console_status:type_name -> relo.core.user.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_school_user_base_proto_init() }
func file_relo_protocol_vo_school_user_base_proto_init() {
	if File_relo_protocol_vo_school_user_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_school_user_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchoolUserBase); i {
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
			RawDescriptor: file_relo_protocol_vo_school_user_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_school_user_base_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_school_user_base_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_school_user_base_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_school_user_base_proto = out.File
	file_relo_protocol_vo_school_user_base_proto_rawDesc = nil
	file_relo_protocol_vo_school_user_base_proto_goTypes = nil
	file_relo_protocol_vo_school_user_base_proto_depIdxs = nil
}
