// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/suspend.proto

package vo

import (
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

// 课节终止信息
type Suspend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 终止信息编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string"`
	// 督课老师编号
	// @gotags: json:"supervisorId,string"
	SupervisorId int64 `protobuf:"varint,3,opt,name=supervisor_id,json=supervisorId,proto3" json:"supervisorId,string"`
	// 督课老师名字
	// @gotags: json:"supervisorName"
	SupervisorName int64 `protobuf:"varint,4,opt,name=supervisor_name,json=supervisorName,proto3" json:"supervisorName"`
	// 中止原因
	// @gotags: json:"suspendReason"
	Reason string `protobuf:"bytes,5,opt,name=reason,proto3" json:"suspendReason"`
	// 中止Unix时间戳
	// @gotags: json:"suspendTimestamp,string"
	Timestamp int64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"suspendTimestamp,string"`
	// 电话
	// @gotags: json:"phone"
	Phone string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone"`
}

func (x *Suspend) Reset() {
	*x = Suspend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_suspend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suspend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suspend) ProtoMessage() {}

func (x *Suspend) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_suspend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suspend.ProtoReflect.Descriptor instead.
func (*Suspend) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_suspend_proto_rawDescGZIP(), []int{0}
}

func (x *Suspend) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Suspend) GetSupervisorId() int64 {
	if x != nil {
		return x.SupervisorId
	}
	return 0
}

func (x *Suspend) GetSupervisorName() int64 {
	if x != nil {
		return x.SupervisorName
	}
	return 0
}

func (x *Suspend) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Suspend) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Suspend) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_relo_protocol_vo_suspend_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_suspend_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_suspend_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_suspend_proto_rawDescData = file_relo_protocol_vo_suspend_proto_rawDesc
)

func file_relo_protocol_vo_suspend_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_suspend_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_suspend_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_suspend_proto_rawDescData)
	})
	return file_relo_protocol_vo_suspend_proto_rawDescData
}

var file_relo_protocol_vo_suspend_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_suspend_proto_goTypes = []interface{}{
	(*Suspend)(nil), // 0: relo.protocol.vo.Suspend
}
var file_relo_protocol_vo_suspend_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_suspend_proto_init() }
func file_relo_protocol_vo_suspend_proto_init() {
	if File_relo_protocol_vo_suspend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_suspend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suspend); i {
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
			RawDescriptor: file_relo_protocol_vo_suspend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_suspend_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_suspend_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_suspend_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_suspend_proto = out.File
	file_relo_protocol_vo_suspend_proto_rawDesc = nil
	file_relo_protocol_vo_suspend_proto_goTypes = nil
	file_relo_protocol_vo_suspend_proto_depIdxs = nil
}
