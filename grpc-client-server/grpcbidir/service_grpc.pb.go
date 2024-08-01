// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/service.proto

package grpc_bidir

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

// YourServiceClient is the client API for YourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YourServiceClient interface {
	RegisterClient(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*ClientResponse, error)
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	ReceiveMessages(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (YourService_ReceiveMessagesClient, error)
}

type yourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYourServiceClient(cc grpc.ClientConnInterface) YourServiceClient {
	return &yourServiceClient{cc}
}

func (c *yourServiceClient) RegisterClient(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/grpc_bidir.YourService/RegisterClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yourServiceClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/grpc_bidir.YourService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yourServiceClient) ReceiveMessages(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (YourService_ReceiveMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &YourService_ServiceDesc.Streams[0], "/grpc_bidir.YourService/ReceiveMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &yourServiceReceiveMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type YourService_ReceiveMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type yourServiceReceiveMessagesClient struct {
	grpc.ClientStream
}

func (x *yourServiceReceiveMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YourServiceServer is the server API for YourService service.
// All implementations must embed UnimplementedYourServiceServer
// for forward compatibility
type YourServiceServer interface {
	RegisterClient(context.Context, *ClientInfo) (*ClientResponse, error)
	SendMessage(context.Context, *MessageRequest) (*MessageResponse, error)
	ReceiveMessages(*ReceiveRequest, YourService_ReceiveMessagesServer) error
	mustEmbedUnimplementedYourServiceServer()
}

// UnimplementedYourServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYourServiceServer struct {
}

func (UnimplementedYourServiceServer) RegisterClient(context.Context, *ClientInfo) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClient not implemented")
}
func (UnimplementedYourServiceServer) SendMessage(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedYourServiceServer) ReceiveMessages(*ReceiveRequest, YourService_ReceiveMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessages not implemented")
}
func (UnimplementedYourServiceServer) mustEmbedUnimplementedYourServiceServer() {}

// UnsafeYourServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YourServiceServer will
// result in compilation errors.
type UnsafeYourServiceServer interface {
	mustEmbedUnimplementedYourServiceServer()
}

func RegisterYourServiceServer(s grpc.ServiceRegistrar, srv YourServiceServer) {
	s.RegisterService(&YourService_ServiceDesc, srv)
}

func _YourService_RegisterClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).RegisterClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_bidir.YourService/RegisterClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).RegisterClient(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _YourService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_bidir.YourService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).SendMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YourService_ReceiveMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(YourServiceServer).ReceiveMessages(m, &yourServiceReceiveMessagesServer{stream})
}

type YourService_ReceiveMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type yourServiceReceiveMessagesServer struct {
	grpc.ServerStream
}

func (x *yourServiceReceiveMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// YourService_ServiceDesc is the grpc.ServiceDesc for YourService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YourService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_bidir.YourService",
	HandlerType: (*YourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterClient",
			Handler:    _YourService_RegisterClient_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _YourService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveMessages",
			Handler:       _YourService_ReceiveMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
