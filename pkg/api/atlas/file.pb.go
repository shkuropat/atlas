// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

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

// File represents abstract file.
type File struct {
	// Filename
	Filename *Filename `protobuf:"bytes,100,opt,name=filename,proto3" json:"filename,omitempty"`
	// Data
	Data                 []byte   `protobuf:"bytes,200,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()      { *m = File{} }
func (*File) ProtoMessage() {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetFilename() *Filename {
	if m != nil {
		return m.Filename
	}
	return nil
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*File)(nil), "atlas.File")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0xe2, 0x03,
	0x09, 0xe5, 0x25, 0xe6, 0x42, 0x85, 0x95, 0x3c, 0xb8, 0x58, 0xdc, 0x32, 0x73, 0x52, 0x85, 0xb4,
	0xb9, 0x38, 0x60, 0x32, 0x12, 0x29, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xfc, 0x7a, 0x60, 0x1d, 0x7a,
	0x6e, 0x50, 0xe1, 0x20, 0xb8, 0x02, 0x21, 0x61, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x89, 0x13,
	0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0x4e, 0x12, 0x1b, 0xd8, 0x40, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0xa2, 0xda, 0x63, 0x75, 0x00, 0x00, 0x00,
}