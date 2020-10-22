// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_control_plane.proto

package atlas

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("service_control_plane.proto", fileDescriptor_144c4e5d826a1f85) }

var fileDescriptor_144c4e5d826a1f85 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x5b, 0xc1, 0x3f, 0x84, 0x2a, 0x3a, 0x54, 0xa4, 0xf1, 0x20, 0xf4, 0x24, 0x08, 0x4b,
	0x51, 0xf0, 0xe0, 0x75, 0x15, 0x4f, 0x05, 0xd1, 0x0f, 0x10, 0xc6, 0x6d, 0xc0, 0xc5, 0x6c, 0x36,
	0x26, 0xb3, 0xc2, 0x7e, 0x25, 0x3f, 0xa5, 0x24, 0x19, 0xc4, 0x6e, 0xf7, 0xf8, 0x7e, 0xef, 0xcd,
	0xf0, 0x66, 0xc4, 0x65, 0xd0, 0xfe, 0xbb, 0xae, 0xb4, 0xaa, 0x5a, 0x4b, 0xbe, 0x35, 0xca, 0x19,
	0xb4, 0xba, 0x70, 0xbe, 0xa5, 0x16, 0xf6, 0x91, 0x0c, 0x06, 0x09, 0xd4, 0xbb, 0x18, 0x68, 0x1a,
	0xb4, 0x9b, 0x6c, 0xc9, 0xf3, 0xc4, 0x36, 0x48, 0xa8, 0xaa, 0x8f, 0xce, 0x7e, 0x32, 0x3e, 0x4b,
	0xb8, 0xd1, 0xe4, 0xeb, 0x8a, 0xd1, 0x22, 0xa1, 0x40, 0x48, 0x5d, 0x50, 0x5e, 0x7f, 0x75, 0x3a,
	0x10, 0x5b, 0x57, 0x23, 0x96, 0x6a, 0x3a, 0x43, 0x35, 0x07, 0x2e, 0xb6, 0x03, 0xce, 0xf4, 0xd9,
	0xb8, 0xfd, 0xd9, 0x13, 0xb3, 0x32, 0x37, 0x7e, 0x89, 0x85, 0x61, 0x25, 0x8e, 0xca, 0x5c, 0x30,
	0xc0, 0x49, 0x91, 0x7a, 0x17, 0x0c, 0xe4, 0x40, 0x2f, 0x27, 0xd7, 0xd3, 0xd5, 0x14, 0xee, 0x85,
	0x78, 0x44, 0xc2, 0x32, 0xb6, 0x0f, 0x70, 0xca, 0x99, 0x3f, 0x24, 0x77, 0x08, 0xcf, 0xdd, 0x88,
	0xc3, 0x75, 0xba, 0x2f, 0xc0, 0x31, 0x47, 0xb2, 0x96, 0xdb, 0x32, 0xc6, 0xe1, 0x41, 0xcc, 0x9e,
	0x2c, 0xd5, 0xd4, 0xbf, 0xa5, 0x1b, 0x60, 0xce, 0x91, 0x2c, 0x5f, 0xf3, 0xc9, 0x12, 0x06, 0xd4,
	0x99, 0x7e, 0x39, 0x81, 0x67, 0x31, 0xff, 0x3f, 0xbb, 0x8e, 0x7f, 0x71, 0x46, 0xc3, 0x62, 0x6c,
	0x47, 0x72, 0xc7, 0x17, 0xbd, 0x1f, 0xa4, 0x9f, 0xdd, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x1a, 0x52, 0x72, 0xec, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControlPlaneClient is the client API for ControlPlane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlPlaneClient interface {
	// Bi-directional Commands stream
	//
	// Commands are sent from service to client and from client to server
	Commands(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_CommandsClient, error)
	// Bi-directional Data stream
	//
	// Some commands may be followed by data load. Be it logs, dumps, etc.
	DataChunks(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_DataChunksClient, error)
	// Metrics stream
	//
	// Some commands may be followed by metrics stream.
	Metrics(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_MetricsClient, error)
	// EntityStatus checks status of the entity on the server
	EntityStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	// EntityStatusMulti checks status of the multiple entities on server
	EntityStatusMultiple(ctx context.Context, in *StatusRequestMulti, opts ...grpc.CallOption) (*StatusReply, error)
}

type controlPlaneClient struct {
	cc *grpc.ClientConn
}

func NewControlPlaneClient(cc *grpc.ClientConn) ControlPlaneClient {
	return &controlPlaneClient{cc}
}

func (c *controlPlaneClient) Commands(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_CommandsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ControlPlane_serviceDesc.Streams[0], "/atlas.ControlPlane/Commands", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlPlaneCommandsClient{stream}
	return x, nil
}

type ControlPlane_CommandsClient interface {
	Send(*Command) error
	Recv() (*Command, error)
	grpc.ClientStream
}

type controlPlaneCommandsClient struct {
	grpc.ClientStream
}

func (x *controlPlaneCommandsClient) Send(m *Command) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controlPlaneCommandsClient) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlPlaneClient) DataChunks(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_DataChunksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ControlPlane_serviceDesc.Streams[1], "/atlas.ControlPlane/DataChunks", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlPlaneDataChunksClient{stream}
	return x, nil
}

type ControlPlane_DataChunksClient interface {
	Send(*DataChunk) error
	Recv() (*DataChunk, error)
	grpc.ClientStream
}

type controlPlaneDataChunksClient struct {
	grpc.ClientStream
}

func (x *controlPlaneDataChunksClient) Send(m *DataChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controlPlaneDataChunksClient) Recv() (*DataChunk, error) {
	m := new(DataChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlPlaneClient) Metrics(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_MetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ControlPlane_serviceDesc.Streams[2], "/atlas.ControlPlane/Metrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlPlaneMetricsClient{stream}
	return x, nil
}

type ControlPlane_MetricsClient interface {
	Send(*Metric) error
	CloseAndRecv() (*Metric, error)
	grpc.ClientStream
}

type controlPlaneMetricsClient struct {
	grpc.ClientStream
}

func (x *controlPlaneMetricsClient) Send(m *Metric) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controlPlaneMetricsClient) CloseAndRecv() (*Metric, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Metric)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlPlaneClient) EntityStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/atlas.ControlPlane/EntityStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) EntityStatusMultiple(ctx context.Context, in *StatusRequestMulti, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/atlas.ControlPlane/EntityStatusMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlPlaneServer is the server API for ControlPlane service.
type ControlPlaneServer interface {
	// Bi-directional Commands stream
	//
	// Commands are sent from service to client and from client to server
	Commands(ControlPlane_CommandsServer) error
	// Bi-directional Data stream
	//
	// Some commands may be followed by data load. Be it logs, dumps, etc.
	DataChunks(ControlPlane_DataChunksServer) error
	// Metrics stream
	//
	// Some commands may be followed by metrics stream.
	Metrics(ControlPlane_MetricsServer) error
	// EntityStatus checks status of the entity on the server
	EntityStatus(context.Context, *StatusRequest) (*StatusReply, error)
	// EntityStatusMulti checks status of the multiple entities on server
	EntityStatusMultiple(context.Context, *StatusRequestMulti) (*StatusReply, error)
}

// UnimplementedControlPlaneServer can be embedded to have forward compatible implementations.
type UnimplementedControlPlaneServer struct {
}

func (*UnimplementedControlPlaneServer) Commands(srv ControlPlane_CommandsServer) error {
	return status.Errorf(codes.Unimplemented, "method Commands not implemented")
}
func (*UnimplementedControlPlaneServer) DataChunks(srv ControlPlane_DataChunksServer) error {
	return status.Errorf(codes.Unimplemented, "method DataChunks not implemented")
}
func (*UnimplementedControlPlaneServer) Metrics(srv ControlPlane_MetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}
func (*UnimplementedControlPlaneServer) EntityStatus(ctx context.Context, req *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityStatus not implemented")
}
func (*UnimplementedControlPlaneServer) EntityStatusMultiple(ctx context.Context, req *StatusRequestMulti) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityStatusMultiple not implemented")
}

func RegisterControlPlaneServer(s *grpc.Server, srv ControlPlaneServer) {
	s.RegisterService(&_ControlPlane_serviceDesc, srv)
}

func _ControlPlane_Commands_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControlPlaneServer).Commands(&controlPlaneCommandsServer{stream})
}

type ControlPlane_CommandsServer interface {
	Send(*Command) error
	Recv() (*Command, error)
	grpc.ServerStream
}

type controlPlaneCommandsServer struct {
	grpc.ServerStream
}

func (x *controlPlaneCommandsServer) Send(m *Command) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controlPlaneCommandsServer) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ControlPlane_DataChunks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControlPlaneServer).DataChunks(&controlPlaneDataChunksServer{stream})
}

type ControlPlane_DataChunksServer interface {
	Send(*DataChunk) error
	Recv() (*DataChunk, error)
	grpc.ServerStream
}

type controlPlaneDataChunksServer struct {
	grpc.ServerStream
}

func (x *controlPlaneDataChunksServer) Send(m *DataChunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controlPlaneDataChunksServer) Recv() (*DataChunk, error) {
	m := new(DataChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ControlPlane_Metrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControlPlaneServer).Metrics(&controlPlaneMetricsServer{stream})
}

type ControlPlane_MetricsServer interface {
	SendAndClose(*Metric) error
	Recv() (*Metric, error)
	grpc.ServerStream
}

type controlPlaneMetricsServer struct {
	grpc.ServerStream
}

func (x *controlPlaneMetricsServer) SendAndClose(m *Metric) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controlPlaneMetricsServer) Recv() (*Metric, error) {
	m := new(Metric)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ControlPlane_EntityStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).EntityStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.ControlPlane/EntityStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).EntityStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_EntityStatusMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequestMulti)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).EntityStatusMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.ControlPlane/EntityStatusMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).EntityStatusMultiple(ctx, req.(*StatusRequestMulti))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlPlane_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atlas.ControlPlane",
	HandlerType: (*ControlPlaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EntityStatus",
			Handler:    _ControlPlane_EntityStatus_Handler,
		},
		{
			MethodName: "EntityStatusMultiple",
			Handler:    _ControlPlane_EntityStatusMultiple_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Commands",
			Handler:       _ControlPlane_Commands_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DataChunks",
			Handler:       _ControlPlane_DataChunks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Metrics",
			Handler:       _ControlPlane_Metrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service_control_plane.proto",
}
