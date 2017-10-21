// Code generated by protoc-gen-go.
// source: song.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Song struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Song) Reset()                    { *m = Song{} }
func (m *Song) String() string            { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()               {}
func (*Song) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *Song) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Song)(nil), "SeppoService.Song")
}

func init() { proto.RegisterFile("song.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 79 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xce, 0xcf, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x09, 0x4e, 0x2d, 0x28, 0xc8, 0x0f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0x12, 0xe3, 0x62, 0x09, 0xce, 0xcf, 0x4b, 0x17, 0xe2, 0xe3, 0x62,
	0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x62, 0xca, 0x4c, 0x49, 0x62, 0x03, 0x2b,
	0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x89, 0xf8, 0xbb, 0x3a, 0x00, 0x00, 0x00,
}
