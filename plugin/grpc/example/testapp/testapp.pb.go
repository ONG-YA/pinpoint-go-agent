// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testapp.proto

package testapp

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

type Greeting struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_98d4e818d9f182b1, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "Greeting")
}

func init() {
	proto.RegisterFile("testapp.proto", fileDescriptor_98d4e818d9f182b1)
}

var fileDescriptor_98d4e818d9f182b1 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x49, 0x2d, 0x2e,
	0x49, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe1, 0xe2, 0x70, 0x2f, 0x4a,
	0x4d, 0x2d, 0xc9, 0xcc, 0x4b, 0x17, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x8d, 0x2e, 0x33, 0x72, 0xb1, 0x7a, 0xa4, 0xe6, 0xe4, 0xe4, 0x0b,
	0xe9, 0x71, 0x89, 0x84, 0xe6, 0x25, 0x16, 0x55, 0x3a, 0x27, 0xe6, 0xe4, 0x80, 0x19, 0x41, 0xa9,
	0x25, 0xa5, 0x45, 0x79, 0x42, 0x9c, 0x7a, 0x30, 0xed, 0x52, 0x08, 0xa6, 0x12, 0x83, 0x90, 0x21,
	0x97, 0x28, 0x5c, 0x7d, 0x70, 0x49, 0x51, 0x6a, 0x62, 0x2e, 0x7e, 0x0d, 0x06, 0x8c, 0x20, 0x2d,
	0x10, 0x95, 0x44, 0xda, 0xa1, 0xc1, 0x28, 0x64, 0xc2, 0x25, 0x86, 0xd0, 0x42, 0x8c, 0x35, 0x1a,
	0x8c, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0xaf, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x80,
	0x6f, 0xce, 0x0b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	UnaryCallUnaryReturn(ctx context.Context, in *Greeting, opts ...grpc.CallOption) (*Greeting, error)
	UnaryCallStreamReturn(ctx context.Context, in *Greeting, opts ...grpc.CallOption) (Hello_UnaryCallStreamReturnClient, error)
	StreamCallUnaryReturn(ctx context.Context, opts ...grpc.CallOption) (Hello_StreamCallUnaryReturnClient, error)
	StreamCallStreamReturn(ctx context.Context, opts ...grpc.CallOption) (Hello_StreamCallStreamReturnClient, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) UnaryCallUnaryReturn(ctx context.Context, in *Greeting, opts ...grpc.CallOption) (*Greeting, error) {
	out := new(Greeting)
	err := c.cc.Invoke(ctx, "/Hello/UnaryCallUnaryReturn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) UnaryCallStreamReturn(ctx context.Context, in *Greeting, opts ...grpc.CallOption) (Hello_UnaryCallStreamReturnClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[0], "/Hello/UnaryCallStreamReturn", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloUnaryCallStreamReturnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hello_UnaryCallStreamReturnClient interface {
	Recv() (*Greeting, error)
	grpc.ClientStream
}

type helloUnaryCallStreamReturnClient struct {
	grpc.ClientStream
}

func (x *helloUnaryCallStreamReturnClient) Recv() (*Greeting, error) {
	m := new(Greeting)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) StreamCallUnaryReturn(ctx context.Context, opts ...grpc.CallOption) (Hello_StreamCallUnaryReturnClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[1], "/Hello/StreamCallUnaryReturn", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloStreamCallUnaryReturnClient{stream}
	return x, nil
}

type Hello_StreamCallUnaryReturnClient interface {
	Send(*Greeting) error
	CloseAndRecv() (*Greeting, error)
	grpc.ClientStream
}

type helloStreamCallUnaryReturnClient struct {
	grpc.ClientStream
}

func (x *helloStreamCallUnaryReturnClient) Send(m *Greeting) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloStreamCallUnaryReturnClient) CloseAndRecv() (*Greeting, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Greeting)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) StreamCallStreamReturn(ctx context.Context, opts ...grpc.CallOption) (Hello_StreamCallStreamReturnClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[2], "/Hello/StreamCallStreamReturn", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloStreamCallStreamReturnClient{stream}
	return x, nil
}

type Hello_StreamCallStreamReturnClient interface {
	Send(*Greeting) error
	Recv() (*Greeting, error)
	grpc.ClientStream
}

type helloStreamCallStreamReturnClient struct {
	grpc.ClientStream
}

func (x *helloStreamCallStreamReturnClient) Send(m *Greeting) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloStreamCallStreamReturnClient) Recv() (*Greeting, error) {
	m := new(Greeting)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	UnaryCallUnaryReturn(context.Context, *Greeting) (*Greeting, error)
	UnaryCallStreamReturn(*Greeting, Hello_UnaryCallStreamReturnServer) error
	StreamCallUnaryReturn(Hello_StreamCallUnaryReturnServer) error
	StreamCallStreamReturn(Hello_StreamCallStreamReturnServer) error
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) UnaryCallUnaryReturn(ctx context.Context, req *Greeting) (*Greeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCallUnaryReturn not implemented")
}
func (*UnimplementedHelloServer) UnaryCallStreamReturn(req *Greeting, srv Hello_UnaryCallStreamReturnServer) error {
	return status.Errorf(codes.Unimplemented, "method UnaryCallStreamReturn not implemented")
}
func (*UnimplementedHelloServer) StreamCallUnaryReturn(srv Hello_StreamCallUnaryReturnServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCallUnaryReturn not implemented")
}
func (*UnimplementedHelloServer) StreamCallStreamReturn(srv Hello_StreamCallStreamReturnServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCallStreamReturn not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_UnaryCallUnaryReturn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Greeting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).UnaryCallUnaryReturn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello/UnaryCallUnaryReturn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).UnaryCallUnaryReturn(ctx, req.(*Greeting))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_UnaryCallStreamReturn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Greeting)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServer).UnaryCallStreamReturn(m, &helloUnaryCallStreamReturnServer{stream})
}

type Hello_UnaryCallStreamReturnServer interface {
	Send(*Greeting) error
	grpc.ServerStream
}

type helloUnaryCallStreamReturnServer struct {
	grpc.ServerStream
}

func (x *helloUnaryCallStreamReturnServer) Send(m *Greeting) error {
	return x.ServerStream.SendMsg(m)
}

func _Hello_StreamCallUnaryReturn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).StreamCallUnaryReturn(&helloStreamCallUnaryReturnServer{stream})
}

type Hello_StreamCallUnaryReturnServer interface {
	SendAndClose(*Greeting) error
	Recv() (*Greeting, error)
	grpc.ServerStream
}

type helloStreamCallUnaryReturnServer struct {
	grpc.ServerStream
}

func (x *helloStreamCallUnaryReturnServer) SendAndClose(m *Greeting) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloStreamCallUnaryReturnServer) Recv() (*Greeting, error) {
	m := new(Greeting)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Hello_StreamCallStreamReturn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).StreamCallStreamReturn(&helloStreamCallStreamReturnServer{stream})
}

type Hello_StreamCallStreamReturnServer interface {
	Send(*Greeting) error
	Recv() (*Greeting, error)
	grpc.ServerStream
}

type helloStreamCallStreamReturnServer struct {
	grpc.ServerStream
}

func (x *helloStreamCallStreamReturnServer) Send(m *Greeting) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloStreamCallStreamReturnServer) Recv() (*Greeting, error) {
	m := new(Greeting)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCallUnaryReturn",
			Handler:    _Hello_UnaryCallUnaryReturn_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UnaryCallStreamReturn",
			Handler:       _Hello_UnaryCallStreamReturn_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamCallUnaryReturn",
			Handler:       _Hello_StreamCallUnaryReturn_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamCallStreamReturn",
			Handler:       _Hello_StreamCallStreamReturn_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "testapp.proto",
}