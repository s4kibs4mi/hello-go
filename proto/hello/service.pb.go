// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello/service.proto

package hello // import "github.com/msyrus/hello-go/proto/hello"

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

type ReqGreet struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqGreet) Reset()         { *m = ReqGreet{} }
func (m *ReqGreet) String() string { return proto.CompactTextString(m) }
func (*ReqGreet) ProtoMessage()    {}
func (*ReqGreet) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_0430b84f4a3a657f, []int{0}
}
func (m *ReqGreet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqGreet.Unmarshal(m, b)
}
func (m *ReqGreet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqGreet.Marshal(b, m, deterministic)
}
func (dst *ReqGreet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqGreet.Merge(dst, src)
}
func (m *ReqGreet) XXX_Size() int {
	return xxx_messageInfo_ReqGreet.Size(m)
}
func (m *ReqGreet) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqGreet.DiscardUnknown(m)
}

var xxx_messageInfo_ReqGreet proto.InternalMessageInfo

func (m *ReqGreet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqGreet)(nil), "hello.ReqGreet")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetingClient is the client API for Greeting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetingClient interface {
	DefaultGreeting(ctx context.Context, in *ReqGreet, opts ...grpc.CallOption) (*Greet, error)
}

type greetingClient struct {
	cc *grpc.ClientConn
}

func NewGreetingClient(cc *grpc.ClientConn) GreetingClient {
	return &greetingClient{cc}
}

func (c *greetingClient) DefaultGreeting(ctx context.Context, in *ReqGreet, opts ...grpc.CallOption) (*Greet, error) {
	out := new(Greet)
	err := c.cc.Invoke(ctx, "/hello.Greeting/DefaultGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServer is the server API for Greeting service.
type GreetingServer interface {
	DefaultGreeting(context.Context, *ReqGreet) (*Greet, error)
}

func RegisterGreetingServer(s *grpc.Server, srv GreetingServer) {
	s.RegisterService(&_Greeting_serviceDesc, srv)
}

func _Greeting_DefaultGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGreet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServer).DefaultGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Greeting/DefaultGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServer).DefaultGreeting(ctx, req.(*ReqGreet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Greeting",
	HandlerType: (*GreetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DefaultGreeting",
			Handler:    _Greeting_DefaultGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello/service.proto",
}

func init() { proto.RegisterFile("proto/hello/service.proto", fileDescriptor_service_0430b84f4a3a657f) }

var fileDescriptor_service_0430b84f4a3a657f = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x03, 0x8b, 0x09, 0xb1, 0x82, 0x05, 0xa5, 0xc4, 0x91, 0x55, 0xa4, 0x17, 0xa5, 0xa6, 0x96, 0x40,
	0xe4, 0x95, 0xe4, 0xb8, 0x38, 0x82, 0x52, 0x0b, 0xdd, 0x41, 0x22, 0x42, 0x42, 0x5c, 0x2c, 0x79,
	0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x91, 0x0d, 0x17, 0x07,
	0x58, 0x32, 0x33, 0x2f, 0x5d, 0xc8, 0x80, 0x8b, 0xdf, 0x25, 0x35, 0x2d, 0xb1, 0x34, 0xa7, 0x04,
	0x2e, 0xc4, 0xaf, 0x07, 0x36, 0x52, 0x0f, 0x66, 0x86, 0x14, 0x0f, 0x54, 0x00, 0xcc, 0x73, 0xd2,
	0x88, 0x52, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xae,
	0x2c, 0x2a, 0x2d, 0x86, 0x38, 0x42, 0x37, 0x3d, 0x5f, 0x1f, 0xc9, 0x4d, 0x49, 0x6c, 0x60, 0x8e,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x98, 0x51, 0xee, 0x3b, 0xcb, 0x00, 0x00, 0x00,
}
