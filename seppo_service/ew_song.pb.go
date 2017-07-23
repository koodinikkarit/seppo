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
	Id            uint32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title         string     `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Author        string     `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Copyright     string     `protobuf:"bytes,4,opt,name=copyright" json:"copyright,omitempty"`
	Administrator string     `protobuf:"bytes,5,opt,name=administrator" json:"administrator,omitempty"`
	Description   string     `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Tags          string     `protobuf:"bytes,7,opt,name=tags" json:"tags,omitempty"`
	Verses        []*EwVerse `protobuf:"bytes,8,rep,name=verses" json:"verses,omitempty"`
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

func (m *EwSong) GetVerses() []*EwVerse {
	if m != nil {
		return m.Verses
	}
	return nil
}

func init() {
	proto.RegisterType((*EwSong)(nil), "SeppoService.EwSong")
}

func init() { proto.RegisterFile("ew_song.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xda, 0xae, 0x76, 0x6a, 0x7a, 0x18, 0x54, 0x16, 0xf1, 0x10, 0xc4, 0x43, 0x2e,
	0xe6, 0xa0, 0xcf, 0xd0, 0x17, 0x48, 0xc0, 0x6b, 0x89, 0xc9, 0x90, 0x0e, 0x68, 0x66, 0xd9, 0x1d,
	0xbb, 0xf8, 0xd4, 0xbe, 0x82, 0xb8, 0x2d, 0x18, 0x6f, 0x33, 0xdf, 0xf7, 0x5f, 0x3e, 0x28, 0x28,
	0xee, 0x83, 0x4c, 0x63, 0xed, 0xbc, 0xa8, 0xe0, 0x55, 0x4b, 0xce, 0x49, 0x4b, 0xfe, 0xc8, 0x3d,
	0xdd, 0x6d, 0x29, 0xee, 0x8f, 0xe4, 0x03, 0x9d, 0xec, 0xc3, 0x77, 0x06, 0x66, 0x17, 0x5b, 0x99,
	0x46, 0xdc, 0x42, 0xce, 0x83, 0xcd, 0xca, 0xac, 0x2a, 0x9a, 0x9c, 0x07, 0xbc, 0x86, 0x95, 0xb2,
	0xbe, 0x93, 0xcd, 0xcb, 0xac, 0x5a, 0x37, 0xa7, 0x07, 0x6f, 0xc1, 0x74, 0x9f, 0x7a, 0x10, 0x6f,
	0x17, 0x09, 0x9f, 0x3f, 0xbc, 0x87, 0x75, 0x2f, 0xee, 0xcb, 0xf3, 0x78, 0x50, 0xbb, 0x4c, 0xea,
	0x0f, 0xe0, 0x23, 0x14, 0xdd, 0xf0, 0xc1, 0x13, 0x07, 0xf5, 0x9d, 0x8a, 0xb7, 0xab, 0xb4, 0xf8,
	0x0f, 0xb1, 0x84, 0xcd, 0x40, 0xa1, 0xf7, 0xec, 0x94, 0x65, 0xb2, 0x26, 0x6d, 0xe6, 0x08, 0x11,
	0x96, 0xda, 0x8d, 0xc1, 0x5e, 0x24, 0x95, 0x6e, 0x7c, 0x02, 0x93, 0x8a, 0x82, 0xbd, 0x2c, 0x17,
	0xd5, 0xe6, 0xf9, 0xa6, 0x9e, 0x17, 0xd7, 0xbb, 0xf8, 0xfa, 0x6b, 0x9b, 0xf3, 0xe8, 0xcd, 0xa4,
	0xf0, 0x97, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x3b, 0x94, 0xff, 0x27, 0x01, 0x00, 0x00,
}
