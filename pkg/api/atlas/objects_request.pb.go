// Code generated by protoc-gen-go. DO NOT EDIT.
// source: objects_request.proto

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

type ObjectsRequest struct {
	// Types that are valid to be assigned to RequestDomainOptional:
	//	*ObjectsRequest_RequestDomain
	RequestDomainOptional isObjectsRequest_RequestDomainOptional `protobuf_oneof:"request_domain_optional"`
	// Types that are valid to be assigned to ResultDomainOptional:
	//	*ObjectsRequest_ResultDomain
	ResultDomainOptional isObjectsRequest_ResultDomainOptional `protobuf_oneof:"result_domain_optional"`
	Requests             []*ObjectRequest                      `protobuf:"bytes,200,rep,name=requests,proto3" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ObjectsRequest) Reset()      { *m = ObjectsRequest{} }
func (*ObjectsRequest) ProtoMessage() {}
func (*ObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aa47e742b67e756, []int{0}
}

func (m *ObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectsRequest.Unmarshal(m, b)
}
func (m *ObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectsRequest.Marshal(b, m, deterministic)
}
func (m *ObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectsRequest.Merge(m, src)
}
func (m *ObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ObjectsRequest.Size(m)
}
func (m *ObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectsRequest proto.InternalMessageInfo

type isObjectsRequest_RequestDomainOptional interface {
	isObjectsRequest_RequestDomainOptional()
}

type ObjectsRequest_RequestDomain struct {
	RequestDomain *Domain `protobuf:"bytes,100,opt,name=request_domain,json=requestDomain,proto3,oneof"`
}

func (*ObjectsRequest_RequestDomain) isObjectsRequest_RequestDomainOptional() {}

func (m *ObjectsRequest) GetRequestDomainOptional() isObjectsRequest_RequestDomainOptional {
	if m != nil {
		return m.RequestDomainOptional
	}
	return nil
}

func (m *ObjectsRequest) GetRequestDomain() *Domain {
	if x, ok := m.GetRequestDomainOptional().(*ObjectsRequest_RequestDomain); ok {
		return x.RequestDomain
	}
	return nil
}

type isObjectsRequest_ResultDomainOptional interface {
	isObjectsRequest_ResultDomainOptional()
}

type ObjectsRequest_ResultDomain struct {
	ResultDomain *Domain `protobuf:"bytes,300,opt,name=result_domain,json=resultDomain,proto3,oneof"`
}

func (*ObjectsRequest_ResultDomain) isObjectsRequest_ResultDomainOptional() {}

func (m *ObjectsRequest) GetResultDomainOptional() isObjectsRequest_ResultDomainOptional {
	if m != nil {
		return m.ResultDomainOptional
	}
	return nil
}

func (m *ObjectsRequest) GetResultDomain() *Domain {
	if x, ok := m.GetResultDomainOptional().(*ObjectsRequest_ResultDomain); ok {
		return x.ResultDomain
	}
	return nil
}

func (m *ObjectsRequest) GetRequests() []*ObjectRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ObjectsRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ObjectsRequest_RequestDomain)(nil),
		(*ObjectsRequest_ResultDomain)(nil),
	}
}

func init() {
	proto.RegisterType((*ObjectsRequest)(nil), "atlas.ObjectsRequest")
}

func init() { proto.RegisterFile("objects_request.proto", fileDescriptor_5aa47e742b67e756) }

var fileDescriptor_5aa47e742b67e756 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4f, 0xca, 0x4a,
	0x4d, 0x2e, 0x29, 0x8e, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0xe2, 0x49, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc,
	0x83, 0x08, 0x4a, 0x89, 0x40, 0xd4, 0xa2, 0x2a, 0x55, 0x7a, 0xce, 0xc8, 0xc5, 0xe7, 0x0f, 0x31,
	0x24, 0x08, 0x22, 0x21, 0x64, 0xc6, 0xc5, 0x07, 0x55, 0x13, 0x0f, 0x31, 0x40, 0x22, 0x45, 0x81,
	0x51, 0x83, 0xdb, 0x88, 0x57, 0x0f, 0x6c, 0xac, 0x9e, 0x0b, 0x58, 0xd0, 0x83, 0x21, 0x88, 0x17,
	0xaa, 0x0c, 0x22, 0x20, 0x64, 0xca, 0xc5, 0x5b, 0x94, 0x5a, 0x5c, 0x9a, 0x03, 0xd7, 0xb6, 0x86,
	0x09, 0x9b, 0x3e, 0xc6, 0x20, 0x1e, 0x88, 0x32, 0xa8, 0x36, 0x43, 0x2e, 0x0e, 0xa8, 0x39, 0xc5,
	0x12, 0x27, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x44, 0xa0, 0x3a, 0x20, 0x0e, 0x83, 0xba, 0x2b,
	0x08, 0xae, 0xcc, 0x49, 0x92, 0x4b, 0x1c, 0xd5, 0x85, 0xf1, 0xf9, 0x05, 0x25, 0x99, 0xf9, 0x79,
	0x89, 0x39, 0x4e, 0x12, 0x5c, 0x62, 0x28, 0x8e, 0x80, 0xcb, 0x24, 0xb1, 0x81, 0x3d, 0x6c, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x53, 0xb6, 0x05, 0x34, 0x01, 0x00, 0x00,
}
