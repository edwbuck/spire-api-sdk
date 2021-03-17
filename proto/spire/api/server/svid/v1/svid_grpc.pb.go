// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package svidv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SVIDClient is the client API for SVID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SVIDClient interface {
	// Mints a one-off X509-SVID outside of the normal node/workload
	// registration process.
	//
	// The caller must be local or present an admin X509-SVID.
	MintX509SVID(ctx context.Context, in *MintX509SVIDRequest, opts ...grpc.CallOption) (*MintX509SVIDResponse, error)
	// Mints a one-off JWT-SVID outside of the normal node/workload
	// registration process.
	//
	// The caller must be local or present an admin X509-SVID.
	MintJWTSVID(ctx context.Context, in *MintJWTSVIDRequest, opts ...grpc.CallOption) (*MintJWTSVIDResponse, error)
	// Creates one or more X509-SVIDs from registration entries.
	//
	// The caller must present an active agent X509-SVID that is authorized
	// to mint the requested entries. See the Entry GetAuthorizedEntries RPC.
	BatchNewX509SVID(ctx context.Context, in *BatchNewX509SVIDRequest, opts ...grpc.CallOption) (*BatchNewX509SVIDResponse, error)
	// Creates an JWT-SVID from a registration entry.
	//
	// The caller must present an active agent X509-SVID that is authorized
	// to mint the requested entry. See the Entry GetAuthorizedEntries RPC.
	NewJWTSVID(ctx context.Context, in *NewJWTSVIDRequest, opts ...grpc.CallOption) (*NewJWTSVIDResponse, error)
	// Creates an X509 CA certificate appropriate for use by a downstream
	// entity to mint X509-SVIDs.
	//
	// The caller must present a downstream X509-SVID.
	NewDownstreamX509CA(ctx context.Context, in *NewDownstreamX509CARequest, opts ...grpc.CallOption) (*NewDownstreamX509CAResponse, error)
}

type sVIDClient struct {
	cc grpc.ClientConnInterface
}

func NewSVIDClient(cc grpc.ClientConnInterface) SVIDClient {
	return &sVIDClient{cc}
}

func (c *sVIDClient) MintX509SVID(ctx context.Context, in *MintX509SVIDRequest, opts ...grpc.CallOption) (*MintX509SVIDResponse, error) {
	out := new(MintX509SVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/MintX509SVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) MintJWTSVID(ctx context.Context, in *MintJWTSVIDRequest, opts ...grpc.CallOption) (*MintJWTSVIDResponse, error) {
	out := new(MintJWTSVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/MintJWTSVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) BatchNewX509SVID(ctx context.Context, in *BatchNewX509SVIDRequest, opts ...grpc.CallOption) (*BatchNewX509SVIDResponse, error) {
	out := new(BatchNewX509SVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/BatchNewX509SVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) NewJWTSVID(ctx context.Context, in *NewJWTSVIDRequest, opts ...grpc.CallOption) (*NewJWTSVIDResponse, error) {
	out := new(NewJWTSVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/NewJWTSVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) NewDownstreamX509CA(ctx context.Context, in *NewDownstreamX509CARequest, opts ...grpc.CallOption) (*NewDownstreamX509CAResponse, error) {
	out := new(NewDownstreamX509CAResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.svid.v1.SVID/NewDownstreamX509CA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SVIDServer is the server API for SVID service.
// All implementations must embed UnimplementedSVIDServer
// for forward compatibility
type SVIDServer interface {
	// Mints a one-off X509-SVID outside of the normal node/workload
	// registration process.
	//
	// The caller must be local or present an admin X509-SVID.
	MintX509SVID(context.Context, *MintX509SVIDRequest) (*MintX509SVIDResponse, error)
	// Mints a one-off JWT-SVID outside of the normal node/workload
	// registration process.
	//
	// The caller must be local or present an admin X509-SVID.
	MintJWTSVID(context.Context, *MintJWTSVIDRequest) (*MintJWTSVIDResponse, error)
	// Creates one or more X509-SVIDs from registration entries.
	//
	// The caller must present an active agent X509-SVID that is authorized
	// to mint the requested entries. See the Entry GetAuthorizedEntries RPC.
	BatchNewX509SVID(context.Context, *BatchNewX509SVIDRequest) (*BatchNewX509SVIDResponse, error)
	// Creates an JWT-SVID from a registration entry.
	//
	// The caller must present an active agent X509-SVID that is authorized
	// to mint the requested entry. See the Entry GetAuthorizedEntries RPC.
	NewJWTSVID(context.Context, *NewJWTSVIDRequest) (*NewJWTSVIDResponse, error)
	// Creates an X509 CA certificate appropriate for use by a downstream
	// entity to mint X509-SVIDs.
	//
	// The caller must present a downstream X509-SVID.
	NewDownstreamX509CA(context.Context, *NewDownstreamX509CARequest) (*NewDownstreamX509CAResponse, error)
	mustEmbedUnimplementedSVIDServer()
}

// UnimplementedSVIDServer must be embedded to have forward compatible implementations.
type UnimplementedSVIDServer struct {
}

func (UnimplementedSVIDServer) MintX509SVID(context.Context, *MintX509SVIDRequest) (*MintX509SVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintX509SVID not implemented")
}
func (UnimplementedSVIDServer) MintJWTSVID(context.Context, *MintJWTSVIDRequest) (*MintJWTSVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintJWTSVID not implemented")
}
func (UnimplementedSVIDServer) BatchNewX509SVID(context.Context, *BatchNewX509SVIDRequest) (*BatchNewX509SVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchNewX509SVID not implemented")
}
func (UnimplementedSVIDServer) NewJWTSVID(context.Context, *NewJWTSVIDRequest) (*NewJWTSVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewJWTSVID not implemented")
}
func (UnimplementedSVIDServer) NewDownstreamX509CA(context.Context, *NewDownstreamX509CARequest) (*NewDownstreamX509CAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewDownstreamX509CA not implemented")
}
func (UnimplementedSVIDServer) mustEmbedUnimplementedSVIDServer() {}

// UnsafeSVIDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SVIDServer will
// result in compilation errors.
type UnsafeSVIDServer interface {
	mustEmbedUnimplementedSVIDServer()
}

func RegisterSVIDServer(s grpc.ServiceRegistrar, srv SVIDServer) {
	s.RegisterService(&_SVID_serviceDesc, srv)
}

func _SVID_MintX509SVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintX509SVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).MintX509SVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/MintX509SVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).MintX509SVID(ctx, req.(*MintX509SVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_MintJWTSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintJWTSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).MintJWTSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/MintJWTSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).MintJWTSVID(ctx, req.(*MintJWTSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_BatchNewX509SVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchNewX509SVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).BatchNewX509SVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/BatchNewX509SVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).BatchNewX509SVID(ctx, req.(*BatchNewX509SVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_NewJWTSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewJWTSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).NewJWTSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/NewJWTSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).NewJWTSVID(ctx, req.(*NewJWTSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_NewDownstreamX509CA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewDownstreamX509CARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).NewDownstreamX509CA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.svid.v1.SVID/NewDownstreamX509CA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).NewDownstreamX509CA(ctx, req.(*NewDownstreamX509CARequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SVID_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.server.svid.v1.SVID",
	HandlerType: (*SVIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintX509SVID",
			Handler:    _SVID_MintX509SVID_Handler,
		},
		{
			MethodName: "MintJWTSVID",
			Handler:    _SVID_MintJWTSVID_Handler,
		},
		{
			MethodName: "BatchNewX509SVID",
			Handler:    _SVID_BatchNewX509SVID_Handler,
		},
		{
			MethodName: "NewJWTSVID",
			Handler:    _SVID_NewJWTSVID_Handler,
		},
		{
			MethodName: "NewDownstreamX509CA",
			Handler:    _SVID_NewDownstreamX509CA_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/api/server/svid/v1/svid.proto",
}
