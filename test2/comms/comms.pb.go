// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comms.proto

package comms

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type Response struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "comms.Request")
	proto.RegisterType((*Response)(nil), "comms.Response")
}

func init() {
	proto.RegisterFile("comms.proto", fileDescriptor_db39efb7717b7d47)
}

var fileDescriptor_db39efb7717b7d47 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xce, 0xcf, 0xcd,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x54, 0xb9, 0xd8, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0xd2, 0x8b, 0x52, 0x53, 0x4b, 0x32,
	0xf3, 0xd2, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x35, 0x2e, 0x8e, 0xa0,
	0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x7c, 0xea, 0x8c, 0xcc, 0xb8, 0x58, 0x9d, 0x41, 0xe6,
	0x0a, 0xe9, 0x72, 0x71, 0x04, 0x27, 0x56, 0x7a, 0xa4, 0xe6, 0xe4, 0xe4, 0x0b, 0xf1, 0xe9, 0x41,
	0x2c, 0x86, 0x5a, 0x24, 0xc5, 0x0f, 0xe7, 0x43, 0x4c, 0x54, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xca,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x33, 0xad, 0x58, 0x77, 0xa3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommsClient is the client API for Comms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommsClient interface {
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type commsClient struct {
	cc grpc.ClientConnInterface
}

func NewCommsClient(cc grpc.ClientConnInterface) CommsClient {
	return &commsClient{cc}
}

func (c *commsClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/comms.Comms/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommsServer is the server API for Comms service.
type CommsServer interface {
	SayHello(context.Context, *Request) (*Response, error)
}

// UnimplementedCommsServer can be embedded to have forward compatible implementations.
type UnimplementedCommsServer struct {
}

func (*UnimplementedCommsServer) SayHello(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterCommsServer(s *grpc.Server, srv CommsServer) {
	s.RegisterService(&_Comms_serviceDesc, srv)
}

func _Comms_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Comms/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Comms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comms.Comms",
	HandlerType: (*CommsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Comms_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comms.proto",
}
