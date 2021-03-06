// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc2
// source: configserver.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConfigServerClient is the client API for ConfigServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServerClient interface {
	// Get Flow configuration from config server.
	GetFlow(ctx context.Context, in *GetFlowRequest, opts ...grpc.CallOption) (*GetFlowResponse, error)
}

type configServerClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServerClient(cc grpc.ClientConnInterface) ConfigServerClient {
	return &configServerClient{cc}
}

func (c *configServerClient) GetFlow(ctx context.Context, in *GetFlowRequest, opts ...grpc.CallOption) (*GetFlowResponse, error) {
	out := new(GetFlowResponse)
	err := c.cc.Invoke(ctx, "/configserver.ConfigServer/GetFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServerServer is the server API for ConfigServer service.
// All implementations must embed UnimplementedConfigServerServer
// for forward compatibility
type ConfigServerServer interface {
	// Get Flow configuration from config server.
	GetFlow(context.Context, *GetFlowRequest) (*GetFlowResponse, error)
	mustEmbedUnimplementedConfigServerServer()
}

// UnimplementedConfigServerServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServerServer struct {
}

func (UnimplementedConfigServerServer) GetFlow(context.Context, *GetFlowRequest) (*GetFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlow not implemented")
}
func (UnimplementedConfigServerServer) mustEmbedUnimplementedConfigServerServer() {}

// UnsafeConfigServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServerServer will
// result in compilation errors.
type UnsafeConfigServerServer interface {
	mustEmbedUnimplementedConfigServerServer()
}

func RegisterConfigServerServer(s grpc.ServiceRegistrar, srv ConfigServerServer) {
	s.RegisterService(&ConfigServer_ServiceDesc, srv)
}

func _ConfigServer_GetFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServerServer).GetFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configserver.ConfigServer/GetFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServerServer).GetFlow(ctx, req.(*GetFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigServer_ServiceDesc is the grpc.ServiceDesc for ConfigServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "configserver.ConfigServer",
	HandlerType: (*ConfigServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlow",
			Handler:    _ConfigServer_GetFlow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configserver.proto",
}
