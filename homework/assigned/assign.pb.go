// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/homework/assigned/assign.proto

package assigned

import (
	homework "github.com/reloclass/core/homework"
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

// 发布作业请求
type AssignReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程编号
	// @gotags: json:"courseId,string" validate:"required"
	CourseId int64 `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"courseId,string" validate:"required"`
	// 作业标题
	// @gotags: json:"title" validate:"required,without_special_symbol,min=2,max=30"
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title" validate:"required,without_special_symbol,min=2,max=30"`
	// 开始时间
	// @gotags: json:"startTime"
	StartTime string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"startTime"`
	// 结束时间
	// @gotags: json:"endTime"
	EndTime string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"endTime"`
	// 评分类型
	// @gotags: json:"gradingType" validate:"required_with=0 1 2 3"
	GradingType homework.GradingType `protobuf:"varint,6,opt,name=grading_type,json=gradingType,proto3,enum=relo.core.homework.GradingType" json:"gradingType" validate:"required_with=0 1 2 3"`
	// 作业内容
	// @gotags: json:"content" validate:"omitempty,max=5000"
	Content string `protobuf:"bytes,7,opt,name=content,proto3" json:"content" validate:"omitempty,max=5000"`
	// 附件
	// @gotags: json:"attachmentIds"
	AttachmentIds []string `protobuf:"bytes,8,rep,name=attachment_ids,json=attachmentIds,proto3" json:"attachmentIds"`
}

func (x *AssignReq) Reset() {
	*x = AssignReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_homework_assigned_assign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignReq) ProtoMessage() {}

func (x *AssignReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_homework_assigned_assign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignReq.ProtoReflect.Descriptor instead.
func (*AssignReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_homework_assigned_assign_proto_rawDescGZIP(), []int{0}
}

func (x *AssignReq) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *AssignReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AssignReq) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *AssignReq) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *AssignReq) GetGradingType() homework.GradingType {
	if x != nil {
		return x.GradingType
	}
	return homework.GradingType_GRADING_NO
}

func (x *AssignReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AssignReq) GetAttachmentIds() []string {
	if x != nil {
		return x.AttachmentIds
	}
	return nil
}

// 发布作业响应
type AssignRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Homework *vo.AssignedHomework `protobuf:"bytes,2,opt,name=homework,proto3" json:"homework,omitempty"`
}

func (x *AssignRsp) Reset() {
	*x = AssignRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_homework_assigned_assign_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRsp) ProtoMessage() {}

func (x *AssignRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_homework_assigned_assign_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRsp.ProtoReflect.Descriptor instead.
func (*AssignRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_homework_assigned_assign_proto_rawDescGZIP(), []int{1}
}

func (x *AssignRsp) GetHomework() *vo.AssignedHomework {
	if x != nil {
		return x.Homework
	}
	return nil
}

var File_relo_protocol_homework_assigned_assign_proto protoreflect.FileDescriptor

var file_relo_protocol_homework_assigned_assign_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x1a,
	0x25, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfd, 0x01, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x67,
	0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73,
	0x22, 0x4b, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x3e, 0x0a,
	0x08, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68,
	0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x3b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_relo_protocol_homework_assigned_assign_proto_rawDescOnce sync.Once
	file_relo_protocol_homework_assigned_assign_proto_rawDescData = file_relo_protocol_homework_assigned_assign_proto_rawDesc
)

func file_relo_protocol_homework_assigned_assign_proto_rawDescGZIP() []byte {
	file_relo_protocol_homework_assigned_assign_proto_rawDescOnce.Do(func() {
		file_relo_protocol_homework_assigned_assign_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_homework_assigned_assign_proto_rawDescData)
	})
	return file_relo_protocol_homework_assigned_assign_proto_rawDescData
}

var file_relo_protocol_homework_assigned_assign_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_homework_assigned_assign_proto_goTypes = []interface{}{
	(*AssignReq)(nil),           // 0: relo.protocol.homework.assigned.AssignReq
	(*AssignRsp)(nil),           // 1: relo.protocol.homework.assigned.AssignRsp
	(homework.GradingType)(0),   // 2: relo.core.homework.GradingType
	(*vo.AssignedHomework)(nil), // 3: relo.protocol.vo.AssignedHomework
}
var file_relo_protocol_homework_assigned_assign_proto_depIdxs = []int32{
	2, // 0: relo.protocol.homework.assigned.AssignReq.grading_type:type_name -> relo.core.homework.GradingType
	3, // 1: relo.protocol.homework.assigned.AssignRsp.homework:type_name -> relo.protocol.vo.AssignedHomework
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_relo_protocol_homework_assigned_assign_proto_init() }
func file_relo_protocol_homework_assigned_assign_proto_init() {
	if File_relo_protocol_homework_assigned_assign_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_homework_assigned_assign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignReq); i {
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
		file_relo_protocol_homework_assigned_assign_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignRsp); i {
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
			RawDescriptor: file_relo_protocol_homework_assigned_assign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_homework_assigned_assign_proto_goTypes,
		DependencyIndexes: file_relo_protocol_homework_assigned_assign_proto_depIdxs,
		MessageInfos:      file_relo_protocol_homework_assigned_assign_proto_msgTypes,
	}.Build()
	File_relo_protocol_homework_assigned_assign_proto = out.File
	file_relo_protocol_homework_assigned_assign_proto_rawDesc = nil
	file_relo_protocol_homework_assigned_assign_proto_goTypes = nil
	file_relo_protocol_homework_assigned_assign_proto_depIdxs = nil
}