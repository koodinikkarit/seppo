// Code generated by protoc-gen-go.
// source: song_database.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SongDatabase struct {
	Id      uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Version uint64 `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
}

func (m *SongDatabase) Reset()                    { *m = SongDatabase{} }
func (m *SongDatabase) String() string            { return proto.CompactTextString(m) }
func (*SongDatabase) ProtoMessage()               {}
func (*SongDatabase) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SongDatabase) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SongDatabase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SongDatabase) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type SongDatabaseEdge struct {
	Node   *SongDatabase `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Cursor uint32        `protobuf:"varint,2,opt,name=cursor" json:"cursor,omitempty"`
}

func (m *SongDatabaseEdge) Reset()                    { *m = SongDatabaseEdge{} }
func (m *SongDatabaseEdge) String() string            { return proto.CompactTextString(m) }
func (*SongDatabaseEdge) ProtoMessage()               {}
func (*SongDatabaseEdge) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *SongDatabaseEdge) GetNode() *SongDatabase {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *SongDatabaseEdge) GetCursor() uint32 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

type SongDatabasesConnection struct {
	Edges      []*SongDatabaseEdge `protobuf:"bytes,1,rep,name=edges" json:"edges,omitempty"`
	TotalCount uint32              `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *SongDatabasesConnection) Reset()                    { *m = SongDatabasesConnection{} }
func (m *SongDatabasesConnection) String() string            { return proto.CompactTextString(m) }
func (*SongDatabasesConnection) ProtoMessage()               {}
func (*SongDatabasesConnection) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *SongDatabasesConnection) GetEdges() []*SongDatabaseEdge {
	if m != nil {
		return m.Edges
	}
	return nil
}

func (m *SongDatabasesConnection) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type FetchSongDatabasesRequest struct {
	After  uint32 `protobuf:"varint,1,opt,name=after" json:"after,omitempty"`
	Before uint32 `protobuf:"varint,2,opt,name=before" json:"before,omitempty"`
	First  uint32 `protobuf:"varint,3,opt,name=first" json:"first,omitempty"`
	Last   uint32 `protobuf:"varint,4,opt,name=last" json:"last,omitempty"`
}

func (m *FetchSongDatabasesRequest) Reset()                    { *m = FetchSongDatabasesRequest{} }
func (m *FetchSongDatabasesRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchSongDatabasesRequest) ProtoMessage()               {}
func (*FetchSongDatabasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *FetchSongDatabasesRequest) GetAfter() uint32 {
	if m != nil {
		return m.After
	}
	return 0
}

func (m *FetchSongDatabasesRequest) GetBefore() uint32 {
	if m != nil {
		return m.Before
	}
	return 0
}

func (m *FetchSongDatabasesRequest) GetFirst() uint32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *FetchSongDatabasesRequest) GetLast() uint32 {
	if m != nil {
		return m.Last
	}
	return 0
}

type FetchSongDatabaseByIdRequest struct {
	SongDatabaseIds []uint32 `protobuf:"varint,1,rep,packed,name=songDatabaseIds" json:"songDatabaseIds,omitempty"`
}

func (m *FetchSongDatabaseByIdRequest) Reset()                    { *m = FetchSongDatabaseByIdRequest{} }
func (m *FetchSongDatabaseByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchSongDatabaseByIdRequest) ProtoMessage()               {}
func (*FetchSongDatabaseByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *FetchSongDatabaseByIdRequest) GetSongDatabaseIds() []uint32 {
	if m != nil {
		return m.SongDatabaseIds
	}
	return nil
}

type FetchSongDatabaseByIdResponse struct {
	SongDatabases []*SongDatabase `protobuf:"bytes,1,rep,name=songDatabases" json:"songDatabases,omitempty"`
}

func (m *FetchSongDatabaseByIdResponse) Reset()                    { *m = FetchSongDatabaseByIdResponse{} }
func (m *FetchSongDatabaseByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchSongDatabaseByIdResponse) ProtoMessage()               {}
func (*FetchSongDatabaseByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *FetchSongDatabaseByIdResponse) GetSongDatabases() []*SongDatabase {
	if m != nil {
		return m.SongDatabases
	}
	return nil
}

type CreateSongDatabaseRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CreateSongDatabaseRequest) Reset()                    { *m = CreateSongDatabaseRequest{} }
func (m *CreateSongDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSongDatabaseRequest) ProtoMessage()               {}
func (*CreateSongDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *CreateSongDatabaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateSongDatabaseResponse struct {
	SongDatabase *SongDatabase `protobuf:"bytes,1,opt,name=songDatabase" json:"songDatabase,omitempty"`
}

func (m *CreateSongDatabaseResponse) Reset()                    { *m = CreateSongDatabaseResponse{} }
func (m *CreateSongDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateSongDatabaseResponse) ProtoMessage()               {}
func (*CreateSongDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *CreateSongDatabaseResponse) GetSongDatabase() *SongDatabase {
	if m != nil {
		return m.SongDatabase
	}
	return nil
}

type EditSongDatabaseRequest struct {
	SongDatabaseId uint32 `protobuf:"varint,1,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *EditSongDatabaseRequest) Reset()                    { *m = EditSongDatabaseRequest{} }
func (m *EditSongDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*EditSongDatabaseRequest) ProtoMessage()               {}
func (*EditSongDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *EditSongDatabaseRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *EditSongDatabaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EditSongDatabaseResponse struct {
	SongDatabase *SongDatabase `protobuf:"bytes,1,opt,name=songDatabase" json:"songDatabase,omitempty"`
}

func (m *EditSongDatabaseResponse) Reset()                    { *m = EditSongDatabaseResponse{} }
func (m *EditSongDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*EditSongDatabaseResponse) ProtoMessage()               {}
func (*EditSongDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *EditSongDatabaseResponse) GetSongDatabase() *SongDatabase {
	if m != nil {
		return m.SongDatabase
	}
	return nil
}

type RemoveSongDatabaseRequest struct {
	SongDatabaseId uint32 `protobuf:"varint,1,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
}

func (m *RemoveSongDatabaseRequest) Reset()                    { *m = RemoveSongDatabaseRequest{} }
func (m *RemoveSongDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveSongDatabaseRequest) ProtoMessage()               {}
func (*RemoveSongDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *RemoveSongDatabaseRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

type RemoveSongDatabaseResponse struct {
}

func (m *RemoveSongDatabaseResponse) Reset()                    { *m = RemoveSongDatabaseResponse{} }
func (m *RemoveSongDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveSongDatabaseResponse) ProtoMessage()               {}
func (*RemoveSongDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

func init() {
	proto.RegisterType((*SongDatabase)(nil), "SeppoService.SongDatabase")
	proto.RegisterType((*SongDatabaseEdge)(nil), "SeppoService.SongDatabaseEdge")
	proto.RegisterType((*SongDatabasesConnection)(nil), "SeppoService.SongDatabasesConnection")
	proto.RegisterType((*FetchSongDatabasesRequest)(nil), "SeppoService.FetchSongDatabasesRequest")
	proto.RegisterType((*FetchSongDatabaseByIdRequest)(nil), "SeppoService.FetchSongDatabaseByIdRequest")
	proto.RegisterType((*FetchSongDatabaseByIdResponse)(nil), "SeppoService.FetchSongDatabaseByIdResponse")
	proto.RegisterType((*CreateSongDatabaseRequest)(nil), "SeppoService.CreateSongDatabaseRequest")
	proto.RegisterType((*CreateSongDatabaseResponse)(nil), "SeppoService.CreateSongDatabaseResponse")
	proto.RegisterType((*EditSongDatabaseRequest)(nil), "SeppoService.EditSongDatabaseRequest")
	proto.RegisterType((*EditSongDatabaseResponse)(nil), "SeppoService.EditSongDatabaseResponse")
	proto.RegisterType((*RemoveSongDatabaseRequest)(nil), "SeppoService.RemoveSongDatabaseRequest")
	proto.RegisterType((*RemoveSongDatabaseResponse)(nil), "SeppoService.RemoveSongDatabaseResponse")
}

func init() { proto.RegisterFile("song_database.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0x26, 0xd7, 0x6b, 0xc5, 0xe9, 0xa5, 0xca, 0x2a, 0x36, 0x77, 0xd4, 0x72, 0xec, 0x83, 0xe4,
	0x29, 0x42, 0xf5, 0x59, 0xc4, 0x58, 0xb1, 0xe0, 0xd3, 0x1e, 0xbe, 0x14, 0x41, 0xf6, 0xb2, 0x73,
	0x31, 0xd0, 0xee, 0xc4, 0xdd, 0xbd, 0x03, 0xff, 0xbd, 0x64, 0x93, 0xe0, 0xa6, 0xcd, 0x89, 0x60,
	0xdf, 0xf6, 0x9b, 0x7c, 0xf3, 0x7d, 0xdf, 0xce, 0x4e, 0xe0, 0x99, 0x25, 0x5d, 0x7e, 0x57, 0xd2,
	0xc9, 0xb5, 0xb4, 0x98, 0xd5, 0x86, 0x1c, 0xb1, 0xd9, 0x0a, 0xeb, 0x9a, 0x56, 0x68, 0x76, 0x55,
	0x81, 0xfc, 0x0b, 0xcc, 0x56, 0xa4, 0xcb, 0x8f, 0x1d, 0x87, 0x9d, 0xc0, 0xa4, 0x52, 0x49, 0xb4,
	0x8c, 0xd2, 0x58, 0x4c, 0x2a, 0xc5, 0x18, 0x4c, 0xb5, 0xbc, 0xc5, 0x64, 0xb2, 0x8c, 0xd2, 0xc7,
	0xc2, 0x9f, 0x59, 0x02, 0x8f, 0x76, 0x68, 0x6c, 0x45, 0x3a, 0x39, 0x58, 0x46, 0xe9, 0x54, 0xf4,
	0x90, 0x5f, 0xc3, 0xd3, 0x50, 0xed, 0x52, 0x95, 0xc8, 0x32, 0x98, 0x6a, 0x52, 0xe8, 0x35, 0x8f,
	0x2f, 0x16, 0x59, 0x68, 0x9f, 0x85, 0x6c, 0xe1, 0x79, 0xec, 0x05, 0x1c, 0x15, 0x5b, 0x63, 0xc9,
	0x78, 0xcf, 0x58, 0x74, 0x88, 0x13, 0x9c, 0x86, 0x6c, 0x9b, 0x93, 0xd6, 0x58, 0xb8, 0x8a, 0x34,
	0x7b, 0x0b, 0x87, 0xa8, 0x4a, 0xb4, 0x49, 0xb4, 0x3c, 0x48, 0x8f, 0x2f, 0xce, 0xf7, 0x7b, 0x34,
	0x89, 0x44, 0x4b, 0x66, 0xe7, 0x00, 0x8e, 0x9c, 0xbc, 0xc9, 0x69, 0xab, 0x9d, 0xbf, 0x49, 0x2c,
	0x82, 0x0a, 0xb7, 0x30, 0xff, 0x84, 0xae, 0xf8, 0x31, 0x70, 0x15, 0xf8, 0x73, 0x8b, 0xd6, 0xb1,
	0xe7, 0x70, 0x28, 0x37, 0x0e, 0x4d, 0x37, 0xaa, 0x16, 0x34, 0xd9, 0xd7, 0xb8, 0x21, 0x83, 0x7d,
	0xf6, 0x16, 0x35, 0xec, 0x4d, 0x65, 0x6c, 0xef, 0xd2, 0x82, 0x66, 0xb6, 0x37, 0xd2, 0xba, 0x64,
	0xea, 0x8b, 0xfe, 0xcc, 0x3f, 0xc3, 0xd9, 0x3d, 0xd3, 0x0f, 0xbf, 0xae, 0x54, 0xef, 0x9b, 0xc2,
	0x13, 0x1b, 0x7c, 0xba, 0x52, 0xed, 0xa5, 0x63, 0x71, 0xb7, 0xcc, 0x25, 0xbc, 0xdc, 0xa3, 0x64,
	0x6b, 0xd2, 0x16, 0xd9, 0x7b, 0x88, 0xc3, 0x9e, 0x7e, 0x7a, 0x7f, 0x7b, 0xa1, 0x61, 0x03, 0x7f,
	0x0d, 0xf3, 0xdc, 0xa0, 0x74, 0x38, 0x20, 0x75, 0x49, 0xfb, 0xcd, 0x89, 0xfe, 0x6c, 0x0e, 0xff,
	0x06, 0x8b, 0xb1, 0x86, 0x2e, 0xd0, 0x3b, 0x98, 0x85, 0xfa, 0xff, 0xb0, 0x31, 0x03, 0x3e, 0xff,
	0x0a, 0xa7, 0x97, 0xaa, 0x72, 0x63, 0x61, 0x5e, 0xc1, 0xc9, 0x70, 0x3e, 0xdd, 0xbb, 0xdd, 0xa9,
	0x8e, 0xad, 0x3b, 0xbf, 0x86, 0xe4, 0xbe, 0xec, 0x03, 0x45, 0xce, 0x61, 0x2e, 0xf0, 0x96, 0x76,
	0xf8, 0x1f, 0xa1, 0xf9, 0x19, 0x2c, 0xc6, 0x44, 0xda, 0x88, 0xeb, 0x23, 0xff, 0xdb, 0xbf, 0xf9,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x0c, 0x7c, 0x81, 0x0d, 0x04, 0x00, 0x00,
}
