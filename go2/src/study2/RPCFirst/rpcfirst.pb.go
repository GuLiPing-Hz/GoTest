// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcfirst.proto

//指定包名

package RPCFirst

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

//请求结构类型
type ReqHello struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHello) Reset()         { *m = ReqHello{} }
func (m *ReqHello) String() string { return proto.CompactTextString(m) }
func (*ReqHello) ProtoMessage()    {}
func (*ReqHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4119f066d4fb2aa, []int{0}
}

func (m *ReqHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHello.Unmarshal(m, b)
}
func (m *ReqHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHello.Marshal(b, m, deterministic)
}
func (m *ReqHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHello.Merge(m, src)
}
func (m *ReqHello) XXX_Size() int {
	return xxx_messageInfo_ReqHello.Size(m)
}
func (m *ReqHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHello.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHello proto.InternalMessageInfo

func (m *ReqHello) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqHello) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

//返回结构类型
type RespHello struct {
	Hi                   string   `protobuf:"bytes,1,opt,name=hi,proto3" json:"hi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespHello) Reset()         { *m = RespHello{} }
func (m *RespHello) String() string { return proto.CompactTextString(m) }
func (*RespHello) ProtoMessage()    {}
func (*RespHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4119f066d4fb2aa, []int{1}
}

func (m *RespHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespHello.Unmarshal(m, b)
}
func (m *RespHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespHello.Marshal(b, m, deterministic)
}
func (m *RespHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespHello.Merge(m, src)
}
func (m *RespHello) XXX_Size() int {
	return xxx_messageInfo_RespHello.Size(m)
}
func (m *RespHello) XXX_DiscardUnknown() {
	xxx_messageInfo_RespHello.DiscardUnknown(m)
}

var xxx_messageInfo_RespHello proto.InternalMessageInfo

func (m *RespHello) GetHi() string {
	if m != nil {
		return m.Hi
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqHello)(nil), "RPCFirst.ReqHello")
	proto.RegisterType((*RespHello)(nil), "RPCFirst.RespHello")
}

func init() { proto.RegisterFile("rpcfirst.proto", fileDescriptor_a4119f066d4fb2aa) }

var fileDescriptor_a4119f066d4fb2aa = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2a, 0x48, 0x4e,
	0xcb, 0x2c, 0x2a, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x08, 0x0a, 0x70, 0x76,
	0x03, 0xf1, 0x95, 0x0c, 0xb8, 0x38, 0x82, 0x52, 0x0b, 0x3d, 0x52, 0x73, 0x72, 0xf2, 0x85, 0x84,
	0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21,
	0x01, 0x2e, 0xe6, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x10, 0x53, 0x49,
	0x9a, 0x8b, 0x33, 0x28, 0xb5, 0xb8, 0x00, 0xa2, 0x85, 0x8f, 0x8b, 0x29, 0x23, 0x13, 0xaa, 0x81,
	0x29, 0x23, 0xd3, 0xe8, 0x3b, 0x23, 0x17, 0x8b, 0x4b, 0x6a, 0x6e, 0xbe, 0x90, 0x29, 0x17, 0x47,
	0x70, 0x62, 0x25, 0xd4, 0x5c, 0x3d, 0x98, 0x75, 0x7a, 0x30, 0xbb, 0xa4, 0x84, 0x91, 0xc5, 0xa0,
	0xa6, 0x29, 0x31, 0x08, 0xd9, 0x70, 0xf1, 0xfa, 0xe4, 0x97, 0x14, 0xfb, 0xa7, 0x05, 0xa5, 0x16,
	0xe4, 0x64, 0xa6, 0x16, 0x93, 0xa0, 0xd7, 0x80, 0x51, 0xc8, 0x8e, 0x8b, 0x1f, 0xa2, 0xdb, 0xbd,
	0x28, 0x35, 0xb5, 0x24, 0x33, 0x2f, 0x9d, 0x14, 0xfd, 0x1a, 0x8c, 0x42, 0x56, 0x5c, 0x9c, 0x4e,
	0x99, 0x29, 0x99, 0xa4, 0xba, 0x5a, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0xb2, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x41, 0x12, 0x54, 0x6b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoClient is the client API for Demo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoClient interface {
	//RPC普通方法，一次调用一次返回
	SayHello(ctx context.Context, in *ReqHello, opts ...grpc.CallOption) (*RespHello, error)
	//RPC 一次请求，流式返回
	LotsOfReplies(ctx context.Context, in *ReqHello, opts ...grpc.CallOption) (Demo_LotsOfRepliesClient, error)
	//RPC 流式请求，一次返回
	LotsOfGreetings(ctx context.Context, opts ...grpc.CallOption) (Demo_LotsOfGreetingsClient, error)
	//RPC 流式请求，流式返回
	BidiHello(ctx context.Context, opts ...grpc.CallOption) (Demo_BidiHelloClient, error)
}

type demoClient struct {
	cc *grpc.ClientConn
}

func NewDemoClient(cc *grpc.ClientConn) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) SayHello(ctx context.Context, in *ReqHello, opts ...grpc.CallOption) (*RespHello, error) {
	out := new(RespHello)
	err := c.cc.Invoke(ctx, "/RPCFirst.Demo/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) LotsOfReplies(ctx context.Context, in *ReqHello, opts ...grpc.CallOption) (Demo_LotsOfRepliesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Demo_serviceDesc.Streams[0], "/RPCFirst.Demo/LotsOfReplies", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoLotsOfRepliesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Demo_LotsOfRepliesClient interface {
	Recv() (*RespHello, error)
	grpc.ClientStream
}

type demoLotsOfRepliesClient struct {
	grpc.ClientStream
}

func (x *demoLotsOfRepliesClient) Recv() (*RespHello, error) {
	m := new(RespHello)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoClient) LotsOfGreetings(ctx context.Context, opts ...grpc.CallOption) (Demo_LotsOfGreetingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Demo_serviceDesc.Streams[1], "/RPCFirst.Demo/LotsOfGreetings", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoLotsOfGreetingsClient{stream}
	return x, nil
}

type Demo_LotsOfGreetingsClient interface {
	Send(*ReqHello) error
	CloseAndRecv() (*RespHello, error)
	grpc.ClientStream
}

type demoLotsOfGreetingsClient struct {
	grpc.ClientStream
}

func (x *demoLotsOfGreetingsClient) Send(m *ReqHello) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoLotsOfGreetingsClient) CloseAndRecv() (*RespHello, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RespHello)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoClient) BidiHello(ctx context.Context, opts ...grpc.CallOption) (Demo_BidiHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Demo_serviceDesc.Streams[2], "/RPCFirst.Demo/BidiHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoBidiHelloClient{stream}
	return x, nil
}

type Demo_BidiHelloClient interface {
	Send(*ReqHello) error
	Recv() (*RespHello, error)
	grpc.ClientStream
}

type demoBidiHelloClient struct {
	grpc.ClientStream
}

func (x *demoBidiHelloClient) Send(m *ReqHello) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoBidiHelloClient) Recv() (*RespHello, error) {
	m := new(RespHello)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemoServer is the server API for Demo service.
type DemoServer interface {
	//RPC普通方法，一次调用一次返回
	SayHello(context.Context, *ReqHello) (*RespHello, error)
	//RPC 一次请求，流式返回
	LotsOfReplies(*ReqHello, Demo_LotsOfRepliesServer) error
	//RPC 流式请求，一次返回
	LotsOfGreetings(Demo_LotsOfGreetingsServer) error
	//RPC 流式请求，流式返回
	BidiHello(Demo_BidiHelloServer) error
}

func RegisterDemoServer(s *grpc.Server, srv DemoServer) {
	s.RegisterService(&_Demo_serviceDesc, srv)
}

func _Demo_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqHello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPCFirst.Demo/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).SayHello(ctx, req.(*ReqHello))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_LotsOfReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReqHello)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServer).LotsOfReplies(m, &demoLotsOfRepliesServer{stream})
}

type Demo_LotsOfRepliesServer interface {
	Send(*RespHello) error
	grpc.ServerStream
}

type demoLotsOfRepliesServer struct {
	grpc.ServerStream
}

func (x *demoLotsOfRepliesServer) Send(m *RespHello) error {
	return x.ServerStream.SendMsg(m)
}

func _Demo_LotsOfGreetings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServer).LotsOfGreetings(&demoLotsOfGreetingsServer{stream})
}

type Demo_LotsOfGreetingsServer interface {
	SendAndClose(*RespHello) error
	Recv() (*ReqHello, error)
	grpc.ServerStream
}

type demoLotsOfGreetingsServer struct {
	grpc.ServerStream
}

func (x *demoLotsOfGreetingsServer) SendAndClose(m *RespHello) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoLotsOfGreetingsServer) Recv() (*ReqHello, error) {
	m := new(ReqHello)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Demo_BidiHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServer).BidiHello(&demoBidiHelloServer{stream})
}

type Demo_BidiHelloServer interface {
	Send(*RespHello) error
	Recv() (*ReqHello, error)
	grpc.ServerStream
}

type demoBidiHelloServer struct {
	grpc.ServerStream
}

func (x *demoBidiHelloServer) Send(m *RespHello) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoBidiHelloServer) Recv() (*ReqHello, error) {
	m := new(ReqHello)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Demo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RPCFirst.Demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Demo_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LotsOfReplies",
			Handler:       _Demo_LotsOfReplies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LotsOfGreetings",
			Handler:       _Demo_LotsOfGreetings_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidiHello",
			Handler:       _Demo_BidiHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpcfirst.proto",
}
