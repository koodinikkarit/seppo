// Code generated by protoc-gen-go.
// source: variation_text.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VariationText struct {
	Id          uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	VariationId uint32 `protobuf:"varint,2,opt,name=variationId" json:"variationId,omitempty"`
	Text        string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
}

func (m *VariationText) Reset()                    { *m = VariationText{} }
func (m *VariationText) String() string            { return proto.CompactTextString(m) }
func (*VariationText) ProtoMessage()               {}
func (*VariationText) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *VariationText) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VariationText) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

func (m *VariationText) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type FetchVariationTextByVariationIdRequest struct {
	VariationIds []uint32 `protobuf:"varint,1,rep,packed,name=variationIds" json:"variationIds,omitempty"`
}

func (m *FetchVariationTextByVariationIdRequest) Reset() {
	*m = FetchVariationTextByVariationIdRequest{}
}
func (m *FetchVariationTextByVariationIdRequest) String() string { return proto.CompactTextString(m) }
func (*FetchVariationTextByVariationIdRequest) ProtoMessage()    {}
func (*FetchVariationTextByVariationIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor16, []int{1}
}

func (m *FetchVariationTextByVariationIdRequest) GetVariationIds() []uint32 {
	if m != nil {
		return m.VariationIds
	}
	return nil
}

type FetchVariationTextByVariationIdResponse struct {
	VariationTexts []*VariationText `protobuf:"bytes,1,rep,name=variationTexts" json:"variationTexts,omitempty"`
}

func (m *FetchVariationTextByVariationIdResponse) Reset() {
	*m = FetchVariationTextByVariationIdResponse{}
}
func (m *FetchVariationTextByVariationIdResponse) String() string { return proto.CompactTextString(m) }
func (*FetchVariationTextByVariationIdResponse) ProtoMessage()    {}
func (*FetchVariationTextByVariationIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor16, []int{2}
}

func (m *FetchVariationTextByVariationIdResponse) GetVariationTexts() []*VariationText {
	if m != nil {
		return m.VariationTexts
	}
	return nil
}

func init() {
	proto.RegisterType((*VariationText)(nil), "SeppoService.VariationText")
	proto.RegisterType((*FetchVariationTextByVariationIdRequest)(nil), "SeppoService.FetchVariationTextByVariationIdRequest")
	proto.RegisterType((*FetchVariationTextByVariationIdResponse)(nil), "SeppoService.FetchVariationTextByVariationIdResponse")
}

func init() { proto.RegisterFile("variation_text.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4b, 0x2c, 0xca,
	0x4c, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0x2f, 0x49, 0xad, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x09, 0x4e, 0x2d, 0x28, 0xc8, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0x0a,
	0xe5, 0xe2, 0x0d, 0x83, 0xa9, 0x0a, 0x49, 0xad, 0x28, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x52, 0xe0, 0xe2, 0x86, 0x1b, 0xe3,
	0x99, 0x22, 0xc1, 0x04, 0x96, 0x40, 0x16, 0x12, 0x12, 0xe2, 0x62, 0x01, 0x19, 0x2f, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x2b, 0xf9, 0x70, 0xa9, 0xb9, 0xa5, 0x96, 0x24, 0x67, 0xa0,
	0x98, 0xed, 0x54, 0x19, 0x86, 0xd0, 0x16, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc4,
	0xc5, 0x83, 0x64, 0x58, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x6f, 0x10, 0x8a, 0x98, 0x52, 0x1e,
	0x97, 0x3a, 0x41, 0xd3, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x9c, 0xb9, 0xf8, 0xca, 0x90,
	0x55, 0x41, 0x0c, 0xe4, 0x36, 0x92, 0xd6, 0x43, 0xf6, 0xb6, 0x1e, 0x8a, 0x49, 0x41, 0x68, 0x5a,
	0x92, 0xd8, 0xc0, 0x21, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xb0, 0x5f, 0x9c, 0x41,
	0x01, 0x00, 0x00,
}
