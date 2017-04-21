// Code generated by protoc-gen-go.
// source: proto/nested_proto/kitten2.proto
// DO NOT EDIT!

/*
Package kitten2 is a generated protocol buffer package.

It is generated from these files:
	proto/nested_proto/kitten2.proto

It has these top-level messages:
	Kitten
	GetKittenRequest
	GetKittenResponse
*/
package kitten2

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

type Breed int32

const (
	Breed_HALP         Breed = 0
	Breed_PERSIAN      Breed = 1
	Breed_BENGAL       Breed = 2
	Breed_RUSSIAN_BLUE Breed = 3
	Breed_CORNISH_REX  Breed = 4
)

var Breed_name = map[int32]string{
	0: "HALP",
	1: "PERSIAN",
	2: "BENGAL",
	3: "RUSSIAN_BLUE",
	4: "CORNISH_REX",
}
var Breed_value = map[string]int32{
	"HALP":         0,
	"PERSIAN":      1,
	"BENGAL":       2,
	"RUSSIAN_BLUE": 3,
	"CORNISH_REX":  4,
}

func (x Breed) String() string {
	return proto.EnumName(Breed_name, int32(x))
}
func (Breed) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Kitten struct {
	Name   string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Breed  Breed   `protobuf:"varint,2,opt,name=breed,enum=Breed" json:"breed,omitempty"`
	Weight float32 `protobuf:"fixed32,3,opt,name=weight" json:"weight,omitempty"`
	Age    int32   `protobuf:"varint,4,opt,name=age" json:"age,omitempty"`
}

func (m *Kitten) Reset()                    { *m = Kitten{} }
func (m *Kitten) String() string            { return proto.CompactTextString(m) }
func (*Kitten) ProtoMessage()               {}
func (*Kitten) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetKittenRequest struct {
	// Breed of kitten requested. If not given then the server
	// may provide a kitten of any breed.
	Breed Breed `protobuf:"varint,1,opt,name=breed,enum=Breed" json:"breed,omitempty"`
}

func (m *GetKittenRequest) Reset()                    { *m = GetKittenRequest{} }
func (m *GetKittenRequest) String() string            { return proto.CompactTextString(m) }
func (*GetKittenRequest) ProtoMessage()               {}
func (*GetKittenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetKittenResponse struct {
	Kitten *Kitten `protobuf:"bytes,1,opt,name=kitten" json:"kitten,omitempty"`
}

func (m *GetKittenResponse) Reset()                    { *m = GetKittenResponse{} }
func (m *GetKittenResponse) String() string            { return proto.CompactTextString(m) }
func (*GetKittenResponse) ProtoMessage()               {}
func (*GetKittenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetKittenResponse) GetKitten() *Kitten {
	if m != nil {
		return m.Kitten
	}
	return nil
}

func init() {
	proto.RegisterType((*Kitten)(nil), "Kitten")
	proto.RegisterType((*GetKittenRequest)(nil), "GetKittenRequest")
	proto.RegisterType((*GetKittenResponse)(nil), "GetKittenResponse")
	proto.RegisterEnum("Breed", Breed_name, Breed_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PetShop service

type PetShopClient interface {
	// Supplies one kitten to the caller.
	GetKitten(ctx context.Context, in *GetKittenRequest, opts ...grpc.CallOption) (*GetKittenResponse, error)
}

type petShopClient struct {
	cc *grpc.ClientConn
}

func NewPetShopClient(cc *grpc.ClientConn) PetShopClient {
	return &petShopClient{cc}
}

func (c *petShopClient) GetKitten(ctx context.Context, in *GetKittenRequest, opts ...grpc.CallOption) (*GetKittenResponse, error) {
	out := new(GetKittenResponse)
	err := grpc.Invoke(ctx, "/PetShop/GetKitten", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PetShop service

type PetShopServer interface {
	// Supplies one kitten to the caller.
	GetKitten(context.Context, *GetKittenRequest) (*GetKittenResponse, error)
}

func RegisterPetShopServer(s *grpc.Server, srv PetShopServer) {
	s.RegisterService(&_PetShop_serviceDesc, srv)
}

func _PetShop_GetKitten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKittenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetShopServer).GetKitten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PetShop/GetKitten",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetShopServer).GetKitten(ctx, req.(*GetKittenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PetShop_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PetShop",
	HandlerType: (*PetShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKitten",
			Handler:    _PetShop_GetKitten_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/nested_proto/kitten2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xf3, 0x40,
	0x10, 0xc5, 0xbf, 0x6d, 0xd3, 0xf4, 0xeb, 0x44, 0x74, 0x3b, 0x07, 0x09, 0x22, 0x18, 0x72, 0x0a,
	0x1e, 0x56, 0x8d, 0x5e, 0x3d, 0x34, 0x12, 0xda, 0x62, 0xad, 0x61, 0x43, 0xc1, 0x5b, 0x49, 0xed,
	0xd0, 0x14, 0xdb, 0x24, 0x36, 0x5b, 0xf4, 0xcf, 0x97, 0x6c, 0x82, 0x8a, 0x7a, 0x9b, 0xdf, 0xdb,
	0x7d, 0x6f, 0xdf, 0x32, 0xe0, 0x14, 0xbb, 0x5c, 0xe5, 0x17, 0x19, 0x95, 0x8a, 0x96, 0xf3, 0x1a,
	0x5e, 0xd6, 0x4a, 0x51, 0xe6, 0x0b, 0x4d, 0xee, 0x12, 0xcc, 0x7b, 0x2d, 0x20, 0x82, 0x91, 0x25,
	0x5b, 0xb2, 0x99, 0xc3, 0xbc, 0x9e, 0xd4, 0x33, 0x9e, 0x42, 0x67, 0xb1, 0x23, 0x5a, 0xda, 0x2d,
	0x87, 0x79, 0x87, 0xbe, 0x29, 0x82, 0x8a, 0x64, 0x2d, 0xe2, 0x31, 0x98, 0x6f, 0xb4, 0x5e, 0xa5,
	0xca, 0x6e, 0x3b, 0xcc, 0x6b, 0xc9, 0x86, 0x90, 0x43, 0x3b, 0x59, 0x91, 0x6d, 0x38, 0xcc, 0xeb,
	0xc8, 0x6a, 0x74, 0x2f, 0x81, 0x0f, 0x49, 0xd5, 0x0f, 0x49, 0x7a, 0xdd, 0x53, 0xa9, 0xbe, 0xb2,
	0xd9, 0x1f, 0xd9, 0xee, 0x0d, 0xf4, 0xbf, 0x39, 0xca, 0x22, 0xcf, 0x4a, 0xc2, 0x33, 0x30, 0xeb,
	0xf6, 0xda, 0x63, 0xf9, 0x5d, 0xd1, 0x5c, 0x68, 0xe4, 0xf3, 0x07, 0xe8, 0xe8, 0x14, 0xfc, 0x0f,
	0xc6, 0x68, 0x30, 0x89, 0xf8, 0x3f, 0xb4, 0xa0, 0x1b, 0x85, 0x32, 0x1e, 0x0f, 0xa6, 0x9c, 0x21,
	0x80, 0x19, 0x84, 0xd3, 0xe1, 0x60, 0xc2, 0x5b, 0xc8, 0xe1, 0x40, 0xce, 0xe2, 0xea, 0x60, 0x1e,
	0x4c, 0x66, 0x21, 0x6f, 0xe3, 0x11, 0x58, 0x77, 0x8f, 0x72, 0x3a, 0x8e, 0x47, 0x73, 0x19, 0x3e,
	0x71, 0xc3, 0xbf, 0x85, 0x6e, 0x44, 0x2a, 0x4e, 0xf3, 0x02, 0x7d, 0xe8, 0x7d, 0xf6, 0xc1, 0xbe,
	0xf8, 0xf9, 0x9b, 0x13, 0x14, 0xbf, 0xea, 0x06, 0x57, 0xe0, 0x66, 0xa4, 0x84, 0x4a, 0xf3, 0xfd,
	0x2a, 0x55, 0xdb, 0xe4, 0x39, 0x5d, 0x67, 0x24, 0x8a, 0x0d, 0x25, 0x25, 0x09, 0x7a, 0x4f, 0xb6,
	0xc5, 0x86, 0xca, 0xc0, 0xaa, 0x5d, 0x51, 0xb5, 0x8e, 0x85, 0xa9, 0xb7, 0x72, 0xfd, 0x11, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0x05, 0x3b, 0x46, 0xb9, 0x01, 0x00, 0x00,
}