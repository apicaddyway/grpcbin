// Code generated by protoc-gen-go.
// source: grpcbin.proto
// DO NOT EDIT!

/*
Package grpcbin is a generated protocol buffer package.

It is generated from these files:
	grpcbin.proto

It has these top-level messages:
	SpecificErrorRequest
	EmptyMessage
	DummyMessage
	IndexReply
*/
package grpcbin

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

type DummyMessage_Enum int32

const (
	DummyMessage_ENUM_0 DummyMessage_Enum = 0
	DummyMessage_ENUM_1 DummyMessage_Enum = 1
	DummyMessage_ENUM_2 DummyMessage_Enum = 2
)

var DummyMessage_Enum_name = map[int32]string{
	0: "ENUM_0",
	1: "ENUM_1",
	2: "ENUM_2",
}
var DummyMessage_Enum_value = map[string]int32{
	"ENUM_0": 0,
	"ENUM_1": 1,
	"ENUM_2": 2,
}

func (x DummyMessage_Enum) String() string {
	return proto.EnumName(DummyMessage_Enum_name, int32(x))
}
func (DummyMessage_Enum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type SpecificErrorRequest struct {
	Code   uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *SpecificErrorRequest) Reset()                    { *m = SpecificErrorRequest{} }
func (m *SpecificErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*SpecificErrorRequest) ProtoMessage()               {}
func (*SpecificErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpecificErrorRequest) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SpecificErrorRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DummyMessage struct {
	FString  string              `protobuf:"bytes,1,opt,name=f_string,json=fString" json:"f_string,omitempty"`
	FStrings []string            `protobuf:"bytes,2,rep,name=f_strings,json=fStrings" json:"f_strings,omitempty"`
	FInt32   int32               `protobuf:"varint,3,opt,name=f_int32,json=fInt32" json:"f_int32,omitempty"`
	FInt32S  []int32             `protobuf:"varint,4,rep,packed,name=f_int32s,json=fInt32s" json:"f_int32s,omitempty"`
	FEnum    DummyMessage_Enum   `protobuf:"varint,5,opt,name=f_enum,json=fEnum,enum=grpcbin.DummyMessage_Enum" json:"f_enum,omitempty"`
	FEnums   []DummyMessage_Enum `protobuf:"varint,6,rep,packed,name=f_enums,json=fEnums,enum=grpcbin.DummyMessage_Enum" json:"f_enums,omitempty"`
	FSub     *DummyMessage_Sub   `protobuf:"bytes,7,opt,name=f_sub,json=fSub" json:"f_sub,omitempty"`
	FSubs    []*DummyMessage_Sub `protobuf:"bytes,8,rep,name=f_subs,json=fSubs" json:"f_subs,omitempty"`
	FBool    bool                `protobuf:"varint,9,opt,name=f_bool,json=fBool" json:"f_bool,omitempty"`
	FBools   []bool              `protobuf:"varint,10,rep,packed,name=f_bools,json=fBools" json:"f_bools,omitempty"`
	FInt64   int64               `protobuf:"varint,11,opt,name=f_int64,json=fInt64" json:"f_int64,omitempty"`
	FInt64S  []int64             `protobuf:"varint,12,rep,packed,name=f_int64s,json=fInt64s" json:"f_int64s,omitempty"`
	FBytes   []byte              `protobuf:"bytes,13,opt,name=f_bytes,json=fBytes,proto3" json:"f_bytes,omitempty"`
	FBytess  [][]byte            `protobuf:"bytes,14,rep,name=f_bytess,json=fBytess,proto3" json:"f_bytess,omitempty"`
	FFloat   float32             `protobuf:"fixed32,15,opt,name=f_float,json=fFloat" json:"f_float,omitempty"`
	FFloats  []float32           `protobuf:"fixed32,16,rep,packed,name=f_floats,json=fFloats" json:"f_floats,omitempty"`
}

func (m *DummyMessage) Reset()                    { *m = DummyMessage{} }
func (m *DummyMessage) String() string            { return proto.CompactTextString(m) }
func (*DummyMessage) ProtoMessage()               {}
func (*DummyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DummyMessage) GetFString() string {
	if m != nil {
		return m.FString
	}
	return ""
}

func (m *DummyMessage) GetFStrings() []string {
	if m != nil {
		return m.FStrings
	}
	return nil
}

func (m *DummyMessage) GetFInt32() int32 {
	if m != nil {
		return m.FInt32
	}
	return 0
}

func (m *DummyMessage) GetFInt32S() []int32 {
	if m != nil {
		return m.FInt32S
	}
	return nil
}

func (m *DummyMessage) GetFEnum() DummyMessage_Enum {
	if m != nil {
		return m.FEnum
	}
	return DummyMessage_ENUM_0
}

func (m *DummyMessage) GetFEnums() []DummyMessage_Enum {
	if m != nil {
		return m.FEnums
	}
	return nil
}

func (m *DummyMessage) GetFSub() *DummyMessage_Sub {
	if m != nil {
		return m.FSub
	}
	return nil
}

func (m *DummyMessage) GetFSubs() []*DummyMessage_Sub {
	if m != nil {
		return m.FSubs
	}
	return nil
}

func (m *DummyMessage) GetFBool() bool {
	if m != nil {
		return m.FBool
	}
	return false
}

func (m *DummyMessage) GetFBools() []bool {
	if m != nil {
		return m.FBools
	}
	return nil
}

func (m *DummyMessage) GetFInt64() int64 {
	if m != nil {
		return m.FInt64
	}
	return 0
}

func (m *DummyMessage) GetFInt64S() []int64 {
	if m != nil {
		return m.FInt64S
	}
	return nil
}

func (m *DummyMessage) GetFBytes() []byte {
	if m != nil {
		return m.FBytes
	}
	return nil
}

func (m *DummyMessage) GetFBytess() [][]byte {
	if m != nil {
		return m.FBytess
	}
	return nil
}

func (m *DummyMessage) GetFFloat() float32 {
	if m != nil {
		return m.FFloat
	}
	return 0
}

func (m *DummyMessage) GetFFloats() []float32 {
	if m != nil {
		return m.FFloats
	}
	return nil
}

type DummyMessage_Sub struct {
	FString string `protobuf:"bytes,1,opt,name=f_string,json=fString" json:"f_string,omitempty"`
}

func (m *DummyMessage_Sub) Reset()                    { *m = DummyMessage_Sub{} }
func (m *DummyMessage_Sub) String() string            { return proto.CompactTextString(m) }
func (*DummyMessage_Sub) ProtoMessage()               {}
func (*DummyMessage_Sub) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *DummyMessage_Sub) GetFString() string {
	if m != nil {
		return m.FString
	}
	return ""
}

type IndexReply struct {
	Description string                 `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Endpoints   []*IndexReply_Endpoint `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *IndexReply) Reset()                    { *m = IndexReply{} }
func (m *IndexReply) String() string            { return proto.CompactTextString(m) }
func (*IndexReply) ProtoMessage()               {}
func (*IndexReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IndexReply) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *IndexReply) GetEndpoints() []*IndexReply_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type IndexReply_Endpoint struct {
	Path        string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *IndexReply_Endpoint) Reset()                    { *m = IndexReply_Endpoint{} }
func (m *IndexReply_Endpoint) String() string            { return proto.CompactTextString(m) }
func (*IndexReply_Endpoint) ProtoMessage()               {}
func (*IndexReply_Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *IndexReply_Endpoint) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *IndexReply_Endpoint) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*SpecificErrorRequest)(nil), "grpcbin.SpecificErrorRequest")
	proto.RegisterType((*EmptyMessage)(nil), "grpcbin.EmptyMessage")
	proto.RegisterType((*DummyMessage)(nil), "grpcbin.DummyMessage")
	proto.RegisterType((*DummyMessage_Sub)(nil), "grpcbin.DummyMessage.Sub")
	proto.RegisterType((*IndexReply)(nil), "grpcbin.IndexReply")
	proto.RegisterType((*IndexReply_Endpoint)(nil), "grpcbin.IndexReply.Endpoint")
	proto.RegisterEnum("grpcbin.DummyMessage_Enum", DummyMessage_Enum_name, DummyMessage_Enum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GRPCBin service

type GRPCBinClient interface {
	Index(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*IndexReply, error)
	Empty(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
	DummyUnary(ctx context.Context, in *DummyMessage, opts ...grpc.CallOption) (*DummyMessage, error)
	DummyServerStream(ctx context.Context, in *DummyMessage, opts ...grpc.CallOption) (GRPCBin_DummyServerStreamClient, error)
	DummyClientStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBin_DummyClientStreamClient, error)
	DummyBidirectionalStreamStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBin_DummyBidirectionalStreamStreamClient, error)
	SpecificError(ctx context.Context, in *SpecificErrorRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	RandomError(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type gRPCBinClient struct {
	cc *grpc.ClientConn
}

func NewGRPCBinClient(cc *grpc.ClientConn) GRPCBinClient {
	return &gRPCBinClient{cc}
}

func (c *gRPCBinClient) Index(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*IndexReply, error) {
	out := new(IndexReply)
	err := grpc.Invoke(ctx, "/grpcbin.GRPCBin/Index", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCBinClient) Empty(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/grpcbin.GRPCBin/Empty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCBinClient) DummyUnary(ctx context.Context, in *DummyMessage, opts ...grpc.CallOption) (*DummyMessage, error) {
	out := new(DummyMessage)
	err := grpc.Invoke(ctx, "/grpcbin.GRPCBin/DummyUnary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCBinClient) DummyServerStream(ctx context.Context, in *DummyMessage, opts ...grpc.CallOption) (GRPCBin_DummyServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GRPCBin_serviceDesc.Streams[0], c.cc, "/grpcbin.GRPCBin/DummyServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCBinDummyServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GRPCBin_DummyServerStreamClient interface {
	Recv() (*DummyMessage, error)
	grpc.ClientStream
}

type gRPCBinDummyServerStreamClient struct {
	grpc.ClientStream
}

func (x *gRPCBinDummyServerStreamClient) Recv() (*DummyMessage, error) {
	m := new(DummyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCBinClient) DummyClientStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBin_DummyClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GRPCBin_serviceDesc.Streams[1], c.cc, "/grpcbin.GRPCBin/DummyClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCBinDummyClientStreamClient{stream}
	return x, nil
}

type GRPCBin_DummyClientStreamClient interface {
	Send(*DummyMessage) error
	CloseAndRecv() (*DummyMessage, error)
	grpc.ClientStream
}

type gRPCBinDummyClientStreamClient struct {
	grpc.ClientStream
}

func (x *gRPCBinDummyClientStreamClient) Send(m *DummyMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCBinDummyClientStreamClient) CloseAndRecv() (*DummyMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DummyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCBinClient) DummyBidirectionalStreamStream(ctx context.Context, opts ...grpc.CallOption) (GRPCBin_DummyBidirectionalStreamStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GRPCBin_serviceDesc.Streams[2], c.cc, "/grpcbin.GRPCBin/DummyBidirectionalStreamStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCBinDummyBidirectionalStreamStreamClient{stream}
	return x, nil
}

type GRPCBin_DummyBidirectionalStreamStreamClient interface {
	Send(*DummyMessage) error
	Recv() (*DummyMessage, error)
	grpc.ClientStream
}

type gRPCBinDummyBidirectionalStreamStreamClient struct {
	grpc.ClientStream
}

func (x *gRPCBinDummyBidirectionalStreamStreamClient) Send(m *DummyMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCBinDummyBidirectionalStreamStreamClient) Recv() (*DummyMessage, error) {
	m := new(DummyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCBinClient) SpecificError(ctx context.Context, in *SpecificErrorRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/grpcbin.GRPCBin/SpecificError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCBinClient) RandomError(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/grpcbin.GRPCBin/RandomError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GRPCBin service

type GRPCBinServer interface {
	Index(context.Context, *EmptyMessage) (*IndexReply, error)
	Empty(context.Context, *EmptyMessage) (*EmptyMessage, error)
	DummyUnary(context.Context, *DummyMessage) (*DummyMessage, error)
	DummyServerStream(*DummyMessage, GRPCBin_DummyServerStreamServer) error
	DummyClientStream(GRPCBin_DummyClientStreamServer) error
	DummyBidirectionalStreamStream(GRPCBin_DummyBidirectionalStreamStreamServer) error
	SpecificError(context.Context, *SpecificErrorRequest) (*EmptyMessage, error)
	RandomError(context.Context, *EmptyMessage) (*EmptyMessage, error)
}

func RegisterGRPCBinServer(s *grpc.Server, srv GRPCBinServer) {
	s.RegisterService(&_GRPCBin_serviceDesc, srv)
}

func _GRPCBin_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCBinServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcbin.GRPCBin/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCBinServer).Index(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCBin_Empty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCBinServer).Empty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcbin.GRPCBin/Empty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCBinServer).Empty(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCBin_DummyUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DummyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCBinServer).DummyUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcbin.GRPCBin/DummyUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCBinServer).DummyUnary(ctx, req.(*DummyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCBin_DummyServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DummyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GRPCBinServer).DummyServerStream(m, &gRPCBinDummyServerStreamServer{stream})
}

type GRPCBin_DummyServerStreamServer interface {
	Send(*DummyMessage) error
	grpc.ServerStream
}

type gRPCBinDummyServerStreamServer struct {
	grpc.ServerStream
}

func (x *gRPCBinDummyServerStreamServer) Send(m *DummyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _GRPCBin_DummyClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCBinServer).DummyClientStream(&gRPCBinDummyClientStreamServer{stream})
}

type GRPCBin_DummyClientStreamServer interface {
	SendAndClose(*DummyMessage) error
	Recv() (*DummyMessage, error)
	grpc.ServerStream
}

type gRPCBinDummyClientStreamServer struct {
	grpc.ServerStream
}

func (x *gRPCBinDummyClientStreamServer) SendAndClose(m *DummyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCBinDummyClientStreamServer) Recv() (*DummyMessage, error) {
	m := new(DummyMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GRPCBin_DummyBidirectionalStreamStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCBinServer).DummyBidirectionalStreamStream(&gRPCBinDummyBidirectionalStreamStreamServer{stream})
}

type GRPCBin_DummyBidirectionalStreamStreamServer interface {
	Send(*DummyMessage) error
	Recv() (*DummyMessage, error)
	grpc.ServerStream
}

type gRPCBinDummyBidirectionalStreamStreamServer struct {
	grpc.ServerStream
}

func (x *gRPCBinDummyBidirectionalStreamStreamServer) Send(m *DummyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCBinDummyBidirectionalStreamStreamServer) Recv() (*DummyMessage, error) {
	m := new(DummyMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GRPCBin_SpecificError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpecificErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCBinServer).SpecificError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcbin.GRPCBin/SpecificError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCBinServer).SpecificError(ctx, req.(*SpecificErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCBin_RandomError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCBinServer).RandomError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcbin.GRPCBin/RandomError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCBinServer).RandomError(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCBin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcbin.GRPCBin",
	HandlerType: (*GRPCBinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _GRPCBin_Index_Handler,
		},
		{
			MethodName: "Empty",
			Handler:    _GRPCBin_Empty_Handler,
		},
		{
			MethodName: "DummyUnary",
			Handler:    _GRPCBin_DummyUnary_Handler,
		},
		{
			MethodName: "SpecificError",
			Handler:    _GRPCBin_SpecificError_Handler,
		},
		{
			MethodName: "RandomError",
			Handler:    _GRPCBin_RandomError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DummyServerStream",
			Handler:       _GRPCBin_DummyServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DummyClientStream",
			Handler:       _GRPCBin_DummyClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DummyBidirectionalStreamStream",
			Handler:       _GRPCBin_DummyBidirectionalStreamStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcbin.proto",
}

func init() { proto.RegisterFile("grpcbin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x6f, 0xda, 0x3a,
	0x18, 0xc5, 0x84, 0x24, 0xf0, 0x01, 0xbd, 0x5c, 0xdf, 0xdb, 0x7b, 0x5d, 0xf6, 0x43, 0x56, 0x9e,
	0xac, 0x3d, 0x20, 0x4a, 0x19, 0x93, 0xa6, 0x4d, 0x9a, 0xe8, 0x58, 0xd5, 0x87, 0x4e, 0x93, 0x59,
	0x9f, 0x51, 0x02, 0x4e, 0x17, 0x09, 0x9c, 0x2c, 0x4e, 0xa6, 0xf5, 0xef, 0xd9, 0xeb, 0xfe, 0xa6,
	0xfd, 0x2d, 0x93, 0x9d, 0x90, 0xb2, 0x8d, 0x6e, 0x6a, 0x9f, 0xf0, 0x77, 0x8e, 0xcf, 0xf9, 0xcc,
	0xc9, 0x67, 0x43, 0xf7, 0x2a, 0x4d, 0x96, 0x41, 0x24, 0x07, 0x49, 0x1a, 0x67, 0x31, 0x76, 0xcb,
	0xd2, 0x9b, 0xc2, 0xbf, 0xf3, 0x44, 0x2c, 0xa3, 0x30, 0x5a, 0xce, 0xd2, 0x34, 0x4e, 0xb9, 0xf8,
	0x98, 0x0b, 0x95, 0x61, 0x0c, 0x8d, 0x65, 0xbc, 0x12, 0x04, 0x51, 0xc4, 0xba, 0xdc, 0xac, 0xf1,
	0x7f, 0xe0, 0xa4, 0xc2, 0x57, 0xb1, 0x24, 0x75, 0x8a, 0x58, 0x8b, 0x97, 0x95, 0x77, 0x00, 0x9d,
	0xd9, 0x26, 0xc9, 0xae, 0x2f, 0x84, 0x52, 0xfe, 0x95, 0xf0, 0xbe, 0x35, 0xa0, 0xf3, 0x3a, 0xdf,
	0x6c, 0xb6, 0x00, 0x3e, 0x82, 0x66, 0xb8, 0x50, 0x59, 0x1a, 0xc9, 0x2b, 0x63, 0xd8, 0xe2, 0x6e,
	0x38, 0x37, 0x25, 0x7e, 0x00, 0xad, 0x2d, 0xa5, 0x48, 0x9d, 0x5a, 0xac, 0xc5, 0x9b, 0x25, 0xa7,
	0xf0, 0xff, 0xe0, 0x86, 0x8b, 0x48, 0x66, 0x27, 0x23, 0x62, 0x51, 0xc4, 0x6c, 0xee, 0x84, 0xe7,
	0xba, 0x2a, 0x0c, 0x0d, 0xa1, 0x48, 0x83, 0x5a, 0xcc, 0xe6, 0x6e, 0xc1, 0x28, 0x7c, 0x0c, 0x4e,
	0xb8, 0x10, 0x32, 0xdf, 0x10, 0x9b, 0x22, 0x76, 0x30, 0xea, 0x0f, 0xb6, 0xff, 0x7c, 0xf7, 0x48,
	0x83, 0x99, 0xcc, 0x37, 0xdc, 0x0e, 0xf5, 0x0f, 0x3e, 0xd1, 0x6d, 0xb4, 0x44, 0x11, 0x87, 0x5a,
	0x7f, 0xd0, 0x38, 0x46, 0xa3, 0xf0, 0x00, 0xec, 0x70, 0xa1, 0xf2, 0x80, 0xb8, 0x14, 0xb1, 0xf6,
	0xe8, 0x68, 0xbf, 0x64, 0x9e, 0x07, 0xbc, 0x11, 0xce, 0xf3, 0x00, 0x0f, 0xf5, 0xb9, 0x54, 0x1e,
	0x28, 0xd2, 0xa4, 0xd6, 0xef, 0x05, 0xb6, 0x16, 0x28, 0x7c, 0xa8, 0x15, 0x41, 0x1c, 0xaf, 0x49,
	0x8b, 0x22, 0xd6, 0xe4, 0x76, 0x38, 0x8d, 0xe3, 0x75, 0x11, 0x8a, 0x86, 0x15, 0x01, 0x6a, 0xb1,
	0x26, 0x77, 0x0c, 0x7e, 0x93, 0xd6, 0x64, 0x4c, 0xda, 0x14, 0x31, 0xab, 0x48, 0x6b, 0x32, 0xae,
	0xd2, 0x9a, 0x8c, 0x15, 0xe9, 0x50, 0x8b, 0x59, 0x45, 0x5a, 0x93, 0x71, 0xa9, 0x09, 0xae, 0x33,
	0xa1, 0x48, 0x97, 0x22, 0xd6, 0xd1, 0x66, 0xba, 0x2a, 0x34, 0x86, 0x50, 0xe4, 0x80, 0x5a, 0xac,
	0xc3, 0xdd, 0x82, 0x29, 0x35, 0xe1, 0x3a, 0xf6, 0x33, 0xf2, 0x17, 0x45, 0xac, 0xce, 0x9d, 0xf0,
	0x8d, 0xae, 0x0a, 0x8d, 0x21, 0x14, 0xe9, 0x51, 0x8b, 0xd5, 0xb9, 0x5b, 0x30, 0xaa, 0x4f, 0xc1,
	0xd2, 0x21, 0xdc, 0x3e, 0x08, 0xde, 0x13, 0x68, 0x98, 0x8f, 0x01, 0xe0, 0xcc, 0xde, 0x5e, 0x5e,
	0x2c, 0x86, 0xbd, 0x5a, 0xb5, 0x3e, 0xee, 0xa1, 0x6a, 0x3d, 0xea, 0xd5, 0xbd, 0xaf, 0x08, 0xe0,
	0x5c, 0xae, 0xc4, 0x67, 0x2e, 0x92, 0xf5, 0x35, 0xa6, 0xd0, 0x5e, 0x09, 0xb5, 0x4c, 0xa3, 0x24,
	0x8b, 0x62, 0x59, 0x1a, 0xef, 0x42, 0xf8, 0x39, 0xb4, 0x84, 0x5c, 0x25, 0x71, 0x24, 0xb3, 0x62,
	0xca, 0xda, 0xa3, 0x87, 0x55, 0xfe, 0x37, 0x4e, 0x83, 0x59, 0xb9, 0x89, 0xdf, 0x6c, 0xef, 0xbf,
	0x82, 0xe6, 0x16, 0xd6, 0xb7, 0x22, 0xf1, 0xb3, 0x0f, 0x65, 0x0b, 0xb3, 0xfe, 0xb9, 0x7b, 0xfd,
	0x97, 0xee, 0xa3, 0x2f, 0x0d, 0x70, 0xcf, 0xf8, 0xbb, 0xd3, 0x69, 0x24, 0xf1, 0x53, 0xb0, 0x4d,
	0x3f, 0x7c, 0x58, 0xf5, 0xdf, 0xbd, 0x3b, 0xfd, 0x7f, 0xf6, 0x1c, 0xcb, 0xab, 0xe1, 0x67, 0x60,
	0x9b, 0x6d, 0xb7, 0xc9, 0xf6, 0xc3, 0x5e, 0x0d, 0xbf, 0x00, 0x30, 0xf3, 0x75, 0x29, 0xfd, 0x74,
	0x57, 0xbd, 0x3b, 0x74, 0xfd, 0xfd, 0xb0, 0x57, 0xc3, 0x33, 0xf8, 0xdb, 0x20, 0x73, 0x91, 0x7e,
	0x12, 0xe9, 0x3c, 0x4b, 0x85, 0xbf, 0xb9, 0xab, 0xc9, 0x10, 0x55, 0x36, 0xa7, 0xeb, 0x48, 0xc8,
	0xec, 0x7e, 0x36, 0x0c, 0xe1, 0xf7, 0xf0, 0xd8, 0x60, 0xd3, 0x68, 0x15, 0xa5, 0x62, 0xa9, 0xb3,
	0xf5, 0xd7, 0x85, 0xdb, 0x7d, 0x3d, 0x87, 0x08, 0x9f, 0x41, 0xf7, 0x87, 0x17, 0x10, 0x3f, 0xaa,
	0x76, 0xef, 0x7b, 0x19, 0x6f, 0x8f, 0xfa, 0x25, 0xb4, 0xb9, 0x2f, 0x57, 0xf1, 0xa6, 0xb0, 0xb9,
	0xe3, 0x97, 0x0a, 0x1c, 0xf3, 0x32, 0x9f, 0x7c, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x35, 0x4c,
	0xc2, 0xaa, 0x05, 0x00, 0x00,
}
