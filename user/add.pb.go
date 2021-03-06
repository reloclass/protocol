// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/user/add.proto

package user

import (
	user "github.com/reloclass/core/user"
	rpc "github.com/reloclass/protocol/user/rpc"
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

// 添加用户请求
type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户列表
	// @gotags: json:"users,omitempty"
	Users []*user.Base `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	// 代注册
	// @gotags:  json:"register" validate:"omitempty,required_with= 0 1"
	Register user.RegistrationType `protobuf:"varint,3,opt,name=register,proto3,enum=relo.core.user.RegistrationType" json:"register" validate:"omitempty,required_with= 0 1"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetUsers() []*user.Base {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *AddReq) GetRegister() user.RegistrationType {
	if x != nil {
		return x.Register
	}
	return user.RegistrationType_YES
}

// 添加学校用户响应
type AddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 未注册用户
	// @gotags: json:"unregisterUsers"
	UnregisterUsers []*user.Base `protobuf:"bytes,2,rep,name=unregister_users,json=unregisterUsers,proto3" json:"unregisterUsers"`
	// 已存在的用户
	// @gotags: json:"existUsers"
	ExistUsers []*user.Base `protobuf:"bytes,3,rep,name=exist_users,json=existUsers,proto3" json:"existUsers"`
}

func (x *AddRsp) Reset() {
	*x = AddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRsp) ProtoMessage() {}

func (x *AddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRsp.ProtoReflect.Descriptor instead.
func (*AddRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddRsp) GetUnregisterUsers() []*user.Base {
	if x != nil {
		return x.UnregisterUsers
	}
	return nil
}

func (x *AddRsp) GetExistUsers() []*user.Base {
	if x != nil {
		return x.ExistUsers
	}
	return nil
}

// 批量添加用户响应
type BatchAddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件重复用户
	// @gotags: json:"unregisterUsers"
	UnregisterUsers []*user.Base `protobuf:"bytes,2,rep,name=unregister_users,json=unregisterUsers,proto3" json:"unregisterUsers"`
	// 文件重复用户
	// @gotags: json:"repeatUsers"
	RepeatUsers []*user.Base `protobuf:"bytes,3,rep,name=repeat_users,json=repeatUsers,proto3" json:"repeatUsers"`
	// 文件信息错误用户
	// @gotags: json:"errUsers"
	ErrUsers []*user.Base `protobuf:"bytes,4,rep,name=err_users,json=errUsers,proto3" json:"errUsers"`
	// 文件中已注册用户
	// @gotags: json:"existUsers"
	ExistUsers []*user.Base `protobuf:"bytes,5,rep,name=exist_users,json=existUsers,proto3" json:"existUsers"`
}

func (x *BatchAddRsp) Reset() {
	*x = BatchAddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_add_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchAddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchAddRsp) ProtoMessage() {}

func (x *BatchAddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_add_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchAddRsp.ProtoReflect.Descriptor instead.
func (*BatchAddRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_add_proto_rawDescGZIP(), []int{2}
}

func (x *BatchAddRsp) GetUnregisterUsers() []*user.Base {
	if x != nil {
		return x.UnregisterUsers
	}
	return nil
}

func (x *BatchAddRsp) GetRepeatUsers() []*user.Base {
	if x != nil {
		return x.RepeatUsers
	}
	return nil
}

func (x *BatchAddRsp) GetErrUsers() []*user.Base {
	if x != nil {
		return x.ErrUsers
	}
	return nil
}

func (x *BatchAddRsp) GetExistUsers() []*user.Base {
	if x != nil {
		return x.ExistUsers
	}
	return nil
}

type AddUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"phone"
	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone"` // 手机号
	// @gotags: json:"nickname"
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"` // 昵称
	// @gotags: json:"password" validate:"required,printascii,min=8,max=30"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" validate:"required,printascii,min=8,max=30"` // 密码
	// @gotags:  json:"code"   validate:"omitempty"
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code" validate:"omitempty"` // 验证码
}

func (x *AddUserReq) Reset() {
	*x = AddUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_add_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserReq) ProtoMessage() {}

func (x *AddUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_add_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserReq.ProtoReflect.Descriptor instead.
func (*AddUserReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_add_proto_rawDescGZIP(), []int{3}
}

func (x *AddUserReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddUserReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AddUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddUserReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AddUserRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"newAdd"
	NewAdd bool `protobuf:"varint,1,opt,name=new_add,json=newAdd,proto3" json:"newAdd"` // 是否新增
	// @gotags: json:"userInfo"
	User *rpc.User `protobuf:"bytes,2,opt,name=user,proto3" json:"userInfo"` // 用户
}

func (x *AddUserRsp) Reset() {
	*x = AddUserRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_add_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRsp) ProtoMessage() {}

func (x *AddUserRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_add_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRsp.ProtoReflect.Descriptor instead.
func (*AddUserRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_add_proto_rawDescGZIP(), []int{4}
}

func (x *AddUserRsp) GetNewAdd() bool {
	if x != nil {
		return x.NewAdd
	}
	return false
}

func (x *AddUserRsp) GetUser() *rpc.User {
	if x != nil {
		return x.User
	}
	return nil
}

// 添加管理员请求
type AddConsoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	// @gotags: json:"phone" validate:"required,mobile"
	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone" validate:"required,mobile"`
	// 名称
	// @gotags: json:"name" validate:"required,min=2,max=32"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" validate:"required,min=2,max=32"`
	// 角色id
	// @gotags: json:"roleIds" validate:"required,min=1"
	RoleIds []*RoleId `protobuf:"bytes,3,rep,name=role_ids,json=roleIds,proto3" json:"roleIds" validate:"required,min=1"`
}

func (x *AddConsoleReq) Reset() {
	*x = AddConsoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_add_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddConsoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddConsoleReq) ProtoMessage() {}

func (x *AddConsoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_user_add_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddConsoleReq.ProtoReflect.Descriptor instead.
func (*AddConsoleReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_add_proto_rawDescGZIP(), []int{5}
}

func (x *AddConsoleReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddConsoleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddConsoleReq) GetRoleIds() []*RoleId {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

var File_relo_protocol_user_add_proto protoreflect.FileDescriptor

var file_relo_protocol_user_add_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x2a, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x80, 0x01, 0x0a, 0x06,
	0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x10, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x0f, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72,
	0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0xf1,
	0x01, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x12, 0x3f,
	0x0a, 0x10, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x0f,
	0x75, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x37, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65,
	0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x08, 0x65, 0x72, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x6e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x57, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x73, 0x70,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x0d, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_user_add_proto_rawDescOnce sync.Once
	file_relo_protocol_user_add_proto_rawDescData = file_relo_protocol_user_add_proto_rawDesc
)

func file_relo_protocol_user_add_proto_rawDescGZIP() []byte {
	file_relo_protocol_user_add_proto_rawDescOnce.Do(func() {
		file_relo_protocol_user_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_user_add_proto_rawDescData)
	})
	return file_relo_protocol_user_add_proto_rawDescData
}

var file_relo_protocol_user_add_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_relo_protocol_user_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),             // 0: relo.protocol.user.AddReq
	(*AddRsp)(nil),             // 1: relo.protocol.user.AddRsp
	(*BatchAddRsp)(nil),        // 2: relo.protocol.user.BatchAddRsp
	(*AddUserReq)(nil),         // 3: relo.protocol.user.AddUserReq
	(*AddUserRsp)(nil),         // 4: relo.protocol.user.AddUserRsp
	(*AddConsoleReq)(nil),      // 5: relo.protocol.user.AddConsoleReq
	(*user.Base)(nil),          // 6: relo.core.user.Base
	(user.RegistrationType)(0), // 7: relo.core.user.RegistrationType
	(*rpc.User)(nil),           // 8: relo.protocol.user.rpc.User
	(*RoleId)(nil),             // 9: relo.protocol.user.RoleId
}
var file_relo_protocol_user_add_proto_depIdxs = []int32{
	6,  // 0: relo.protocol.user.AddReq.users:type_name -> relo.core.user.Base
	7,  // 1: relo.protocol.user.AddReq.register:type_name -> relo.core.user.RegistrationType
	6,  // 2: relo.protocol.user.AddRsp.unregister_users:type_name -> relo.core.user.Base
	6,  // 3: relo.protocol.user.AddRsp.exist_users:type_name -> relo.core.user.Base
	6,  // 4: relo.protocol.user.BatchAddRsp.unregister_users:type_name -> relo.core.user.Base
	6,  // 5: relo.protocol.user.BatchAddRsp.repeat_users:type_name -> relo.core.user.Base
	6,  // 6: relo.protocol.user.BatchAddRsp.err_users:type_name -> relo.core.user.Base
	6,  // 7: relo.protocol.user.BatchAddRsp.exist_users:type_name -> relo.core.user.Base
	8,  // 8: relo.protocol.user.AddUserRsp.user:type_name -> relo.protocol.user.rpc.User
	9,  // 9: relo.protocol.user.AddConsoleReq.role_ids:type_name -> relo.protocol.user.RoleId
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_relo_protocol_user_add_proto_init() }
func file_relo_protocol_user_add_proto_init() {
	if File_relo_protocol_user_add_proto != nil {
		return
	}
	file_relo_protocol_user_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_user_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_relo_protocol_user_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRsp); i {
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
		file_relo_protocol_user_add_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchAddRsp); i {
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
		file_relo_protocol_user_add_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserReq); i {
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
		file_relo_protocol_user_add_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRsp); i {
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
		file_relo_protocol_user_add_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddConsoleReq); i {
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
			RawDescriptor: file_relo_protocol_user_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_user_add_proto_goTypes,
		DependencyIndexes: file_relo_protocol_user_add_proto_depIdxs,
		MessageInfos:      file_relo_protocol_user_add_proto_msgTypes,
	}.Build()
	File_relo_protocol_user_add_proto = out.File
	file_relo_protocol_user_add_proto_rawDesc = nil
	file_relo_protocol_user_add_proto_goTypes = nil
	file_relo_protocol_user_add_proto_depIdxs = nil
}
