// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package pb

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

type Hi struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hi) Reset()         { *m = Hi{} }
func (m *Hi) String() string { return proto.CompactTextString(m) }
func (*Hi) ProtoMessage()    {}
func (*Hi) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *Hi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hi.Unmarshal(m, b)
}
func (m *Hi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hi.Marshal(b, m, deterministic)
}
func (m *Hi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hi.Merge(m, src)
}
func (m *Hi) XXX_Size() int {
	return xxx_messageInfo_Hi.Size(m)
}
func (m *Hi) XXX_DiscardUnknown() {
	xxx_messageInfo_Hi.DiscardUnknown(m)
}

var xxx_messageInfo_Hi proto.InternalMessageInfo

func (m *Hi) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hi) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Bye struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bye) Reset()         { *m = Bye{} }
func (m *Bye) String() string { return proto.CompactTextString(m) }
func (*Bye) ProtoMessage()    {}
func (*Bye) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *Bye) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bye.Unmarshal(m, b)
}
func (m *Bye) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bye.Marshal(b, m, deterministic)
}
func (m *Bye) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bye.Merge(m, src)
}
func (m *Bye) XXX_Size() int {
	return xxx_messageInfo_Bye.Size(m)
}
func (m *Bye) XXX_DiscardUnknown() {
	xxx_messageInfo_Bye.DiscardUnknown(m)
}

var xxx_messageInfo_Bye proto.InternalMessageInfo

func (m *Bye) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Bye) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Hi)(nil), "pb.Hi")
	proto.RegisterType((*Bye)(nil), "pb.Bye")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x52, 0xe3, 0x62, 0xf2, 0xc8, 0x14, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x02,
	0x0b, 0x80, 0x98, 0x4a, 0xea, 0x5c, 0xcc, 0x4e, 0x95, 0xa9, 0x84, 0x15, 0x1a, 0x69, 0x73, 0xf1,
	0x78, 0x80, 0x2c, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x92, 0xe6, 0xe2, 0x08, 0x4e,
	0xac, 0x04, 0x0b, 0x09, 0xb1, 0xe9, 0x15, 0x24, 0xe9, 0x79, 0x64, 0x4a, 0xb1, 0x83, 0x68, 0xa7,
	0xca, 0xd4, 0x24, 0x36, 0xb0, 0x43, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x6a, 0xfe,
	0x48, 0x9c, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *Hi, opts ...grpc.CallOption) (*Bye, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *Hi, opts ...grpc.CallOption) (*Bye, error) {
	out := new(Bye)
	err := c.cc.Invoke(ctx, "/pb.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	SayHello(context.Context, *Hi) (*Bye, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hi)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*Hi))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
