// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/proto/season/season.proto

package season

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Season struct {
	Id                   int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsCurrent            *wrappers.BoolValue `protobuf:"bytes,3,opt,name=is_current,json=isCurrent,proto3" json:"is_current,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Season) Reset()         { *m = Season{} }
func (m *Season) String() string { return proto.CompactTextString(m) }
func (*Season) ProtoMessage()    {}
func (*Season) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dc5e7fcc83b3508, []int{0}
}

func (m *Season) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Season.Unmarshal(m, b)
}
func (m *Season) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Season.Marshal(b, m, deterministic)
}
func (m *Season) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Season.Merge(m, src)
}
func (m *Season) XXX_Size() int {
	return xxx_messageInfo_Season.Size(m)
}
func (m *Season) XXX_DiscardUnknown() {
	xxx_messageInfo_Season.DiscardUnknown(m)
}

var xxx_messageInfo_Season proto.InternalMessageInfo

func (m *Season) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Season) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Season) GetIsCurrent() *wrappers.BoolValue {
	if m != nil {
		return m.IsCurrent
	}
	return nil
}

func init() {
	proto.RegisterType((*Season)(nil), "season.Season")
}

func init() { proto.RegisterFile("internal/proto/season/season.proto", fileDescriptor_0dc5e7fcc83b3508) }

var fileDescriptor_0dc5e7fcc83b3508 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0x4b, 0xc5, 0x30,
	0x14, 0x85, 0x49, 0x9f, 0x14, 0x5e, 0x04, 0x87, 0x4c, 0xe5, 0x0d, 0x52, 0xde, 0xd4, 0xc5, 0x04,
	0x74, 0x12, 0x71, 0xa9, 0xff, 0xa0, 0x82, 0x83, 0x8b, 0xdc, 0xb6, 0x31, 0x5e, 0x48, 0x73, 0x4b,
	0x72, 0x83, 0x7f, 0x5f, 0x48, 0x14, 0x1c, 0xde, 0x74, 0xcf, 0xb9, 0x7c, 0xf0, 0x71, 0xe4, 0x19,
	0x03, 0xdb, 0x18, 0xc0, 0x9b, 0x3d, 0x12, 0x93, 0x49, 0x16, 0x12, 0x85, 0xdf, 0xa3, 0xcb, 0x4f,
	0xb5, 0xb5, 0x9d, 0x6e, 0x1d, 0x91, 0xf3, 0xb6, 0x92, 0x73, 0xfe, 0x34, 0xdf, 0x11, 0xf6, 0xdd,
	0xc6, 0x54, 0xb9, 0xb3, 0x93, 0xed, 0x6b, 0x21, 0xd5, 0x8d, 0x6c, 0x70, 0xed, 0x44, 0x2f, 0x86,
	0xc3, 0xd4, 0xe0, 0xaa, 0x94, 0xbc, 0x0a, 0xb0, 0xd9, 0xae, 0xe9, 0xc5, 0x70, 0x9c, 0x4a, 0x56,
	0x8f, 0x52, 0x62, 0xfa, 0x58, 0x72, 0x8c, 0x36, 0x70, 0x77, 0xe8, 0xc5, 0x70, 0x7d, 0x7f, 0xd2,
	0x55, 0xa1, 0xff, 0x14, 0x7a, 0x24, 0xf2, 0x6f, 0xe0, 0xb3, 0x9d, 0x8e, 0x98, 0x5e, 0x2a, 0x3c,
	0x3e, 0xbf, 0x3f, 0x39, 0xe4, 0xaf, 0x3c, 0xeb, 0x85, 0x36, 0x93, 0x18, 0x18, 0x13, 0xe3, 0x42,
	0xff, 0xe2, 0xdd, 0x0a, 0x0c, 0xe6, 0xe2, 0xb8, 0xb9, 0x2d, 0xed, 0xe1, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x7f, 0x7b, 0xdf, 0x02, 0xfc, 0x00, 0x00, 0x00,
}
