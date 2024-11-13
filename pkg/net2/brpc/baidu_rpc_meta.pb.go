// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: baidu_rpc_meta.proto

package brpc

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

type RpcMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request        *RpcRequestMeta  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response       *RpcResponseMeta `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	CompressType   int32            `protobuf:"varint,3,opt,name=compress_type,json=compressType,proto3" json:"compress_type,omitempty"`
	CorrelationId  int64            `protobuf:"varint,4,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	AttachmentSize int32            `protobuf:"varint,5,opt,name=attachment_size,json=attachmentSize,proto3" json:"attachment_size,omitempty"`
	MtprotoMeta    *MTProtoMeta     `protobuf:"bytes,9,opt,name=mtproto_meta,json=mtprotoMeta,proto3" json:"mtproto_meta,omitempty"`
}

func (x *RpcMeta) Reset() {
	*x = RpcMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baidu_rpc_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcMeta) ProtoMessage() {}

func (x *RpcMeta) ProtoReflect() protoreflect.Message {
	mi := &file_baidu_rpc_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcMeta.ProtoReflect.Descriptor instead.
func (*RpcMeta) Descriptor() ([]byte, []int) {
	return file_baidu_rpc_meta_proto_rawDescGZIP(), []int{0}
}

func (x *RpcMeta) GetRequest() *RpcRequestMeta {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *RpcMeta) GetResponse() *RpcResponseMeta {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *RpcMeta) GetCompressType() int32 {
	if x != nil {
		return x.CompressType
	}
	return 0
}

func (x *RpcMeta) GetCorrelationId() int64 {
	if x != nil {
		return x.CorrelationId
	}
	return 0
}

func (x *RpcMeta) GetAttachmentSize() int32 {
	if x != nil {
		return x.AttachmentSize
	}
	return 0
}

func (x *RpcMeta) GetMtprotoMeta() *MTProtoMeta {
	if x != nil {
		return x.MtprotoMeta
	}
	return nil
}

type RpcRequestMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName  string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	MethodName   string `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	LogId        int64  `protobuf:"varint,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	TraceId      int64  `protobuf:"varint,4,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId       int64  `protobuf:"varint,5,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	ParentSpanId int64  `protobuf:"varint,6,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
}

func (x *RpcRequestMeta) Reset() {
	*x = RpcRequestMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baidu_rpc_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcRequestMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcRequestMeta) ProtoMessage() {}

func (x *RpcRequestMeta) ProtoReflect() protoreflect.Message {
	mi := &file_baidu_rpc_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcRequestMeta.ProtoReflect.Descriptor instead.
func (*RpcRequestMeta) Descriptor() ([]byte, []int) {
	return file_baidu_rpc_meta_proto_rawDescGZIP(), []int{1}
}

func (x *RpcRequestMeta) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RpcRequestMeta) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *RpcRequestMeta) GetLogId() int64 {
	if x != nil {
		return x.LogId
	}
	return 0
}

func (x *RpcRequestMeta) GetTraceId() int64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (x *RpcRequestMeta) GetSpanId() int64 {
	if x != nil {
		return x.SpanId
	}
	return 0
}

func (x *RpcRequestMeta) GetParentSpanId() int64 {
	if x != nil {
		return x.ParentSpanId
	}
	return 0
}

type RpcResponseMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorText string `protobuf:"bytes,2,opt,name=error_text,json=errorText,proto3" json:"error_text,omitempty"`
}

func (x *RpcResponseMeta) Reset() {
	*x = RpcResponseMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baidu_rpc_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcResponseMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcResponseMeta) ProtoMessage() {}

func (x *RpcResponseMeta) ProtoReflect() protoreflect.Message {
	mi := &file_baidu_rpc_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcResponseMeta.ProtoReflect.Descriptor instead.
func (*RpcResponseMeta) Descriptor() ([]byte, []int) {
	return file_baidu_rpc_meta_proto_rawDescGZIP(), []int{2}
}

func (x *RpcResponseMeta) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *RpcResponseMeta) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

var File_baidu_rpc_meta_proto protoreflect.FileDescriptor

var file_baidu_rpc_meta_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x61, 0x69, 0x64, 0x75, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x72, 0x70, 0x63, 0x1a, 0x12, 0x6d, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x02, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x72, 0x70,
	0x63, 0x2e, 0x4d, 0x54, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0b, 0x6d,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x22, 0xc5, 0x01, 0x0a, 0x0e, 0x52,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e,
	0x49, 0x64, 0x22, 0x4f, 0x0a, 0x0f, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54,
	0x65, 0x78, 0x74, 0x42, 0x53, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x0d, 0x42, 0x61, 0x69, 0x64, 0x75, 0x52, 0x70, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x67,
	0x72, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e,
	0x65, 0x74, 0x32, 0x2f, 0x62, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_baidu_rpc_meta_proto_rawDescOnce sync.Once
	file_baidu_rpc_meta_proto_rawDescData = file_baidu_rpc_meta_proto_rawDesc
)

func file_baidu_rpc_meta_proto_rawDescGZIP() []byte {
	file_baidu_rpc_meta_proto_rawDescOnce.Do(func() {
		file_baidu_rpc_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_baidu_rpc_meta_proto_rawDescData)
	})
	return file_baidu_rpc_meta_proto_rawDescData
}

var file_baidu_rpc_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_baidu_rpc_meta_proto_goTypes = []any{
	(*RpcMeta)(nil),         // 0: brpc.RpcMeta
	(*RpcRequestMeta)(nil),  // 1: brpc.RpcRequestMeta
	(*RpcResponseMeta)(nil), // 2: brpc.RpcResponseMeta
	(*MTProtoMeta)(nil),     // 3: brpc.MTProtoMeta
}
var file_baidu_rpc_meta_proto_depIdxs = []int32{
	1, // 0: brpc.RpcMeta.request:type_name -> brpc.RpcRequestMeta
	2, // 1: brpc.RpcMeta.response:type_name -> brpc.RpcResponseMeta
	3, // 2: brpc.RpcMeta.mtproto_meta:type_name -> brpc.MTProtoMeta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_baidu_rpc_meta_proto_init() }
func file_baidu_rpc_meta_proto_init() {
	if File_baidu_rpc_meta_proto != nil {
		return
	}
	file_mtproto_meta_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_baidu_rpc_meta_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RpcMeta); i {
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
		file_baidu_rpc_meta_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RpcRequestMeta); i {
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
		file_baidu_rpc_meta_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RpcResponseMeta); i {
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
			RawDescriptor: file_baidu_rpc_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_baidu_rpc_meta_proto_goTypes,
		DependencyIndexes: file_baidu_rpc_meta_proto_depIdxs,
		MessageInfos:      file_baidu_rpc_meta_proto_msgTypes,
	}.Build()
	File_baidu_rpc_meta_proto = out.File
	file_baidu_rpc_meta_proto_rawDesc = nil
	file_baidu_rpc_meta_proto_goTypes = nil
	file_baidu_rpc_meta_proto_depIdxs = nil
}