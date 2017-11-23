// Code generated by protoc-gen-go.
// source: matias_service.proto
// DO NOT EDIT!

package MatiasService

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

type RequestMatiasKeyRequest struct {
}

func (m *RequestMatiasKeyRequest) Reset()                    { *m = RequestMatiasKeyRequest{} }
func (m *RequestMatiasKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestMatiasKeyRequest) ProtoMessage()               {}
func (*RequestMatiasKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type RequestMatiasKeyResponse struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *RequestMatiasKeyResponse) Reset()                    { *m = RequestMatiasKeyResponse{} }
func (m *RequestMatiasKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestMatiasKeyResponse) ProtoMessage()               {}
func (*RequestMatiasKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *RequestMatiasKeyResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestMatiasKeyRequest)(nil), "MatiasService.RequestMatiasKeyRequest")
	proto.RegisterType((*RequestMatiasKeyResponse)(nil), "MatiasService.RequestMatiasKeyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Matias service

type MatiasClient interface {
	RequestMatiasKey(ctx context.Context, in *RequestMatiasKeyRequest, opts ...grpc.CallOption) (*RequestMatiasKeyResponse, error)
	SyncEwDatabase(ctx context.Context, in *SyncEwDatabaseRequest, opts ...grpc.CallOption) (*SyncEwDatabaseResponse, error)
	InsertEwSongIds(ctx context.Context, in *InsertEwSongIdsRequest, opts ...grpc.CallOption) (*InsertEwSongIdsResponse, error)
	RequestEwChanges(ctx context.Context, in *RequestEwDatabaseChangesRequest, opts ...grpc.CallOption) (Matias_RequestEwChangesClient, error)
}

type matiasClient struct {
	cc *grpc.ClientConn
}

func NewMatiasClient(cc *grpc.ClientConn) MatiasClient {
	return &matiasClient{cc}
}

func (c *matiasClient) RequestMatiasKey(ctx context.Context, in *RequestMatiasKeyRequest, opts ...grpc.CallOption) (*RequestMatiasKeyResponse, error) {
	out := new(RequestMatiasKeyResponse)
	err := grpc.Invoke(ctx, "/MatiasService.Matias/requestMatiasKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matiasClient) SyncEwDatabase(ctx context.Context, in *SyncEwDatabaseRequest, opts ...grpc.CallOption) (*SyncEwDatabaseResponse, error) {
	out := new(SyncEwDatabaseResponse)
	err := grpc.Invoke(ctx, "/MatiasService.Matias/syncEwDatabase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matiasClient) InsertEwSongIds(ctx context.Context, in *InsertEwSongIdsRequest, opts ...grpc.CallOption) (*InsertEwSongIdsResponse, error) {
	out := new(InsertEwSongIdsResponse)
	err := grpc.Invoke(ctx, "/MatiasService.Matias/insertEwSongIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matiasClient) RequestEwChanges(ctx context.Context, in *RequestEwDatabaseChangesRequest, opts ...grpc.CallOption) (Matias_RequestEwChangesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Matias_serviceDesc.Streams[0], c.cc, "/MatiasService.Matias/requestEwChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &matiasRequestEwChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Matias_RequestEwChangesClient interface {
	Recv() (*EwDatabaseChange, error)
	grpc.ClientStream
}

type matiasRequestEwChangesClient struct {
	grpc.ClientStream
}

func (x *matiasRequestEwChangesClient) Recv() (*EwDatabaseChange, error) {
	m := new(EwDatabaseChange)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Matias service

type MatiasServer interface {
	RequestMatiasKey(context.Context, *RequestMatiasKeyRequest) (*RequestMatiasKeyResponse, error)
	SyncEwDatabase(context.Context, *SyncEwDatabaseRequest) (*SyncEwDatabaseResponse, error)
	InsertEwSongIds(context.Context, *InsertEwSongIdsRequest) (*InsertEwSongIdsResponse, error)
	RequestEwChanges(*RequestEwDatabaseChangesRequest, Matias_RequestEwChangesServer) error
}

func RegisterMatiasServer(s *grpc.Server, srv MatiasServer) {
	s.RegisterService(&_Matias_serviceDesc, srv)
}

func _Matias_RequestMatiasKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMatiasKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatiasServer).RequestMatiasKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatiasService.Matias/RequestMatiasKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatiasServer).RequestMatiasKey(ctx, req.(*RequestMatiasKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matias_SyncEwDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncEwDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatiasServer).SyncEwDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatiasService.Matias/SyncEwDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatiasServer).SyncEwDatabase(ctx, req.(*SyncEwDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matias_InsertEwSongIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertEwSongIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatiasServer).InsertEwSongIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatiasService.Matias/InsertEwSongIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatiasServer).InsertEwSongIds(ctx, req.(*InsertEwSongIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matias_RequestEwChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestEwDatabaseChangesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatiasServer).RequestEwChanges(m, &matiasRequestEwChangesServer{stream})
}

type Matias_RequestEwChangesServer interface {
	Send(*EwDatabaseChange) error
	grpc.ServerStream
}

type matiasRequestEwChangesServer struct {
	grpc.ServerStream
}

func (x *matiasRequestEwChangesServer) Send(m *EwDatabaseChange) error {
	return x.ServerStream.SendMsg(m)
}

var _Matias_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MatiasService.Matias",
	HandlerType: (*MatiasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "requestMatiasKey",
			Handler:    _Matias_RequestMatiasKey_Handler,
		},
		{
			MethodName: "syncEwDatabase",
			Handler:    _Matias_SyncEwDatabase_Handler,
		},
		{
			MethodName: "insertEwSongIds",
			Handler:    _Matias_InsertEwSongIds_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "requestEwChanges",
			Handler:       _Matias_RequestEwChanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "matias_service.proto",
}

func init() { proto.RegisterFile("matias_service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0x2c, 0xc9,
	0x4c, 0x2c, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0xf5, 0x05, 0x8b, 0x06, 0x43, 0x04, 0xa5, 0x04, 0x53, 0xcb, 0xe3, 0x53, 0x12, 0x4b,
	0x12, 0x93, 0x12, 0x8b, 0xa1, 0x2a, 0x94, 0x24, 0xb9, 0xc4, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x20, 0x4a, 0xbd, 0x53, 0x2b, 0xa1, 0x7c, 0x25, 0x1d, 0x2e, 0x09, 0x4c, 0xa9, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x10, 0xd3, 0x68, 0x2e, 0x33, 0x17, 0x1b, 0x44, 0x9d, 0x50, 0x2a, 0x97, 0x40, 0x11,
	0x9a, 0x46, 0x21, 0x35, 0x3d, 0x14, 0xa7, 0xe8, 0xe1, 0xb0, 0x54, 0x4a, 0x9d, 0xa0, 0x3a, 0x88,
	0x0b, 0x94, 0x18, 0x84, 0xe2, 0xb9, 0xf8, 0x8a, 0x2b, 0xf3, 0x92, 0x5d, 0xcb, 0x5d, 0xa0, 0x5e,
	0x12, 0x52, 0x41, 0xd3, 0x1c, 0x8c, 0x22, 0x0d, 0xb3, 0x42, 0x95, 0x80, 0x2a, 0xb8, 0x05, 0x49,
	0x5c, 0xfc, 0x99, 0x79, 0xc5, 0xa9, 0x45, 0x25, 0xae, 0xe5, 0xc1, 0xf9, 0x79, 0xe9, 0x9e, 0x29,
	0xc5, 0x42, 0xe8, 0x7a, 0x3d, 0x51, 0xe5, 0x61, 0x56, 0xa8, 0x11, 0x52, 0x06, 0xb7, 0x23, 0x1d,
	0x1e, 0x56, 0xae, 0xe5, 0xce, 0x19, 0x89, 0x79, 0xe9, 0xa9, 0xc5, 0x42, 0x7a, 0xd8, 0xc3, 0x00,
	0xe1, 0x46, 0xa8, 0x42, 0x98, 0x6d, 0xf2, 0x68, 0xea, 0xd1, 0x15, 0x2a, 0x31, 0x18, 0x30, 0x26,
	0xb1, 0x81, 0xe3, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x35, 0x93, 0x33, 0x29, 0x02,
	0x00, 0x00,
}
