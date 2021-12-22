// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/user/add.proto

package user

import (
	vo "github.com/reloclass/protocol/vo"
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
	Users []*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
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

func (x *AddReq) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	// @gotags: json:"phone" validate:"required,printascii,mobile"
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone" validate:"required,printascii,mobile"`
	// 昵称
	// @gotags: json:"nickname"
	NickName string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nickname"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_user_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_relo_protocol_user_add_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

// 批量添加用户响应
type BatchAddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件重复用户
	// @gotags: json:"repeatUsers"
	RepeatUsers []*vo.User `protobuf:"bytes,2,rep,name=repeat_users,json=repeatUsers,proto3" json:"repeatUsers"`
	// 文件信息错误用户
	// @gotags: json:"errUsers"
	ErrUsers []*vo.User `protobuf:"bytes,3,rep,name=err_users,json=errUsers,proto3" json:"errUsers"`
	// 文件中已注册用户
	// @gotags: json:"existUsers"
	ExistUsers []*vo.User `protobuf:"bytes,4,rep,name=exist_users,json=existUsers,proto3" json:"existUsers"`
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

func (x *BatchAddRsp) GetRepeatUsers() []*vo.User {
	if x != nil {
		return x.RepeatUsers
	}
	return nil
}

func (x *BatchAddRsp) GetErrUsers() []*vo.User {
	if x != nil {
		return x.ErrUsers
	}
	return nil
}

func (x *BatchAddRsp) GetExistUsers() []*vo.User {
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
	UserInfo *vo.UserInfo `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo"` // 用户
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

func (x *AddUserRsp) GetUserInfo() *vo.UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_relo_protocol_user_add_proto protoreflect.FileDescriptor

var file_relo_protocol_user_add_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x38, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x39, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x64,
	0x64, 0x52, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x33, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x65, 0x72, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x65, 0x78, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6c, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x6e, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5d, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x65, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x65,
	0x77, 0x41, 0x64, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_relo_protocol_user_add_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_relo_protocol_user_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),      // 0: relo.protocol.user.AddReq
	(*User)(nil),        // 1: relo.protocol.user.User
	(*BatchAddRsp)(nil), // 2: relo.protocol.user.BatchAddRsp
	(*AddUserReq)(nil),  // 3: relo.protocol.user.AddUserReq
	(*AddUserRsp)(nil),  // 4: relo.protocol.user.AddUserRsp
	(*vo.User)(nil),     // 5: relo.protocol.vo.User
	(*vo.UserInfo)(nil), // 6: relo.protocol.vo.UserInfo
}
var file_relo_protocol_user_add_proto_depIdxs = []int32{
	1, // 0: relo.protocol.user.AddReq.users:type_name -> relo.protocol.user.User
	5, // 1: relo.protocol.user.BatchAddRsp.repeat_users:type_name -> relo.protocol.vo.User
	5, // 2: relo.protocol.user.BatchAddRsp.err_users:type_name -> relo.protocol.vo.User
	5, // 3: relo.protocol.user.BatchAddRsp.exist_users:type_name -> relo.protocol.vo.User
	6, // 4: relo.protocol.user.AddUserRsp.userInfo:type_name -> relo.protocol.vo.UserInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_relo_protocol_user_add_proto_init() }
func file_relo_protocol_user_add_proto_init() {
	if File_relo_protocol_user_add_proto != nil {
		return
	}
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
			switch v := v.(*User); i {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relo_protocol_user_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
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
