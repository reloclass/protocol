// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/school.proto

package vo

import (
	school "github.com/reloclass/core/school"
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

// 学校信息
type School struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 学校编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string"`
	// 学校名
	// @gotags: json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// 学校状态
	// @gotags: json:"status"
	Status school.Status `protobuf:"varint,4,opt,name=status,proto3,enum=relo.core.school.Status" json:"status"`
	// 负责人手机号
	// @gotags: json:"principalPhone"
	PrincipalPhone string `protobuf:"bytes,5,opt,name=principal_phone,json=principalPhone,proto3" json:"principalPhone"`
	// 负责人手机号
	// @gotags: json:"principalName"
	PrincipalName string `protobuf:"bytes,6,opt,name=principal_name,json=principalName,proto3" json:"principalName"`
	// 文件编号
	// @gotags: json:"logoUrl"
	LogoUrl string `protobuf:"bytes,7,opt,name=logo_url,json=logoUrl,proto3" json:"logoUrl"`
	// 详细地址
	// @gotags: json:"address"
	Address string `protobuf:"bytes,8,opt,name=address,proto3" json:"address"`
	// 详细地址
	// @gotags: json:"areaCode"
	AreaCode string `protobuf:"bytes,9,opt,name=area_code,json=areaCode,proto3" json:"areaCode"`
	// 经度
	// @gotags: json:"longitude"
	Longitude float32 `protobuf:"fixed32,10,opt,name=longitude,proto3" json:"longitude"`
	// 纬度
	// @gotags: json:"latitude"
	Latitude float32 `protobuf:"fixed32,11,opt,name=latitude,proto3" json:"latitude"`
	// 教师人数
	// @gotags: json:"teacherNumber"
	TeacherNumber int32 `protobuf:"varint,12,opt,name=teacher_number,json=teacherNumber,proto3" json:"teacherNumber"`
	// 学生人数
	// @gotags: json:"studentNumber"
	StudentNumber int32 `protobuf:"varint,13,opt,name=student_number,json=studentNumber,proto3" json:"studentNumber"`
	// 区域地址
	// @gotags：json:"areaAddress"
	AreaAddress string `protobuf:"bytes,14,opt,name=area_address,json=areaAddress,proto3" json:"area_address,omitempty"`
	// 学校logo设置
	// @gotags：json:"logoFileId"
	LogoFileId string `protobuf:"bytes,15,opt,name=logo_file_id,json=logoFileId,proto3" json:"logo_file_id,omitempty"`
	// 开设课程
	// @gotags: json:"setUpCourses"
	Courses []*school.Course `protobuf:"bytes,16,rep,name=courses,proto3" json:"setUpCourses"`
}

func (x *School) Reset() {
	*x = School{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_school_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *School) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*School) ProtoMessage() {}

func (x *School) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_school_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use School.ProtoReflect.Descriptor instead.
func (*School) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_school_proto_rawDescGZIP(), []int{0}
}

func (x *School) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *School) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *School) GetStatus() school.Status {
	if x != nil {
		return x.Status
	}
	return school.Status_AVAILABLE
}

func (x *School) GetPrincipalPhone() string {
	if x != nil {
		return x.PrincipalPhone
	}
	return ""
}

func (x *School) GetPrincipalName() string {
	if x != nil {
		return x.PrincipalName
	}
	return ""
}

func (x *School) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *School) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *School) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *School) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *School) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *School) GetTeacherNumber() int32 {
	if x != nil {
		return x.TeacherNumber
	}
	return 0
}

func (x *School) GetStudentNumber() int32 {
	if x != nil {
		return x.StudentNumber
	}
	return 0
}

func (x *School) GetAreaAddress() string {
	if x != nil {
		return x.AreaAddress
	}
	return ""
}

func (x *School) GetLogoFileId() string {
	if x != nil {
		return x.LogoFileId
	}
	return ""
}

func (x *School) GetCourses() []*school.Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

// 学校基本信息
type BaseSchool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 学校id
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string"`
	// 学校名
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 学校logoUrl
	// @gotags: json:"logoUrl"
	LogoUrl string `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logoUrl"`
}

func (x *BaseSchool) Reset() {
	*x = BaseSchool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_school_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseSchool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseSchool) ProtoMessage() {}

func (x *BaseSchool) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_school_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseSchool.ProtoReflect.Descriptor instead.
func (*BaseSchool) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_school_proto_rawDescGZIP(), []int{1}
}

func (x *BaseSchool) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseSchool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BaseSchool) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

var File_relo_protocol_vo_school_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_school_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76,
	0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x81, 0x04, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72,
	0x65, 0x61, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a,
	0x0c, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x32, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x53, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_school_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_school_proto_rawDescData = file_relo_protocol_vo_school_proto_rawDesc
)

func file_relo_protocol_vo_school_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_school_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_school_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_school_proto_rawDescData)
	})
	return file_relo_protocol_vo_school_proto_rawDescData
}

var file_relo_protocol_vo_school_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_vo_school_proto_goTypes = []interface{}{
	(*School)(nil),        // 0: relo.protocol.vo.School
	(*BaseSchool)(nil),    // 1: relo.protocol.vo.BaseSchool
	(school.Status)(0),    // 2: relo.core.school.Status
	(*school.Course)(nil), // 3: relo.core.school.Course
}
var file_relo_protocol_vo_school_proto_depIdxs = []int32{
	2, // 0: relo.protocol.vo.School.status:type_name -> relo.core.school.Status
	3, // 1: relo.protocol.vo.School.courses:type_name -> relo.core.school.Course
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_school_proto_init() }
func file_relo_protocol_vo_school_proto_init() {
	if File_relo_protocol_vo_school_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_school_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*School); i {
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
		file_relo_protocol_vo_school_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseSchool); i {
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
			RawDescriptor: file_relo_protocol_vo_school_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_school_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_school_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_school_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_school_proto = out.File
	file_relo_protocol_vo_school_proto_rawDesc = nil
	file_relo_protocol_vo_school_proto_goTypes = nil
	file_relo_protocol_vo_school_proto_depIdxs = nil
}
