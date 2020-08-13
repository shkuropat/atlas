// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_uuid.proto

package atlas

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UUID struct {
	// May contain any arbitrary sequence of bytes no longer than 2^32
	Data                 []byte   `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}
func (*UUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_999f4f27d257c4c8, []int{0}
}

func (m *UUID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUID.Unmarshal(m, b)
}
func (m *UUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUID.Marshal(b, m, deterministic)
}
func (m *UUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUID.Merge(m, src)
}
func (m *UUID) XXX_Size() int {
	return xxx_messageInfo_UUID.Size(m)
}
func (m *UUID) XXX_DiscardUnknown() {
	xxx_messageInfo_UUID.DiscardUnknown(m)
}

var xxx_messageInfo_UUID proto.InternalMessageInfo

func (m *UUID) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UUID)(nil), "atlas.UUID")
}

func init() {
	proto.RegisterFile("type_uuid.proto", fileDescriptor_999f4f27d257c4c8)
}

var fileDescriptor_999f4f27d257c4c8 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xa9, 0x2c, 0x48,
	0x8d, 0x2f, 0x2d, 0xcd, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9,
	0x49, 0x2c, 0x56, 0x92, 0xe2, 0x62, 0x09, 0x0d, 0xf5, 0x74, 0x11, 0x12, 0xe2, 0x62, 0x49, 0x49,
	0x2c, 0x49, 0x94, 0x48, 0x51, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x93, 0xd8, 0xc0, 0x2a, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xe2, 0xc9, 0xc4, 0x3c, 0x00, 0x00, 0x00,
}
