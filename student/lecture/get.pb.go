// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/student/lecture/get.proto

package lecture

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

// 分页获取学生录播课请求
type GetsByPagingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前页
	// @gotags: default:"1" query:"page"
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" default:"1" query:"page"`
	// 每页大小
	// @gotags: default:"20" query:"perPage"
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty" default:"20" query:"perPage"`
	// 查询关键字
	// @gotags: query:"keyword"
	Keyword string `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty" query:"keyword"`
	// 排序规则
	// @gotags: default:"desc" query:"sortOrder" validate:"oneof=desc DESC asc ASC"
	Order string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty" default:"desc" query:"sortOrder" validate:"oneof=desc DESC asc ASC"`
	// 排序字段
	// @gotags: default:"student_id" query:"sortField"
	Sort string `protobuf:"bytes,6,opt,name=sort,proto3" json:"sort,omitempty" default:"student_id" query:"sortField"`
}

func (x *GetsByPagingReq) Reset() {
	*x = GetsByPagingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_student_lecture_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByPagingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByPagingReq) ProtoMessage() {}

func (x *GetsByPagingReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_student_lecture_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByPagingReq.ProtoReflect.Descriptor instead.
func (*GetsByPagingReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_student_lecture_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetsByPagingReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetsByPagingReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetsByPagingReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *GetsByPagingReq) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *GetsByPagingReq) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

// 分页获取学生录播课响应
type GetsByPagingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	// @gotags: json:"totalNum"
	TotalNum int64 `protobuf:"varint,2,opt,name=total_num,json=totalNum,proto3" json:"totalNum"`
	// 数据
	Items []*vo.StudentLecture `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetsByPagingRsp) Reset() {
	*x = GetsByPagingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_student_lecture_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByPagingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByPagingRsp) ProtoMessage() {}

func (x *GetsByPagingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_student_lecture_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByPagingRsp.ProtoReflect.Descriptor instead.
func (*GetsByPagingRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_student_lecture_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetsByPagingRsp) GetTotalNum() int64 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *GetsByPagingRsp) GetItems() []*vo.StudentLecture {
	if x != nil {
		return x.Items
	}
	return nil
}

// 获取学生录播课详情请求
type GetDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程编号
	// @gotags: param:"id"
	CourseId int64 `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty" param:"id"`
}

func (x *GetDetailReq) Reset() {
	*x = GetDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_student_lecture_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailReq) ProtoMessage() {}

func (x *GetDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_student_lecture_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailReq.ProtoReflect.Descriptor instead.
func (*GetDetailReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_student_lecture_get_proto_rawDescGZIP(), []int{2}
}

func (x *GetDetailReq) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

// 获取学生录播课详情响应
type GetDetailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程信息
	Course *vo.Course `protobuf:"bytes,2,opt,name=course,proto3" json:"course,omitempty"`
	// 章节列表
	Chapters []*vo.LectureChapter `protobuf:"bytes,3,rep,name=chapters,proto3" json:"chapters,omitempty"`
	// 完成情况
	Finished []*vo.FinishedLectureSection `protobuf:"bytes,4,rep,name=finished,proto3" json:"finished,omitempty"`
}

func (x *GetDetailRsp) Reset() {
	*x = GetDetailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_student_lecture_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailRsp) ProtoMessage() {}

func (x *GetDetailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_student_lecture_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailRsp.ProtoReflect.Descriptor instead.
func (*GetDetailRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_student_lecture_get_proto_rawDescGZIP(), []int{3}
}

func (x *GetDetailRsp) GetCourse() *vo.Course {
	if x != nil {
		return x.Course
	}
	return nil
}

func (x *GetDetailRsp) GetChapters() []*vo.LectureChapter {
	if x != nil {
		return x.Chapters
	}
	return nil
}

func (x *GetDetailRsp) GetFinished() []*vo.FinishedLectureSection {
	if x != nil {
		return x.Finished
	}
	return nil
}

var File_relo_protocol_student_lecture_get_proto protoreflect.FileDescriptor

var file_relo_protocol_student_lecture_get_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x26, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x6f, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x73,
	0x42, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x66, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x2b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0xc4, 0x01, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x30, 0x0a,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12,
	0x3c, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x44, 0x0a,
	0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x3b, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_student_lecture_get_proto_rawDescOnce sync.Once
	file_relo_protocol_student_lecture_get_proto_rawDescData = file_relo_protocol_student_lecture_get_proto_rawDesc
)

func file_relo_protocol_student_lecture_get_proto_rawDescGZIP() []byte {
	file_relo_protocol_student_lecture_get_proto_rawDescOnce.Do(func() {
		file_relo_protocol_student_lecture_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_student_lecture_get_proto_rawDescData)
	})
	return file_relo_protocol_student_lecture_get_proto_rawDescData
}

var file_relo_protocol_student_lecture_get_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_relo_protocol_student_lecture_get_proto_goTypes = []interface{}{
	(*GetsByPagingReq)(nil),           // 0: relo.protocol.student.lecture.GetsByPagingReq
	(*GetsByPagingRsp)(nil),           // 1: relo.protocol.student.lecture.GetsByPagingRsp
	(*GetDetailReq)(nil),              // 2: relo.protocol.student.lecture.GetDetailReq
	(*GetDetailRsp)(nil),              // 3: relo.protocol.student.lecture.GetDetailRsp
	(*vo.StudentLecture)(nil),         // 4: relo.protocol.vo.StudentLecture
	(*vo.Course)(nil),                 // 5: relo.protocol.vo.Course
	(*vo.LectureChapter)(nil),         // 6: relo.protocol.vo.LectureChapter
	(*vo.FinishedLectureSection)(nil), // 7: relo.protocol.vo.FinishedLectureSection
}
var file_relo_protocol_student_lecture_get_proto_depIdxs = []int32{
	4, // 0: relo.protocol.student.lecture.GetsByPagingRsp.items:type_name -> relo.protocol.vo.StudentLecture
	5, // 1: relo.protocol.student.lecture.GetDetailRsp.course:type_name -> relo.protocol.vo.Course
	6, // 2: relo.protocol.student.lecture.GetDetailRsp.chapters:type_name -> relo.protocol.vo.LectureChapter
	7, // 3: relo.protocol.student.lecture.GetDetailRsp.finished:type_name -> relo.protocol.vo.FinishedLectureSection
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_relo_protocol_student_lecture_get_proto_init() }
func file_relo_protocol_student_lecture_get_proto_init() {
	if File_relo_protocol_student_lecture_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_student_lecture_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByPagingReq); i {
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
		file_relo_protocol_student_lecture_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByPagingRsp); i {
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
		file_relo_protocol_student_lecture_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailReq); i {
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
		file_relo_protocol_student_lecture_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailRsp); i {
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
			RawDescriptor: file_relo_protocol_student_lecture_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_student_lecture_get_proto_goTypes,
		DependencyIndexes: file_relo_protocol_student_lecture_get_proto_depIdxs,
		MessageInfos:      file_relo_protocol_student_lecture_get_proto_msgTypes,
	}.Build()
	File_relo_protocol_student_lecture_get_proto = out.File
	file_relo_protocol_student_lecture_get_proto_rawDesc = nil
	file_relo_protocol_student_lecture_get_proto_goTypes = nil
	file_relo_protocol_student_lecture_get_proto_depIdxs = nil
}
