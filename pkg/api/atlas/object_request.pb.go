// Code generated by protoc-gen-go. DO NOT EDIT.
// source: object_request.proto

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

// ObjectRequest represents request for the object(s)
type ObjectRequest struct {
	// Types that are valid to be assigned to RequestDomainOptional:
	//	*ObjectRequest_RequestDomain
	RequestDomainOptional isObjectRequest_RequestDomainOptional `protobuf_oneof:"request_domain_optional"`
	// Types that are valid to be assigned to ResultDomainOptional:
	//	*ObjectRequest_ResultDomain
	ResultDomainOptional isObjectRequest_ResultDomainOptional `protobuf_oneof:"result_domain_optional"`
	// Address of the entity the request is made about
	Address              *Address `protobuf:"bytes,300,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectRequest) Reset()      { *m = ObjectRequest{} }
func (*ObjectRequest) ProtoMessage() {}
func (*ObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_075bd3731b60164f, []int{0}
}

func (m *ObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectRequest.Unmarshal(m, b)
}
func (m *ObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectRequest.Marshal(b, m, deterministic)
}
func (m *ObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectRequest.Merge(m, src)
}
func (m *ObjectRequest) XXX_Size() int {
	return xxx_messageInfo_ObjectRequest.Size(m)
}
func (m *ObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectRequest proto.InternalMessageInfo

type isObjectRequest_RequestDomainOptional interface {
	isObjectRequest_RequestDomainOptional()
}

type ObjectRequest_RequestDomain struct {
	RequestDomain *Domain `protobuf:"bytes,100,opt,name=request_domain,json=requestDomain,proto3,oneof"`
}

func (*ObjectRequest_RequestDomain) isObjectRequest_RequestDomainOptional() {}

func (m *ObjectRequest) GetRequestDomainOptional() isObjectRequest_RequestDomainOptional {
	if m != nil {
		return m.RequestDomainOptional
	}
	return nil
}

func (m *ObjectRequest) GetRequestDomain() *Domain {
	if x, ok := m.GetRequestDomainOptional().(*ObjectRequest_RequestDomain); ok {
		return x.RequestDomain
	}
	return nil
}

type isObjectRequest_ResultDomainOptional interface {
	isObjectRequest_ResultDomainOptional()
}

type ObjectRequest_ResultDomain struct {
	ResultDomain *Domain `protobuf:"bytes,200,opt,name=result_domain,json=resultDomain,proto3,oneof"`
}

func (*ObjectRequest_ResultDomain) isObjectRequest_ResultDomainOptional() {}

func (m *ObjectRequest) GetResultDomainOptional() isObjectRequest_ResultDomainOptional {
	if m != nil {
		return m.ResultDomainOptional
	}
	return nil
}

func (m *ObjectRequest) GetResultDomain() *Domain {
	if x, ok := m.GetResultDomainOptional().(*ObjectRequest_ResultDomain); ok {
		return x.ResultDomain
	}
	return nil
}

func (m *ObjectRequest) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ObjectRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ObjectRequest_RequestDomain)(nil),
		(*ObjectRequest_ResultDomain)(nil),
	}
}

func init() {
	proto.RegisterType((*ObjectRequest)(nil), "atlas.ObjectRequest")
}

func init() { proto.RegisterFile("object_request.proto", fileDescriptor_075bd3731b60164f) }

var fileDescriptor_075bd3731b60164f = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4f, 0xca, 0x4a,
	0x4d, 0x2e, 0x89, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0xe2, 0x49, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x83,
	0x08, 0x4a, 0xf1, 0x26, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x43, 0xb8, 0x4a, 0xf7, 0x18, 0xb9,
	0x78, 0xfd, 0xc1, 0x9a, 0x83, 0x20, 0x7a, 0x85, 0xcc, 0xb8, 0xf8, 0xa0, 0xc6, 0xc4, 0x43, 0x34,
	0x4a, 0xa4, 0x28, 0x30, 0x6a, 0x70, 0x1b, 0xf1, 0xea, 0x81, 0x8d, 0xd3, 0x73, 0x01, 0x0b, 0x7a,
	0x30, 0x04, 0xf1, 0x42, 0x95, 0x41, 0x04, 0x84, 0x4c, 0xb9, 0x78, 0x8b, 0x52, 0x8b, 0x4b, 0x73,
	0xe0, 0xda, 0x4e, 0x30, 0x62, 0xd3, 0xc7, 0x18, 0xc4, 0x03, 0x51, 0x06, 0xd5, 0xa6, 0xc9, 0xc5,
	0x0e, 0x75, 0x91, 0xc4, 0x1a, 0x26, 0xb0, 0x06, 0x3e, 0xa8, 0x06, 0x47, 0x88, 0x70, 0x10, 0x4c,
	0xde, 0x49, 0x92, 0x4b, 0x1c, 0xd5, 0x65, 0xf1, 0xf9, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x89, 0x39,
	0x4e, 0x12, 0x5c, 0x62, 0x28, 0x96, 0xc3, 0x65, 0x92, 0xd8, 0xc0, 0xfe, 0x34, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x20, 0x2c, 0x02, 0x2d, 0x23, 0x01, 0x00, 0x00,
}
