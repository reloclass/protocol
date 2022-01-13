// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/homework/finished/review.proto

package finished

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

// 老师/助教批阅学生作业请求
type ReviewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id,string" validate:"required"
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,string" validate:"required"`
	// 学生作业内容与批阅内容
	// @gotags: json:"content" validate:"max=5000"
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" validate:"max=5000"`
	// 评分
	// @gotags: json:"grading" validate:"omitempty,max=32"
	Grading string `protobuf:"bytes,4,opt,name=grading,proto3" json:"grading" validate:"omitempty,max=32"`
	// 评语
	// @gotags: json:"comment" validate:"omitempty,max=300"
	Comment string `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment" validate:"omitempty,max=300"`
	// 订正状态
	// @gotags: json:"state" validate:"omitempty,required_with=0 1"
	State homework.State `protobuf:"varint,6,opt,name=state,proto3,enum=relo.core.homework.State" json:"state" validate:"omitempty,required_with=0 1"`
}

func (x *ReviewReq) Reset() {
	*x = ReviewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_homework_finished_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewReq) ProtoMessage() {}

func (x *ReviewReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_homework_finished_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewReq.ProtoReflect.Descriptor instead.
func (*ReviewReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_homework_finished_review_proto_rawDescGZIP(), []int{0}
}

func (x *ReviewReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReviewReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReviewReq) GetGrading() string {
	if x != nil {
		return x.Grading
	}
	return ""
}

func (x *ReviewReq) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *ReviewReq) GetState() homework.State {
	if x != nil {
		return x.State
	}
	return homework.State_UNREVIEW
}

var File_relo_protocol_homework_finished_review_proto protoreflect.FileDescriptor

var file_relo_protocol_homework_finished_review_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x1a,
	0x1e, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9a, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x6c,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x3b,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_homework_finished_review_proto_rawDescOnce sync.Once
	file_relo_protocol_homework_finished_review_proto_rawDescData = file_relo_protocol_homework_finished_review_proto_rawDesc
)

func file_relo_protocol_homework_finished_review_proto_rawDescGZIP() []byte {
	file_relo_protocol_homework_finished_review_proto_rawDescOnce.Do(func() {
		file_relo_protocol_homework_finished_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_homework_finished_review_proto_rawDescData)
	})
	return file_relo_protocol_homework_finished_review_proto_rawDescData
}

var file_relo_protocol_homework_finished_review_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relo_protocol_homework_finished_review_proto_goTypes = []interface{}{
	(*ReviewReq)(nil),   // 0: relo.protocol.homework.finished.ReviewReq
	(homework.State)(0), // 1: relo.core.homework.State
}
var file_relo_protocol_homework_finished_review_proto_depIdxs = []int32{
	1, // 0: relo.protocol.homework.finished.ReviewReq.state:type_name -> relo.core.homework.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relo_protocol_homework_finished_review_proto_init() }
func file_relo_protocol_homework_finished_review_proto_init() {
	if File_relo_protocol_homework_finished_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_homework_finished_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewReq); i {
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
			RawDescriptor: file_relo_protocol_homework_finished_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_homework_finished_review_proto_goTypes,
		DependencyIndexes: file_relo_protocol_homework_finished_review_proto_depIdxs,
		MessageInfos:      file_relo_protocol_homework_finished_review_proto_msgTypes,
	}.Build()
	File_relo_protocol_homework_finished_review_proto = out.File
	file_relo_protocol_homework_finished_review_proto_rawDesc = nil
	file_relo_protocol_homework_finished_review_proto_goTypes = nil
	file_relo_protocol_homework_finished_review_proto_depIdxs = nil
}
