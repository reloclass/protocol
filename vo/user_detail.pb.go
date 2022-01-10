// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/user_detail.proto

package vo

import (
	role "github.com/reloclass/core/role"
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

// 学校用户详细信息
type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"phone" validate:"required,mobile"
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone" validate:"required,mobile"`
	// @gotags: json:"name" validate:"required,min=2,max=32"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" validate:"required,min=2,max=32"`
	// 用户id
	// @gotags: json:"userId,string"
	UserId int64 `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"userId,string"`
	// 更新时间
	// @gotags: json:"updated_at"
	Updated string `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated_at"`
	// 图像链接
	// @gotags:  json:"avatarUrl"
	AvatarUrl string `protobuf:"bytes,7,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatarUrl"`
	// 状态
	Status user.Status `protobuf:"varint,8,opt,name=status,proto3,enum=relo.core.user.Status" json:"status,omitempty"`
	// 角色信息
	Roles []*role.Base `protobuf:"bytes,9,rep,name=roles,proto3" json:"roles,omitempty"`
	// 出勤统计
	// @gotags: json:"attendanceStatistics"
	AttendanceSummary *AttendanceSummary `protobuf:"bytes,10,opt,name=attendance_summary,json=attendanceSummary,proto3" json:"attendanceStatistics"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_user_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_user_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_user_detail_proto_rawDescGZIP(), []int{0}
}

func (x *UserDetail) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDetail) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserDetail) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *UserDetail) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserDetail) GetStatus() user.Status {
	if x != nil {
		return x.Status
	}
	return user.Status_NORMAL
}

func (x *UserDetail) GetRoles() []*role.Base {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UserDetail) GetAttendanceSummary() *AttendanceSummary {
	if x != nil {
		return x.AttendanceSummary
	}
	return nil
}

var File_relo_protocol_vo_user_detail_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_user_detail_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f,
	0x2f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x02, 0x0a, 0x0a, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x52, 0x0a, 0x12, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x11, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_user_detail_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_user_detail_proto_rawDescData = file_relo_protocol_vo_user_detail_proto_rawDesc
)

func file_relo_protocol_vo_user_detail_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_user_detail_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_user_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_user_detail_proto_rawDescData)
	})
	return file_relo_protocol_vo_user_detail_proto_rawDescData
}

var file_relo_protocol_vo_user_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_user_detail_proto_goTypes = []interface{}{
	(*UserDetail)(nil),        // 0: relo.protocol.vo.UserDetail
	(user.Status)(0),          // 1: relo.core.user.Status
	(*role.Base)(nil),         // 2: relo.core.role.Base
	(*AttendanceSummary)(nil), // 3: relo.protocol.vo.AttendanceSummary
}
var file_relo_protocol_vo_user_detail_proto_depIdxs = []int32{
	1, // 0: relo.protocol.vo.UserDetail.status:type_name -> relo.core.user.Status
	2, // 1: relo.protocol.vo.UserDetail.roles:type_name -> relo.core.role.Base
	3, // 2: relo.protocol.vo.UserDetail.attendance_summary:type_name -> relo.protocol.vo.AttendanceSummary
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_user_detail_proto_init() }
func file_relo_protocol_vo_user_detail_proto_init() {
	if File_relo_protocol_vo_user_detail_proto != nil {
		return
	}
	file_relo_protocol_vo_attendance_summary_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_user_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
			RawDescriptor: file_relo_protocol_vo_user_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_user_detail_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_user_detail_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_user_detail_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_user_detail_proto = out.File
	file_relo_protocol_vo_user_detail_proto_rawDesc = nil
	file_relo_protocol_vo_user_detail_proto_goTypes = nil
	file_relo_protocol_vo_user_detail_proto_depIdxs = nil
}