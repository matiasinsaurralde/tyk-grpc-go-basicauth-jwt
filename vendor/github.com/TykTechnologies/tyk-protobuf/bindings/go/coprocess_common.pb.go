// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coprocess_common.proto

package coprocess

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HookType int32

const (
	HookType_Unknown        HookType = 0
	HookType_Pre            HookType = 1
	HookType_Post           HookType = 2
	HookType_PostKeyAuth    HookType = 3
	HookType_CustomKeyCheck HookType = 4
)

var HookType_name = map[int32]string{
	0: "Unknown",
	1: "Pre",
	2: "Post",
	3: "PostKeyAuth",
	4: "CustomKeyCheck",
}

var HookType_value = map[string]int32{
	"Unknown":        0,
	"Pre":            1,
	"Post":           2,
	"PostKeyAuth":    3,
	"CustomKeyCheck": 4,
}

func (x HookType) String() string {
	return proto.EnumName(HookType_name, int32(x))
}

func (HookType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad9b17a8ddc1be7d, []int{0}
}

type StringSlice struct {
	Items                []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringSlice) Reset()         { *m = StringSlice{} }
func (m *StringSlice) String() string { return proto.CompactTextString(m) }
func (*StringSlice) ProtoMessage()    {}
func (*StringSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad9b17a8ddc1be7d, []int{0}
}

func (m *StringSlice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringSlice.Unmarshal(m, b)
}
func (m *StringSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringSlice.Marshal(b, m, deterministic)
}
func (m *StringSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringSlice.Merge(m, src)
}
func (m *StringSlice) XXX_Size() int {
	return xxx_messageInfo_StringSlice.Size(m)
}
func (m *StringSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_StringSlice.DiscardUnknown(m)
}

var xxx_messageInfo_StringSlice proto.InternalMessageInfo

func (m *StringSlice) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterEnum("coprocess.HookType", HookType_name, HookType_value)
	proto.RegisterType((*StringSlice)(nil), "coprocess.StringSlice")
}

func init() { proto.RegisterFile("coprocess_common.proto", fileDescriptor_ad9b17a8ddc1be7d) }

var fileDescriptor_ad9b17a8ddc1be7d = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x40, 0xad, 0xad, 0xb6, 0xbd, 0x82, 0x86, 0x43, 0xc4, 0x51, 0x74, 0x11, 0x07, 0x17, 0xbf,
	0x40, 0xba, 0x08, 0x1d, 0x2c, 0x56, 0x67, 0xc1, 0x70, 0xd8, 0x52, 0x93, 0x2b, 0x49, 0x8a, 0xe4,
	0xef, 0x45, 0x05, 0xb7, 0xf7, 0xde, 0xf4, 0x60, 0x2e, 0xb9, 0x33, 0x2c, 0xc9, 0xda, 0x9b, 0x64,
	0xa5, 0x58, 0xef, 0x3a, 0xc3, 0x8e, 0x31, 0xfd, 0xf7, 0xd5, 0x1a, 0xb2, 0xca, 0x99, 0x46, 0x3f,
	0xaa, 0x67, 0x23, 0x09, 0x67, 0x30, 0x6a, 0x1c, 0x29, 0xbb, 0x08, 0x96, 0xe1, 0x26, 0x3d, 0xff,
	0x64, 0x7b, 0x82, 0xe4, 0xc8, 0xdc, 0x5e, 0x7c, 0x47, 0x98, 0x41, 0x7c, 0xd5, 0xad, 0xe6, 0x97,
	0x16, 0x03, 0x8c, 0x21, 0x2c, 0x0d, 0x89, 0x00, 0x13, 0x88, 0x4a, 0xb6, 0x4e, 0x0c, 0x71, 0x0a,
	0xd9, 0x87, 0x0a, 0xf2, 0x87, 0xde, 0xd5, 0x22, 0x44, 0x84, 0x49, 0xde, 0x5b, 0xc7, 0xaa, 0x20,
	0x9f, 0xd7, 0x24, 0x5b, 0x11, 0xdd, 0xc7, 0xdf, 0x8f, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0x7b, 0x4f, 0xde, 0xa1, 0x00, 0x00, 0x00,
}
