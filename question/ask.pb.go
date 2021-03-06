// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/question/ask.proto

package question

import (
	question "github.com/reloclass/core/question"
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

// 提问请求
type AskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 问题提问类型
	// @gotags: json:"type" validate:"required_with=0 1"
	Type question.Type `protobuf:"varint,2,opt,name=type,proto3,enum=relo.core.question.Type" json:"type" validate:"required_with=0 1"`
	// 课节编号
	// @gotags: json:"courseTimeId,string" validate:"required"
	LessonId int64 `protobuf:"varint,3,opt,name=lesson_id,json=lessonId,proto3" json:"courseTimeId,string" validate:"required"`
	// 提问内容
	// @gotags: json:"askContent" validate:"required"
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"askContent" validate:"required"`
	// 如果Type为0 表示讲师,助教的编号
	// @gotags: json:"teacherId,string" validate:"omitempty"
	TeacherId int64 `protobuf:"varint,5,opt,name=teacher_id,json=teacherId,proto3" json:"teacherId,string" validate:"omitempty"`
}

func (x *AskReq) Reset() {
	*x = AskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_question_ask_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskReq) ProtoMessage() {}

func (x *AskReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_question_ask_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskReq.ProtoReflect.Descriptor instead.
func (*AskReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_question_ask_proto_rawDescGZIP(), []int{0}
}

func (x *AskReq) GetType() question.Type {
	if x != nil {
		return x.Type
	}
	return question.Type_TEACHER
}

func (x *AskReq) GetLessonId() int64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

func (x *AskReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AskReq) GetTeacherId() int64 {
	if x != nil {
		return x.TeacherId
	}
	return 0
}

var File_relo_protocol_question_ask_proto protoreflect.FileDescriptor

var file_relo_protocol_question_ask_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x06, 0x41, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x3b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_question_ask_proto_rawDescOnce sync.Once
	file_relo_protocol_question_ask_proto_rawDescData = file_relo_protocol_question_ask_proto_rawDesc
)

func file_relo_protocol_question_ask_proto_rawDescGZIP() []byte {
	file_relo_protocol_question_ask_proto_rawDescOnce.Do(func() {
		file_relo_protocol_question_ask_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_question_ask_proto_rawDescData)
	})
	return file_relo_protocol_question_ask_proto_rawDescData
}

var file_relo_protocol_question_ask_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_question_ask_proto_goTypes = []interface{}{
	(*AskReq)(nil),     // 0: relo.protocol.question.AskReq
	(question.Type)(0), // 1: relo.core.question.Type
}
var file_relo_protocol_question_ask_proto_depIdxs = []int32{
	1, // 0: relo.protocol.question.AskReq.type:type_name -> relo.core.question.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relo_protocol_question_ask_proto_init() }
func file_relo_protocol_question_ask_proto_init() {
	if File_relo_protocol_question_ask_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_question_ask_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskReq); i {
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
			RawDescriptor: file_relo_protocol_question_ask_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_question_ask_proto_goTypes,
		DependencyIndexes: file_relo_protocol_question_ask_proto_depIdxs,
		MessageInfos:      file_relo_protocol_question_ask_proto_msgTypes,
	}.Build()
	File_relo_protocol_question_ask_proto = out.File
	file_relo_protocol_question_ask_proto_rawDesc = nil
	file_relo_protocol_question_ask_proto_goTypes = nil
	file_relo_protocol_question_ask_proto_depIdxs = nil
}
