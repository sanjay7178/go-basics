// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service.proto

package machineinfo

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

// MachineInfoServiceClient is the client API for MachineInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineInfoServiceClient interface {
	GetMachineInfo(ctx context.Context, opts ...grpc.CallOption) (MachineInfoService_GetMachineInfoClient, error)
}

type machineInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineInfoServiceClient(cc grpc.ClientConnInterface) MachineInfoServiceClient {
	return &machineInfoServiceClient{cc}
}

func (c *machineInfoServiceClient) GetMachineInfo(ctx context.Context, opts ...grpc.CallOption) (MachineInfoService_GetMachineInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &MachineInfoService_ServiceDesc.Streams[0], "/machineinfo.MachineInfoService/GetMachineInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &machineInfoServiceGetMachineInfoClient{stream}
	return x, nil
}

type MachineInfoService_GetMachineInfoClient interface {
	Send(*MachineInfoRequest) error
	Recv() (*MachineInfoResponse, error)
	grpc.ClientStream
}

type machineInfoServiceGetMachineInfoClient struct {
	grpc.ClientStream
}

func (x *machineInfoServiceGetMachineInfoClient) Send(m *MachineInfoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *machineInfoServiceGetMachineInfoClient) Recv() (*MachineInfoResponse, error) {
	m := new(MachineInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MachineInfoServiceServer is the server API for MachineInfoService service.
// All implementations must embed UnimplementedMachineInfoServiceServer
// for forward compatibility
type MachineInfoServiceServer interface {
	GetMachineInfo(MachineInfoService_GetMachineInfoServer) error
	mustEmbedUnimplementedMachineInfoServiceServer()
}

// UnimplementedMachineInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMachineInfoServiceServer struct {
}

func (UnimplementedMachineInfoServiceServer) GetMachineInfo(MachineInfoService_GetMachineInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMachineInfo not implemented")
}
func (UnimplementedMachineInfoServiceServer) mustEmbedUnimplementedMachineInfoServiceServer() {}

// UnsafeMachineInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineInfoServiceServer will
// result in compilation errors.
type UnsafeMachineInfoServiceServer interface {
	mustEmbedUnimplementedMachineInfoServiceServer()
}

func RegisterMachineInfoServiceServer(s grpc.ServiceRegistrar, srv MachineInfoServiceServer) {
	s.RegisterService(&MachineInfoService_ServiceDesc, srv)
}

func _MachineInfoService_GetMachineInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MachineInfoServiceServer).GetMachineInfo(&machineInfoServiceGetMachineInfoServer{stream})
}

type MachineInfoService_GetMachineInfoServer interface {
	Send(*MachineInfoResponse) error
	Recv() (*MachineInfoRequest, error)
	grpc.ServerStream
}

type machineInfoServiceGetMachineInfoServer struct {
	grpc.ServerStream
}

func (x *machineInfoServiceGetMachineInfoServer) Send(m *MachineInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *machineInfoServiceGetMachineInfoServer) Recv() (*MachineInfoRequest, error) {
	m := new(MachineInfoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MachineInfoService_ServiceDesc is the grpc.ServiceDesc for MachineInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MachineInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "machineinfo.MachineInfoService",
	HandlerType: (*MachineInfoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMachineInfo",
			Handler:       _MachineInfoService_GetMachineInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
