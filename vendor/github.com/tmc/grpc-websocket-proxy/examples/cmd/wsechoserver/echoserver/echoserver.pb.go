// Code generated by protoc-gen-go.
// source: echoserver.proto
// DO NOT EDIT!

/*
Package echoserver is a generated protocol buffer package.

It is generated from these files:
	echoserver.proto

It has these top-level messages:
	EchoRequest
	EchoResponse
	Heartbeat
	Empty
*/
package echoserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Heartbeat_Status int32

const (
	Heartbeat_UNKNOWN Heartbeat_Status = 0
	Heartbeat_OK      Heartbeat_Status = 1
)

var Heartbeat_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
}
var Heartbeat_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"OK":      1,
}

func (x Heartbeat_Status) String() string {
	return proto.EnumName(Heartbeat_Status_name, int32(x))
}
func (Heartbeat_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type EchoRequest struct {
	Message string `protobuf:"bytes,1,opt,name=Message,json=message" json:"Message,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EchoResponse struct {
	Message string `protobuf:"bytes,1,opt,name=Message,json=message" json:"Message,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Heartbeat struct {
	Status Heartbeat_Status `protobuf:"varint,1,opt,name=status,enum=echoserver.Heartbeat_Status" json:"status,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*EchoRequest)(nil), "echoserver.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "echoserver.EchoResponse")
	proto.RegisterType((*Heartbeat)(nil), "echoserver.Heartbeat")
	proto.RegisterType((*Empty)(nil), "echoserver.Empty")
	proto.RegisterEnum("echoserver.Heartbeat_Status", Heartbeat_Status_name, Heartbeat_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for EchoService service

type EchoServiceClient interface {
	Echo(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClient, error)
	Stream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (EchoService_StreamClient, error)
	Heartbeats(ctx context.Context, opts ...grpc.CallOption) (EchoService_HeartbeatsClient, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoService_serviceDesc.Streams[0], c.cc, "/echoserver.EchoService/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoClient{stream}
	return x, nil
}

type EchoService_EchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServiceEchoClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) Stream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (EchoService_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoService_serviceDesc.Streams[1], c.cc, "/echoserver.EchoService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoService_StreamClient interface {
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServiceStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceStreamClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) Heartbeats(ctx context.Context, opts ...grpc.CallOption) (EchoService_HeartbeatsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoService_serviceDesc.Streams[2], c.cc, "/echoserver.EchoService/Heartbeats", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceHeartbeatsClient{stream}
	return x, nil
}

type EchoService_HeartbeatsClient interface {
	Send(*Empty) error
	Recv() (*Heartbeat, error)
	grpc.ClientStream
}

type echoServiceHeartbeatsClient struct {
	grpc.ClientStream
}

func (x *echoServiceHeartbeatsClient) Send(m *Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceHeartbeatsClient) Recv() (*Heartbeat, error) {
	m := new(Heartbeat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Echo(EchoService_EchoServer) error
	Stream(*Empty, EchoService_StreamServer) error
	Heartbeats(EchoService_HeartbeatsServer) error
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).Echo(&echoServiceEchoServer{stream})
}

type EchoService_EchoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoServiceEchoServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EchoService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServiceServer).Stream(m, &echoServiceStreamServer{stream})
}

type EchoService_StreamServer interface {
	Send(*EchoResponse) error
	grpc.ServerStream
}

type echoServiceStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceStreamServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EchoService_Heartbeats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).Heartbeats(&echoServiceHeartbeatsServer{stream})
}

type EchoService_HeartbeatsServer interface {
	Send(*Heartbeat) error
	Recv() (*Empty, error)
	grpc.ServerStream
}

type echoServiceHeartbeatsServer struct {
	grpc.ServerStream
}

func (x *echoServiceHeartbeatsServer) Send(m *Heartbeat) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceHeartbeatsServer) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echoserver.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _EchoService_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Stream",
			Handler:       _EchoService_Stream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Heartbeats",
			Handler:       _EchoService_Heartbeats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("echoserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0x62, 0x13, 0x3a, 0xb1, 0x12, 0x57, 0xc4, 0x50, 0x2a, 0xc8, 0x5e, 0x0c, 0x1e,
	0x12, 0xa9, 0x9e, 0xbc, 0x57, 0x84, 0x62, 0x02, 0x29, 0xe2, 0xd5, 0x6d, 0x19, 0xd2, 0x80, 0xc9,
	0xc6, 0xec, 0x56, 0xf0, 0xea, 0x2b, 0xf8, 0x68, 0xbe, 0x82, 0xef, 0xe0, 0xd5, 0x64, 0x57, 0x62,
	0xc4, 0xd2, 0xe3, 0x4c, 0xbe, 0xf9, 0xe6, 0x9f, 0x2c, 0xb8, 0xb8, 0x5c, 0x09, 0x89, 0xd5, 0x0b,
	0x56, 0x41, 0x59, 0x09, 0x25, 0x28, 0xfc, 0x76, 0x46, 0xe3, 0x54, 0x88, 0xf4, 0x09, 0x43, 0x5e,
	0x66, 0x21, 0x2f, 0x0a, 0xa1, 0xb8, 0xca, 0x44, 0x21, 0x0d, 0xc9, 0xce, 0xc0, 0x99, 0xd6, 0x6c,
	0x82, 0xcf, 0x6b, 0x94, 0x8a, 0x7a, 0x60, 0xdf, 0xa1, 0x94, 0x3c, 0x45, 0x8f, 0x9c, 0x12, 0x7f,
	0x90, 0xd8, 0xb9, 0x29, 0x99, 0x0f, 0x7b, 0x06, 0x94, 0x65, 0x3d, 0x8d, 0x5b, 0xc8, 0x47, 0x18,
	0xdc, 0x22, 0xaf, 0xd4, 0x02, 0xb9, 0xa2, 0x57, 0x60, 0xc9, 0x7a, 0xe3, 0x5a, 0x6a, 0x6a, 0x7f,
	0x32, 0x0e, 0x3a, 0x61, 0x5b, 0x2c, 0x98, 0x6b, 0x26, 0xf9, 0x61, 0xd9, 0x09, 0x58, 0xa6, 0x43,
	0x1d, 0xb0, 0xef, 0xa3, 0x59, 0x14, 0x3f, 0x44, 0xee, 0x0e, 0xb5, 0xa0, 0x17, 0xcf, 0x5c, 0xc2,
	0x6c, 0xe8, 0x4f, 0xf3, 0x52, 0xbd, 0x4e, 0xbe, 0x88, 0x89, 0x3f, 0xaf, 0x7d, 0xd9, 0x12, 0x69,
	0x0c, 0xbb, 0x4d, 0x49, 0x8f, 0xbb, 0x5b, 0x3a, 0xf7, 0x8d, 0xbc, 0xff, 0x1f, 0xcc, 0x3d, 0xcc,
	0x7d, 0xfb, 0xf8, 0x7c, 0xef, 0x01, 0xeb, 0x87, 0x0d, 0x71, 0x4d, 0xce, 0x7d, 0x72, 0x41, 0xe8,
	0x4d, 0x13, 0xa4, 0x42, 0x9e, 0xd3, 0x83, 0x3f, 0x93, 0xcd, 0xf6, 0x2d, 0xb2, 0xa1, 0x96, 0xd9,
	0xd4, 0xc8, 0x6a, 0x4f, 0x0c, 0xd0, 0x1e, 0x2b, 0x37, 0xb9, 0x8e, 0x36, 0xfe, 0x17, 0x76, 0xa8,
	0x45, 0x43, 0xe6, 0x84, 0xab, 0x76, 0xbc, 0x09, 0xb6, 0xb0, 0xf4, 0xf3, 0x5d, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xad, 0x4a, 0x13, 0x1e, 0xfc, 0x01, 0x00, 0x00,
}
