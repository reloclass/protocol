// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/vo/finished_homework.proto

package vo

import (
	homework "github.com/reloclass/core/homework"
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

// 学生提交的作业
type FinishedHomework struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string"`
	// 学生编号
	// @gotags: json:"studentId,string"
	StudentId int64 `protobuf:"varint,3,opt,name=student_id,json=studentId,proto3" json:"studentId,string"`
	// 作业编号
	// @gotags: json:"homeworkId,string"
	AssignedId int64 `protobuf:"varint,4,opt,name=assigned_id,json=assignedId,proto3" json:"homeworkId,string"`
	// 提交状态
	// @gotags: json:"submittedState"
	SubmittedState homework.SubmittedState `protobuf:"varint,5,opt,name=submitted_state,json=submittedState,proto3,enum=relo.core.homework.SubmittedState" json:"submittedState"`
	// 学生作业提交时间
	// @gotags: json:"submissionTime"
	SubmittedTime string `protobuf:"bytes,6,opt,name=submitted_time,json=submittedTime,proto3" json:"submissionTime"`
	// 作业内容
	Content string `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
	// 评分
	Grading string `protobuf:"bytes,20,opt,name=grading,proto3" json:"grading,omitempty"`
	// 创建时间
	// @gotags: json:"createdAt"
	Created string `protobuf:"bytes,11,opt,name=created,proto3" json:"createdAt"`
	// 评语
	Comment string `protobuf:"bytes,12,opt,name=comment,proto3" json:"comment,omitempty"`
	// 作业状态
	State homework.State `protobuf:"varint,13,opt,name=state,proto3,enum=relo.core.homework.State" json:"state,omitempty"`
	// 学生作业附件
	// @gotags: json:"studentHomeworkAttachmentList"
	Attachments *FinishedHomeworkAttachment `protobuf:"bytes,14,opt,name=attachments,proto3" json:"studentHomeworkAttachmentList"`
	// 学生信息
	// @gotags: json:"studentInfo"
	Student *UserDetail `protobuf:"bytes,15,opt,name=student,proto3" json:"studentInfo"`
}

func (x *FinishedHomework) Reset() {
	*x = FinishedHomework{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_vo_finished_homework_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishedHomework) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishedHomework) ProtoMessage() {}

func (x *FinishedHomework) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_vo_finished_homework_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishedHomework.ProtoReflect.Descriptor instead.
func (*FinishedHomework) Descriptor() ([]byte, []int) {
	return file_relo_protocol_vo_finished_homework_proto_rawDescGZIP(), []int{0}
}

func (x *FinishedHomework) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FinishedHomework) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *FinishedHomework) GetAssignedId() int64 {
	if x != nil {
		return x.AssignedId
	}
	return 0
}

func (x *FinishedHomework) GetSubmittedState() homework.SubmittedState {
	if x != nil {
		return x.SubmittedState
	}
	return homework.SubmittedState_SUBMITTED_NOT
}

func (x *FinishedHomework) GetSubmittedTime() string {
	if x != nil {
		return x.SubmittedTime
	}
	return ""
}

func (x *FinishedHomework) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *FinishedHomework) GetGrading() string {
	if x != nil {
		return x.Grading
	}
	return ""
}

func (x *FinishedHomework) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *FinishedHomework) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *FinishedHomework) GetState() homework.State {
	if x != nil {
		return x.State
	}
	return homework.State_UNREVIEW
}

func (x *FinishedHomework) GetAttachments() *FinishedHomeworkAttachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *FinishedHomework) GetStudent() *UserDetail {
	if x != nil {
		return x.Student
	}
	return nil
}

var File_relo_protocol_vo_finished_homework_proto protoreflect.FileDescriptor

var file_relo_protocol_vo_finished_homework_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x6c, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x1e, 0x72, 0x65,
	0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x72, 0x65,
	0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x72, 0x65, 0x6c,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf7, 0x03, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x48, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_vo_finished_homework_proto_rawDescOnce sync.Once
	file_relo_protocol_vo_finished_homework_proto_rawDescData = file_relo_protocol_vo_finished_homework_proto_rawDesc
)

func file_relo_protocol_vo_finished_homework_proto_rawDescGZIP() []byte {
	file_relo_protocol_vo_finished_homework_proto_rawDescOnce.Do(func() {
		file_relo_protocol_vo_finished_homework_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_vo_finished_homework_proto_rawDescData)
	})
	return file_relo_protocol_vo_finished_homework_proto_rawDescData
}

var file_relo_protocol_vo_finished_homework_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_vo_finished_homework_proto_goTypes = []interface{}{
	(*FinishedHomework)(nil),           // 0: relo.protocol.vo.FinishedHomework
	(homework.SubmittedState)(0),       // 1: relo.core.homework.SubmittedState
	(homework.State)(0),                // 2: relo.core.homework.State
	(*FinishedHomeworkAttachment)(nil), // 3: relo.protocol.vo.FinishedHomeworkAttachment
	(*UserDetail)(nil),                 // 4: relo.protocol.vo.UserDetail
}
var file_relo_protocol_vo_finished_homework_proto_depIdxs = []int32{
	1, // 0: relo.protocol.vo.FinishedHomework.submitted_state:type_name -> relo.core.homework.SubmittedState
	2, // 1: relo.protocol.vo.FinishedHomework.state:type_name -> relo.core.homework.State
	3, // 2: relo.protocol.vo.FinishedHomework.attachments:type_name -> relo.protocol.vo.FinishedHomeworkAttachment
	4, // 3: relo.protocol.vo.FinishedHomework.student:type_name -> relo.protocol.vo.UserDetail
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_relo_protocol_vo_finished_homework_proto_init() }
func file_relo_protocol_vo_finished_homework_proto_init() {
	if File_relo_protocol_vo_finished_homework_proto != nil {
		return
	}
	file_relo_protocol_vo_finished_homework_attachment_proto_init()
	file_relo_protocol_vo_user_detail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_vo_finished_homework_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishedHomework); i {
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
			RawDescriptor: file_relo_protocol_vo_finished_homework_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_vo_finished_homework_proto_goTypes,
		DependencyIndexes: file_relo_protocol_vo_finished_homework_proto_depIdxs,
		MessageInfos:      file_relo_protocol_vo_finished_homework_proto_msgTypes,
	}.Build()
	File_relo_protocol_vo_finished_homework_proto = out.File
	file_relo_protocol_vo_finished_homework_proto_rawDesc = nil
	file_relo_protocol_vo_finished_homework_proto_goTypes = nil
	file_relo_protocol_vo_finished_homework_proto_depIdxs = nil
}
