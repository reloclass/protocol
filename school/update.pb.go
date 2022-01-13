// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/school/update.proto

package school

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

// 更新学校请求
type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 学校编号
	// @gotags: param:"id"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty" param:"id"`
	// 学校名
	// @gotags: json:"name" validate:"required,min=2,max=32"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" validate:"required,min=2,max=32"`
	// 学校状态
	// @gotags: json:"status" validate:"required_with=0 1"
	Status school.Status `protobuf:"varint,4,opt,name=status,proto3,enum=relo.core.school.Status" json:"status" validate:"required_with=0 1"`
	// 负责人手机号
	// @gotags: json:"principalPhone" validate:"required,mobile"
	PrincipalPhone string `protobuf:"bytes,5,opt,name=principal_phone,json=principalPhone,proto3" json:"principalPhone" validate:"required,mobile"`
	// 负责人手机号
	// @gotags: json:"principalName" validate:"required,min=2,max=32"
	PrincipalName string `protobuf:"bytes,6,opt,name=principal_name,json=principalName,proto3" json:"principalName" validate:"required,min=2,max=32"`
	// 文件编号
	// @gotags: json:"fileId" validate:"omitempty,len=20"
	FileId string `protobuf:"bytes,7,opt,name=file_id,json=fileId,proto3" json:"fileId" validate:"omitempty,len=20"`
	// 详细地址
	// @gotags: json:"address" validate:"omitempty,min=2,max=255"
	Address string `protobuf:"bytes,8,opt,name=address,proto3" json:"address" validate:"omitempty,min=2,max=255"`
	// 详细地址
	// @gotags: json:"areaCode" validate:"required,number,len=6"
	AreaCode string `protobuf:"bytes,9,opt,name=area_code,json=areaCode,proto3" json:"areaCode" validate:"required,number,len=6"`
	// 经度
	// @gotags: json:"longitude" validate:"omitempty"
	Longitude float32 `protobuf:"fixed32,10,opt,name=longitude,proto3" json:"longitude" validate:"omitempty"`
	// 纬度
	// @gotags: json:"latitude" validate:"omitempty"
	Latitude float32 `protobuf:"fixed32,11,opt,name=latitude,proto3" json:"latitude" validate:"omitempty"`
	// 验证码
	// @gotags: json:"verificationCode" validate:"required"
	VerificationCode string `protobuf:"bytes,12,opt,name=verification_code,json=verificationCode,proto3" json:"verificationCode" validate:"required"`
	// 开设课程
	// @gotags: json:"setUpCourses" validate:"required"
	Courses []*school.Course `protobuf:"bytes,13,rep,name=courses,proto3" json:"setUpCourses" validate:"required"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_school_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_school_update_proto_msgTypes[0]
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
	return file_relo_protocol_school_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateReq) GetStatus() school.Status {
	if x != nil {
		return x.Status
	}
	return school.Status_AVAILABLE
}

func (x *UpdateReq) GetPrincipalPhone() string {
	if x != nil {
		return x.PrincipalPhone
	}
	return ""
}

func (x *UpdateReq) GetPrincipalName() string {
	if x != nil {
		return x.PrincipalName
	}
	return ""
}

func (x *UpdateReq) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UpdateReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateReq) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *UpdateReq) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UpdateReq) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UpdateReq) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

func (x *UpdateReq) GetCourses() []*school.Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

// 更新学校状态请求
type UpdateStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 学校编号
	// @gotags: param:"id"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty" param:"id"`
	// 学校状态
	// @gotags: json:"status" validate:"required_with=0 1"
	Status school.Status `protobuf:"varint,3,opt,name=status,proto3,enum=relo.core.school.Status" json:"status" validate:"required_with=0 1"`
	// 密码
	// @gotags: json:"password" validate:"required,printascii,min=8,max=30"
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password" validate:"required,printascii,min=8,max=30"`
}

func (x *UpdateStatusReq) Reset() {
	*x = UpdateStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_school_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusReq) ProtoMessage() {}

func (x *UpdateStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_school_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateStatusReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_school_update_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateStatusReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStatusReq) GetStatus() school.Status {
	if x != nil {
		return x.Status
	}
	return school.Status_AVAILABLE
}

func (x *UpdateStatusReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 更新学校名字和图标请求
type UpdateLogoAndNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 学校编号
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// 学校名
	// @gotags: json:"name" validate:"required,min=2,max=32"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" validate:"required,min=2,max=32"`
	// 文件编号
	// @gotags: json:"fileId" validate:"omitempty,len=20"
	FileId string `protobuf:"bytes,4,opt,name=file_id,json=fileId,proto3" json:"fileId" validate:"omitempty,len=20"`
}

func (x *UpdateLogoAndNameReq) Reset() {
	*x = UpdateLogoAndNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_school_update_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLogoAndNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLogoAndNameReq) ProtoMessage() {}

func (x *UpdateLogoAndNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_school_update_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLogoAndNameReq.ProtoReflect.Descriptor instead.
func (*UpdateLogoAndNameReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_school_update_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateLogoAndNameReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateLogoAndNameReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateLogoAndNameReq) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

var File_relo_protocol_school_update_proto protoreflect.FileDescriptor

var file_relo_protocol_school_update_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x03, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x22, 0x6f, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x53, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x3b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_school_update_proto_rawDescOnce sync.Once
	file_relo_protocol_school_update_proto_rawDescData = file_relo_protocol_school_update_proto_rawDesc
)

func file_relo_protocol_school_update_proto_rawDescGZIP() []byte {
	file_relo_protocol_school_update_proto_rawDescOnce.Do(func() {
		file_relo_protocol_school_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_school_update_proto_rawDescData)
	})
	return file_relo_protocol_school_update_proto_rawDescData
}

var file_relo_protocol_school_update_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_relo_protocol_school_update_proto_goTypes = []interface{}{
	(*UpdateReq)(nil),            // 0: relo.protocol.school.UpdateReq
	(*UpdateStatusReq)(nil),      // 1: relo.protocol.school.UpdateStatusReq
	(*UpdateLogoAndNameReq)(nil), // 2: relo.protocol.school.UpdateLogoAndNameReq
	(school.Status)(0),           // 3: relo.core.school.Status
	(*school.Course)(nil),        // 4: relo.core.school.Course
}
var file_relo_protocol_school_update_proto_depIdxs = []int32{
	3, // 0: relo.protocol.school.UpdateReq.status:type_name -> relo.core.school.Status
	4, // 1: relo.protocol.school.UpdateReq.courses:type_name -> relo.core.school.Course
	3, // 2: relo.protocol.school.UpdateStatusReq.status:type_name -> relo.core.school.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relo_protocol_school_update_proto_init() }
func file_relo_protocol_school_update_proto_init() {
	if File_relo_protocol_school_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_school_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_relo_protocol_school_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatusReq); i {
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
		file_relo_protocol_school_update_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLogoAndNameReq); i {
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
			RawDescriptor: file_relo_protocol_school_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_school_update_proto_goTypes,
		DependencyIndexes: file_relo_protocol_school_update_proto_depIdxs,
		MessageInfos:      file_relo_protocol_school_update_proto_msgTypes,
	}.Build()
	File_relo_protocol_school_update_proto = out.File
	file_relo_protocol_school_update_proto_rawDesc = nil
	file_relo_protocol_school_update_proto_goTypes = nil
	file_relo_protocol_school_update_proto_depIdxs = nil
}
