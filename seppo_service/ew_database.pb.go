// Code generated by protoc-gen-go.
// source: ew_database.proto
// DO NOT EDIT!

/*
Package SeppoService is a generated protocol buffer package.

It is generated from these files:
	ew_database.proto
	ew_song.proto
	ew_verse.proto
	seppo_service.proto
	song.proto
	song_database.proto
	song_database_variation.proto
	variation.proto

It has these top-level messages:
	EwDatabase
	EwDatabasesConnection
	ListenForChangedEwSongRequest
	FetchEwDatabasesRequest
	FetchEwDatabasesResponse
	FetchEwDatabaseByIdRequest
	FetchEwDatabaseByIdResponse
	CreateEwDatabaseRequest
	CreateEwDatabaseResponse
	EditEwDatabaseRequest
	EditEwDatabaseResponse
	RemoveEwDatabaseRequest
	RemoveEwDatabaseResponse
	SyncEwDatabaseRequest
	SyncEwDatabaseResponse
	EwSong
	EwVerse
	Song
	SongDatabase
	SongDatabaseEdge
	SongDatabasesConnection
	FetchSongDatabasesRequest
	FetchSongDatabaseByIdRequest
	FetchSongDatabaseByIdResponse
	CreateSongDatabaseRequest
	CreateSongDatabaseResponse
	EditSongDatabaseRequest
	EditSongDatabaseResponse
	RemoveSongDatabaseRequest
	RemoveSongDatabaseResponse
	SongDatabaseVariation
	SongDatabaseVariations
	FetchVariationsBySongDatabaseIdRequest
	FetchVariationsBySongDatabaseIdResponse
	AddVariationToSongDatabaseRequest
	AddVariationToSongDatabaseResponse
	RemoveVariationFromSongDatabaseRequest
	RemoveVariationFromSongDatabaseResponse
	Variation
	CreateVariationRequest
	CreateVariationResponse
	EditVariationRequest
	EditVariationResponse
	RemoveVariationRequest
	RemoveVariationResponse
	FetchVariationByIdRequest
	FetchVariationByIdResponse
	SearchVariationsRequest
	SearchVariationsResponse
*/
package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EwDatabase struct {
	Id             uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SongDatabaseId uint32 `protobuf:"varint,3,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	Version        uint64 `protobuf:"varint,4,opt,name=version" json:"version,omitempty"`
}

func (m *EwDatabase) Reset()                    { *m = EwDatabase{} }
func (m *EwDatabase) String() string            { return proto.CompactTextString(m) }
func (*EwDatabase) ProtoMessage()               {}
func (*EwDatabase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EwDatabase) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EwDatabase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EwDatabase) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *EwDatabase) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type EwDatabasesConnection struct {
	EwDatabases []*EwDatabase `protobuf:"bytes,1,rep,name=ewDatabases" json:"ewDatabases,omitempty"`
}

func (m *EwDatabasesConnection) Reset()                    { *m = EwDatabasesConnection{} }
func (m *EwDatabasesConnection) String() string            { return proto.CompactTextString(m) }
func (*EwDatabasesConnection) ProtoMessage()               {}
func (*EwDatabasesConnection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EwDatabasesConnection) GetEwDatabases() []*EwDatabase {
	if m != nil {
		return m.EwDatabases
	}
	return nil
}

type ListenForChangedEwSongRequest struct {
	EwDatabaseId uint32 `protobuf:"varint,1,opt,name=ewDatabaseId" json:"ewDatabaseId,omitempty"`
}

func (m *ListenForChangedEwSongRequest) Reset()                    { *m = ListenForChangedEwSongRequest{} }
func (m *ListenForChangedEwSongRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenForChangedEwSongRequest) ProtoMessage()               {}
func (*ListenForChangedEwSongRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListenForChangedEwSongRequest) GetEwDatabaseId() uint32 {
	if m != nil {
		return m.EwDatabaseId
	}
	return 0
}

type FetchEwDatabasesRequest struct {
}

func (m *FetchEwDatabasesRequest) Reset()                    { *m = FetchEwDatabasesRequest{} }
func (m *FetchEwDatabasesRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchEwDatabasesRequest) ProtoMessage()               {}
func (*FetchEwDatabasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type FetchEwDatabasesResponse struct {
	EwDatabases []*EwDatabase `protobuf:"bytes,1,rep,name=ewDatabases" json:"ewDatabases,omitempty"`
}

func (m *FetchEwDatabasesResponse) Reset()                    { *m = FetchEwDatabasesResponse{} }
func (m *FetchEwDatabasesResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchEwDatabasesResponse) ProtoMessage()               {}
func (*FetchEwDatabasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FetchEwDatabasesResponse) GetEwDatabases() []*EwDatabase {
	if m != nil {
		return m.EwDatabases
	}
	return nil
}

type FetchEwDatabaseByIdRequest struct {
	EwDatabaseIds []uint32 `protobuf:"varint,1,rep,packed,name=ewDatabaseIds" json:"ewDatabaseIds,omitempty"`
}

func (m *FetchEwDatabaseByIdRequest) Reset()                    { *m = FetchEwDatabaseByIdRequest{} }
func (m *FetchEwDatabaseByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchEwDatabaseByIdRequest) ProtoMessage()               {}
func (*FetchEwDatabaseByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FetchEwDatabaseByIdRequest) GetEwDatabaseIds() []uint32 {
	if m != nil {
		return m.EwDatabaseIds
	}
	return nil
}

type FetchEwDatabaseByIdResponse struct {
	EwDatabases []*EwDatabase `protobuf:"bytes,1,rep,name=ewDatabases" json:"ewDatabases,omitempty"`
}

func (m *FetchEwDatabaseByIdResponse) Reset()                    { *m = FetchEwDatabaseByIdResponse{} }
func (m *FetchEwDatabaseByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchEwDatabaseByIdResponse) ProtoMessage()               {}
func (*FetchEwDatabaseByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FetchEwDatabaseByIdResponse) GetEwDatabases() []*EwDatabase {
	if m != nil {
		return m.EwDatabases
	}
	return nil
}

type CreateEwDatabaseRequest struct {
	SongDatabaseId uint32 `protobuf:"varint,1,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *CreateEwDatabaseRequest) Reset()                    { *m = CreateEwDatabaseRequest{} }
func (m *CreateEwDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateEwDatabaseRequest) ProtoMessage()               {}
func (*CreateEwDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateEwDatabaseRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *CreateEwDatabaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateEwDatabaseResponse struct {
	EwDatabase *EwDatabase `protobuf:"bytes,1,opt,name=ewDatabase" json:"ewDatabase,omitempty"`
}

func (m *CreateEwDatabaseResponse) Reset()                    { *m = CreateEwDatabaseResponse{} }
func (m *CreateEwDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateEwDatabaseResponse) ProtoMessage()               {}
func (*CreateEwDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateEwDatabaseResponse) GetEwDatabase() *EwDatabase {
	if m != nil {
		return m.EwDatabase
	}
	return nil
}

type EditEwDatabaseRequest struct {
	EwDatabaseId   uint32 `protobuf:"varint,1,opt,name=ewDatabaseId" json:"ewDatabaseId,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SongDatabaseId uint32 `protobuf:"varint,3,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
}

func (m *EditEwDatabaseRequest) Reset()                    { *m = EditEwDatabaseRequest{} }
func (m *EditEwDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*EditEwDatabaseRequest) ProtoMessage()               {}
func (*EditEwDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *EditEwDatabaseRequest) GetEwDatabaseId() uint32 {
	if m != nil {
		return m.EwDatabaseId
	}
	return 0
}

func (m *EditEwDatabaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EditEwDatabaseRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

type EditEwDatabaseResponse struct {
	EwDatabase *EwDatabase `protobuf:"bytes,1,opt,name=ewDatabase" json:"ewDatabase,omitempty"`
}

func (m *EditEwDatabaseResponse) Reset()                    { *m = EditEwDatabaseResponse{} }
func (m *EditEwDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*EditEwDatabaseResponse) ProtoMessage()               {}
func (*EditEwDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *EditEwDatabaseResponse) GetEwDatabase() *EwDatabase {
	if m != nil {
		return m.EwDatabase
	}
	return nil
}

type RemoveEwDatabaseRequest struct {
	EwDatabaseId uint32 `protobuf:"varint,1,opt,name=ewDatabaseId" json:"ewDatabaseId,omitempty"`
}

func (m *RemoveEwDatabaseRequest) Reset()                    { *m = RemoveEwDatabaseRequest{} }
func (m *RemoveEwDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveEwDatabaseRequest) ProtoMessage()               {}
func (*RemoveEwDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RemoveEwDatabaseRequest) GetEwDatabaseId() uint32 {
	if m != nil {
		return m.EwDatabaseId
	}
	return 0
}

type RemoveEwDatabaseResponse struct {
}

func (m *RemoveEwDatabaseResponse) Reset()                    { *m = RemoveEwDatabaseResponse{} }
func (m *RemoveEwDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveEwDatabaseResponse) ProtoMessage()               {}
func (*RemoveEwDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type SyncEwDatabaseRequest struct {
	EwDatabaseId uint32    `protobuf:"varint,1,opt,name=EwDatabaseId" json:"EwDatabaseId,omitempty"`
	EwSongs      []*EwSong `protobuf:"bytes,2,rep,name=ewSongs" json:"ewSongs,omitempty"`
}

func (m *SyncEwDatabaseRequest) Reset()                    { *m = SyncEwDatabaseRequest{} }
func (m *SyncEwDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncEwDatabaseRequest) ProtoMessage()               {}
func (*SyncEwDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SyncEwDatabaseRequest) GetEwDatabaseId() uint32 {
	if m != nil {
		return m.EwDatabaseId
	}
	return 0
}

func (m *SyncEwDatabaseRequest) GetEwSongs() []*EwSong {
	if m != nil {
		return m.EwSongs
	}
	return nil
}

type SyncEwDatabaseResponse struct {
	EwSongs []*EwSong `protobuf:"bytes,1,rep,name=ewSongs" json:"ewSongs,omitempty"`
}

func (m *SyncEwDatabaseResponse) Reset()                    { *m = SyncEwDatabaseResponse{} }
func (m *SyncEwDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*SyncEwDatabaseResponse) ProtoMessage()               {}
func (*SyncEwDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SyncEwDatabaseResponse) GetEwSongs() []*EwSong {
	if m != nil {
		return m.EwSongs
	}
	return nil
}

func init() {
	proto.RegisterType((*EwDatabase)(nil), "SeppoService.EwDatabase")
	proto.RegisterType((*EwDatabasesConnection)(nil), "SeppoService.EwDatabasesConnection")
	proto.RegisterType((*ListenForChangedEwSongRequest)(nil), "SeppoService.ListenForChangedEwSongRequest")
	proto.RegisterType((*FetchEwDatabasesRequest)(nil), "SeppoService.FetchEwDatabasesRequest")
	proto.RegisterType((*FetchEwDatabasesResponse)(nil), "SeppoService.FetchEwDatabasesResponse")
	proto.RegisterType((*FetchEwDatabaseByIdRequest)(nil), "SeppoService.FetchEwDatabaseByIdRequest")
	proto.RegisterType((*FetchEwDatabaseByIdResponse)(nil), "SeppoService.FetchEwDatabaseByIdResponse")
	proto.RegisterType((*CreateEwDatabaseRequest)(nil), "SeppoService.CreateEwDatabaseRequest")
	proto.RegisterType((*CreateEwDatabaseResponse)(nil), "SeppoService.CreateEwDatabaseResponse")
	proto.RegisterType((*EditEwDatabaseRequest)(nil), "SeppoService.EditEwDatabaseRequest")
	proto.RegisterType((*EditEwDatabaseResponse)(nil), "SeppoService.EditEwDatabaseResponse")
	proto.RegisterType((*RemoveEwDatabaseRequest)(nil), "SeppoService.RemoveEwDatabaseRequest")
	proto.RegisterType((*RemoveEwDatabaseResponse)(nil), "SeppoService.RemoveEwDatabaseResponse")
	proto.RegisterType((*SyncEwDatabaseRequest)(nil), "SeppoService.SyncEwDatabaseRequest")
	proto.RegisterType((*SyncEwDatabaseResponse)(nil), "SeppoService.SyncEwDatabaseResponse")
}

func init() { proto.RegisterFile("ew_database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0xab, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x5e, 0xf1, 0xe1, 0xbc, 0xa6, 0xe0, 0x62, 0xdb, 0xb5, 0x22, 0x84, 0x45, 0x24,
	0xa7, 0x1c, 0xf4, 0x22, 0x82, 0x97, 0xc6, 0x14, 0x0b, 0x9e, 0x36, 0x2a, 0x78, 0x2a, 0x69, 0x32,
	0xb4, 0x41, 0xba, 0x1b, 0xb3, 0x6b, 0x43, 0xff, 0x7b, 0x31, 0x4d, 0xc8, 0x36, 0x09, 0x52, 0xec,
	0xbb, 0xb5, 0xb3, 0xdf, 0xf7, 0xcd, 0x6f, 0x92, 0xd9, 0xc0, 0x33, 0x2c, 0x37, 0x69, 0xac, 0xe3,
	0x6d, 0xac, 0xd0, 0xcf, 0x0b, 0xa9, 0x25, 0x19, 0x47, 0x98, 0xe7, 0x32, 0xc2, 0xe2, 0x98, 0x25,
	0xb8, 0x70, 0xb0, 0xdc, 0x28, 0x29, 0x76, 0xe7, 0x43, 0x56, 0x00, 0x84, 0xe5, 0xa7, 0xda, 0x40,
	0x26, 0x60, 0x67, 0x29, 0xb5, 0x5c, 0xcb, 0x73, 0xb8, 0x9d, 0xa5, 0x84, 0xc0, 0x48, 0xc4, 0x07,
	0xa4, 0xb6, 0x6b, 0x79, 0x4f, 0x79, 0xf5, 0x9b, 0xbc, 0x81, 0xc9, 0x5f, 0x7f, 0xe3, 0x59, 0xa7,
	0xf4, 0xae, 0xd2, 0x77, 0xaa, 0x84, 0xc2, 0xfd, 0x11, 0x0b, 0x95, 0x49, 0x41, 0x47, 0xae, 0xe5,
	0x8d, 0x78, 0xf3, 0x97, 0x45, 0x30, 0x6d, 0x7b, 0xaa, 0x40, 0x0a, 0x81, 0x89, 0xce, 0xa4, 0x20,
	0x1f, 0xe0, 0x01, 0xdb, 0x03, 0x6a, 0xb9, 0x77, 0xde, 0xc3, 0x5b, 0xea, 0x9b, 0xfc, 0x7e, 0xeb,
	0xe4, 0xa6, 0x98, 0x05, 0xf0, 0xea, 0x4b, 0xa6, 0x34, 0x8a, 0x95, 0x2c, 0x82, 0x7d, 0x2c, 0x76,
	0x98, 0x86, 0x65, 0x24, 0xc5, 0x8e, 0xe3, 0xaf, 0xdf, 0xa8, 0x34, 0x61, 0x30, 0x6e, 0xf5, 0xeb,
	0x66, 0xca, 0x8b, 0x1a, 0x7b, 0x01, 0xf3, 0x15, 0xea, 0x64, 0x6f, 0xe0, 0xd5, 0x76, 0xf6, 0x1d,
	0x68, 0xff, 0x48, 0xe5, 0x52, 0x28, 0xbc, 0x89, 0x7b, 0x09, 0x8b, 0x4e, 0xee, 0xf2, 0xb4, 0x4e,
	0x1b, 0xe8, 0xd7, 0xe0, 0x98, 0x80, 0xe7, 0x6c, 0x87, 0x5f, 0x16, 0xd9, 0x0f, 0x78, 0x39, 0x98,
	0xf1, 0x08, 0x78, 0xdf, 0x60, 0x1e, 0x14, 0x18, 0x6b, 0x34, 0x04, 0x35, 0x5b, 0x7f, 0x11, 0xac,
	0xc1, 0x45, 0x18, 0x58, 0x22, 0xf6, 0x15, 0x68, 0x3f, 0xb6, 0xc6, 0x7d, 0x0f, 0xd0, 0x12, 0x54,
	0x99, 0xff, 0xa2, 0x35, 0xb4, 0xac, 0x84, 0x69, 0x98, 0x66, 0xba, 0x8f, 0x7a, 0xc5, 0xbb, 0xbf,
	0x65, 0xd7, 0x19, 0x87, 0x59, 0xb7, 0xf1, 0xcd, 0xc3, 0x7c, 0x84, 0x39, 0xc7, 0x83, 0x3c, 0xe2,
	0x7f, 0x8d, 0xc3, 0x16, 0x40, 0xfb, 0xf6, 0x33, 0x14, 0xfb, 0x09, 0xd3, 0xe8, 0x24, 0x92, 0xc1,
	0xe0, 0x70, 0x20, 0xd8, 0xac, 0x11, 0x1f, 0xee, 0xb1, 0xba, 0x58, 0x8a, 0xda, 0xd5, 0x26, 0x3d,
	0xef, 0x8e, 0x53, 0xdd, 0xba, 0x46, 0xc4, 0x3e, 0xc3, 0xac, 0xdb, 0xac, 0x7e, 0x36, 0x46, 0x92,
	0x75, 0x45, 0xd2, 0xf6, 0x49, 0xf5, 0xc9, 0x7a, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd8,
	0xd4, 0x05, 0xe4, 0x04, 0x00, 0x00,
}
