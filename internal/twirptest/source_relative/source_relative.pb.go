// Code generated by protoc-gen-go. DO NOT EDIT.
// source: source_relative.proto

package source_relative

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

type Msg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8a2a00319471c84, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Msg)(nil), "twirk.internal.twirktest.source_relative.Msg")
}

func init() { proto.RegisterFile("source_relative.proto", fileDescriptor_f8a2a00319471c84) }

var fileDescriptor_f8a2a00319471c84 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0xce, 0x2f, 0x2d,
	0x4a, 0x4e, 0x8d, 0x2f, 0x4a, 0xcd, 0x49, 0x2c, 0xc9, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xd2, 0x28, 0x29, 0xcf, 0x2c, 0x2a, 0xd0, 0xcb, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b,
	0xcc, 0xd1, 0x03, 0x73, 0x4b, 0x52, 0x8b, 0x4b, 0xf4, 0xd0, 0xd4, 0x2b, 0xb1, 0x72, 0x31, 0xfb,
	0x16, 0xa7, 0x1b, 0xe5, 0x72, 0x31, 0x07, 0x97, 0x25, 0x0b, 0xa5, 0x71, 0xb1, 0xf9, 0xa6, 0x96,
	0x64, 0xe4, 0xa7, 0x08, 0xe9, 0xea, 0x11, 0x6b, 0x84, 0x9e, 0x6f, 0x71, 0xba, 0x14, 0x69, 0xca,
	0x9d, 0xec, 0xa2, 0x6c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x4b,
	0xca, 0x33, 0x4b, 0x92, 0x33, 0x4a, 0xca, 0xf4, 0xc1, 0x9a, 0xf4, 0x61, 0x66, 0xe8, 0xc3, 0xcd,
	0xd0, 0x47, 0x33, 0x23, 0x89, 0x0d, 0xec, 0x4d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc,
	0x3b, 0x72, 0xb8, 0xff, 0x00, 0x00, 0x00,
}
