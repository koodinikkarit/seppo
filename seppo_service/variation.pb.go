// Code generated by protoc-gen-go.
// source: variation.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Variation struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SongId      uint64 `protobuf:"varint,3,opt,name=songId" json:"songId,omitempty"`
	LanguageId  uint64 `protobuf:"varint,4,opt,name=languageId" json:"languageId,omitempty"`
	AuthorId    uint64 `protobuf:"varint,5,opt,name=authorId" json:"authorId,omitempty"`
	CopyrightId uint64 `protobuf:"varint,6,opt,name=copyrightId" json:"copyrightId,omitempty"`
}

func (m *Variation) Reset()                    { *m = Variation{} }
func (m *Variation) String() string            { return proto.CompactTextString(m) }
func (*Variation) ProtoMessage()               {}
func (*Variation) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *Variation) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Variation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Variation) GetSongId() uint64 {
	if m != nil {
		return m.SongId
	}
	return 0
}

func (m *Variation) GetLanguageId() uint64 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *Variation) GetAuthorId() uint64 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func (m *Variation) GetCopyrightId() uint64 {
	if m != nil {
		return m.CopyrightId
	}
	return 0
}

type SearchVariationsRequest struct {
	SearchWord           string   `protobuf:"bytes,1,opt,name=searchWord" json:"searchWord,omitempty"`
	SongDatabaseId       uint64   `protobuf:"varint,2,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	SongDatabaseFilterId uint64   `protobuf:"varint,3,opt,name=songDatabaseFilterId" json:"songDatabaseFilterId,omitempty"`
	TagId                uint64   `protobuf:"varint,4,opt,name=tagId" json:"tagId,omitempty"`
	LanguageId           uint64   `protobuf:"varint,5,opt,name=languageId" json:"languageId,omitempty"`
	Offset               uint32   `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,7,opt,name=limit" json:"limit,omitempty"`
	ScheduleId           uint64   `protobuf:"varint,8,opt,name=scheduleId" json:"scheduleId,omitempty"`
	SkipVariationIds     []uint32 `protobuf:"varint,9,rep,packed,name=skipVariationIds" json:"skipVariationIds,omitempty"`
	OrderBy              uint32   `protobuf:"varint,10,opt,name=orderBy" json:"orderBy,omitempty"`
	SearchFrom           uint32   `protobuf:"varint,11,opt,name=searchFrom" json:"searchFrom,omitempty"`
}

func (m *SearchVariationsRequest) Reset()                    { *m = SearchVariationsRequest{} }
func (m *SearchVariationsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchVariationsRequest) ProtoMessage()               {}
func (*SearchVariationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *SearchVariationsRequest) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

func (m *SearchVariationsRequest) GetSongDatabaseId() uint64 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *SearchVariationsRequest) GetSongDatabaseFilterId() uint64 {
	if m != nil {
		return m.SongDatabaseFilterId
	}
	return 0
}

func (m *SearchVariationsRequest) GetTagId() uint64 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *SearchVariationsRequest) GetLanguageId() uint64 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *SearchVariationsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchVariationsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchVariationsRequest) GetScheduleId() uint64 {
	if m != nil {
		return m.ScheduleId
	}
	return 0
}

func (m *SearchVariationsRequest) GetSkipVariationIds() []uint32 {
	if m != nil {
		return m.SkipVariationIds
	}
	return nil
}

func (m *SearchVariationsRequest) GetOrderBy() uint32 {
	if m != nil {
		return m.OrderBy
	}
	return 0
}

func (m *SearchVariationsRequest) GetSearchFrom() uint32 {
	if m != nil {
		return m.SearchFrom
	}
	return 0
}

type SearchVariationsResponse struct {
	Variations    []*Variation `protobuf:"bytes,1,rep,name=variations" json:"variations,omitempty"`
	MaxVariations uint64       `protobuf:"varint,2,opt,name=maxVariations" json:"maxVariations,omitempty"`
}

func (m *SearchVariationsResponse) Reset()                    { *m = SearchVariationsResponse{} }
func (m *SearchVariationsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchVariationsResponse) ProtoMessage()               {}
func (*SearchVariationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *SearchVariationsResponse) GetVariations() []*Variation {
	if m != nil {
		return m.Variations
	}
	return nil
}

func (m *SearchVariationsResponse) GetMaxVariations() uint64 {
	if m != nil {
		return m.MaxVariations
	}
	return 0
}

type FetchVariationByIdRequest struct {
	VariationIds []uint64 `protobuf:"varint,1,rep,packed,name=variationIds" json:"variationIds,omitempty"`
}

func (m *FetchVariationByIdRequest) Reset()                    { *m = FetchVariationByIdRequest{} }
func (m *FetchVariationByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchVariationByIdRequest) ProtoMessage()               {}
func (*FetchVariationByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *FetchVariationByIdRequest) GetVariationIds() []uint64 {
	if m != nil {
		return m.VariationIds
	}
	return nil
}

type FetchVariationByIdResponse struct {
	Variations []*Variation `protobuf:"bytes,1,rep,name=variations" json:"variations,omitempty"`
}

func (m *FetchVariationByIdResponse) Reset()                    { *m = FetchVariationByIdResponse{} }
func (m *FetchVariationByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchVariationByIdResponse) ProtoMessage()               {}
func (*FetchVariationByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *FetchVariationByIdResponse) GetVariations() []*Variation {
	if m != nil {
		return m.Variations
	}
	return nil
}

type CreateVariationRequest struct {
	Name            string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Text            string   `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	TagIds          []uint64 `protobuf:"varint,3,rep,packed,name=tagIds" json:"tagIds,omitempty"`
	SongDatabaseIds []uint64 `protobuf:"varint,4,rep,packed,name=songDatabaseIds" json:"songDatabaseIds,omitempty"`
}

func (m *CreateVariationRequest) Reset()                    { *m = CreateVariationRequest{} }
func (m *CreateVariationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVariationRequest) ProtoMessage()               {}
func (*CreateVariationRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *CreateVariationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateVariationRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CreateVariationRequest) GetTagIds() []uint64 {
	if m != nil {
		return m.TagIds
	}
	return nil
}

func (m *CreateVariationRequest) GetSongDatabaseIds() []uint64 {
	if m != nil {
		return m.SongDatabaseIds
	}
	return nil
}

type CreateVariationResponse struct {
	Variation *Variation `protobuf:"bytes,1,opt,name=variation" json:"variation,omitempty"`
}

func (m *CreateVariationResponse) Reset()                    { *m = CreateVariationResponse{} }
func (m *CreateVariationResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateVariationResponse) ProtoMessage()               {}
func (*CreateVariationResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *CreateVariationResponse) GetVariation() *Variation {
	if m != nil {
		return m.Variation
	}
	return nil
}

type UpdateVariationRequest struct {
	VariationId           uint64   `protobuf:"varint,1,opt,name=variationId" json:"variationId,omitempty"`
	Name                  string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Text                  string   `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	SongId                uint64   `protobuf:"varint,4,opt,name=songId" json:"songId,omitempty"`
	LanguageId            uint64   `protobuf:"varint,5,opt,name=languageId" json:"languageId,omitempty"`
	AddTagIds             []uint64 `protobuf:"varint,6,rep,packed,name=addTagIds" json:"addTagIds,omitempty"`
	RemoveTagIds          []uint64 `protobuf:"varint,7,rep,packed,name=removeTagIds" json:"removeTagIds,omitempty"`
	AddSongDatabaseIds    []uint64 `protobuf:"varint,8,rep,packed,name=addSongDatabaseIds" json:"addSongDatabaseIds,omitempty"`
	RemoveSongDatabaseIds []uint64 `protobuf:"varint,9,rep,packed,name=removeSongDatabaseIds" json:"removeSongDatabaseIds,omitempty"`
}

func (m *UpdateVariationRequest) Reset()                    { *m = UpdateVariationRequest{} }
func (m *UpdateVariationRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateVariationRequest) ProtoMessage()               {}
func (*UpdateVariationRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *UpdateVariationRequest) GetVariationId() uint64 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

func (m *UpdateVariationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateVariationRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *UpdateVariationRequest) GetSongId() uint64 {
	if m != nil {
		return m.SongId
	}
	return 0
}

func (m *UpdateVariationRequest) GetLanguageId() uint64 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *UpdateVariationRequest) GetAddTagIds() []uint64 {
	if m != nil {
		return m.AddTagIds
	}
	return nil
}

func (m *UpdateVariationRequest) GetRemoveTagIds() []uint64 {
	if m != nil {
		return m.RemoveTagIds
	}
	return nil
}

func (m *UpdateVariationRequest) GetAddSongDatabaseIds() []uint64 {
	if m != nil {
		return m.AddSongDatabaseIds
	}
	return nil
}

func (m *UpdateVariationRequest) GetRemoveSongDatabaseIds() []uint64 {
	if m != nil {
		return m.RemoveSongDatabaseIds
	}
	return nil
}

type UpdateVariationResponse struct {
	Variation *Variation `protobuf:"bytes,1,opt,name=variation" json:"variation,omitempty"`
	Success   bool       `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *UpdateVariationResponse) Reset()                    { *m = UpdateVariationResponse{} }
func (m *UpdateVariationResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateVariationResponse) ProtoMessage()               {}
func (*UpdateVariationResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *UpdateVariationResponse) GetVariation() *Variation {
	if m != nil {
		return m.Variation
	}
	return nil
}

func (m *UpdateVariationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RemoveVariationRequest struct {
	VariationId uint64 `protobuf:"varint,1,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *RemoveVariationRequest) Reset()                    { *m = RemoveVariationRequest{} }
func (m *RemoveVariationRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveVariationRequest) ProtoMessage()               {}
func (*RemoveVariationRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *RemoveVariationRequest) GetVariationId() uint64 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type RemoveVariationResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *RemoveVariationResponse) Reset()                    { *m = RemoveVariationResponse{} }
func (m *RemoveVariationResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveVariationResponse) ProtoMessage()               {}
func (*RemoveVariationResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

func (m *RemoveVariationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Variation)(nil), "SeppoService.Variation")
	proto.RegisterType((*SearchVariationsRequest)(nil), "SeppoService.SearchVariationsRequest")
	proto.RegisterType((*SearchVariationsResponse)(nil), "SeppoService.SearchVariationsResponse")
	proto.RegisterType((*FetchVariationByIdRequest)(nil), "SeppoService.FetchVariationByIdRequest")
	proto.RegisterType((*FetchVariationByIdResponse)(nil), "SeppoService.FetchVariationByIdResponse")
	proto.RegisterType((*CreateVariationRequest)(nil), "SeppoService.CreateVariationRequest")
	proto.RegisterType((*CreateVariationResponse)(nil), "SeppoService.CreateVariationResponse")
	proto.RegisterType((*UpdateVariationRequest)(nil), "SeppoService.UpdateVariationRequest")
	proto.RegisterType((*UpdateVariationResponse)(nil), "SeppoService.UpdateVariationResponse")
	proto.RegisterType((*RemoveVariationRequest)(nil), "SeppoService.RemoveVariationRequest")
	proto.RegisterType((*RemoveVariationResponse)(nil), "SeppoService.RemoveVariationResponse")
}

func init() { proto.RegisterFile("variation.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xd2, 0xac, 0x6b, 0x4e, 0xd7, 0x0d, 0x59, 0xa3, 0x35, 0x13, 0x42, 0x91, 0x85, 0x50,
	0xc5, 0x45, 0x2f, 0x36, 0x10, 0x12, 0x37, 0x48, 0x03, 0x55, 0xca, 0x1d, 0x4a, 0x19, 0x5c, 0x7b,
	0xb1, 0xd7, 0x06, 0xda, 0x3a, 0xd8, 0x6e, 0xb5, 0xde, 0xf2, 0x2e, 0xbc, 0x08, 0x2f, 0x06, 0x8a,
	0xf3, 0xe7, 0x36, 0xd1, 0x24, 0xb4, 0xbb, 0x9c, 0xef, 0x1c, 0xfb, 0x7c, 0xdf, 0xf9, 0x71, 0xe0,
	0x6c, 0x4b, 0x65, 0x42, 0x75, 0x22, 0xd6, 0x93, 0x54, 0x0a, 0x2d, 0xd0, 0xc9, 0x8c, 0xa7, 0xa9,
	0x98, 0x71, 0xb9, 0x4d, 0x62, 0x4e, 0x7e, 0x3b, 0xe0, 0x7f, 0x2d, 0x23, 0xd0, 0x29, 0xb8, 0x09,
	0xc3, 0x4e, 0xe0, 0x8c, 0xbd, 0xc8, 0x4d, 0x18, 0x42, 0xe0, 0xad, 0xe9, 0x8a, 0x63, 0x37, 0x70,
	0xc6, 0x7e, 0x64, 0xbe, 0xd1, 0x10, 0xba, 0x4a, 0xac, 0xe7, 0x21, 0xc3, 0x1d, 0x13, 0x57, 0x58,
	0xe8, 0x05, 0xc0, 0x92, 0xae, 0xe7, 0x1b, 0x3a, 0xe7, 0x21, 0xc3, 0x9e, 0xf1, 0x59, 0x08, 0xba,
	0x80, 0x1e, 0xdd, 0xe8, 0x85, 0x90, 0x21, 0xc3, 0x47, 0xc6, 0x5b, 0xd9, 0x28, 0x80, 0x7e, 0x2c,
	0xd2, 0x9d, 0x4c, 0xe6, 0x0b, 0x1d, 0x32, 0xdc, 0x35, 0x6e, 0x1b, 0x22, 0x7f, 0x5d, 0x18, 0xcd,
	0x38, 0x95, 0xf1, 0xa2, 0x62, 0xab, 0x22, 0xfe, 0x73, 0xc3, 0x95, 0xce, 0x32, 0x2b, 0xe3, 0xfa,
	0x26, 0x64, 0xce, 0xde, 0x8f, 0x2c, 0x04, 0xbd, 0x82, 0xd3, 0x8c, 0xe3, 0x27, 0xaa, 0xe9, 0x2d,
	0x55, 0x19, 0x3b, 0xd7, 0x24, 0x38, 0x40, 0xd1, 0x25, 0x9c, 0xdb, 0xc8, 0x34, 0x59, 0x6a, 0x2e,
	0x2b, 0x9d, 0xad, 0x3e, 0x74, 0x0e, 0x47, 0x9a, 0xce, 0x2b, 0xc1, 0xb9, 0x71, 0x50, 0x8b, 0xa3,
	0x46, 0x2d, 0x86, 0xd0, 0x15, 0x77, 0x77, 0x8a, 0x6b, 0x23, 0x75, 0x10, 0x15, 0x56, 0x76, 0xdb,
	0x32, 0x59, 0x25, 0x1a, 0x1f, 0x1b, 0x38, 0x37, 0x8c, 0xbe, 0x78, 0xc1, 0xd9, 0x66, 0x99, 0xdd,
	0xd6, 0xcb, 0x6f, 0xab, 0x11, 0xf4, 0x1a, 0x9e, 0xa8, 0x1f, 0x49, 0x5a, 0x15, 0x26, 0x64, 0x0a,
	0xfb, 0x41, 0x67, 0x3c, 0x88, 0x1a, 0x38, 0xc2, 0x70, 0x2c, 0x24, 0xe3, 0xf2, 0x7a, 0x87, 0xc1,
	0xe4, 0x28, 0xcd, 0xba, 0x8a, 0x53, 0x29, 0x56, 0xb8, 0x6f, 0x9c, 0x16, 0x42, 0x76, 0x80, 0x9b,
	0x0d, 0x50, 0xa9, 0x58, 0x2b, 0x8e, 0xde, 0x01, 0x54, 0x63, 0xa6, 0xb0, 0x13, 0x74, 0xc6, 0xfd,
	0xcb, 0xd1, 0xc4, 0x1e, 0xb4, 0x49, 0x75, 0x2a, 0xb2, 0x42, 0xd1, 0x4b, 0x18, 0xac, 0xe8, 0x7d,
	0x7d, 0x63, 0xd1, 0x99, 0x7d, 0x90, 0x7c, 0x80, 0x67, 0x53, 0xae, 0xad, 0xcc, 0xd7, 0xbb, 0x90,
	0x95, 0xdd, 0x27, 0x70, 0xb2, 0xb5, 0x95, 0x67, 0xd9, 0xbd, 0x68, 0x0f, 0x23, 0x37, 0x70, 0xd1,
	0x76, 0xc1, 0x23, 0xd9, 0x93, 0x5f, 0x0e, 0x0c, 0x3f, 0x4a, 0x4e, 0x35, 0xaf, 0xfd, 0x05, 0xab,
	0x72, 0x73, 0x1c, 0x6b, 0x73, 0x10, 0x78, 0x9a, 0xdf, 0xeb, 0x72, 0x9b, 0xb2, 0xef, 0x6c, 0x12,
	0xcc, 0xc8, 0x28, 0xdc, 0x31, 0xbc, 0x0b, 0x0b, 0x8d, 0xe1, 0x6c, 0x7f, 0x3a, 0x15, 0xf6, 0x4c,
	0xc0, 0x21, 0x4c, 0x3e, 0xc3, 0xa8, 0xc1, 0xa1, 0x10, 0xf6, 0x16, 0xfc, 0x8a, 0xad, 0x61, 0xf2,
	0x80, 0xae, 0x3a, 0x92, 0xfc, 0x71, 0x61, 0x78, 0x93, 0xb2, 0x36, 0x59, 0x01, 0xf4, 0xad, 0xc2,
	0x16, 0x2f, 0x85, 0x0d, 0xb5, 0x3e, 0x19, 0xa5, 0xf0, 0xce, 0xbe, 0xf0, 0xe2, 0x19, 0xf1, 0x1e,
	0x78, 0x46, 0x9a, 0xab, 0xf3, 0x1c, 0x7c, 0xca, 0xd8, 0x97, 0xbc, 0x66, 0x5d, 0x53, 0x92, 0x1a,
	0xc8, 0x86, 0x41, 0xf2, 0x95, 0xd8, 0xf2, 0x22, 0xe0, 0x38, 0x1f, 0x06, 0x1b, 0x43, 0x13, 0x40,
	0x94, 0xb1, 0xd9, 0x41, 0x75, 0x7b, 0x26, 0xb2, 0xc5, 0x83, 0xde, 0xc0, 0xd3, 0xfc, 0xfc, 0xe1,
	0x11, 0xdf, 0x1c, 0x69, 0x77, 0x92, 0xef, 0x30, 0x6a, 0xd4, 0xf0, 0x51, 0x6d, 0xc9, 0x56, 0x57,
	0x6d, 0xe2, 0x98, 0xab, 0x7c, 0x4b, 0x7a, 0x51, 0x69, 0x92, 0xf7, 0x30, 0x8c, 0x0c, 0x89, 0xff,
	0xef, 0x17, 0xb9, 0x82, 0x51, 0xe3, 0x6c, 0xc1, 0xd3, 0x4a, 0xe8, 0xec, 0x25, 0xbc, 0xed, 0x9a,
	0x5f, 0xc9, 0xd5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x49, 0x9f, 0x40, 0x5d, 0x06, 0x00,
	0x00,
}
