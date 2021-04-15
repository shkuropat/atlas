// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

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

type Task struct {
	// Header of the task
	Header *Metadata `protobuf:"bytes,100,opt,name=header,proto3" json:"header,omitempty"`
	// Optional. Any arbitrary sequence of bytes no longer than 2^32
	Bytes []byte `protobuf:"bytes,200,opt,name=bytes,proto3" json:"bytes,omitempty"`
	// Optional. Multiple task's subjects.
	Subjects []*Metadata `protobuf:"bytes,300,rep,name=subjects,proto3" json:"subjects,omitempty"`
	// Optional. Recursive chain of tasks
	Tasks                []*Task  `protobuf:"bytes,400,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()      { *m = Task{} }
func (*Task) ProtoMessage() {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetHeader() *Metadata {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Task) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *Task) GetSubjects() []*Metadata {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *Task) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "atlas.Task")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0xe2, 0xcb,
	0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x84, 0x08, 0x2b, 0xcd, 0x60, 0xe4, 0x62, 0x09, 0x49,
	0x2c, 0xce, 0x16, 0x52, 0xe7, 0x62, 0xcb, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x92, 0x48, 0x51, 0x60,
	0xd4, 0xe0, 0x36, 0xe2, 0xd7, 0x03, 0x6b, 0xd0, 0xf3, 0x85, 0xaa, 0x0f, 0x82, 0x4a, 0x0b, 0x89,
	0x72, 0xb1, 0x26, 0x55, 0x96, 0xa4, 0x16, 0x4b, 0x9c, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82,
	0xf0, 0x84, 0x74, 0xb8, 0x38, 0x8a, 0x4b, 0x93, 0xb2, 0x52, 0x93, 0x4b, 0x8a, 0x25, 0xd6, 0x30,
	0x29, 0x30, 0x63, 0x33, 0x02, 0xae, 0x42, 0x48, 0x89, 0x8b, 0x15, 0xe4, 0xb6, 0x62, 0x89, 0x09,
	0xcc, 0x60, 0xa5, 0xdc, 0x50, 0xa5, 0x20, 0xa7, 0x04, 0x41, 0xa4, 0x92, 0xd8, 0xc0, 0x2e, 0x34,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xaa, 0x94, 0x55, 0xc6, 0x00, 0x00, 0x00,
}