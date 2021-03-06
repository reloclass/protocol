// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/record/get.proto

package record

import (
	course "github.com/reloclass/core/course"
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

// 分页获取录课资源请求
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
	// @gotags: default:"created_at" query:"sortField" validate:"oneof=updated_at created_at"
	Sort string `protobuf:"bytes,6,opt,name=sort,proto3" json:"sort,omitempty" default:"created_at" query:"sortField" validate:"oneof=updated_at created_at"`
	// 课程属于学校的编号
	// @gotags: json:"schoolId,string"
	SchoolId int64 `protobuf:"varint,7,opt,name=school_id,json=schoolId,proto3" json:"schoolId,string"`
	// 课程名称关键字
	// @gotags: query:"courseNameKey" json:"courseNameKey"
	CourseNameKey string `protobuf:"bytes,8,opt,name=course_name_key,json=courseNameKey,proto3" json:"courseNameKey" query:"courseNameKey"`
	// 课节名称关键字
	// @gotags: query:"nameKey" json:"nameKey"
	NameKey string `protobuf:"bytes,9,opt,name=name_key,json=nameKey,proto3" json:"nameKey" query:"nameKey"`
	// 课程类型
	// @gotags: query:"courseType" json:"courseType"
	CourseType course.Type `protobuf:"varint,10,opt,name=course_type,json=courseType,proto3,enum=relo.core.course.Type" json:"courseType" query:"courseType"`
	// 上课模式
	// @gotags: query:"classMode" json:"classMode"
	ClassMode course.ClassMode `protobuf:"varint,11,opt,name=class_mode,json=classMode,proto3,enum=relo.core.course.ClassMode" json:"classMode" query:"classMode"`
	// 开始时间条件
	// @gotags: query:"startTime" json:"startTime"
	StartTime string `protobuf:"bytes,12,opt,name=start_time,json=startTime,proto3" json:"startTime" query:"startTime"`
	// 结束时间条件
	// @gotags: query:"endTime" json:"endTime"
	EndTime string `protobuf:"bytes,13,opt,name=end_time,json=endTime,proto3" json:"endTime" query:"endTime"`
}

func (x *GetsByPagingReq) Reset() {
	*x = GetsByPagingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_record_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByPagingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByPagingReq) ProtoMessage() {}

func (x *GetsByPagingReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_record_get_proto_msgTypes[0]
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
	return file_relo_protocol_record_get_proto_rawDescGZIP(), []int{0}
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

func (x *GetsByPagingReq) GetSchoolId() int64 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *GetsByPagingReq) GetCourseNameKey() string {
	if x != nil {
		return x.CourseNameKey
	}
	return ""
}

func (x *GetsByPagingReq) GetNameKey() string {
	if x != nil {
		return x.NameKey
	}
	return ""
}

func (x *GetsByPagingReq) GetCourseType() course.Type {
	if x != nil {
		return x.CourseType
	}
	return course.Type_COURSE_TYPE_UNSPECIFIED
}

func (x *GetsByPagingReq) GetClassMode() course.ClassMode {
	if x != nil {
		return x.ClassMode
	}
	return course.ClassMode_CLASS_MODE_UNSPECIFIED
}

func (x *GetsByPagingReq) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *GetsByPagingReq) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

// 分页获取录课资源响应
type GetsByPagingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	TotalNum int64 `protobuf:"varint,2,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`
	// 数据
	Items []*vo.EsRecord `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetsByPagingRsp) Reset() {
	*x = GetsByPagingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_record_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByPagingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByPagingRsp) ProtoMessage() {}

func (x *GetsByPagingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_record_get_proto_msgTypes[1]
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
	return file_relo_protocol_record_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetsByPagingRsp) GetTotalNum() int64 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *GetsByPagingRsp) GetItems() []*vo.EsRecord {
	if x != nil {
		return x.Items
	}
	return nil
}

// 获取课程录播资源占用空间响应
type GetSizeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程资源占用大小
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GetSizeRsp) Reset() {
	*x = GetSizeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_record_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSizeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSizeRsp) ProtoMessage() {}

func (x *GetSizeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_record_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSizeRsp.ProtoReflect.Descriptor instead.
func (*GetSizeRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_record_get_proto_rawDescGZIP(), []int{2}
}

func (x *GetSizeRsp) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// 获取课节录课资源
type GetByLessonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课节编号
	// @gotags: param:"id" validate:"required"
	LessonId int64 `protobuf:"varint,2,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty" param:"id" validate:"required"`
}

func (x *GetByLessonReq) Reset() {
	*x = GetByLessonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_record_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByLessonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByLessonReq) ProtoMessage() {}

func (x *GetByLessonReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_record_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByLessonReq.ProtoReflect.Descriptor instead.
func (*GetByLessonReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_record_get_proto_rawDescGZIP(), []int{3}
}

func (x *GetByLessonReq) GetLessonId() int64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

var File_relo_protocol_record_get_proto protoreflect.FileDescriptor

var file_relo_protocol_record_get_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1b, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x73, 0x42, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x61, 0x6d, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x45, 0x73, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x2d, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x3b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_relo_protocol_record_get_proto_rawDescOnce sync.Once
	file_relo_protocol_record_get_proto_rawDescData = file_relo_protocol_record_get_proto_rawDesc
)

func file_relo_protocol_record_get_proto_rawDescGZIP() []byte {
	file_relo_protocol_record_get_proto_rawDescOnce.Do(func() {
		file_relo_protocol_record_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_record_get_proto_rawDescData)
	})
	return file_relo_protocol_record_get_proto_rawDescData
}

var file_relo_protocol_record_get_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_relo_protocol_record_get_proto_goTypes = []interface{}{
	(*GetsByPagingReq)(nil), // 0: relo.protocol.record.GetsByPagingReq
	(*GetsByPagingRsp)(nil), // 1: relo.protocol.record.GetsByPagingRsp
	(*GetSizeRsp)(nil),      // 2: relo.protocol.record.GetSizeRsp
	(*GetByLessonReq)(nil),  // 3: relo.protocol.record.GetByLessonReq
	(course.Type)(0),        // 4: relo.core.course.Type
	(course.ClassMode)(0),   // 5: relo.core.course.ClassMode
	(*vo.EsRecord)(nil),     // 6: relo.protocol.vo.EsRecord
}
var file_relo_protocol_record_get_proto_depIdxs = []int32{
	4, // 0: relo.protocol.record.GetsByPagingReq.course_type:type_name -> relo.core.course.Type
	5, // 1: relo.protocol.record.GetsByPagingReq.class_mode:type_name -> relo.core.course.ClassMode
	6, // 2: relo.protocol.record.GetsByPagingRsp.items:type_name -> relo.protocol.vo.EsRecord
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relo_protocol_record_get_proto_init() }
func file_relo_protocol_record_get_proto_init() {
	if File_relo_protocol_record_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_record_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_relo_protocol_record_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_relo_protocol_record_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSizeRsp); i {
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
		file_relo_protocol_record_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByLessonReq); i {
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
			RawDescriptor: file_relo_protocol_record_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_record_get_proto_goTypes,
		DependencyIndexes: file_relo_protocol_record_get_proto_depIdxs,
		MessageInfos:      file_relo_protocol_record_get_proto_msgTypes,
	}.Build()
	File_relo_protocol_record_get_proto = out.File
	file_relo_protocol_record_get_proto_rawDesc = nil
	file_relo_protocol_record_get_proto_goTypes = nil
	file_relo_protocol_record_get_proto_depIdxs = nil
}
