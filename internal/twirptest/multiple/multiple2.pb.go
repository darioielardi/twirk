// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multiple2.proto

// test to make sure that multiple proto files in one package works

package multiple

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Msg2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg2) Reset()         { *m = Msg2{} }
func (m *Msg2) String() string { return proto.CompactTextString(m) }
func (*Msg2) ProtoMessage()    {}
func (*Msg2) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1bf26928564f54, []int{0}
}

func (m *Msg2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg2.Unmarshal(m, b)
}
func (m *Msg2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg2.Marshal(b, m, deterministic)
}
func (m *Msg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg2.Merge(m, src)
}
func (m *Msg2) XXX_Size() int {
	return xxx_messageInfo_Msg2.Size(m)
}
func (m *Msg2) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg2.DiscardUnknown(m)
}

var xxx_messageInfo_Msg2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Msg2)(nil), "twirk.internal.twirktest.multiple.Msg2")
}

func init() { proto.RegisterFile("multiple2.proto", fileDescriptor_2f1bf26928564f54) }

var fileDescriptor_2f1bf26928564f54 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2d, 0xcd, 0x29,
	0xc9, 0x2c, 0xc8, 0x49, 0x35, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2c, 0x29, 0xcf,
	0x2c, 0x2a, 0xd0, 0xcb, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x03, 0x73, 0x4b, 0x52,
	0x8b, 0x4b, 0xf4, 0x60, 0x2a, 0xa5, 0xe0, 0x7a, 0x0c, 0x21, 0x7a, 0x94, 0xd8, 0xb8, 0x58, 0x7c,
	0x8b, 0xd3, 0x8d, 0x8c, 0xce, 0x30, 0x72, 0xb1, 0x04, 0x97, 0x25, 0x1b, 0x09, 0x45, 0x70, 0xb1,
	0x04, 0xa7, 0xe6, 0xa5, 0x08, 0xa9, 0xeb, 0x11, 0x34, 0x4d, 0x0f, 0xa4, 0x53, 0x8a, 0x58, 0x85,
	0x42, 0x59, 0x5c, 0x62, 0xc1, 0x89, 0xb9, 0xa9, 0x01, 0x89, 0xc9, 0xd9, 0x89, 0xe9, 0xa9, 0x01,
	0x20, 0xeb, 0x3d, 0x73, 0x0b, 0xf2, 0x8b, 0x4a, 0x88, 0xb5, 0xcb, 0x90, 0x58, 0xbb, 0x0c, 0x9d,
	0xb8, 0xa2, 0x38, 0x60, 0x02, 0x49, 0x6c, 0x60, 0x9f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0x66, 0x12, 0x3c, 0x30, 0x01, 0x00, 0x00,
}
