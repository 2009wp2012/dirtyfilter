// Code generated by protoc-gen-go.
// source: wordfilter.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	wordfilter.proto

It has these top-level messages:
	WordFilter
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type WordFilter struct {
}

func (m *WordFilter) Reset()                    { *m = WordFilter{} }
func (m *WordFilter) String() string            { return proto1.CompactTextString(m) }
func (*WordFilter) ProtoMessage()               {}
func (*WordFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type WordFilter_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *WordFilter_Text) Reset()                    { *m = WordFilter_Text{} }
func (m *WordFilter_Text) String() string            { return proto1.CompactTextString(m) }
func (*WordFilter_Text) ProtoMessage()               {}
func (*WordFilter_Text) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func init() {
	proto1.RegisterType((*WordFilter)(nil), "proto.WordFilter")
	proto1.RegisterType((*WordFilter_Text)(nil), "proto.WordFilter.Text")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for WordFilterService service

type WordFilterServiceClient interface {
	Filter(ctx context.Context, in *WordFilter_Text, opts ...grpc.CallOption) (*WordFilter_Text, error)
}

type wordFilterServiceClient struct {
	cc *grpc.ClientConn
}

func NewWordFilterServiceClient(cc *grpc.ClientConn) WordFilterServiceClient {
	return &wordFilterServiceClient{cc}
}

func (c *wordFilterServiceClient) Filter(ctx context.Context, in *WordFilter_Text, opts ...grpc.CallOption) (*WordFilter_Text, error) {
	out := new(WordFilter_Text)
	err := grpc.Invoke(ctx, "/proto.WordFilterService/Filter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WordFilterService service

type WordFilterServiceServer interface {
	Filter(context.Context, *WordFilter_Text) (*WordFilter_Text, error)
}

func RegisterWordFilterServiceServer(s *grpc.Server, srv WordFilterServiceServer) {
	s.RegisterService(&_WordFilterService_serviceDesc, srv)
}

func _WordFilterService_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WordFilter_Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordFilterServiceServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WordFilterService/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordFilterServiceServer).Filter(ctx, req.(*WordFilter_Text))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordFilterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WordFilterService",
	HandlerType: (*WordFilterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Filter",
			Handler:    _WordFilterService_Filter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("wordfilter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xcf, 0x2f, 0x4a,
	0x49, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x4a, 0x4a, 0x5c, 0x5c, 0xe1, 0xf9, 0x45, 0x29, 0x6e, 0x60, 0x29, 0x29, 0x11, 0x2e, 0x96, 0x90,
	0xd4, 0x8a, 0x12, 0x21, 0x1e, 0x2e, 0x96, 0x92, 0xd4, 0x8a, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x4e, 0x23, 0x7f, 0x2e, 0x41, 0x84, 0x9a, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x2b,
	0x2e, 0x36, 0x88, 0x80, 0x90, 0x18, 0xc4, 0x44, 0x3d, 0x84, 0x1a, 0x3d, 0x90, 0x21, 0x52, 0x38,
	0xc4, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x12, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x3e,
	0x91, 0xe4, 0x96, 0x00, 0x00, 0x00,
}