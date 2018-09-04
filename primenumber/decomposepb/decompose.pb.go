// Code generated by protoc-gen-go. DO NOT EDIT.
// source: primenumber/decomposepb/decompose.proto

package decomposepb

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

type Decompose struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decompose) Reset()         { *m = Decompose{} }
func (m *Decompose) String() string { return proto.CompactTextString(m) }
func (*Decompose) ProtoMessage()    {}
func (*Decompose) Descriptor() ([]byte, []int) {
	return fileDescriptor_decompose_ade7e6d362745697, []int{0}
}
func (m *Decompose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decompose.Unmarshal(m, b)
}
func (m *Decompose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decompose.Marshal(b, m, deterministic)
}
func (dst *Decompose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decompose.Merge(dst, src)
}
func (m *Decompose) XXX_Size() int {
	return xxx_messageInfo_Decompose.Size(m)
}
func (m *Decompose) XXX_DiscardUnknown() {
	xxx_messageInfo_Decompose.DiscardUnknown(m)
}

var xxx_messageInfo_Decompose proto.InternalMessageInfo

func (m *Decompose) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type DecomposeRequest struct {
	Request              *Decompose `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DecomposeRequest) Reset()         { *m = DecomposeRequest{} }
func (m *DecomposeRequest) String() string { return proto.CompactTextString(m) }
func (*DecomposeRequest) ProtoMessage()    {}
func (*DecomposeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_decompose_ade7e6d362745697, []int{1}
}
func (m *DecomposeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecomposeRequest.Unmarshal(m, b)
}
func (m *DecomposeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecomposeRequest.Marshal(b, m, deterministic)
}
func (dst *DecomposeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecomposeRequest.Merge(dst, src)
}
func (m *DecomposeRequest) XXX_Size() int {
	return xxx_messageInfo_DecomposeRequest.Size(m)
}
func (m *DecomposeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecomposeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecomposeRequest proto.InternalMessageInfo

func (m *DecomposeRequest) GetRequest() *Decompose {
	if m != nil {
		return m.Request
	}
	return nil
}

type DecomposeResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecomposeResponse) Reset()         { *m = DecomposeResponse{} }
func (m *DecomposeResponse) String() string { return proto.CompactTextString(m) }
func (*DecomposeResponse) ProtoMessage()    {}
func (*DecomposeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_decompose_ade7e6d362745697, []int{2}
}
func (m *DecomposeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecomposeResponse.Unmarshal(m, b)
}
func (m *DecomposeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecomposeResponse.Marshal(b, m, deterministic)
}
func (dst *DecomposeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecomposeResponse.Merge(dst, src)
}
func (m *DecomposeResponse) XXX_Size() int {
	return xxx_messageInfo_DecomposeResponse.Size(m)
}
func (m *DecomposeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecomposeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecomposeResponse proto.InternalMessageInfo

func (m *DecomposeResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type DecomposeManyTimesRequest struct {
	Request              *Decompose `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DecomposeManyTimesRequest) Reset()         { *m = DecomposeManyTimesRequest{} }
func (m *DecomposeManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*DecomposeManyTimesRequest) ProtoMessage()    {}
func (*DecomposeManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_decompose_ade7e6d362745697, []int{3}
}
func (m *DecomposeManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecomposeManyTimesRequest.Unmarshal(m, b)
}
func (m *DecomposeManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecomposeManyTimesRequest.Marshal(b, m, deterministic)
}
func (dst *DecomposeManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecomposeManyTimesRequest.Merge(dst, src)
}
func (m *DecomposeManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_DecomposeManyTimesRequest.Size(m)
}
func (m *DecomposeManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecomposeManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecomposeManyTimesRequest proto.InternalMessageInfo

func (m *DecomposeManyTimesRequest) GetRequest() *Decompose {
	if m != nil {
		return m.Request
	}
	return nil
}

type DecomposeManyTimesResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecomposeManyTimesResponse) Reset()         { *m = DecomposeManyTimesResponse{} }
func (m *DecomposeManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*DecomposeManyTimesResponse) ProtoMessage()    {}
func (*DecomposeManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_decompose_ade7e6d362745697, []int{4}
}
func (m *DecomposeManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecomposeManyTimesResponse.Unmarshal(m, b)
}
func (m *DecomposeManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecomposeManyTimesResponse.Marshal(b, m, deterministic)
}
func (dst *DecomposeManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecomposeManyTimesResponse.Merge(dst, src)
}
func (m *DecomposeManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_DecomposeManyTimesResponse.Size(m)
}
func (m *DecomposeManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecomposeManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecomposeManyTimesResponse proto.InternalMessageInfo

func (m *DecomposeManyTimesResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Decompose)(nil), "primenumber.Decompose")
	proto.RegisterType((*DecomposeRequest)(nil), "primenumber.DecomposeRequest")
	proto.RegisterType((*DecomposeResponse)(nil), "primenumber.DecomposeResponse")
	proto.RegisterType((*DecomposeManyTimesRequest)(nil), "primenumber.DecomposeManyTimesRequest")
	proto.RegisterType((*DecomposeManyTimesResponse)(nil), "primenumber.DecomposeManyTimesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DecomposeServiceClient is the client API for DecomposeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DecomposeServiceClient interface {
	// Unary
	Decompose(ctx context.Context, in *DecomposeRequest, opts ...grpc.CallOption) (*DecomposeResponse, error)
	DecomposeManyTimes(ctx context.Context, in *DecomposeManyTimesRequest, opts ...grpc.CallOption) (DecomposeService_DecomposeManyTimesClient, error)
}

type decomposeServiceClient struct {
	cc *grpc.ClientConn
}

func NewDecomposeServiceClient(cc *grpc.ClientConn) DecomposeServiceClient {
	return &decomposeServiceClient{cc}
}

func (c *decomposeServiceClient) Decompose(ctx context.Context, in *DecomposeRequest, opts ...grpc.CallOption) (*DecomposeResponse, error) {
	out := new(DecomposeResponse)
	err := c.cc.Invoke(ctx, "/primenumber.DecomposeService/Decompose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decomposeServiceClient) DecomposeManyTimes(ctx context.Context, in *DecomposeManyTimesRequest, opts ...grpc.CallOption) (DecomposeService_DecomposeManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DecomposeService_serviceDesc.Streams[0], "/primenumber.DecomposeService/DecomposeManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &decomposeServiceDecomposeManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DecomposeService_DecomposeManyTimesClient interface {
	Recv() (*DecomposeManyTimesResponse, error)
	grpc.ClientStream
}

type decomposeServiceDecomposeManyTimesClient struct {
	grpc.ClientStream
}

func (x *decomposeServiceDecomposeManyTimesClient) Recv() (*DecomposeManyTimesResponse, error) {
	m := new(DecomposeManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DecomposeServiceServer is the server API for DecomposeService service.
type DecomposeServiceServer interface {
	// Unary
	Decompose(context.Context, *DecomposeRequest) (*DecomposeResponse, error)
	DecomposeManyTimes(*DecomposeManyTimesRequest, DecomposeService_DecomposeManyTimesServer) error
}

func RegisterDecomposeServiceServer(s *grpc.Server, srv DecomposeServiceServer) {
	s.RegisterService(&_DecomposeService_serviceDesc, srv)
}

func _DecomposeService_Decompose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecomposeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecomposeServiceServer).Decompose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/primenumber.DecomposeService/Decompose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecomposeServiceServer).Decompose(ctx, req.(*DecomposeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecomposeService_DecomposeManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DecomposeManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DecomposeServiceServer).DecomposeManyTimes(m, &decomposeServiceDecomposeManyTimesServer{stream})
}

type DecomposeService_DecomposeManyTimesServer interface {
	Send(*DecomposeManyTimesResponse) error
	grpc.ServerStream
}

type decomposeServiceDecomposeManyTimesServer struct {
	grpc.ServerStream
}

func (x *decomposeServiceDecomposeManyTimesServer) Send(m *DecomposeManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _DecomposeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "primenumber.DecomposeService",
	HandlerType: (*DecomposeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Decompose",
			Handler:    _DecomposeService_Decompose_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DecomposeManyTimes",
			Handler:       _DecomposeService_DecomposeManyTimes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "primenumber/decomposepb/decompose.proto",
}

func init() {
	proto.RegisterFile("primenumber/decomposepb/decompose.proto", fileDescriptor_decompose_ade7e6d362745697)
}

var fileDescriptor_decompose_ade7e6d362745697 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x28, 0xca, 0xcc,
	0x4d, 0xcd, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x4f, 0x49, 0x4d, 0xce, 0xcf, 0x2d, 0xc8, 0x2f,
	0x4e, 0x2d, 0x48, 0x42, 0xb0, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0x91, 0x14, 0x2a,
	0x29, 0x73, 0x71, 0xba, 0xc0, 0xe4, 0x85, 0xc4, 0xb8, 0xd8, 0x20, 0xc2, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xac, 0x41, 0x50, 0x9e, 0x92, 0x0b, 0x97, 0x00, 0x5c, 0x51, 0x50, 0x6a, 0x61, 0x69, 0x6a,
	0x71, 0x89, 0x90, 0x01, 0x17, 0x7b, 0x11, 0x84, 0x09, 0x56, 0xcc, 0x6d, 0x24, 0xa6, 0x87, 0x64,
	0xae, 0x1e, 0x42, 0x3d, 0x4c, 0x99, 0x92, 0x36, 0x97, 0x20, 0x92, 0x29, 0xc5, 0x05, 0xf9, 0x79,
	0x10, 0x2b, 0x8b, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x60, 0x56, 0x42, 0x78, 0x4a, 0xbe, 0x5c, 0x92,
	0x70, 0xc5, 0xbe, 0x89, 0x79, 0x95, 0x21, 0x99, 0xb9, 0xa9, 0xc5, 0xe4, 0xdb, 0x6d, 0xc2, 0x25,
	0x85, 0xcd, 0x38, 0xfc, 0x8e, 0x30, 0x3a, 0xcd, 0x88, 0xe4, 0xf1, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc,
	0xe4, 0x54, 0x21, 0x1f, 0xe4, 0x10, 0x93, 0xc5, 0x61, 0x31, 0xc4, 0x5a, 0x29, 0x39, 0x5c, 0xd2,
	0x10, 0x8b, 0x95, 0x18, 0x84, 0x32, 0xb9, 0x84, 0x30, 0x1d, 0x26, 0xa4, 0x86, 0x5d, 0x1f, 0x7a,
	0x40, 0x48, 0xa9, 0x13, 0x54, 0x07, 0xb3, 0xc8, 0x80, 0xd1, 0x89, 0x37, 0x8a, 0x1b, 0x29, 0x59,
	0x24, 0xb1, 0x81, 0x53, 0x83, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x75, 0x9c, 0x7f, 0x89, 0x38,
	0x02, 0x00, 0x00,
}