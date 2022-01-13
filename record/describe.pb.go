// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relo/protocol/record/describe.proto

package record

import (
	record "github.com/reloclass/core/record"
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

// 课节录制操作请求
type DescribeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务编号
	// @gotags: json:"taskId" validate:"required"
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"taskId" validate:"required"`
}

func (x *DescribeReq) Reset() {
	*x = DescribeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_record_describe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeReq) ProtoMessage() {}

func (x *DescribeReq) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_record_describe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeReq.ProtoReflect.Descriptor instead.
func (*DescribeReq) Descriptor() ([]byte, []int) {
	return file_relo_protocol_record_describe_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeReq) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

// 课节录制操作响应
type DescribeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误类型，只有请求出错的时候才会返回此字段
	// @gotags: json:"ErrorCode"
	ErrorCode string `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"ErrorCode"`
	// 错误说明，只有请求出错的时候才会返回此字段
	// @gotags: json:"ErrorMessage"
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"ErrorMessage"`
	// 请求ID
	// @gotags: json:"RequestID"
	RequestId string `protobuf:"bytes,4,opt,name=request_id,json=requestId,proto3" json:"RequestID"`
	// 录制任务ID，只有开始录制成功的时候才会返回此字段
	// @gotags: json:"TaskID"
	TaskId string `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3" json:"TaskID"`
	// 录制任务当前状态，可能的取值及其含义如下：
	// normal -- 录制任务创建成功，但录制任务还没
	// 有开始录制
	// recording -- 录制任务正在录制中
	// canceled -- 录制任务已经被取消
	// transcode -- 录制任务正在对录制视频进⾏拼接 转码等操作
	// upload -- 录制任务正在上传录制视频到云存储
	// callback -- 录制任务正在把录制结果通过发起
	// 录制时配置的回调地址进⾏通知
	// finished -- 录制任务的所有操作已经完成
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	// 录制任务创建时间戳 单位秒
	// @gotags: json:"CreateTime"
	CreateTime int64 `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"CreateTime"`
	// 录制任务开始录制时间戳，单位s，只有开始录制后才会返回此字段
	// @gotags: json:"StartTime"
	StartTime int64 `protobuf:"varint,8,opt,name=start_time,json=startTime,proto3" json:"StartTime"`
	// 录制任务被取消的时间戳，单位s，只有发起过取消录制请求才会返回此字段
	// @gotags: json:"CancelTime"
	CancelTime int64 `protobuf:"varint,9,opt,name=cancel_time,json=cancelTime,proto3" json:"CancelTime"`
	// 录制任务结束录制时间戳，单位s，只有录制任务停⽌后才会返回此字段
	// @gotags: json:"StopTime"
	StopTime int64 `protobuf:"varint,10,opt,name=stop_time,json=stopTime,proto3" json:"StopTime"`
	// 录制任务完成录制时间戳，单位s，只有录制任务完成所有录制处理后才会返回此字段
	// @gotags: json:"FinishTime"
	FinishTime int64 `protobuf:"varint,11,opt,name=finish_time,json=finishTime,proto3" json:"FinishTime"`
	// 录制任务结果
	// @gotags: json:"Result"
	Result *record.Result `protobuf:"bytes,12,opt,name=result,proto3" json:"Result"`
}

func (x *DescribeRsp) Reset() {
	*x = DescribeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relo_protocol_record_describe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRsp) ProtoMessage() {}

func (x *DescribeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_relo_protocol_record_describe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRsp.ProtoReflect.Descriptor instead.
func (*DescribeRsp) Descriptor() ([]byte, []int) {
	return file_relo_protocol_record_describe_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeRsp) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *DescribeRsp) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *DescribeRsp) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *DescribeRsp) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *DescribeRsp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DescribeRsp) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *DescribeRsp) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *DescribeRsp) GetCancelTime() int64 {
	if x != nil {
		return x.CancelTime
	}
	return 0
}

func (x *DescribeRsp) GetStopTime() int64 {
	if x != nil {
		return x.StopTime
	}
	return 0
}

func (x *DescribeRsp) GetFinishTime() int64 {
	if x != nil {
		return x.FinishTime
	}
	return 0
}

func (x *DescribeRsp) GetResult() *record.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_relo_protocol_record_describe_proto protoreflect.FileDescriptor

var file_relo_protocol_record_describe_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x65, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1d, 0x72, 0x65, 0x6c,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x22, 0xf2, 0x02, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x3b,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relo_protocol_record_describe_proto_rawDescOnce sync.Once
	file_relo_protocol_record_describe_proto_rawDescData = file_relo_protocol_record_describe_proto_rawDesc
)

func file_relo_protocol_record_describe_proto_rawDescGZIP() []byte {
	file_relo_protocol_record_describe_proto_rawDescOnce.Do(func() {
		file_relo_protocol_record_describe_proto_rawDescData = protoimpl.X.CompressGZIP(file_relo_protocol_record_describe_proto_rawDescData)
	})
	return file_relo_protocol_record_describe_proto_rawDescData
}

var file_relo_protocol_record_describe_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relo_protocol_record_describe_proto_goTypes = []interface{}{
	(*DescribeReq)(nil),   // 0: relo.protocol.record.DescribeReq
	(*DescribeRsp)(nil),   // 1: relo.protocol.record.DescribeRsp
	(*record.Result)(nil), // 2: relo.core.record.Result
}
var file_relo_protocol_record_describe_proto_depIdxs = []int32{
	2, // 0: relo.protocol.record.DescribeRsp.result:type_name -> relo.core.record.Result
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relo_protocol_record_describe_proto_init() }
func file_relo_protocol_record_describe_proto_init() {
	if File_relo_protocol_record_describe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relo_protocol_record_describe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeReq); i {
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
		file_relo_protocol_record_describe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRsp); i {
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
			RawDescriptor: file_relo_protocol_record_describe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relo_protocol_record_describe_proto_goTypes,
		DependencyIndexes: file_relo_protocol_record_describe_proto_depIdxs,
		MessageInfos:      file_relo_protocol_record_describe_proto_msgTypes,
	}.Build()
	File_relo_protocol_record_describe_proto = out.File
	file_relo_protocol_record_describe_proto_rawDesc = nil
	file_relo_protocol_record_describe_proto_goTypes = nil
	file_relo_protocol_record_describe_proto_depIdxs = nil
}