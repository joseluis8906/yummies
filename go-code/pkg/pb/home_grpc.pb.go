// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: home.proto

package pb

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

const (
	HomeService_Index_FullMethodName = "/yummies.HomeService/Index"
)

// HomeServiceClient is the client API for HomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeServiceClient interface {
	Index(ctx context.Context, in *HomeIndexRequest, opts ...grpc.CallOption) (*HomeIndexResponse, error)
}

type homeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeServiceClient(cc grpc.ClientConnInterface) HomeServiceClient {
	return &homeServiceClient{cc}
}

func (c *homeServiceClient) Index(ctx context.Context, in *HomeIndexRequest, opts ...grpc.CallOption) (*HomeIndexResponse, error) {
	out := new(HomeIndexResponse)
	err := c.cc.Invoke(ctx, HomeService_Index_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeServiceServer is the server API for HomeService service.
// All implementations must embed UnimplementedHomeServiceServer
// for forward compatibility
type HomeServiceServer interface {
	Index(context.Context, *HomeIndexRequest) (*HomeIndexResponse, error)
	mustEmbedUnimplementedHomeServiceServer()
}

// UnimplementedHomeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHomeServiceServer struct {
}

func (UnimplementedHomeServiceServer) Index(context.Context, *HomeIndexRequest) (*HomeIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedHomeServiceServer) mustEmbedUnimplementedHomeServiceServer() {}

// UnsafeHomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeServiceServer will
// result in compilation errors.
type UnsafeHomeServiceServer interface {
	mustEmbedUnimplementedHomeServiceServer()
}

func RegisterHomeServiceServer(s grpc.ServiceRegistrar, srv HomeServiceServer) {
	s.RegisterService(&HomeService_ServiceDesc, srv)
}

func _HomeService_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomeIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServiceServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HomeService_Index_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServiceServer).Index(ctx, req.(*HomeIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeService_ServiceDesc is the grpc.ServiceDesc for HomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yummies.HomeService",
	HandlerType: (*HomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _HomeService_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "home.proto",
}
