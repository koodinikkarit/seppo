// Code generated by protoc-gen-go.
// source: seppo_service.proto
// DO NOT EDIT!

package SeppoService

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Seppo service

type SeppoClient interface {
	FetchVariationById(ctx context.Context, in *FetchVariationByIdRequest, opts ...grpc.CallOption) (*FetchVariationByIdResponse, error)
	SearchVariations(ctx context.Context, in *SearchVariationsRequest, opts ...grpc.CallOption) (*SearchVariationsResponse, error)
	ListenForChangedEwSong(ctx context.Context, in *ListenForChangedEwSongRequest, opts ...grpc.CallOption) (Seppo_ListenForChangedEwSongClient, error)
	CreateVariation(ctx context.Context, in *CreateVariationRequest, opts ...grpc.CallOption) (*CreateVariationResponse, error)
	EditVariation(ctx context.Context, in *EditVariationRequest, opts ...grpc.CallOption) (*EditVariationResponse, error)
	SyncEwDatabase(ctx context.Context, in *SyncEwDatabaseRequest, opts ...grpc.CallOption) (*SyncEwDatabaseResponse, error)
}

type seppoClient struct {
	cc *grpc.ClientConn
}

func NewSeppoClient(cc *grpc.ClientConn) SeppoClient {
	return &seppoClient{cc}
}

func (c *seppoClient) FetchVariationById(ctx context.Context, in *FetchVariationByIdRequest, opts ...grpc.CallOption) (*FetchVariationByIdResponse, error) {
	out := new(FetchVariationByIdResponse)
	err := grpc.Invoke(ctx, "/SeppoService.Seppo/fetchVariationById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seppoClient) SearchVariations(ctx context.Context, in *SearchVariationsRequest, opts ...grpc.CallOption) (*SearchVariationsResponse, error) {
	out := new(SearchVariationsResponse)
	err := grpc.Invoke(ctx, "/SeppoService.Seppo/searchVariations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seppoClient) ListenForChangedEwSong(ctx context.Context, in *ListenForChangedEwSongRequest, opts ...grpc.CallOption) (Seppo_ListenForChangedEwSongClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Seppo_serviceDesc.Streams[0], c.cc, "/SeppoService.Seppo/listenForChangedEwSong", opts...)
	if err != nil {
		return nil, err
	}
	x := &seppoListenForChangedEwSongClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Seppo_ListenForChangedEwSongClient interface {
	Recv() (*EwSong, error)
	grpc.ClientStream
}

type seppoListenForChangedEwSongClient struct {
	grpc.ClientStream
}

func (x *seppoListenForChangedEwSongClient) Recv() (*EwSong, error) {
	m := new(EwSong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seppoClient) CreateVariation(ctx context.Context, in *CreateVariationRequest, opts ...grpc.CallOption) (*CreateVariationResponse, error) {
	out := new(CreateVariationResponse)
	err := grpc.Invoke(ctx, "/SeppoService.Seppo/createVariation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seppoClient) EditVariation(ctx context.Context, in *EditVariationRequest, opts ...grpc.CallOption) (*EditVariationResponse, error) {
	out := new(EditVariationResponse)
	err := grpc.Invoke(ctx, "/SeppoService.Seppo/editVariation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seppoClient) SyncEwDatabase(ctx context.Context, in *SyncEwDatabaseRequest, opts ...grpc.CallOption) (*SyncEwDatabaseResponse, error) {
	out := new(SyncEwDatabaseResponse)
	err := grpc.Invoke(ctx, "/SeppoService.Seppo/syncEwDatabase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Seppo service

type SeppoServer interface {
	FetchVariationById(context.Context, *FetchVariationByIdRequest) (*FetchVariationByIdResponse, error)
	SearchVariations(context.Context, *SearchVariationsRequest) (*SearchVariationsResponse, error)
	ListenForChangedEwSong(*ListenForChangedEwSongRequest, Seppo_ListenForChangedEwSongServer) error
	CreateVariation(context.Context, *CreateVariationRequest) (*CreateVariationResponse, error)
	EditVariation(context.Context, *EditVariationRequest) (*EditVariationResponse, error)
	SyncEwDatabase(context.Context, *SyncEwDatabaseRequest) (*SyncEwDatabaseResponse, error)
}

func RegisterSeppoServer(s *grpc.Server, srv SeppoServer) {
	s.RegisterService(&_Seppo_serviceDesc, srv)
}

func _Seppo_FetchVariationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchVariationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeppoServer).FetchVariationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SeppoService.Seppo/FetchVariationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeppoServer).FetchVariationById(ctx, req.(*FetchVariationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seppo_SearchVariations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchVariationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeppoServer).SearchVariations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SeppoService.Seppo/SearchVariations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeppoServer).SearchVariations(ctx, req.(*SearchVariationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seppo_ListenForChangedEwSong_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenForChangedEwSongRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeppoServer).ListenForChangedEwSong(m, &seppoListenForChangedEwSongServer{stream})
}

type Seppo_ListenForChangedEwSongServer interface {
	Send(*EwSong) error
	grpc.ServerStream
}

type seppoListenForChangedEwSongServer struct {
	grpc.ServerStream
}

func (x *seppoListenForChangedEwSongServer) Send(m *EwSong) error {
	return x.ServerStream.SendMsg(m)
}

func _Seppo_CreateVariation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVariationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeppoServer).CreateVariation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SeppoService.Seppo/CreateVariation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeppoServer).CreateVariation(ctx, req.(*CreateVariationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seppo_EditVariation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditVariationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeppoServer).EditVariation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SeppoService.Seppo/EditVariation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeppoServer).EditVariation(ctx, req.(*EditVariationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seppo_SyncEwDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncEwDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeppoServer).SyncEwDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SeppoService.Seppo/SyncEwDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeppoServer).SyncEwDatabase(ctx, req.(*SyncEwDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Seppo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SeppoService.Seppo",
	HandlerType: (*SeppoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "fetchVariationById",
			Handler:    _Seppo_FetchVariationById_Handler,
		},
		{
			MethodName: "searchVariations",
			Handler:    _Seppo_SearchVariations_Handler,
		},
		{
			MethodName: "createVariation",
			Handler:    _Seppo_CreateVariation_Handler,
		},
		{
			MethodName: "editVariation",
			Handler:    _Seppo_EditVariation_Handler,
		},
		{
			MethodName: "syncEwDatabase",
			Handler:    _Seppo_SyncEwDatabase_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "listenForChangedEwSong",
			Handler:       _Seppo_ListenForChangedEwSong_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "seppo_service.proto",
}

func init() { proto.RegisterFile("seppo_service.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x2b, 0xa8, 0x87, 0xc5, 0xb6, 0x3a, 0x15, 0x0f, 0x39, 0xc6, 0x56, 0x05, 0x21, 0x88,
	0xfe, 0x03, 0x6b, 0x0b, 0x82, 0x27, 0x03, 0x1e, 0x04, 0x89, 0xdb, 0x64, 0x4c, 0x17, 0x64, 0x37,
	0xee, 0xac, 0x09, 0xfd, 0xa1, 0xfe, 0x1f, 0x69, 0x93, 0x8d, 0xd9, 0x58, 0x9a, 0xe3, 0xbe, 0xf7,
	0xcd, 0x7b, 0xb3, 0x30, 0x6c, 0x44, 0x98, 0x65, 0x2a, 0x22, 0xd4, 0xb9, 0x88, 0x31, 0xc8, 0xb4,
	0x32, 0x0a, 0x8e, 0xc2, 0xb5, 0x18, 0x96, 0x9a, 0x37, 0x22, 0x25, 0xd3, 0x28, 0xe1, 0x86, 0x2f,
	0x38, 0x55, 0x88, 0x77, 0x82, 0x45, 0x5b, 0xea, 0x63, 0x11, 0xad, 0xd1, 0xea, 0x39, 0xc0, 0x22,
	0xca, 0x51, 0xd7, 0xf6, 0x30, 0xe7, 0x5a, 0x70, 0x23, 0x94, 0x2c, 0x85, 0xdb, 0x9f, 0x7d, 0x76,
	0xb0, 0x29, 0x02, 0xc1, 0xe0, 0x03, 0x4d, 0xbc, 0x7c, 0xb1, 0xc4, 0xfd, 0xea, 0x31, 0x81, 0xcb,
	0xa0, 0xb9, 0x46, 0x30, 0xff, 0x47, 0x3c, 0xe3, 0xd7, 0x37, 0x92, 0xf1, 0xae, 0xba, 0x41, 0xca,
	0x94, 0x24, 0xf4, 0x7b, 0x10, 0xb3, 0x63, 0x42, 0xae, 0x1b, 0x00, 0xc1, 0xc4, 0x9d, 0x0f, 0x5b,
	0xbe, 0xad, 0xb9, 0xe8, 0xc2, 0xea, 0x92, 0x88, 0x9d, 0x7d, 0x0a, 0x32, 0x28, 0xe7, 0x4a, 0x4f,
	0x97, 0x5c, 0xa6, 0x98, 0xcc, 0x8a, 0x50, 0xc9, 0x14, 0xae, 0xdd, 0x8c, 0xa7, 0xad, 0x94, 0x2d,
	0x3c, 0x75, 0xe1, 0xd2, 0xf4, 0x7b, 0x37, 0x7b, 0xf0, 0xce, 0x86, 0xb1, 0x46, 0x6e, 0xb0, 0xae,
	0x87, 0xb1, 0x0b, 0x4f, 0x5d, 0xdb, 0x46, 0x4e, 0x3a, 0xa8, 0xfa, 0x0b, 0xaf, 0xac, 0x8f, 0x89,
	0x30, 0x7f, 0xf9, 0x7e, 0x6b, 0x99, 0xa6, 0x69, 0xd3, 0xcf, 0x77, 0x32, 0x75, 0xf6, 0x1b, 0x1b,
	0xd0, 0x4a, 0xc6, 0xb3, 0xe2, 0xa1, 0x3a, 0x20, 0x68, 0x0d, 0x86, 0x8e, 0x6b, 0xd3, 0xc7, 0xbb,
	0x21, 0x1b, 0xbf, 0x38, 0xdc, 0x9c, 0xd7, 0xdd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xdd,
	0x1a, 0x14, 0xdb, 0x02, 0x00, 0x00,
}