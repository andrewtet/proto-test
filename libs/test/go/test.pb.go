// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_dbcb285394730f20, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Greeting             string   `protobuf:"bytes,2,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_dbcb285394730f20, []int{1}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "proto.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "proto.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.Hello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_dbcb285394730f20) }

var fileDescriptor_test_dbcb285394730f20 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0xa9,
	0x39, 0x39, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89,
	0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x2b, 0x17, 0x2f, 0x54,
	0x4d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x71, 0x69, 0x72, 0x72, 0x6a,
	0x71, 0x31, 0x58, 0x1d, 0x47, 0x10, 0x8c, 0x2b, 0x24, 0xc5, 0xc5, 0x91, 0x5e, 0x94, 0x9a, 0x5a,
	0x92, 0x99, 0x97, 0x2e, 0xc1, 0x04, 0x36, 0x02, 0xce, 0x37, 0xb2, 0xe3, 0x62, 0x05, 0x1b, 0x23,
	0x64, 0xca, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0x61, 0x0b, 0x43, 0x9c, 0xa3, 0x87, 0xec, 0x08, 0x29,
	0x11, 0x54, 0x41, 0x88, 0xad, 0x49, 0x6c, 0x60, 0x41, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x31, 0x72, 0x22, 0x8e, 0xc6, 0x00, 0x00, 0x00,
}