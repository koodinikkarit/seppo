// Code generated by protoc-gen-go.
// source: tag_variation.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TagVariation struct {
	Id          uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TagId       uint32 `protobuf:"varint,2,opt,name=tagId" json:"tagId,omitempty"`
	VariationId uint32 `protobuf:"varint,3,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *TagVariation) Reset()                    { *m = TagVariation{} }
func (m *TagVariation) String() string            { return proto.CompactTextString(m) }
func (*TagVariation) ProtoMessage()               {}
func (*TagVariation) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *TagVariation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TagVariation) GetTagId() uint32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *TagVariation) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type VariationTags struct {
	VariationId uint32 `protobuf:"varint,1,opt,name=variationId" json:"variationId,omitempty"`
	Tags        []*Tag `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
}

func (m *VariationTags) Reset()                    { *m = VariationTags{} }
func (m *VariationTags) String() string            { return proto.CompactTextString(m) }
func (*VariationTags) ProtoMessage()               {}
func (*VariationTags) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *VariationTags) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

func (m *VariationTags) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TagVariations struct {
	TagId      uint32       `protobuf:"varint,1,opt,name=tagId" json:"tagId,omitempty"`
	Variations []*Variation `protobuf:"bytes,2,rep,name=variations" json:"variations,omitempty"`
}

func (m *TagVariations) Reset()                    { *m = TagVariations{} }
func (m *TagVariations) String() string            { return proto.CompactTextString(m) }
func (*TagVariations) ProtoMessage()               {}
func (*TagVariations) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *TagVariations) GetTagId() uint32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *TagVariations) GetVariations() []*Variation {
	if m != nil {
		return m.Variations
	}
	return nil
}

type FetchVariationTagsRequest struct {
	VariationIds []uint32 `protobuf:"varint,1,rep,packed,name=variationIds" json:"variationIds,omitempty"`
}

func (m *FetchVariationTagsRequest) Reset()                    { *m = FetchVariationTagsRequest{} }
func (m *FetchVariationTagsRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchVariationTagsRequest) ProtoMessage()               {}
func (*FetchVariationTagsRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *FetchVariationTagsRequest) GetVariationIds() []uint32 {
	if m != nil {
		return m.VariationIds
	}
	return nil
}

type FetchVariationTagsResponse struct {
	VariationTags []*VariationTags `protobuf:"bytes,1,rep,name=variationTags" json:"variationTags,omitempty"`
}

func (m *FetchVariationTagsResponse) Reset()                    { *m = FetchVariationTagsResponse{} }
func (m *FetchVariationTagsResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchVariationTagsResponse) ProtoMessage()               {}
func (*FetchVariationTagsResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *FetchVariationTagsResponse) GetVariationTags() []*VariationTags {
	if m != nil {
		return m.VariationTags
	}
	return nil
}

type FetchTagVariationsRequest struct {
	TagIds []uint32 `protobuf:"varint,1,rep,packed,name=tagIds" json:"tagIds,omitempty"`
}

func (m *FetchTagVariationsRequest) Reset()                    { *m = FetchTagVariationsRequest{} }
func (m *FetchTagVariationsRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchTagVariationsRequest) ProtoMessage()               {}
func (*FetchTagVariationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *FetchTagVariationsRequest) GetTagIds() []uint32 {
	if m != nil {
		return m.TagIds
	}
	return nil
}

type FetchTagVariationsResponse struct {
	TagVariations []*TagVariations `protobuf:"bytes,1,rep,name=tagVariations" json:"tagVariations,omitempty"`
}

func (m *FetchTagVariationsResponse) Reset()                    { *m = FetchTagVariationsResponse{} }
func (m *FetchTagVariationsResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchTagVariationsResponse) ProtoMessage()               {}
func (*FetchTagVariationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *FetchTagVariationsResponse) GetTagVariations() []*TagVariations {
	if m != nil {
		return m.TagVariations
	}
	return nil
}

type AddTagToVariationRequest struct {
	TagId       uint32 `protobuf:"varint,1,opt,name=tagId" json:"tagId,omitempty"`
	VariationId uint32 `protobuf:"varint,2,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *AddTagToVariationRequest) Reset()                    { *m = AddTagToVariationRequest{} }
func (m *AddTagToVariationRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTagToVariationRequest) ProtoMessage()               {}
func (*AddTagToVariationRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *AddTagToVariationRequest) GetTagId() uint32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *AddTagToVariationRequest) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type AddTagToVariationResponse struct {
	Success      bool          `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	TagVariation *TagVariation `protobuf:"bytes,2,opt,name=tagVariation" json:"tagVariation,omitempty"`
}

func (m *AddTagToVariationResponse) Reset()                    { *m = AddTagToVariationResponse{} }
func (m *AddTagToVariationResponse) String() string            { return proto.CompactTextString(m) }
func (*AddTagToVariationResponse) ProtoMessage()               {}
func (*AddTagToVariationResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *AddTagToVariationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AddTagToVariationResponse) GetTagVariation() *TagVariation {
	if m != nil {
		return m.TagVariation
	}
	return nil
}

type RemoveTagFromVariationRequest struct {
	TagId       uint32 `protobuf:"varint,1,opt,name=tagId" json:"tagId,omitempty"`
	VariationId uint32 `protobuf:"varint,2,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *RemoveTagFromVariationRequest) Reset()                    { *m = RemoveTagFromVariationRequest{} }
func (m *RemoveTagFromVariationRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveTagFromVariationRequest) ProtoMessage()               {}
func (*RemoveTagFromVariationRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *RemoveTagFromVariationRequest) GetTagId() uint32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *RemoveTagFromVariationRequest) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type RemoveTagFromVariationResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *RemoveTagFromVariationResponse) Reset()         { *m = RemoveTagFromVariationResponse{} }
func (m *RemoveTagFromVariationResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveTagFromVariationResponse) ProtoMessage()    {}
func (*RemoveTagFromVariationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{10}
}

func (m *RemoveTagFromVariationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*TagVariation)(nil), "SeppoService.TagVariation")
	proto.RegisterType((*VariationTags)(nil), "SeppoService.VariationTags")
	proto.RegisterType((*TagVariations)(nil), "SeppoService.TagVariations")
	proto.RegisterType((*FetchVariationTagsRequest)(nil), "SeppoService.FetchVariationTagsRequest")
	proto.RegisterType((*FetchVariationTagsResponse)(nil), "SeppoService.FetchVariationTagsResponse")
	proto.RegisterType((*FetchTagVariationsRequest)(nil), "SeppoService.FetchTagVariationsRequest")
	proto.RegisterType((*FetchTagVariationsResponse)(nil), "SeppoService.FetchTagVariationsResponse")
	proto.RegisterType((*AddTagToVariationRequest)(nil), "SeppoService.AddTagToVariationRequest")
	proto.RegisterType((*AddTagToVariationResponse)(nil), "SeppoService.AddTagToVariationResponse")
	proto.RegisterType((*RemoveTagFromVariationRequest)(nil), "SeppoService.RemoveTagFromVariationRequest")
	proto.RegisterType((*RemoveTagFromVariationResponse)(nil), "SeppoService.RemoveTagFromVariationResponse")
}

func init() { proto.RegisterFile("tag_variation.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xb1, 0x6f, 0xa3, 0x30,
	0x14, 0x87, 0x05, 0xb9, 0xcb, 0xdd, 0x3d, 0xe0, 0x4e, 0xe7, 0x3b, 0xb5, 0x84, 0xaa, 0x15, 0xb2,
	0x54, 0x29, 0x53, 0x86, 0x64, 0xa8, 0xd4, 0xa1, 0x55, 0x96, 0x48, 0x59, 0x1d, 0x94, 0x76, 0x6a,
	0xe4, 0x82, 0xe5, 0x32, 0x24, 0xa6, 0xd8, 0xe1, 0xef, 0xaf, 0xe4, 0x00, 0xb1, 0x09, 0x51, 0x97,
	0x8e, 0x7e, 0xcf, 0xbf, 0x4f, 0xdf, 0x7b, 0x18, 0xf8, 0xa7, 0x28, 0xdf, 0x54, 0xb4, 0xcc, 0xa9,
	0xca, 0xc5, 0x6e, 0x52, 0x94, 0x42, 0x09, 0xe4, 0xaf, 0x58, 0x51, 0x88, 0x15, 0x2b, 0xab, 0x3c,
	0x65, 0xd1, 0x2f, 0x45, 0xf9, 0xa1, 0x11, 0xfd, 0xe9, 0xdc, 0xc4, 0x6b, 0xf0, 0x13, 0xca, 0xd7,
	0x4d, 0x15, 0xfd, 0x06, 0x37, 0xcf, 0x42, 0x27, 0x76, 0xc6, 0x01, 0x71, 0xf3, 0x0c, 0xfd, 0x87,
	0xef, 0x8a, 0xf2, 0x65, 0x16, 0xba, 0xba, 0x74, 0x38, 0xa0, 0x18, 0xbc, 0x16, 0xb4, 0xcc, 0xc2,
	0x81, 0xee, 0x99, 0x25, 0xfc, 0x0c, 0x41, 0x0b, 0x4d, 0x28, 0x97, 0xdd, 0x88, 0x73, 0x12, 0x41,
	0xb7, 0xf0, 0x4d, 0x51, 0x2e, 0x43, 0x37, 0x1e, 0x8c, 0xbd, 0xe9, 0xdf, 0x89, 0x39, 0xc3, 0x24,
	0xa1, 0x9c, 0xe8, 0x36, 0x7e, 0x81, 0xc0, 0x34, 0x96, 0x47, 0x45, 0xc7, 0x54, 0xbc, 0x03, 0x68,
	0xe1, 0x0d, 0xf3, 0xd2, 0x66, 0xb6, 0x0c, 0x62, 0x5c, 0xc5, 0x8f, 0x30, 0x5a, 0x30, 0x95, 0xbe,
	0x59, 0xfa, 0x84, 0xbd, 0xef, 0x99, 0x54, 0x08, 0x83, 0x6f, 0x28, 0xcb, 0xd0, 0x89, 0x07, 0xe3,
	0x80, 0x58, 0x35, 0xbc, 0x81, 0xa8, 0x0f, 0x20, 0x0b, 0xb1, 0x93, 0x0c, 0xcd, 0x21, 0xa8, 0xcc,
	0x86, 0x46, 0x78, 0xd3, 0xab, 0x33, 0x6a, 0x3a, 0x6b, 0x27, 0xf0, 0xac, 0x36, 0xb4, 0xd6, 0xd0,
	0x18, 0x5e, 0xc0, 0x50, 0x2f, 0xa0, 0x71, 0xab, 0x4f, 0xad, 0x55, 0x27, 0x74, 0xb4, 0x52, 0x66,
	0xa3, 0xdf, 0xca, 0xce, 0xda, 0x09, 0x4c, 0x20, 0x9c, 0x67, 0x59, 0x42, 0x79, 0x22, 0x8e, 0x8b,
	0xad, 0xa5, 0xfa, 0x3f, 0x51, 0xe7, 0x49, 0xb8, 0xa7, 0xaf, 0x68, 0x0f, 0xa3, 0x1e, 0x66, 0xed,
	0x1c, 0xc2, 0x0f, 0xb9, 0x4f, 0x53, 0x26, 0xa5, 0xc6, 0xfe, 0x24, 0xcd, 0x11, 0x3d, 0x80, 0x6f,
	0xba, 0x69, 0xb2, 0x37, 0x8d, 0xce, 0x0f, 0x43, 0xac, 0xfb, 0xf8, 0x09, 0xae, 0x09, 0xdb, 0x8a,
	0x8a, 0x25, 0x94, 0x2f, 0x4a, 0xb1, 0xfd, 0xb2, 0x79, 0xee, 0xe1, 0xe6, 0x1c, 0xf8, 0xb3, 0xa1,
	0x5e, 0x87, 0xfa, 0x87, 0x9d, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x23, 0xbd, 0x56, 0xd8, 0xf1,
	0x03, 0x00, 0x00,
}