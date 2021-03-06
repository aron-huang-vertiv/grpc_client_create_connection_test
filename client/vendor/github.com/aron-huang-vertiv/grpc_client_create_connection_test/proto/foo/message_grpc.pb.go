// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package foo

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

// FooServiceClient is the client API for FooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FooServiceClient interface {
	Foo(ctx context.Context, in *FooReq, opts ...grpc.CallOption) (*FooResp, error)
}

type fooServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFooServiceClient(cc grpc.ClientConnInterface) FooServiceClient {
	return &fooServiceClient{cc}
}

func (c *fooServiceClient) Foo(ctx context.Context, in *FooReq, opts ...grpc.CallOption) (*FooResp, error) {
	out := new(FooResp)
	err := c.cc.Invoke(ctx, "/foo.FooService/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServiceServer is the server API for FooService service.
// All implementations must embed UnimplementedFooServiceServer
// for forward compatibility
type FooServiceServer interface {
	Foo(context.Context, *FooReq) (*FooResp, error)
	mustEmbedUnimplementedFooServiceServer()
}

// UnimplementedFooServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFooServiceServer struct {
}

func (UnimplementedFooServiceServer) Foo(context.Context, *FooReq) (*FooResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}
func (UnimplementedFooServiceServer) mustEmbedUnimplementedFooServiceServer() {}

// UnsafeFooServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FooServiceServer will
// result in compilation errors.
type UnsafeFooServiceServer interface {
	mustEmbedUnimplementedFooServiceServer()
}

func RegisterFooServiceServer(s grpc.ServiceRegistrar, srv FooServiceServer) {
	s.RegisterService(&FooService_ServiceDesc, srv)
}

func _FooService_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServiceServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.FooService/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServiceServer).Foo(ctx, req.(*FooReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FooService_ServiceDesc is the grpc.ServiceDesc for FooService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FooService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "foo.FooService",
	HandlerType: (*FooServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _FooService_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
