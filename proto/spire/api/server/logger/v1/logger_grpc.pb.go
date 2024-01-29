// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package loggerv1

import (
	context "context"
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerClient interface {
	// Count agents.
	//
	// The caller must be local or present an admin X509-SVID.
	CountLoggers(ctx context.Context, in *CountLoggersRequest, opts ...grpc.CallOption) (*CountLoggersResponse, error)
	// Lists loggers.
	//
	// The caller must be local or present an admin X509-SVID.
	ListLoggers(ctx context.Context, in *ListLoggersRequest, opts ...grpc.CallOption) (*ListLoggersResponse, error)
	// Gets an agent.
	//
	// The caller must be local or present an admin X509-SVID.
	GetLogger(ctx context.Context, in *GetLoggerRequest, opts ...grpc.CallOption) (*types.Logger, error)
	// the Issuer AttestAgent RPC.
	//
	// The caller must be local or present an admin X509-SVID.
	AdjustLogger(ctx context.Context, in *AdjustLoggerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type loggerClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerClient(cc grpc.ClientConnInterface) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) CountLoggers(ctx context.Context, in *CountLoggersRequest, opts ...grpc.CallOption) (*CountLoggersResponse, error) {
	out := new(CountLoggersResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.logger.v1.Logger/CountLoggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerClient) ListLoggers(ctx context.Context, in *ListLoggersRequest, opts ...grpc.CallOption) (*ListLoggersResponse, error) {
	out := new(ListLoggersResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.logger.v1.Logger/ListLoggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerClient) GetLogger(ctx context.Context, in *GetLoggerRequest, opts ...grpc.CallOption) (*types.Logger, error) {
	out := new(types.Logger)
	err := c.cc.Invoke(ctx, "/spire.api.server.logger.v1.Logger/GetLogger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerClient) AdjustLogger(ctx context.Context, in *AdjustLoggerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/spire.api.server.logger.v1.Logger/AdjustLogger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
// All implementations must embed UnimplementedLoggerServer
// for forward compatibility
type LoggerServer interface {
	// Count agents.
	//
	// The caller must be local or present an admin X509-SVID.
	CountLoggers(context.Context, *CountLoggersRequest) (*CountLoggersResponse, error)
	// Lists loggers.
	//
	// The caller must be local or present an admin X509-SVID.
	ListLoggers(context.Context, *ListLoggersRequest) (*ListLoggersResponse, error)
	// Gets an agent.
	//
	// The caller must be local or present an admin X509-SVID.
	GetLogger(context.Context, *GetLoggerRequest) (*types.Logger, error)
	// the Issuer AttestAgent RPC.
	//
	// The caller must be local or present an admin X509-SVID.
	AdjustLogger(context.Context, *AdjustLoggerRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedLoggerServer()
}

// UnimplementedLoggerServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerServer struct {
}

func (UnimplementedLoggerServer) CountLoggers(context.Context, *CountLoggersRequest) (*CountLoggersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountLoggers not implemented")
}
func (UnimplementedLoggerServer) ListLoggers(context.Context, *ListLoggersRequest) (*ListLoggersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLoggers not implemented")
}
func (UnimplementedLoggerServer) GetLogger(context.Context, *GetLoggerRequest) (*types.Logger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogger not implemented")
}
func (UnimplementedLoggerServer) AdjustLogger(context.Context, *AdjustLoggerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdjustLogger not implemented")
}
func (UnimplementedLoggerServer) mustEmbedUnimplementedLoggerServer() {}

// UnsafeLoggerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerServer will
// result in compilation errors.
type UnsafeLoggerServer interface {
	mustEmbedUnimplementedLoggerServer()
}

func RegisterLoggerServer(s grpc.ServiceRegistrar, srv LoggerServer) {
	s.RegisterService(&_Logger_serviceDesc, srv)
}

func _Logger_CountLoggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountLoggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).CountLoggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.logger.v1.Logger/CountLoggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).CountLoggers(ctx, req.(*CountLoggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logger_ListLoggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLoggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).ListLoggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.logger.v1.Logger/ListLoggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).ListLoggers(ctx, req.(*ListLoggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logger_GetLogger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).GetLogger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.logger.v1.Logger/GetLogger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).GetLogger(ctx, req.(*GetLoggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logger_AdjustLogger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustLoggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).AdjustLogger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.logger.v1.Logger/AdjustLogger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).AdjustLogger(ctx, req.(*AdjustLoggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.server.logger.v1.Logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountLoggers",
			Handler:    _Logger_CountLoggers_Handler,
		},
		{
			MethodName: "ListLoggers",
			Handler:    _Logger_ListLoggers_Handler,
		},
		{
			MethodName: "GetLogger",
			Handler:    _Logger_GetLogger_Handler,
		},
		{
			MethodName: "AdjustLogger",
			Handler:    _Logger_AdjustLogger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/api/server/logger/v1/logger.proto",
}
