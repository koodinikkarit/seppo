// Code generated by protoc-gen-go.
// source: ew_song.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EwSong struct {
	Id            uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Author        string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Copyright     string `protobuf:"bytes,4,opt,name=copyright" json:"copyright,omitempty"`
	Administrator string `protobuf:"bytes,5,opt,name=administrator" json:"administrator,omitempty"`
	Description   string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Tags          string `protobuf:"bytes,7,opt,name=tags" json:"tags,omitempty"`
	Text          string `protobuf:"bytes,8,opt,name=text" json:"text,omitempty"`
}

func (m *EwSong) Reset()                    { *m = EwSong{} }
func (m *EwSong) String() string            { return proto.CompactTextString(m) }
func (*EwSong) ProtoMessage()               {}
func (*EwSong) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EwSong) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EwSong) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EwSong) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *EwSong) GetCopyright() string {
	if m != nil {
		return m.Copyright
	}
	return ""
}

func (m *EwSong) GetAdministrator() string {
	if m != nil {
		return m.Administrator
	}
	return ""
}

func (m *EwSong) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EwSong) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *EwSong) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*EwSong)(nil), "SeppoService.EwSong")
}

func init() { proto.RegisterFile("ew_song.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcf, 0x31, 0x4e, 0xc3, 0x60,
	0x0c, 0x05, 0x60, 0x25, 0xb4, 0x81, 0x1a, 0xd2, 0xc1, 0x42, 0xc8, 0x42, 0x0c, 0x11, 0x62, 0xe8,
	0xc4, 0xc2, 0x19, 0xb8, 0x40, 0x73, 0x80, 0x2a, 0x24, 0x56, 0x6a, 0x09, 0xe2, 0x5f, 0xfe, 0x4d,
	0x53, 0xee, 0xc9, 0x81, 0x90, 0xdc, 0x4a, 0xc0, 0xe6, 0xf7, 0xbd, 0xb7, 0x18, 0x6a, 0x9e, 0x77,
	0x59, 0xa7, 0xf1, 0x39, 0x99, 0xba, 0xe2, 0x4d, 0xcb, 0x29, 0x69, 0xcb, 0x76, 0x90, 0x9e, 0xef,
	0xd7, 0x3c, 0xef, 0x0e, 0x6c, 0x99, 0x4f, 0xed, 0xe3, 0x77, 0x01, 0xd5, 0xeb, 0xdc, 0xea, 0x34,
	0xe2, 0x1a, 0x4a, 0x19, 0xa8, 0x68, 0x8a, 0x4d, 0xbd, 0x2d, 0x65, 0xc0, 0x5b, 0x58, 0xba, 0xf8,
	0x3b, 0x53, 0xd9, 0x14, 0x9b, 0xd5, 0xf6, 0x14, 0xf0, 0x0e, 0xaa, 0xee, 0xd3, 0xf7, 0x6a, 0x74,
	0x11, 0x7c, 0x4e, 0xf8, 0x00, 0xab, 0x5e, 0xd3, 0x97, 0xc9, 0xb8, 0x77, 0x5a, 0x44, 0xf5, 0x0b,
	0xf8, 0x04, 0x75, 0x37, 0x7c, 0xc8, 0x24, 0xd9, 0xad, 0x73, 0x35, 0x5a, 0xc6, 0xe2, 0x3f, 0x62,
	0x03, 0xd7, 0x03, 0xe7, 0xde, 0x24, 0xb9, 0xe8, 0x44, 0x55, 0x6c, 0xfe, 0x12, 0x22, 0x2c, 0xbc,
	0x1b, 0x33, 0x5d, 0x46, 0x15, 0x77, 0x18, 0x1f, 0x9d, 0xae, 0xce, 0xc6, 0x47, 0x7f, 0xab, 0xe2,
	0xbb, 0x97, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x80, 0xb2, 0x0a, 0x0c, 0x01, 0x00, 0x00,
}
