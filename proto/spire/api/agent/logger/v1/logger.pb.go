// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/api/agent/logger/v1/logger.proto

package loggerv1

import (
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdjustLoggerRequest_LogAdjustment int32

const (
	AdjustLoggerRequest_PANIC_LEVEL   AdjustLoggerRequest_LogAdjustment = 0
	AdjustLoggerRequest_FATAL_LEVEL   AdjustLoggerRequest_LogAdjustment = 1
	AdjustLoggerRequest_ERROR_LEVEL   AdjustLoggerRequest_LogAdjustment = 2
	AdjustLoggerRequest_WARN_LEVEL    AdjustLoggerRequest_LogAdjustment = 3
	AdjustLoggerRequest_INFO_LEVEL    AdjustLoggerRequest_LogAdjustment = 4
	AdjustLoggerRequest_DEBUG_LEVEL   AdjustLoggerRequest_LogAdjustment = 5
	AdjustLoggerRequest_TRACE_LEVEL   AdjustLoggerRequest_LogAdjustment = 6
	AdjustLoggerRequest_DEFAULT_LEVEL AdjustLoggerRequest_LogAdjustment = 7
)

// Enum value maps for AdjustLoggerRequest_LogAdjustment.
var (
	AdjustLoggerRequest_LogAdjustment_name = map[int32]string{
		0: "PANIC_LEVEL",
		1: "FATAL_LEVEL",
		2: "ERROR_LEVEL",
		3: "WARN_LEVEL",
		4: "INFO_LEVEL",
		5: "DEBUG_LEVEL",
		6: "TRACE_LEVEL",
		7: "DEFAULT_LEVEL",
	}
	AdjustLoggerRequest_LogAdjustment_value = map[string]int32{
		"PANIC_LEVEL":   0,
		"FATAL_LEVEL":   1,
		"ERROR_LEVEL":   2,
		"WARN_LEVEL":    3,
		"INFO_LEVEL":    4,
		"DEBUG_LEVEL":   5,
		"TRACE_LEVEL":   6,
		"DEFAULT_LEVEL": 7,
	}
)

func (x AdjustLoggerRequest_LogAdjustment) Enum() *AdjustLoggerRequest_LogAdjustment {
	p := new(AdjustLoggerRequest_LogAdjustment)
	*p = x
	return p
}

func (x AdjustLoggerRequest_LogAdjustment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdjustLoggerRequest_LogAdjustment) Descriptor() protoreflect.EnumDescriptor {
	return file_spire_api_agent_logger_v1_logger_proto_enumTypes[0].Descriptor()
}

func (AdjustLoggerRequest_LogAdjustment) Type() protoreflect.EnumType {
	return &file_spire_api_agent_logger_v1_logger_proto_enumTypes[0]
}

func (x AdjustLoggerRequest_LogAdjustment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdjustLoggerRequest_LogAdjustment.Descriptor instead.
func (AdjustLoggerRequest_LogAdjustment) EnumDescriptor() ([]byte, []int) {
	return file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP(), []int{5, 0}
}

type CountLoggersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountLoggersRequest) Reset() {
	*x = CountLoggersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountLoggersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountLoggersRequest) ProtoMessage() {}

func (x *CountLoggersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountLoggersRequest.ProtoReflect.Descriptor instead.
func (*CountLoggersRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP(), []int{0}
}

type CountLoggersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountLoggersResponse) Reset() {
	*x = CountLoggersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountLoggersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountLoggersResponse) ProtoMessage() {}

func (x *CountLoggersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountLoggersResponse.ProtoReflect.Descriptor instead.
func (*CountLoggersResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP(), []int{1}
}

func (x *CountLoggersResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListLoggersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filters the agents returned by the list operation.
	Filter *ListLoggersRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// An output mask indicating which agent fields are set in the response.
	OutputMask *types.LoggerMask `protobuf:"bytes,2,opt,name=output_mask,json=outputMask,proto3" json:"output_mask,omitempty"`
	// The maximum number of results to return. The server may further
	// constrain this value, or if zero, choose its own.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous request, if any.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListLoggersRequest) Reset() {
	*x = ListLoggersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLoggersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoggersRequest) ProtoMessage() {}

func (x *ListLoggersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoggersRequest.ProtoReflect.Descriptor instead.
func (*ListLoggersRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP(), []int{2}
}

func (x *ListLoggersRequest) GetFilter() *ListLoggersRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListLoggersRequest) GetOutputMask() *types.LoggerMask {
	if x != nil {
		return x.OutputMask
	}
	return nil
}

func (x *ListLoggersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListLoggersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListLoggersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The agents.
	Loggers []*types.Logger `protobuf:"bytes,1,rep,name=loggers,proto3" json:"loggers,omitempty"`
	// The page token for the next request. Empty if there are no more results.
	// This field should be checked by clients even when a page_size was not
	// requested, since the server may choose its own (see page_size).
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListLoggersResponse) Reset() {
	*x = ListLoggersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLoggersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoggersResponse) ProtoMessage() {}

func (x *ListLoggersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoggersResponse.ProtoReflect.Descriptor instead.
func (*ListLoggersResponse) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP(), []int{3}
}

func (x *ListLoggersResponse) GetLoggers() []*types.Logger {
	if x != nil {
		return x.Loggers
	}
	return nil
}

func (x *ListLoggersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetLoggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The SPIFFE ID of the agent.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An output mask indicating which agent fields are set in the response.
	OutputMask *types.LoggerMask `protobuf:"bytes,2,opt,name=output_mask,json=outputMask,proto3" json:"output_mask,omitempty"`
}

func (x *GetLoggerRequest) Reset() {
	*x = GetLoggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoggerRequest) ProtoMessage() {}

func (x *GetLoggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoggerRequest.ProtoReflect.Descriptor instead.
func (*GetLoggerRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP(), []int{4}
}

func (x *GetLoggerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetLoggerRequest) GetOutputMask() *types.LoggerMask {
	if x != nil {
		return x.OutputMask
	}
	return nil
}

type AdjustLoggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The SPIFFE ID of the agent.
	Name             string                            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LogLevel         AdjustLoggerRequest_LogAdjustment `protobuf:"varint,2,opt,name=log_level,json=logLevel,proto3,enum=spire.api.agent.logger.v1.AdjustLoggerRequest_LogAdjustment" json:"log_level,omitempty"`
	AdjustSubloggers bool                              `protobuf:"varint,3,opt,name=adjust_subloggers,json=adjustSubloggers,proto3" json:"adjust_subloggers,omitempty"`
}

func (x *AdjustLoggerRequest) Reset() {
	*x = AdjustLoggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustLoggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustLoggerRequest) ProtoMessage() {}

func (x *AdjustLoggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustLoggerRequest.ProtoReflect.Descriptor instead.
func (*AdjustLoggerRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP(), []int{5}
}

func (x *AdjustLoggerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdjustLoggerRequest) GetLogLevel() AdjustLoggerRequest_LogAdjustment {
	if x != nil {
		return x.LogLevel
	}
	return AdjustLoggerRequest_PANIC_LEVEL
}

func (x *AdjustLoggerRequest) GetAdjustSubloggers() bool {
	if x != nil {
		return x.AdjustSubloggers
	}
	return false
}

type ListLoggersRequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filters agents to those matching the attestation type.
	ByName         string `protobuf:"bytes,1,opt,name=by_name,json=byName,proto3" json:"by_name,omitempty"`
	WithSubloggers bool   `protobuf:"varint,2,opt,name=with_subloggers,json=withSubloggers,proto3" json:"with_subloggers,omitempty"`
}

func (x *ListLoggersRequest_Filter) Reset() {
	*x = ListLoggersRequest_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLoggersRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoggersRequest_Filter) ProtoMessage() {}

func (x *ListLoggersRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_agent_logger_v1_logger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoggersRequest_Filter.ProtoReflect.Descriptor instead.
func (*ListLoggersRequest_Filter) Descriptor() ([]byte, []int) {
	return file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListLoggersRequest_Filter) GetByName() string {
	if x != nil {
		return x.ByName
	}
	return ""
}

func (x *ListLoggersRequest_Filter) GetWithSubloggers() bool {
	if x != nil {
		return x.WithSubloggers
	}
	return false
}

var File_spire_api_agent_logger_v1_logger_proto protoreflect.FileDescriptor

var file_spire_api_agent_logger_v1_logger_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15,
	0x0a, 0x13, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x14, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0xa8, 0x02, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x1a, 0x4a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x75,
	0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x77, 0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x22, 0x70,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x07, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x64, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xcb, 0x02, 0x0a, 0x13, 0x41, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x59, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2b, 0x0a,
	0x11, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x41, 0x4e, 0x49, 0x43, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x57, 0x41, 0x52, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x04, 0x12,
	0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x42, 0x55, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x05,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10,
	0x06, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x10, 0x07, 0x32, 0x92, 0x03, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12,
	0x6f, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12,
	0x2e, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12,
	0x2d, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x12, 0x56, 0x0a, 0x0c, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x12, 0x2e, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x6a, 0x75, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_agent_logger_v1_logger_proto_rawDescOnce sync.Once
	file_spire_api_agent_logger_v1_logger_proto_rawDescData = file_spire_api_agent_logger_v1_logger_proto_rawDesc
)

func file_spire_api_agent_logger_v1_logger_proto_rawDescGZIP() []byte {
	file_spire_api_agent_logger_v1_logger_proto_rawDescOnce.Do(func() {
		file_spire_api_agent_logger_v1_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_agent_logger_v1_logger_proto_rawDescData)
	})
	return file_spire_api_agent_logger_v1_logger_proto_rawDescData
}

var file_spire_api_agent_logger_v1_logger_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spire_api_agent_logger_v1_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spire_api_agent_logger_v1_logger_proto_goTypes = []interface{}{
	(AdjustLoggerRequest_LogAdjustment)(0), // 0: spire.api.agent.logger.v1.AdjustLoggerRequest.LogAdjustment
	(*CountLoggersRequest)(nil),            // 1: spire.api.agent.logger.v1.CountLoggersRequest
	(*CountLoggersResponse)(nil),           // 2: spire.api.agent.logger.v1.CountLoggersResponse
	(*ListLoggersRequest)(nil),             // 3: spire.api.agent.logger.v1.ListLoggersRequest
	(*ListLoggersResponse)(nil),            // 4: spire.api.agent.logger.v1.ListLoggersResponse
	(*GetLoggerRequest)(nil),               // 5: spire.api.agent.logger.v1.GetLoggerRequest
	(*AdjustLoggerRequest)(nil),            // 6: spire.api.agent.logger.v1.AdjustLoggerRequest
	(*ListLoggersRequest_Filter)(nil),      // 7: spire.api.agent.logger.v1.ListLoggersRequest.Filter
	(*types.LoggerMask)(nil),               // 8: spire.api.types.LoggerMask
	(*types.Logger)(nil),                   // 9: spire.api.types.Logger
	(*emptypb.Empty)(nil),                  // 10: google.protobuf.Empty
}
var file_spire_api_agent_logger_v1_logger_proto_depIdxs = []int32{
	7,  // 0: spire.api.agent.logger.v1.ListLoggersRequest.filter:type_name -> spire.api.agent.logger.v1.ListLoggersRequest.Filter
	8,  // 1: spire.api.agent.logger.v1.ListLoggersRequest.output_mask:type_name -> spire.api.types.LoggerMask
	9,  // 2: spire.api.agent.logger.v1.ListLoggersResponse.loggers:type_name -> spire.api.types.Logger
	8,  // 3: spire.api.agent.logger.v1.GetLoggerRequest.output_mask:type_name -> spire.api.types.LoggerMask
	0,  // 4: spire.api.agent.logger.v1.AdjustLoggerRequest.log_level:type_name -> spire.api.agent.logger.v1.AdjustLoggerRequest.LogAdjustment
	1,  // 5: spire.api.agent.logger.v1.Logger.CountLoggers:input_type -> spire.api.agent.logger.v1.CountLoggersRequest
	3,  // 6: spire.api.agent.logger.v1.Logger.ListLoggers:input_type -> spire.api.agent.logger.v1.ListLoggersRequest
	5,  // 7: spire.api.agent.logger.v1.Logger.GetLogger:input_type -> spire.api.agent.logger.v1.GetLoggerRequest
	6,  // 8: spire.api.agent.logger.v1.Logger.AdjustLogger:input_type -> spire.api.agent.logger.v1.AdjustLoggerRequest
	2,  // 9: spire.api.agent.logger.v1.Logger.CountLoggers:output_type -> spire.api.agent.logger.v1.CountLoggersResponse
	4,  // 10: spire.api.agent.logger.v1.Logger.ListLoggers:output_type -> spire.api.agent.logger.v1.ListLoggersResponse
	9,  // 11: spire.api.agent.logger.v1.Logger.GetLogger:output_type -> spire.api.types.Logger
	10, // 12: spire.api.agent.logger.v1.Logger.AdjustLogger:output_type -> google.protobuf.Empty
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_spire_api_agent_logger_v1_logger_proto_init() }
func file_spire_api_agent_logger_v1_logger_proto_init() {
	if File_spire_api_agent_logger_v1_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_agent_logger_v1_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountLoggersRequest); i {
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
		file_spire_api_agent_logger_v1_logger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountLoggersResponse); i {
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
		file_spire_api_agent_logger_v1_logger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLoggersRequest); i {
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
		file_spire_api_agent_logger_v1_logger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLoggersResponse); i {
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
		file_spire_api_agent_logger_v1_logger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoggerRequest); i {
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
		file_spire_api_agent_logger_v1_logger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustLoggerRequest); i {
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
		file_spire_api_agent_logger_v1_logger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLoggersRequest_Filter); i {
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
			RawDescriptor: file_spire_api_agent_logger_v1_logger_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_api_agent_logger_v1_logger_proto_goTypes,
		DependencyIndexes: file_spire_api_agent_logger_v1_logger_proto_depIdxs,
		EnumInfos:         file_spire_api_agent_logger_v1_logger_proto_enumTypes,
		MessageInfos:      file_spire_api_agent_logger_v1_logger_proto_msgTypes,
	}.Build()
	File_spire_api_agent_logger_v1_logger_proto = out.File
	file_spire_api_agent_logger_v1_logger_proto_rawDesc = nil
	file_spire_api_agent_logger_v1_logger_proto_goTypes = nil
	file_spire_api_agent_logger_v1_logger_proto_depIdxs = nil
}
