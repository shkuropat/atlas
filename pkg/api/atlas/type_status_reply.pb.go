// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_status_reply.proto

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

type StatusType int32

const (
	// Due to first enum value has to be zero in proto3
	StatusType_STATUS_RESERVED    StatusType = 0
	StatusType_STATUS_UNSPECIFIED StatusType = 100
	StatusType_STATUS_NOT_FOUND   StatusType = 200
	StatusType_STATUS_FOUND       StatusType = 300
)

var StatusType_name = map[int32]string{
	0:   "STATUS_RESERVED",
	100: "STATUS_UNSPECIFIED",
	200: "STATUS_NOT_FOUND",
	300: "STATUS_FOUND",
}

var StatusType_value = map[string]int32{
	"STATUS_RESERVED":    0,
	"STATUS_UNSPECIFIED": 100,
	"STATUS_NOT_FOUND":   200,
	"STATUS_FOUND":       300,
}

func (x StatusType) String() string {
	return proto.EnumName(StatusType_name, int32(x))
}

func (StatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_80bf857677dab50f, []int{0}
}

type StatusReply struct {
	Status               StatusType `protobuf:"varint,100,opt,name=status,proto3,enum=atlas.StatusType" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StatusReply) Reset()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()    {}
func (*StatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_80bf857677dab50f, []int{0}
}

func (m *StatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReply.Unmarshal(m, b)
}
func (m *StatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReply.Marshal(b, m, deterministic)
}
func (m *StatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReply.Merge(m, src)
}
func (m *StatusReply) XXX_Size() int {
	return xxx_messageInfo_StatusReply.Size(m)
}
func (m *StatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReply proto.InternalMessageInfo

func (m *StatusReply) GetStatus() StatusType {
	if m != nil {
		return m.Status
	}
	return StatusType_STATUS_RESERVED
}

func init() {
	proto.RegisterEnum("atlas.StatusType", StatusType_name, StatusType_value)
	proto.RegisterType((*StatusReply)(nil), "atlas.StatusReply")
}

func init() { proto.RegisterFile("type_status_reply.proto", fileDescriptor_80bf857677dab50f) }

var fileDescriptor_80bf857677dab50f = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xa9, 0x2c, 0x48,
	0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x8e, 0x2f, 0x4a, 0x2d, 0xc8, 0xa9, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x56, 0xb2, 0xe0, 0xe2, 0x0e, 0x06,
	0x4b, 0x06, 0x81, 0xe4, 0x84, 0x34, 0xb9, 0xd8, 0x20, 0x6a, 0x25, 0x52, 0x14, 0x18, 0x35, 0xf8,
	0x8c, 0x04, 0xf5, 0xc0, 0xca, 0xf4, 0x20, 0x6a, 0x42, 0x2a, 0x0b, 0x52, 0x83, 0xa0, 0x0a, 0xb4,
	0x92, 0xb9, 0xb8, 0x10, 0xa2, 0x42, 0xc2, 0x5c, 0xfc, 0xc1, 0x21, 0x8e, 0x21, 0xa1, 0xc1, 0xf1,
	0x41, 0xae, 0xc1, 0xae, 0x41, 0x61, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0x62, 0x5c, 0x42, 0x50, 0xc1,
	0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x81, 0x14, 0x21, 0x51, 0x2e,
	0x01, 0xa8, 0xb8, 0x9f, 0x7f, 0x48, 0xbc, 0x9b, 0x7f, 0xa8, 0x9f, 0x8b, 0xc0, 0x09, 0x46, 0x21,
	0x41, 0x2e, 0x1e, 0xa8, 0x30, 0x44, 0x68, 0x0d, 0x53, 0x12, 0x1b, 0xd8, 0xb1, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfe, 0x1a, 0xf9, 0x78, 0xc7, 0x00, 0x00, 0x00,
}
