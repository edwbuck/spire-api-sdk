// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/api/server/logger/v1/logger.proto

package loggerv1

import (
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
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

// Empty Get Logger Request message for future extension
type GetLoggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLoggerRequest) Reset() {
	*x = GetLoggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_server_logger_v1_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoggerRequest) ProtoMessage() {}

func (x *GetLoggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_server_logger_v1_logger_proto_msgTypes[0]
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
	return file_spire_api_server_logger_v1_logger_proto_rawDescGZIP(), []int{0}
}

// Set Log Level Request message
type SetLogLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new level the logger should assume
	NewLevel types.LogLevel `protobuf:"varint,1,opt,name=new_level,json=newLevel,proto3,enum=spire.api.types.LogLevel" json:"new_level,omitempty"`
}

func (x *SetLogLevelRequest) Reset() {
	*x = SetLogLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_server_logger_v1_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLogLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLogLevelRequest) ProtoMessage() {}

func (x *SetLogLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_server_logger_v1_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLogLevelRequest.ProtoReflect.Descriptor instead.
func (*SetLogLevelRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_server_logger_v1_logger_proto_rawDescGZIP(), []int{1}
}

func (x *SetLogLevelRequest) GetNewLevel() types.LogLevel {
	if x != nil {
		return x.NewLevel
	}
	return types.LogLevel(0)
}

// Empty Reset Log Level Request message for future extension
type ResetLogLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetLogLevelRequest) Reset() {
	*x = ResetLogLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_server_logger_v1_logger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetLogLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetLogLevelRequest) ProtoMessage() {}

func (x *ResetLogLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_server_logger_v1_logger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetLogLevelRequest.ProtoReflect.Descriptor instead.
func (*ResetLogLevelRequest) Descriptor() ([]byte, []int) {
	return file_spire_api_server_logger_v1_logger_proto_rawDescGZIP(), []int{2}
}

var File_spire_api_server_logger_v1_logger_proto protoreflect.FileDescriptor

var file_spire_api_server_logger_v1_logger_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x09, 0x6e, 0x65, 0x77, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x08, 0x6e, 0x65, 0x77,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x90, 0x02,
	0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0b,
	0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x30, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_server_logger_v1_logger_proto_rawDescOnce sync.Once
	file_spire_api_server_logger_v1_logger_proto_rawDescData = file_spire_api_server_logger_v1_logger_proto_rawDesc
)

func file_spire_api_server_logger_v1_logger_proto_rawDescGZIP() []byte {
	file_spire_api_server_logger_v1_logger_proto_rawDescOnce.Do(func() {
		file_spire_api_server_logger_v1_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_server_logger_v1_logger_proto_rawDescData)
	})
	return file_spire_api_server_logger_v1_logger_proto_rawDescData
}

var file_spire_api_server_logger_v1_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spire_api_server_logger_v1_logger_proto_goTypes = []interface{}{
	(*GetLoggerRequest)(nil),     // 0: spire.api.server.logger.v1.GetLoggerRequest
	(*SetLogLevelRequest)(nil),   // 1: spire.api.server.logger.v1.SetLogLevelRequest
	(*ResetLogLevelRequest)(nil), // 2: spire.api.server.logger.v1.ResetLogLevelRequest
	(types.LogLevel)(0),          // 3: spire.api.types.LogLevel
	(*types.Logger)(nil),         // 4: spire.api.types.Logger
}
var file_spire_api_server_logger_v1_logger_proto_depIdxs = []int32{
	3, // 0: spire.api.server.logger.v1.SetLogLevelRequest.new_level:type_name -> spire.api.types.LogLevel
	0, // 1: spire.api.server.logger.v1.Logger.GetLogger:input_type -> spire.api.server.logger.v1.GetLoggerRequest
	1, // 2: spire.api.server.logger.v1.Logger.SetLogLevel:input_type -> spire.api.server.logger.v1.SetLogLevelRequest
	2, // 3: spire.api.server.logger.v1.Logger.ResetLogLevel:input_type -> spire.api.server.logger.v1.ResetLogLevelRequest
	4, // 4: spire.api.server.logger.v1.Logger.GetLogger:output_type -> spire.api.types.Logger
	4, // 5: spire.api.server.logger.v1.Logger.SetLogLevel:output_type -> spire.api.types.Logger
	4, // 6: spire.api.server.logger.v1.Logger.ResetLogLevel:output_type -> spire.api.types.Logger
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spire_api_server_logger_v1_logger_proto_init() }
func file_spire_api_server_logger_v1_logger_proto_init() {
	if File_spire_api_server_logger_v1_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_server_logger_v1_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_spire_api_server_logger_v1_logger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLogLevelRequest); i {
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
		file_spire_api_server_logger_v1_logger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetLogLevelRequest); i {
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
			RawDescriptor: file_spire_api_server_logger_v1_logger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_api_server_logger_v1_logger_proto_goTypes,
		DependencyIndexes: file_spire_api_server_logger_v1_logger_proto_depIdxs,
		MessageInfos:      file_spire_api_server_logger_v1_logger_proto_msgTypes,
	}.Build()
	File_spire_api_server_logger_v1_logger_proto = out.File
	file_spire_api_server_logger_v1_logger_proto_rawDesc = nil
	file_spire_api_server_logger_v1_logger_proto_goTypes = nil
	file_spire_api_server_logger_v1_logger_proto_depIdxs = nil
}
