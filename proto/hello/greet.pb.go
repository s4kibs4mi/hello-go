// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello/greet.proto

package hello // import "github.com/msyrus/hello-go/proto/hello"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Greet struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greet) Reset()         { *m = Greet{} }
func (m *Greet) String() string { return proto.CompactTextString(m) }
func (*Greet) ProtoMessage()    {}
func (*Greet) Descriptor() ([]byte, []int) {
	return fileDescriptor_greet_f05e94bde0dbf4eb, []int{0}
}
func (m *Greet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greet.Unmarshal(m, b)
}
func (m *Greet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greet.Marshal(b, m, deterministic)
}
func (dst *Greet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greet.Merge(dst, src)
}
func (m *Greet) XXX_Size() int {
	return xxx_messageInfo_Greet.Size(m)
}
func (m *Greet) XXX_DiscardUnknown() {
	xxx_messageInfo_Greet.DiscardUnknown(m)
}

var xxx_messageInfo_Greet proto.InternalMessageInfo

func (m *Greet) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Greet)(nil), "hello.Greet")
}

func init() { proto.RegisterFile("proto/hello/greet.proto", fileDescriptor_greet_f05e94bde0dbf4eb) }

var fileDescriptor_greet_f05e94bde0dbf4eb = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0xd1, 0x03, 0x8b,
	0x08, 0xb1, 0x82, 0x85, 0x94, 0x14, 0xb9, 0x58, 0xdd, 0x41, 0xa2, 0x42, 0x12, 0x5c, 0xec, 0xb9,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x93,
	0x46, 0x94, 0x5a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x6e, 0x71,
	0x65, 0x51, 0x69, 0x31, 0xc4, 0x40, 0xdd, 0xf4, 0x7c, 0x7d, 0x24, 0xf3, 0x93, 0xd8, 0xc0, 0x1c,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x01, 0xbe, 0xdc, 0x75, 0x00, 0x00, 0x00,
}
