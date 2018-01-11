// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	Void
	Message
	Session
*/
package grpc

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

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Session struct {
	Id     int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Session) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Session) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Void)(nil), "grpc.Void")
	proto.RegisterType((*Message)(nil), "grpc.Message")
	proto.RegisterType((*Session)(nil), "grpc.Session")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IronMan service

type IronmanClient interface {
	// RPC Calls for IronMan
	Ping(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type ironmanClient struct {
	cc *grpc.ClientConn
}

func NewIronmanClient(cc *grpc.ClientConn) IronmanClient {
	return &ironmanClient{cc}
}

func (c *ironmanClient) Ping(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/grpc.Ironman/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ironman service

type IronmanServer interface {
	// RPC Calls for Ironman
	Ping(context.Context, *Message) (*Message, error)
}

func RegisterIronmanServer(s *grpc.Server, srv IronmanServer) {
	s.RegisterService(&_Ironman_serviceDesc, srv)
}

func _Ironman_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IronmanServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Ironman/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IronmanServer).Ping(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ironman_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Ironman",
	HandlerType: (*IronmanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Ironman_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

// Client API for Batman service

type BatmanClient interface {
	IronmanCallback(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Void, error)
}

type batmanClient struct {
	cc *grpc.ClientConn
}

func NewBatmanClient(cc *grpc.ClientConn) BatmanClient {
	return &batmanClient{cc}
}

func (c *batmanClient) IronmanCallback(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/grpc.Batman/IronmanCallback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Batman service

type BatmanServer interface {
	IronmanCallback(context.Context, *Session) (*Void, error)
}

func RegisterBatmanServer(s *grpc.Server, srv BatmanServer) {
	s.RegisterService(&_Batman_serviceDesc, srv)
}

func _Batman_IronmanCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatmanServer).IronmanCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Batman/IronmanCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatmanServer).IronmanCallback(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

var _Batman_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Batman",
	HandlerType: (*BatmanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IronmanCallback",
			Handler:    _Batman_IronmanCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x6b, 0x83, 0x40,
	0x10, 0x85, 0xab, 0x6c, 0x57, 0x1c, 0xa8, 0x85, 0x39, 0x14, 0x11, 0x5a, 0xca, 0x1e, 0x4a, 0xa1,
	0x20, 0xad, 0xa5, 0x7f, 0xc0, 0x5e, 0x4a, 0xa1, 0x20, 0x16, 0x72, 0x1f, 0x75, 0x31, 0x4b, 0x56,
	0x37, 0xb8, 0x9b, 0x90, 0x9f, 0x1f, 0x5c, 0xcd, 0x21, 0xb7, 0xf7, 0x66, 0xf8, 0x78, 0xef, 0x01,
	0x6c, 0x49, 0xeb, 0x7c, 0x3f, 0x19, 0x67, 0x90, 0xcd, 0x5a, 0x70, 0x60, 0x1b, 0xa3, 0x3a, 0xf1,
	0x08, 0xd1, 0x9f, 0xb4, 0x96, 0x7a, 0x89, 0x08, 0xcc, 0xc9, 0x93, 0x4b, 0x83, 0xe7, 0xe0, 0x35,
	0xae, 0xbd, 0x16, 0x1f, 0x10, 0xfd, 0x4b, 0x6b, 0x95, 0x19, 0x31, 0x81, 0x50, 0x75, 0xfe, 0x79,
	0x5b, 0x87, 0xaa, 0xc3, 0x07, 0xe0, 0xd6, 0x91, 0x3b, 0xd8, 0x34, 0xf4, 0xc0, 0xea, 0x8a, 0x77,
	0xe0, 0xbf, 0x34, 0x1d, 0x95, 0xc5, 0x17, 0x60, 0x95, 0x1a, 0x7b, 0xbc, 0xcb, 0x7d, 0xfc, 0x9a,
	0x93, 0x5d, 0x5b, 0x71, 0x53, 0x7c, 0x01, 0x2f, 0xc9, 0x0d, 0x34, 0xe2, 0x1b, 0x24, 0x0b, 0xfb,
	0x4d, 0x5a, 0x37, 0xd4, 0xee, 0x2e, 0xec, 0x5a, 0x22, 0x83, 0xc5, 0xce, 0xd5, 0xcb, 0x27, 0xb8,
	0x6f, 0xcd, 0x90, 0x93, 0x9a, 0x5a, 0x67, 0xfc, 0xbd, 0x8c, 0x7f, 0x48, 0xeb, 0x6a, 0x9e, 0x59,
	0x05, 0x0d, 0xf7, 0x7b, 0x3f, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x4b, 0xbc, 0x9e, 0xfd,
	0x00, 0x00, 0x00,
}
