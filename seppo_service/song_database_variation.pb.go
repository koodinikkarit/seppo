// Code generated by protoc-gen-go.
// source: song_database_variation.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RemoveVariationFromSongDatabaseResponse_State int32

const (
	RemoveVariationFromSongDatabaseResponse_SUCCESS   RemoveVariationFromSongDatabaseResponse_State = 0
	RemoveVariationFromSongDatabaseResponse_NOT_FOUND RemoveVariationFromSongDatabaseResponse_State = 1
)

var RemoveVariationFromSongDatabaseResponse_State_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_FOUND",
}
var RemoveVariationFromSongDatabaseResponse_State_value = map[string]int32{
	"SUCCESS":   0,
	"NOT_FOUND": 1,
}

func (x RemoveVariationFromSongDatabaseResponse_State) String() string {
	return proto.EnumName(RemoveVariationFromSongDatabaseResponse_State_name, int32(x))
}
func (RemoveVariationFromSongDatabaseResponse_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor11, []int{7, 0}
}

type SongDatabaseVariation struct {
	Id             uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SongDatabaseId uint32 `protobuf:"varint,2,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	VariationId    uint32 `protobuf:"varint,3,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *SongDatabaseVariation) Reset()                    { *m = SongDatabaseVariation{} }
func (m *SongDatabaseVariation) String() string            { return proto.CompactTextString(m) }
func (*SongDatabaseVariation) ProtoMessage()               {}
func (*SongDatabaseVariation) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *SongDatabaseVariation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SongDatabaseVariation) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *SongDatabaseVariation) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type SongDatabaseVariations struct {
	SongDatabaseId uint32       `protobuf:"varint,1,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	Variations     []*Variation `protobuf:"bytes,2,rep,name=variations" json:"variations,omitempty"`
}

func (m *SongDatabaseVariations) Reset()                    { *m = SongDatabaseVariations{} }
func (m *SongDatabaseVariations) String() string            { return proto.CompactTextString(m) }
func (*SongDatabaseVariations) ProtoMessage()               {}
func (*SongDatabaseVariations) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *SongDatabaseVariations) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *SongDatabaseVariations) GetVariations() []*Variation {
	if m != nil {
		return m.Variations
	}
	return nil
}

type FetchVariationsBySongDatabaseIdRequest struct {
	SongDatabaseIds []uint32 `protobuf:"varint,1,rep,packed,name=songDatabaseIds" json:"songDatabaseIds,omitempty"`
}

func (m *FetchVariationsBySongDatabaseIdRequest) Reset() {
	*m = FetchVariationsBySongDatabaseIdRequest{}
}
func (m *FetchVariationsBySongDatabaseIdRequest) String() string { return proto.CompactTextString(m) }
func (*FetchVariationsBySongDatabaseIdRequest) ProtoMessage()    {}
func (*FetchVariationsBySongDatabaseIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{2}
}

func (m *FetchVariationsBySongDatabaseIdRequest) GetSongDatabaseIds() []uint32 {
	if m != nil {
		return m.SongDatabaseIds
	}
	return nil
}

type FetchVariationsBySongDatabaseIdResponse struct {
	SongDatabaseVariations []*SongDatabaseVariations `protobuf:"bytes,1,rep,name=songDatabaseVariations" json:"songDatabaseVariations,omitempty"`
}

func (m *FetchVariationsBySongDatabaseIdResponse) Reset() {
	*m = FetchVariationsBySongDatabaseIdResponse{}
}
func (m *FetchVariationsBySongDatabaseIdResponse) String() string { return proto.CompactTextString(m) }
func (*FetchVariationsBySongDatabaseIdResponse) ProtoMessage()    {}
func (*FetchVariationsBySongDatabaseIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{3}
}

func (m *FetchVariationsBySongDatabaseIdResponse) GetSongDatabaseVariations() []*SongDatabaseVariations {
	if m != nil {
		return m.SongDatabaseVariations
	}
	return nil
}

type AddVariationToSongDatabaseRequest struct {
	SongDatabaseId uint32 `protobuf:"varint,1,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	VariationId    uint32 `protobuf:"varint,2,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *AddVariationToSongDatabaseRequest) Reset()         { *m = AddVariationToSongDatabaseRequest{} }
func (m *AddVariationToSongDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*AddVariationToSongDatabaseRequest) ProtoMessage()    {}
func (*AddVariationToSongDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{4}
}

func (m *AddVariationToSongDatabaseRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *AddVariationToSongDatabaseRequest) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type AddVariationToSongDatabaseResponse struct {
	SongDatabaseVariation *SongDatabaseVariation `protobuf:"bytes,1,opt,name=songDatabaseVariation" json:"songDatabaseVariation,omitempty"`
}

func (m *AddVariationToSongDatabaseResponse) Reset()         { *m = AddVariationToSongDatabaseResponse{} }
func (m *AddVariationToSongDatabaseResponse) String() string { return proto.CompactTextString(m) }
func (*AddVariationToSongDatabaseResponse) ProtoMessage()    {}
func (*AddVariationToSongDatabaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{5}
}

func (m *AddVariationToSongDatabaseResponse) GetSongDatabaseVariation() *SongDatabaseVariation {
	if m != nil {
		return m.SongDatabaseVariation
	}
	return nil
}

type RemoveVariationFromSongDatabaseRequest struct {
	SongDatabaseId uint32 `protobuf:"varint,1,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	VariationId    uint32 `protobuf:"varint,2,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *RemoveVariationFromSongDatabaseRequest) Reset() {
	*m = RemoveVariationFromSongDatabaseRequest{}
}
func (m *RemoveVariationFromSongDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveVariationFromSongDatabaseRequest) ProtoMessage()    {}
func (*RemoveVariationFromSongDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{6}
}

func (m *RemoveVariationFromSongDatabaseRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *RemoveVariationFromSongDatabaseRequest) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type RemoveVariationFromSongDatabaseResponse struct {
	State RemoveVariationFromSongDatabaseResponse_State `protobuf:"varint,1,opt,name=state,enum=SeppoService.RemoveVariationFromSongDatabaseResponse_State" json:"state,omitempty"`
}

func (m *RemoveVariationFromSongDatabaseResponse) Reset() {
	*m = RemoveVariationFromSongDatabaseResponse{}
}
func (m *RemoveVariationFromSongDatabaseResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveVariationFromSongDatabaseResponse) ProtoMessage()    {}
func (*RemoveVariationFromSongDatabaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{7}
}

func (m *RemoveVariationFromSongDatabaseResponse) GetState() RemoveVariationFromSongDatabaseResponse_State {
	if m != nil {
		return m.State
	}
	return RemoveVariationFromSongDatabaseResponse_SUCCESS
}

func init() {
	proto.RegisterType((*SongDatabaseVariation)(nil), "SeppoService.SongDatabaseVariation")
	proto.RegisterType((*SongDatabaseVariations)(nil), "SeppoService.SongDatabaseVariations")
	proto.RegisterType((*FetchVariationsBySongDatabaseIdRequest)(nil), "SeppoService.FetchVariationsBySongDatabaseIdRequest")
	proto.RegisterType((*FetchVariationsBySongDatabaseIdResponse)(nil), "SeppoService.FetchVariationsBySongDatabaseIdResponse")
	proto.RegisterType((*AddVariationToSongDatabaseRequest)(nil), "SeppoService.AddVariationToSongDatabaseRequest")
	proto.RegisterType((*AddVariationToSongDatabaseResponse)(nil), "SeppoService.AddVariationToSongDatabaseResponse")
	proto.RegisterType((*RemoveVariationFromSongDatabaseRequest)(nil), "SeppoService.RemoveVariationFromSongDatabaseRequest")
	proto.RegisterType((*RemoveVariationFromSongDatabaseResponse)(nil), "SeppoService.RemoveVariationFromSongDatabaseResponse")
	proto.RegisterEnum("SeppoService.RemoveVariationFromSongDatabaseResponse_State", RemoveVariationFromSongDatabaseResponse_State_name, RemoveVariationFromSongDatabaseResponse_State_value)
}

func init() { proto.RegisterFile("song_database_variation.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x3f, 0x4f, 0xc2, 0x40,
	0x18, 0xc6, 0x6d, 0x09, 0x1a, 0xdf, 0xca, 0x9f, 0x5c, 0x02, 0x36, 0x26, 0x26, 0xf5, 0x30, 0xd0,
	0xa9, 0x43, 0x1d, 0x1c, 0x9c, 0x14, 0x24, 0x61, 0x81, 0x78, 0x07, 0x26, 0x26, 0x26, 0xa4, 0xd0,
	0x0b, 0x76, 0xa0, 0x57, 0x7a, 0x95, 0x84, 0xc9, 0xd1, 0xcf, 0xe1, 0x37, 0x35, 0x56, 0x28, 0x6d,
	0x3d, 0x85, 0xc5, 0x91, 0xe7, 0x7d, 0xef, 0xf7, 0xe3, 0xb9, 0x4b, 0xe1, 0x5c, 0x70, 0x7f, 0x36,
	0x76, 0x9d, 0xc8, 0x99, 0x38, 0x82, 0x8d, 0x97, 0x4e, 0xe8, 0x39, 0x91, 0xc7, 0x7d, 0x2b, 0x08,
	0x79, 0xc4, 0xd1, 0x09, 0x65, 0x41, 0xc0, 0x29, 0x0b, 0x97, 0xde, 0x94, 0x9d, 0x55, 0x72, 0x63,
	0xbc, 0x80, 0x1a, 0xe5, 0xfe, 0xac, 0xb3, 0x3e, 0xfe, 0xb8, 0x19, 0xa3, 0x32, 0xa8, 0x9e, 0xab,
	0x2b, 0x86, 0x62, 0x96, 0x88, 0xea, 0xb9, 0xa8, 0x09, 0x65, 0x91, 0x5a, 0xec, 0xb9, 0xba, 0x1a,
	0xcf, 0x72, 0x29, 0x32, 0x40, 0x4b, 0x1c, 0x3d, 0x57, 0x2f, 0xc4, 0x4b, 0xe9, 0x08, 0xaf, 0xa0,
	0x2e, 0x55, 0x0a, 0x89, 0x43, 0x91, 0x3a, 0xae, 0x01, 0x12, 0xa0, 0xd0, 0x55, 0xa3, 0x60, 0x6a,
	0xf6, 0xa9, 0x95, 0x2e, 0x6a, 0x25, 0x54, 0x92, 0x5a, 0xc5, 0x04, 0x9a, 0x5d, 0x16, 0x4d, 0x5f,
	0xb6, 0xce, 0xbb, 0x15, 0xcd, 0xb0, 0x09, 0x5b, 0xbc, 0x32, 0x11, 0x21, 0x13, 0x2a, 0x59, 0xa9,
	0xd0, 0x15, 0xa3, 0x60, 0x96, 0x48, 0x3e, 0xc6, 0xef, 0x0a, 0xb4, 0x76, 0x42, 0x45, 0xc0, 0x7d,
	0xc1, 0xd0, 0x33, 0xd4, 0x85, 0xb4, 0x7a, 0x0c, 0xd7, 0xec, 0xcb, 0x6c, 0x09, 0xf9, 0x35, 0x91,
	0x5f, 0x18, 0x78, 0x0e, 0x17, 0xb7, 0xae, 0x9b, 0x04, 0x43, 0x9e, 0x3e, 0xbf, 0x29, 0xb6, 0xef,
	0x1d, 0xe7, 0xde, 0x51, 0xfd, 0xf9, 0x8e, 0x6f, 0x80, 0xff, 0xd2, 0xad, 0x2b, 0x3f, 0x41, 0x4d,
	0xfa, 0x77, 0x63, 0xad, 0x66, 0x37, 0xf6, 0x68, 0x4c, 0xe4, 0x04, 0x1c, 0x42, 0x93, 0xb0, 0x39,
	0x5f, 0x6e, 0xa3, 0x6e, 0xc8, 0xe7, 0xff, 0x5b, 0xfa, 0x43, 0x81, 0xd6, 0x4e, 0xe9, 0xba, 0xfa,
	0x03, 0x14, 0x45, 0xe4, 0x44, 0x2c, 0x96, 0x95, 0xed, 0x9b, 0x6c, 0xd5, 0x3d, 0x29, 0x16, 0xfd,
	0x42, 0x90, 0x6f, 0x12, 0x6e, 0x40, 0x31, 0xfe, 0x8d, 0x34, 0x38, 0xa2, 0xa3, 0x76, 0xfb, 0x9e,
	0xd2, 0xea, 0x01, 0x2a, 0xc1, 0x71, 0x7f, 0x30, 0x1c, 0x77, 0x07, 0xa3, 0x7e, 0xa7, 0xaa, 0x4c,
	0x0e, 0xe3, 0x4f, 0xfb, 0xea, 0x33, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x7b, 0x91, 0xf4, 0x1a, 0x04,
	0x00, 0x00,
}
