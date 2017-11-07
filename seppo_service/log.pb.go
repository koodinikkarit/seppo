// Code generated by protoc-gen-go.
// source: log.proto
// DO NOT EDIT!

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Log struct {
	Id          uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	LogType     uint32 `protobuf:"varint,2,opt,name=logType" json:"logType,omitempty"`
	Message     string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	MessageDate int64  `protobuf:"varint,4,opt,name=messageDate" json:"messageDate,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Log) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Log) GetLogType() uint32 {
	if m != nil {
		return m.LogType
	}
	return 0
}

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Log) GetMessageDate() int64 {
	if m != nil {
		return m.MessageDate
	}
	return 0
}

type SearchLogsRequest struct {
	Offset      uint32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit       uint32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	MessageType uint32 `protobuf:"varint,3,opt,name=messageType" json:"messageType,omitempty"`
	StartDate   int64  `protobuf:"varint,4,opt,name=startDate" json:"startDate,omitempty"`
	EndDate     int64  `protobuf:"varint,5,opt,name=endDate" json:"endDate,omitempty"`
	SearchWord  string `protobuf:"bytes,6,opt,name=searchWord" json:"searchWord,omitempty"`
}

func (m *SearchLogsRequest) Reset()                    { *m = SearchLogsRequest{} }
func (m *SearchLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchLogsRequest) ProtoMessage()               {}
func (*SearchLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SearchLogsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchLogsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchLogsRequest) GetMessageType() uint32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *SearchLogsRequest) GetStartDate() int64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *SearchLogsRequest) GetEndDate() int64 {
	if m != nil {
		return m.EndDate
	}
	return 0
}

func (m *SearchLogsRequest) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

type SearchLogsResponse struct {
	Logs    []*Log `protobuf:"bytes,1,rep,name=logs" json:"logs,omitempty"`
	MaxLogs uint32 `protobuf:"varint,2,opt,name=maxLogs" json:"maxLogs,omitempty"`
}

func (m *SearchLogsResponse) Reset()                    { *m = SearchLogsResponse{} }
func (m *SearchLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchLogsResponse) ProtoMessage()               {}
func (*SearchLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *SearchLogsResponse) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *SearchLogsResponse) GetMaxLogs() uint32 {
	if m != nil {
		return m.MaxLogs
	}
	return 0
}

func init() {
	proto.RegisterType((*Log)(nil), "SeppoService.Log")
	proto.RegisterType((*SearchLogsRequest)(nil), "SeppoService.SearchLogsRequest")
	proto.RegisterType((*SearchLogsResponse)(nil), "SeppoService.SearchLogsResponse")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0xbb, 0x5b, 0xe9, 0xac, 0x0a, 0x1b, 0x44, 0x72, 0x10, 0x29, 0x05, 0xa1, 0xa7,
	0x1e, 0xf4, 0x15, 0x3c, 0xf6, 0xd4, 0x2a, 0x9e, 0xeb, 0x76, 0x36, 0x16, 0xbb, 0x3b, 0xb5, 0x13,
	0x45, 0x9f, 0xcd, 0x97, 0x93, 0x8e, 0x29, 0x9b, 0x5b, 0xbe, 0xff, 0x27, 0xcc, 0x97, 0x09, 0xa4,
	0x03, 0xd9, 0x72, 0x9c, 0xc8, 0x91, 0x3e, 0x6f, 0x70, 0x1c, 0xa9, 0xc1, 0xe9, 0xab, 0xdf, 0x61,
	0xfe, 0x0e, 0x71, 0x45, 0x56, 0x5f, 0x42, 0xd4, 0x77, 0x46, 0x65, 0xaa, 0xb8, 0xa8, 0xa3, 0xbe,
	0xd3, 0x06, 0xce, 0x06, 0xb2, 0x4f, 0x3f, 0x23, 0x9a, 0x48, 0xc2, 0x05, 0xe7, 0xe6, 0x80, 0xcc,
	0xad, 0x45, 0x13, 0x67, 0xaa, 0x48, 0xeb, 0x05, 0x75, 0x06, 0x1b, 0x7f, 0x7c, 0x6c, 0x1d, 0x9a,
	0x55, 0xa6, 0x8a, 0xb8, 0x0e, 0xa3, 0xfc, 0x57, 0xc1, 0xb6, 0xc1, 0x76, 0xda, 0xbd, 0x55, 0x64,
	0xb9, 0xc6, 0x8f, 0x4f, 0x64, 0xa7, 0xaf, 0x21, 0xa1, 0xfd, 0x9e, 0xd1, 0xf9, 0xf9, 0x9e, 0xf4,
	0x15, 0xac, 0x87, 0xfe, 0xd0, 0x3b, 0x6f, 0xf0, 0x0f, 0xc1, 0x14, 0xb1, 0x8b, 0xa5, 0x0b, 0x23,
	0x7d, 0x03, 0x29, 0xbb, 0x76, 0x72, 0x81, 0xc5, 0x29, 0x98, 0xfd, 0xf1, 0xd8, 0x49, 0xb7, 0x96,
	0x6e, 0x41, 0x7d, 0x0b, 0xc0, 0x22, 0xf7, 0x42, 0x53, 0x67, 0x12, 0x79, 0x5c, 0x90, 0xe4, 0xcf,
	0xa0, 0x43, 0x79, 0x1e, 0xe9, 0xc8, 0xa8, 0xef, 0x60, 0x35, 0x90, 0x65, 0xa3, 0xb2, 0xb8, 0xd8,
	0xdc, 0x6f, 0xcb, 0x70, 0xbb, 0x65, 0x45, 0xb6, 0x96, 0x5a, 0xd6, 0xd6, 0x7e, 0xcf, 0x37, 0x97,
	0x85, 0x7a, 0x7c, 0x4d, 0xe4, 0x5b, 0x1e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x88, 0xdd,
	0x2f, 0xa3, 0x01, 0x00, 0x00,
}