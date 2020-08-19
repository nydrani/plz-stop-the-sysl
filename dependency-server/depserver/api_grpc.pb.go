// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package depserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MyserverdepClient is the client API for Myserverdep service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyserverdepClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type myserverdepClient struct {
	cc grpc.ClientConnInterface
}

func NewMyserverdepClient(cc grpc.ClientConnInterface) MyserverdepClient {
	return &myserverdepClient{cc}
}

func (c *myserverdepClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/depserver.myserverdep/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyserverdepServer is the server API for Myserverdep service.
// All implementations must embed UnimplementedMyserverdepServer
// for forward compatibility
type MyserverdepServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedMyserverdepServer()
}

// UnimplementedMyserverdepServer must be embedded to have forward compatible implementations.
type UnimplementedMyserverdepServer struct {
}

func (*UnimplementedMyserverdepServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedMyserverdepServer) mustEmbedUnimplementedMyserverdepServer() {}

func RegisterMyserverdepServer(s *grpc.Server, srv MyserverdepServer) {
	s.RegisterService(&_Myserverdep_serviceDesc, srv)
}

func _Myserverdep_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyserverdepServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/depserver.myserverdep/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyserverdepServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Myserverdep_serviceDesc = grpc.ServiceDesc{
	ServiceName: "depserver.myserverdep",
	HandlerType: (*MyserverdepServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Myserverdep_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}