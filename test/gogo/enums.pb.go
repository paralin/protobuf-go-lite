// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enums.proto

package test

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RegularEnum int32

const (
	REGULAR_UNKNOWN RegularEnum = 0
	REGULAR_A       RegularEnum = 1
	REGULAR_B       RegularEnum = 2
)

var RegularEnum_name = map[int32]string{
	0: "REGULAR_UNKNOWN",
	1: "REGULAR_A",
	2: "REGULAR_B",
}

var RegularEnum_value = map[string]int32{
	"REGULAR_UNKNOWN": 0,
	"REGULAR_A":       1,
	"REGULAR_B":       2,
}

func (RegularEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{0}
}

type CustomEnum int32

const (
	CustomEnum_CUSTOM_UNKNOWN CustomEnum = 0
	CustomEnum_CUSTOM_V1_0    CustomEnum = 1
	CustomEnum_CUSTOM_V1_0_1  CustomEnum = 2
)

var CustomEnum_name = map[int32]string{
	0: "CUSTOM_UNKNOWN",
	1: "CUSTOM_V1_0",
	2: "CUSTOM_V1_0_1",
}

var CustomEnum_value = map[string]int32{
	"CUSTOM_UNKNOWN": 0,
	"CUSTOM_V1_0":    1,
	"CUSTOM_V1_0_1":  2,
}

func (CustomEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{1}
}

type CustomEnumValue struct {
	Value                CustomEnum `protobuf:"varint,1,opt,name=value,proto3,enum=thethings.json.test.CustomEnum" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CustomEnumValue) Reset()         { *m = CustomEnumValue{} }
func (m *CustomEnumValue) String() string { return proto.CompactTextString(m) }
func (*CustomEnumValue) ProtoMessage()    {}
func (*CustomEnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{0}
}
func (m *CustomEnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomEnumValue.Unmarshal(m, b)
}
func (m *CustomEnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomEnumValue.Marshal(b, m, deterministic)
}
func (m *CustomEnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomEnumValue.Merge(m, src)
}
func (m *CustomEnumValue) XXX_Size() int {
	return xxx_messageInfo_CustomEnumValue.Size(m)
}
func (m *CustomEnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomEnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_CustomEnumValue proto.InternalMessageInfo

func (m *CustomEnumValue) GetValue() CustomEnum {
	if m != nil {
		return m.Value
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

type MessageWithEnums struct {
	Regular              RegularEnum        `protobuf:"varint,1,opt,name=regular,proto3,enum=thethings.json.test.RegularEnum" json:"regular,omitempty"`
	Regulars             []RegularEnum      `protobuf:"varint,2,rep,packed,name=regulars,proto3,enum=thethings.json.test.RegularEnum" json:"regulars,omitempty"`
	Custom               CustomEnum         `protobuf:"varint,3,opt,name=custom,proto3,enum=thethings.json.test.CustomEnum" json:"custom,omitempty"`
	Customs              []CustomEnum       `protobuf:"varint,4,rep,packed,name=customs,proto3,enum=thethings.json.test.CustomEnum" json:"customs,omitempty"`
	WrappedCustom        *CustomEnumValue   `protobuf:"bytes,5,opt,name=wrapped_custom,json=wrappedCustom,proto3" json:"wrapped_custom,omitempty"`
	WrappedCustoms       []*CustomEnumValue `protobuf:"bytes,6,rep,name=wrapped_customs,json=wrappedCustoms,proto3" json:"wrapped_customs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MessageWithEnums) Reset()         { *m = MessageWithEnums{} }
func (m *MessageWithEnums) String() string { return proto.CompactTextString(m) }
func (*MessageWithEnums) ProtoMessage()    {}
func (*MessageWithEnums) Descriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{1}
}
func (m *MessageWithEnums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithEnums.Unmarshal(m, b)
}
func (m *MessageWithEnums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithEnums.Marshal(b, m, deterministic)
}
func (m *MessageWithEnums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithEnums.Merge(m, src)
}
func (m *MessageWithEnums) XXX_Size() int {
	return xxx_messageInfo_MessageWithEnums.Size(m)
}
func (m *MessageWithEnums) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithEnums.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithEnums proto.InternalMessageInfo

func (m *MessageWithEnums) GetRegular() RegularEnum {
	if m != nil {
		return m.Regular
	}
	return REGULAR_UNKNOWN
}

func (m *MessageWithEnums) GetRegulars() []RegularEnum {
	if m != nil {
		return m.Regulars
	}
	return nil
}

func (m *MessageWithEnums) GetCustom() CustomEnum {
	if m != nil {
		return m.Custom
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

func (m *MessageWithEnums) GetCustoms() []CustomEnum {
	if m != nil {
		return m.Customs
	}
	return nil
}

func (m *MessageWithEnums) GetWrappedCustom() *CustomEnumValue {
	if m != nil {
		return m.WrappedCustom
	}
	return nil
}

func (m *MessageWithEnums) GetWrappedCustoms() []*CustomEnumValue {
	if m != nil {
		return m.WrappedCustoms
	}
	return nil
}

type MessageWithOneofEnums struct {
	// Types that are valid to be assigned to Value:
	//	*MessageWithOneofEnums_Regular
	//	*MessageWithOneofEnums_Custom
	//	*MessageWithOneofEnums_WrappedCustom
	Value                isMessageWithOneofEnums_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MessageWithOneofEnums) Reset()         { *m = MessageWithOneofEnums{} }
func (m *MessageWithOneofEnums) String() string { return proto.CompactTextString(m) }
func (*MessageWithOneofEnums) ProtoMessage()    {}
func (*MessageWithOneofEnums) Descriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{2}
}
func (m *MessageWithOneofEnums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithOneofEnums.Unmarshal(m, b)
}
func (m *MessageWithOneofEnums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithOneofEnums.Marshal(b, m, deterministic)
}
func (m *MessageWithOneofEnums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithOneofEnums.Merge(m, src)
}
func (m *MessageWithOneofEnums) XXX_Size() int {
	return xxx_messageInfo_MessageWithOneofEnums.Size(m)
}
func (m *MessageWithOneofEnums) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithOneofEnums.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithOneofEnums proto.InternalMessageInfo

type isMessageWithOneofEnums_Value interface {
	isMessageWithOneofEnums_Value()
}

type MessageWithOneofEnums_Regular struct {
	Regular RegularEnum `protobuf:"varint,1,opt,name=regular,proto3,enum=thethings.json.test.RegularEnum,oneof" json:"regular,omitempty"`
}
type MessageWithOneofEnums_Custom struct {
	Custom CustomEnum `protobuf:"varint,2,opt,name=custom,proto3,enum=thethings.json.test.CustomEnum,oneof" json:"custom,omitempty"`
}
type MessageWithOneofEnums_WrappedCustom struct {
	WrappedCustom *CustomEnumValue `protobuf:"bytes,3,opt,name=wrapped_custom,json=wrappedCustom,proto3,oneof" json:"wrapped_custom,omitempty"`
}

func (*MessageWithOneofEnums_Regular) isMessageWithOneofEnums_Value()       {}
func (*MessageWithOneofEnums_Custom) isMessageWithOneofEnums_Value()        {}
func (*MessageWithOneofEnums_WrappedCustom) isMessageWithOneofEnums_Value() {}

func (m *MessageWithOneofEnums) GetValue() isMessageWithOneofEnums_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MessageWithOneofEnums) GetRegular() RegularEnum {
	if x, ok := m.GetValue().(*MessageWithOneofEnums_Regular); ok {
		return x.Regular
	}
	return REGULAR_UNKNOWN
}

func (m *MessageWithOneofEnums) GetCustom() CustomEnum {
	if x, ok := m.GetValue().(*MessageWithOneofEnums_Custom); ok {
		return x.Custom
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

func (m *MessageWithOneofEnums) GetWrappedCustom() *CustomEnumValue {
	if x, ok := m.GetValue().(*MessageWithOneofEnums_WrappedCustom); ok {
		return x.WrappedCustom
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MessageWithOneofEnums) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MessageWithOneofEnums_Regular)(nil),
		(*MessageWithOneofEnums_Custom)(nil),
		(*MessageWithOneofEnums_WrappedCustom)(nil),
	}
}

func init() {
	proto.RegisterEnum("thethings.json.test.RegularEnum", RegularEnum_name, RegularEnum_value)
	proto.RegisterEnum("thethings.json.test.CustomEnum", CustomEnum_name, CustomEnum_value)
	proto.RegisterType((*CustomEnumValue)(nil), "thethings.json.test.CustomEnumValue")
	proto.RegisterType((*MessageWithEnums)(nil), "thethings.json.test.MessageWithEnums")
	proto.RegisterType((*MessageWithOneofEnums)(nil), "thethings.json.test.MessageWithOneofEnums")
}

func init() { proto.RegisterFile("enums.proto", fileDescriptor_888b6bd9597961ff) }

var fileDescriptor_888b6bd9597961ff = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x51, 0x6f, 0xd2, 0x50,
	0x14, 0xc7, 0x7b, 0xe9, 0x06, 0xf3, 0x20, 0x50, 0xef, 0x62, 0x52, 0x48, 0xc4, 0x4a, 0x7c, 0x58,
	0x48, 0x28, 0x74, 0x46, 0xcd, 0x96, 0xbd, 0x8c, 0x65, 0x11, 0x33, 0x01, 0x53, 0x61, 0x4b, 0x7c,
	0x21, 0x1d, 0xbb, 0xb6, 0x35, 0x70, 0x2f, 0xe9, 0xbd, 0xd5, 0xaf, 0xe0, 0x87, 0xf0, 0x4d, 0x9f,
	0x78, 0xf2, 0x33, 0xed, 0x69, 0xe1, 0xc9, 0x8f, 0x60, 0xb8, 0x2d, 0xd0, 0xe9, 0x92, 0xd5, 0xb7,
	0xd3, 0x73, 0xff, 0xff, 0xff, 0x3d, 0xfd, 0x9d, 0x16, 0xf2, 0x84, 0x86, 0x53, 0x6e, 0xce, 0x02,
	0x26, 0x18, 0xde, 0x15, 0x1e, 0x11, 0x9e, 0x4f, 0x5d, 0x6e, 0x7e, 0xe6, 0x8c, 0x9a, 0x82, 0x70,
	0x51, 0x79, 0xe4, 0x50, 0xca, 0x84, 0x23, 0x7c, 0x46, 0x63, 0x5d, 0xe5, 0x89, 0x4f, 0x05, 0x09,
	0xa8, 0x33, 0x69, 0xba, 0xcc, 0x65, 0xb2, 0x27, 0xab, 0xe8, 0xb8, 0xf6, 0x1e, 0x4a, 0x27, 0x21,
	0x17, 0x6c, 0x7a, 0x4a, 0xc3, 0xe9, 0xb9, 0x33, 0x09, 0x09, 0x7e, 0x09, 0xdb, 0x5f, 0x96, 0x85,
	0x8e, 0x0c, 0xb4, 0x57, 0xdc, 0x7f, 0x6a, 0xde, 0x71, 0x93, 0xb9, 0x31, 0xd9, 0x91, 0xfa, 0x30,
	0xbb, 0x98, 0x97, 0x33, 0x3a, 0xaa, 0x7d, 0x57, 0x41, 0xeb, 0x12, 0xce, 0x1d, 0x97, 0x5c, 0xf8,
	0xc2, 0x5b, 0x4a, 0x38, 0x3e, 0x84, 0x5c, 0x40, 0xdc, 0x70, 0xe2, 0x04, 0x71, 0xaa, 0x71, 0x67,
	0xaa, 0x1d, 0x69, 0x64, 0xec, 0xca, 0x80, 0x8f, 0x60, 0x27, 0x2e, 0xb9, 0x9e, 0x31, 0xd4, 0x54,
	0xe6, 0xb5, 0x03, 0xbf, 0x86, 0xec, 0x58, 0xce, 0xaa, 0xab, 0xe9, 0x5e, 0x27, 0x96, 0xe3, 0x03,
	0xc8, 0x45, 0x15, 0xd7, 0xb7, 0xe4, 0xad, 0xf7, 0x3a, 0x57, 0x7a, 0x7c, 0x06, 0xc5, 0xaf, 0x81,
	0x33, 0x9b, 0x91, 0xab, 0x51, 0x7c, 0xf7, 0xb6, 0x81, 0xf6, 0xf2, 0xfb, 0xcf, 0xef, 0x49, 0x90,
	0xfc, 0xed, 0x42, 0xec, 0x8d, 0xfa, 0xb8, 0x0b, 0xa5, 0xdb, 0x61, 0x5c, 0xcf, 0x1a, 0x6a, 0xea,
	0xb4, 0xe2, 0xad, 0x34, 0x5e, 0xbb, 0x41, 0xf0, 0x38, 0xb1, 0x9e, 0x3e, 0x25, 0xec, 0x53, 0xb4,
	0xa3, 0xa3, 0xff, 0xde, 0x51, 0x47, 0xd9, 0x6c, 0xe9, 0x60, 0xcd, 0x39, 0x93, 0x8a, 0x73, 0x47,
	0x59, 0x93, 0xee, 0xfe, 0x83, 0x4b, 0x4d, 0x8f, 0xab, 0xa3, 0xfc, 0x05, 0xac, 0x9d, 0x8b, 0xbf,
	0xdf, 0xba, 0x0d, 0xf9, 0xc4, 0xb0, 0x78, 0x17, 0x4a, 0xf6, 0xe9, 0x9b, 0xe1, 0xbb, 0x63, 0x7b,
	0x34, 0xec, 0x9d, 0xf5, 0xfa, 0x17, 0x3d, 0x4d, 0xc1, 0x05, 0x78, 0xb0, 0x6a, 0x1e, 0x6b, 0x28,
	0xf9, 0xd8, 0xd6, 0x32, 0x15, 0x6d, 0x31, 0x2f, 0x6f, 0xed, 0x28, 0x86, 0xf2, 0xed, 0x47, 0x55,
	0xf9, 0xf5, 0xb3, 0xaa, 0xd4, 0x39, 0xc0, 0x66, 0x00, 0x8c, 0xa1, 0x78, 0x32, 0xfc, 0x30, 0xe8,
	0x77, 0x13, 0x89, 0xcf, 0x20, 0x1f, 0xf7, 0xce, 0xad, 0x51, 0x4b, 0x43, 0x32, 0xe4, 0x21, 0xa8,
	0x96, 0xd9, 0xc2, 0xdb, 0x96, 0xd9, 0x32, 0x5b, 0xb8, 0x0a, 0x85, 0x84, 0x64, 0x64, 0x69, 0x99,
	0x4a, 0x7e, 0x31, 0x2f, 0xe7, 0x40, 0x9e, 0x5b, 0x95, 0xe2, 0x62, 0x5e, 0x06, 0x1d, 0xd5, 0xb3,
	0x91, 0xaa, 0xdd, 0x91, 0x63, 0x20, 0x0d, 0xdd, 0x5c, 0x57, 0x95, 0xdf, 0xd7, 0x55, 0xf4, 0xf1,
	0x95, 0xeb, 0x0b, 0x2f, 0xbc, 0x34, 0xc7, 0x6c, 0xda, 0x1c, 0x78, 0x64, 0x20, 0x21, 0xbd, 0xa5,
	0x57, 0x21, 0x17, 0x81, 0x4f, 0x78, 0x53, 0xfe, 0xdc, 0xe3, 0x86, 0x4b, 0x68, 0xc3, 0x65, 0x8d,
	0x25, 0xbc, 0xe6, 0x12, 0xde, 0x65, 0x56, 0x1e, 0xbc, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x19,
	0x9f, 0x31, 0xfc, 0x4b, 0x04, 0x00, 0x00,
}

func (x CustomEnum) String() string {
	s, ok := CustomEnum_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
