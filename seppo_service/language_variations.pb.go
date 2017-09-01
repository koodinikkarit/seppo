// Code generated by protoc-gen-go.
// source: language_variations.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LanguageVariations struct {
	LanguageId uint32       `protobuf:"varint,1,opt,name=languageId" json:"languageId,omitempty"`
	Variations []*Variation `protobuf:"bytes,2,rep,name=variations" json:"variations,omitempty"`
}

func (m *LanguageVariations) Reset()                    { *m = LanguageVariations{} }
func (m *LanguageVariations) String() string            { return proto.CompactTextString(m) }
func (*LanguageVariations) ProtoMessage()               {}
func (*LanguageVariations) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *LanguageVariations) GetLanguageId() uint32 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *LanguageVariations) GetVariations() []*Variation {
	if m != nil {
		return m.Variations
	}
	return nil
}

type FetchLanguageVariationsRequest struct {
	LanguageIds uint32 `protobuf:"varint,1,opt,name=languageIds" json:"languageIds,omitempty"`
}

func (m *FetchLanguageVariationsRequest) Reset()                    { *m = FetchLanguageVariationsRequest{} }
func (m *FetchLanguageVariationsRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchLanguageVariationsRequest) ProtoMessage()               {}
func (*FetchLanguageVariationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *FetchLanguageVariationsRequest) GetLanguageIds() uint32 {
	if m != nil {
		return m.LanguageIds
	}
	return 0
}

type FetchLanguageVariationsResponse struct {
	LanguageVariations []*LanguageVariations `protobuf:"bytes,1,rep,name=languageVariations" json:"languageVariations,omitempty"`
}

func (m *FetchLanguageVariationsResponse) Reset()                    { *m = FetchLanguageVariationsResponse{} }
func (m *FetchLanguageVariationsResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchLanguageVariationsResponse) ProtoMessage()               {}
func (*FetchLanguageVariationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *FetchLanguageVariationsResponse) GetLanguageVariations() []*LanguageVariations {
	if m != nil {
		return m.LanguageVariations
	}
	return nil
}

func init() {
	proto.RegisterType((*LanguageVariations)(nil), "SeppoService.LanguageVariations")
	proto.RegisterType((*FetchLanguageVariationsRequest)(nil), "SeppoService.FetchLanguageVariationsRequest")
	proto.RegisterType((*FetchLanguageVariationsResponse)(nil), "SeppoService.FetchLanguageVariationsResponse")
}

func init() { proto.RegisterFile("language_variations.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x49, 0xcc, 0x4b,
	0x2f, 0x4d, 0x4c, 0x4f, 0x8d, 0x2f, 0x4b, 0x2c, 0xca, 0x4c, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x09, 0x4e, 0x2d, 0x28, 0xc8, 0x0f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x95, 0xe2, 0x87, 0xcb, 0x43, 0xa4, 0x95, 0x72, 0xb9, 0x84, 0x7c, 0xa0, 0x7a,
	0xc3, 0xe0, 0x5a, 0x85, 0xe4, 0xb8, 0xb8, 0x60, 0x26, 0x7a, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x06, 0x21, 0x89, 0x08, 0x99, 0x73, 0x71, 0x21, 0x2c, 0x92, 0x60, 0x52, 0x60, 0xd6, 0xe0,
	0x36, 0x12, 0xd7, 0x43, 0xb6, 0x49, 0x0f, 0x6e, 0x5a, 0x10, 0x92, 0x52, 0x25, 0x27, 0x2e, 0x39,
	0xb7, 0xd4, 0x92, 0xe4, 0x0c, 0x4c, 0x3b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x14,
	0xb8, 0xb8, 0x11, 0x16, 0x15, 0x43, 0xed, 0x46, 0x16, 0x52, 0x2a, 0xe6, 0x92, 0xc7, 0x69, 0x46,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x50, 0x00, 0x97, 0x50, 0x0e, 0x86, 0xac, 0x04, 0x23, 0xd8,
	0x9d, 0x0a, 0xa8, 0xee, 0xc4, 0x62, 0x0a, 0x16, 0xbd, 0x49, 0x6c, 0xe0, 0xe0, 0x32, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x36, 0x11, 0xde, 0x44, 0x6a, 0x01, 0x00, 0x00,
}
