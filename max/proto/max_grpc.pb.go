// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: max.proto

package proto

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

// MaxServiceClient is the client API for MaxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MaxServiceClient interface {
	Max(ctx context.Context, opts ...grpc.CallOption) (MaxService_MaxClient, error)
}

type maxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaxServiceClient(cc grpc.ClientConnInterface) MaxServiceClient {
	return &maxServiceClient{cc}
}

func (c *maxServiceClient) Max(ctx context.Context, opts ...grpc.CallOption) (MaxService_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &MaxService_ServiceDesc.Streams[0], "/max.MaxService/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &maxServiceMaxClient{stream}
	return x, nil
}

type MaxService_MaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type maxServiceMaxClient struct {
	grpc.ClientStream
}

func (x *maxServiceMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maxServiceMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaxServiceServer is the server API for MaxService service.
// All implementations must embed UnimplementedMaxServiceServer
// for forward compatibility
type MaxServiceServer interface {
	Max(MaxService_MaxServer) error
	mustEmbedUnimplementedMaxServiceServer()
}

// UnimplementedMaxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMaxServiceServer struct {
}

func (UnimplementedMaxServiceServer) Max(MaxService_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedMaxServiceServer) mustEmbedUnimplementedMaxServiceServer() {}

// UnsafeMaxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MaxServiceServer will
// result in compilation errors.
type UnsafeMaxServiceServer interface {
	mustEmbedUnimplementedMaxServiceServer()
}

func RegisterMaxServiceServer(s grpc.ServiceRegistrar, srv MaxServiceServer) {
	s.RegisterService(&MaxService_ServiceDesc, srv)
}

func _MaxService_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaxServiceServer).Max(&maxServiceMaxServer{stream})
}

type MaxService_MaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type maxServiceMaxServer struct {
	grpc.ServerStream
}

func (x *maxServiceMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maxServiceMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaxService_ServiceDesc is the grpc.ServiceDesc for MaxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MaxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "max.MaxService",
	HandlerType: (*MaxServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Max",
			Handler:       _MaxService_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "max.proto",
}
