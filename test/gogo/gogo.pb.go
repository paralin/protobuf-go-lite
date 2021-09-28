// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gogo.proto

package test

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	github_com_TheThingsIndustries_protoc_gen_go_json_test_types "github.com/TheThingsIndustries/protoc-gen-go-json/test/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MessageWithGoGoOptions struct {
	EUIWithCustomName                   []byte                                                               `protobuf:"bytes,1,opt,name=eui_with_custom_name,json=euiWithCustomName,proto3" json:"eui_with_custom_name,omitempty"`
	EUIWithCustomNameAndType            *github_com_TheThingsIndustries_protoc_gen_go_json_test_types.EUI64  `protobuf:"bytes,2,opt,name=eui_with_custom_name_and_type,json=euiWithCustomNameAndType,proto3,customtype=github.com/TheThingsIndustries/protoc-gen-go-json/test/types.EUI64" json:"eui_with_custom_name_and_type,omitempty"`
	NonNullableEUIWithCustomNameAndType github_com_TheThingsIndustries_protoc_gen_go_json_test_types.EUI64   `protobuf:"bytes,3,opt,name=non_nullable_eui_with_custom_name_and_type,json=nonNullableEuiWithCustomNameAndType,proto3,customtype=github.com/TheThingsIndustries/protoc-gen-go-json/test/types.EUI64" json:"non_nullable_eui_with_custom_name_and_type"`
	EUIsWithCustomNameAndType           []github_com_TheThingsIndustries_protoc_gen_go_json_test_types.EUI64 `protobuf:"bytes,4,rep,name=euis_with_custom_name_and_type,json=euisWithCustomNameAndType,proto3,customtype=github.com/TheThingsIndustries/protoc-gen-go-json/test/types.EUI64" json:"euis_with_custom_name_and_type,omitempty"`
	Duration                            *time.Duration                                                       `protobuf:"bytes,5,opt,name=duration,proto3,stdduration" json:"duration,omitempty"`
	NonNullableDuration                 time.Duration                                                        `protobuf:"bytes,6,opt,name=non_nullable_duration,json=nonNullableDuration,proto3,stdduration" json:"non_nullable_duration"`
	Timestamp                           *time.Time                                                           `protobuf:"bytes,7,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
	NonNullableTimestamp                time.Time                                                            `protobuf:"bytes,8,opt,name=non_nullable_timestamp,json=nonNullableTimestamp,proto3,stdtime" json:"non_nullable_timestamp"`
	XXX_NoUnkeyedLiteral                struct{}                                                             `json:"-"`
	XXX_unrecognized                    []byte                                                               `json:"-"`
	XXX_sizecache                       int32                                                                `json:"-"`
}

func (m *MessageWithGoGoOptions) Reset()         { *m = MessageWithGoGoOptions{} }
func (m *MessageWithGoGoOptions) String() string { return proto.CompactTextString(m) }
func (*MessageWithGoGoOptions) ProtoMessage()    {}
func (*MessageWithGoGoOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_592445b5231bc2b9, []int{0}
}
func (m *MessageWithGoGoOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithGoGoOptions.Unmarshal(m, b)
}
func (m *MessageWithGoGoOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithGoGoOptions.Marshal(b, m, deterministic)
}
func (m *MessageWithGoGoOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithGoGoOptions.Merge(m, src)
}
func (m *MessageWithGoGoOptions) XXX_Size() int {
	return xxx_messageInfo_MessageWithGoGoOptions.Size(m)
}
func (m *MessageWithGoGoOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithGoGoOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithGoGoOptions proto.InternalMessageInfo

func (m *MessageWithGoGoOptions) GetEUIWithCustomName() []byte {
	if m != nil {
		return m.EUIWithCustomName
	}
	return nil
}

func (m *MessageWithGoGoOptions) GetDuration() *time.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *MessageWithGoGoOptions) GetNonNullableDuration() time.Duration {
	if m != nil {
		return m.NonNullableDuration
	}
	return 0
}

func (m *MessageWithGoGoOptions) GetTimestamp() *time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *MessageWithGoGoOptions) GetNonNullableTimestamp() time.Time {
	if m != nil {
		return m.NonNullableTimestamp
	}
	return time.Time{}
}

type SubMessage struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubMessage) Reset()         { *m = SubMessage{} }
func (m *SubMessage) String() string { return proto.CompactTextString(m) }
func (*SubMessage) ProtoMessage()    {}
func (*SubMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_592445b5231bc2b9, []int{1}
}
func (m *SubMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubMessage.Unmarshal(m, b)
}
func (m *SubMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubMessage.Marshal(b, m, deterministic)
}
func (m *SubMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubMessage.Merge(m, src)
}
func (m *SubMessage) XXX_Size() int {
	return xxx_messageInfo_SubMessage.Size(m)
}
func (m *SubMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubMessage proto.InternalMessageInfo

func (m *SubMessage) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type SubMessageWithoutMarshalers struct {
	OtherField           string   `protobuf:"bytes,1,opt,name=other_field,json=otherField,proto3" json:"other_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubMessageWithoutMarshalers) Reset()         { *m = SubMessageWithoutMarshalers{} }
func (m *SubMessageWithoutMarshalers) String() string { return proto.CompactTextString(m) }
func (*SubMessageWithoutMarshalers) ProtoMessage()    {}
func (*SubMessageWithoutMarshalers) Descriptor() ([]byte, []int) {
	return fileDescriptor_592445b5231bc2b9, []int{2}
}
func (m *SubMessageWithoutMarshalers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubMessageWithoutMarshalers.Unmarshal(m, b)
}
func (m *SubMessageWithoutMarshalers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubMessageWithoutMarshalers.Marshal(b, m, deterministic)
}
func (m *SubMessageWithoutMarshalers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubMessageWithoutMarshalers.Merge(m, src)
}
func (m *SubMessageWithoutMarshalers) XXX_Size() int {
	return xxx_messageInfo_SubMessageWithoutMarshalers.Size(m)
}
func (m *SubMessageWithoutMarshalers) XXX_DiscardUnknown() {
	xxx_messageInfo_SubMessageWithoutMarshalers.DiscardUnknown(m)
}

var xxx_messageInfo_SubMessageWithoutMarshalers proto.InternalMessageInfo

func (m *SubMessageWithoutMarshalers) GetOtherField() string {
	if m != nil {
		return m.OtherField
	}
	return ""
}

type MessageWithNullable struct {
	Sub                  SubMessage                    `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub"`
	Subs                 []SubMessage                  `protobuf:"bytes,2,rep,name=subs,proto3" json:"subs"`
	OtherSub             SubMessageWithoutMarshalers   `protobuf:"bytes,3,opt,name=other_sub,json=otherSub,proto3" json:"other_sub"`
	OtherSubs            []SubMessageWithoutMarshalers `protobuf:"bytes,4,rep,name=other_subs,json=otherSubs,proto3" json:"other_subs"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MessageWithNullable) Reset()         { *m = MessageWithNullable{} }
func (m *MessageWithNullable) String() string { return proto.CompactTextString(m) }
func (*MessageWithNullable) ProtoMessage()    {}
func (*MessageWithNullable) Descriptor() ([]byte, []int) {
	return fileDescriptor_592445b5231bc2b9, []int{3}
}
func (m *MessageWithNullable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithNullable.Unmarshal(m, b)
}
func (m *MessageWithNullable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithNullable.Marshal(b, m, deterministic)
}
func (m *MessageWithNullable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithNullable.Merge(m, src)
}
func (m *MessageWithNullable) XXX_Size() int {
	return xxx_messageInfo_MessageWithNullable.Size(m)
}
func (m *MessageWithNullable) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithNullable.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithNullable proto.InternalMessageInfo

func (m *MessageWithNullable) GetSub() SubMessage {
	if m != nil {
		return m.Sub
	}
	return SubMessage{}
}

func (m *MessageWithNullable) GetSubs() []SubMessage {
	if m != nil {
		return m.Subs
	}
	return nil
}

func (m *MessageWithNullable) GetOtherSub() SubMessageWithoutMarshalers {
	if m != nil {
		return m.OtherSub
	}
	return SubMessageWithoutMarshalers{}
}

func (m *MessageWithNullable) GetOtherSubs() []SubMessageWithoutMarshalers {
	if m != nil {
		return m.OtherSubs
	}
	return nil
}

type MessageWithEmbedded struct {
	*SubMessage                  `protobuf:"bytes,1,opt,name=sub,proto3,embedded=sub" json:"sub,omitempty"`
	*SubMessageWithoutMarshalers `protobuf:"bytes,2,opt,name=other_sub,json=otherSub,proto3,embedded=other_sub" json:"other_sub,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *MessageWithEmbedded) Reset()         { *m = MessageWithEmbedded{} }
func (m *MessageWithEmbedded) String() string { return proto.CompactTextString(m) }
func (*MessageWithEmbedded) ProtoMessage()    {}
func (*MessageWithEmbedded) Descriptor() ([]byte, []int) {
	return fileDescriptor_592445b5231bc2b9, []int{4}
}
func (m *MessageWithEmbedded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithEmbedded.Unmarshal(m, b)
}
func (m *MessageWithEmbedded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithEmbedded.Marshal(b, m, deterministic)
}
func (m *MessageWithEmbedded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithEmbedded.Merge(m, src)
}
func (m *MessageWithEmbedded) XXX_Size() int {
	return xxx_messageInfo_MessageWithEmbedded.Size(m)
}
func (m *MessageWithEmbedded) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithEmbedded.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithEmbedded proto.InternalMessageInfo

type MessageWithNullableEmbedded struct {
	SubMessage                  `protobuf:"bytes,1,opt,name=sub,proto3,embedded=sub" json:"sub"`
	SubMessageWithoutMarshalers `protobuf:"bytes,2,opt,name=other_sub,json=otherSub,proto3,embedded=other_sub" json:"other_sub"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *MessageWithNullableEmbedded) Reset()         { *m = MessageWithNullableEmbedded{} }
func (m *MessageWithNullableEmbedded) String() string { return proto.CompactTextString(m) }
func (*MessageWithNullableEmbedded) ProtoMessage()    {}
func (*MessageWithNullableEmbedded) Descriptor() ([]byte, []int) {
	return fileDescriptor_592445b5231bc2b9, []int{5}
}
func (m *MessageWithNullableEmbedded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithNullableEmbedded.Unmarshal(m, b)
}
func (m *MessageWithNullableEmbedded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithNullableEmbedded.Marshal(b, m, deterministic)
}
func (m *MessageWithNullableEmbedded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithNullableEmbedded.Merge(m, src)
}
func (m *MessageWithNullableEmbedded) XXX_Size() int {
	return xxx_messageInfo_MessageWithNullableEmbedded.Size(m)
}
func (m *MessageWithNullableEmbedded) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithNullableEmbedded.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithNullableEmbedded proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MessageWithGoGoOptions)(nil), "thethings.json.test.MessageWithGoGoOptions")
	proto.RegisterType((*SubMessage)(nil), "thethings.json.test.SubMessage")
	proto.RegisterType((*SubMessageWithoutMarshalers)(nil), "thethings.json.test.SubMessageWithoutMarshalers")
	proto.RegisterType((*MessageWithNullable)(nil), "thethings.json.test.MessageWithNullable")
	proto.RegisterType((*MessageWithEmbedded)(nil), "thethings.json.test.MessageWithEmbedded")
	proto.RegisterType((*MessageWithNullableEmbedded)(nil), "thethings.json.test.MessageWithNullableEmbedded")
}

func init() { proto.RegisterFile("gogo.proto", fileDescriptor_592445b5231bc2b9) }

var fileDescriptor_592445b5231bc2b9 = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0x3f, 0x4f, 0xdb, 0x4e,
	0x18, 0xc7, 0xb1, 0x13, 0xf8, 0x25, 0x17, 0x06, 0x38, 0xfe, 0x28, 0x09, 0x02, 0x47, 0x61, 0x41,
	0x3f, 0x09, 0xa7, 0xa2, 0x15, 0x55, 0x8b, 0x54, 0x95, 0xb4, 0x01, 0x22, 0x95, 0x54, 0x0a, 0x89,
	0xa8, 0x58, 0x2c, 0x1b, 0x1f, 0xb6, 0xab, 0xf8, 0x2e, 0xf2, 0xdd, 0xa9, 0xe2, 0x0d, 0x30, 0x74,
	0xea, 0xd0, 0xa1, 0x53, 0x97, 0x6e, 0xac, 0x55, 0xc7, 0xee, 0x55, 0xa7, 0xce, 0x1d, 0x52, 0x09,
	0x75, 0xe2, 0x2d, 0x74, 0xa9, 0xee, 0x92, 0xd8, 0x49, 0x13, 0xa0, 0x40, 0x86, 0x8e, 0xbe, 0xe7,
	0x9e, 0x8f, 0xbf, 0xf7, 0xf1, 0xe5, 0x09, 0x00, 0x0e, 0x71, 0x88, 0xde, 0x0c, 0x08, 0x23, 0x70,
	0x86, 0xb9, 0x88, 0xb9, 0x1e, 0x76, 0xa8, 0xfe, 0x92, 0x12, 0xac, 0x33, 0x44, 0x59, 0x76, 0xda,
	0xc4, 0x98, 0x30, 0x93, 0x79, 0x04, 0xd3, 0xf6, 0xbe, 0xec, 0xa2, 0x87, 0x19, 0x0a, 0xb0, 0xd9,
	0x28, 0x88, 0x66, 0xb9, 0x56, 0x88, 0x30, 0xd9, 0x25, 0x87, 0x10, 0xa7, 0x81, 0x0a, 0xf2, 0xc9,
	0xe2, 0x47, 0x05, 0x9b, 0x07, 0xb2, 0xbf, 0x53, 0xd7, 0xfe, 0xac, 0x33, 0xcf, 0x47, 0x94, 0x99,
	0x7e, 0xb3, 0xbd, 0x21, 0x7f, 0x92, 0x02, 0xf3, 0xbb, 0x88, 0x52, 0xd3, 0x41, 0xfb, 0x1e, 0x73,
	0xb7, 0xc9, 0x36, 0x79, 0xde, 0x94, 0x01, 0xe0, 0x16, 0x98, 0x45, 0xdc, 0x33, 0x5e, 0x79, 0xcc,
	0x35, 0x0e, 0x39, 0x65, 0xc4, 0x37, 0xb0, 0xe9, 0xa3, 0xb4, 0x92, 0x53, 0x56, 0x26, 0x8b, 0x73,
	0x67, 0x2d, 0x6d, 0xba, 0x54, 0x2f, 0x8b, 0xae, 0x27, 0xb2, 0x5a, 0x31, 0x7d, 0x54, 0x9d, 0x46,
	0xdc, 0xeb, 0x5f, 0x82, 0x9f, 0x55, 0xb0, 0x38, 0x0c, 0x64, 0x98, 0xd8, 0x36, 0xd8, 0x71, 0x13,
	0xa5, 0x55, 0x49, 0xfc, 0xa5, 0x9c, 0x9f, 0x66, 0xde, 0x2a, 0x60, 0xdb, 0xf1, 0x98, 0xcb, 0x2d,
	0xfd, 0x90, 0xf8, 0x85, 0x9a, 0x8b, 0x6a, 0xd2, 0x52, 0x19, 0xdb, 0x9c, 0xb2, 0xc0, 0x43, 0xb4,
	0x7d, 0x96, 0xc3, 0x55, 0x07, 0xe1, 0x55, 0x87, 0xac, 0x0a, 0x7b, 0x05, 0x61, 0xaf, 0x20, 0x50,
	0x54, 0xdf, 0x35, 0x03, 0xea, 0x9a, 0x8d, 0x9d, 0xd2, 0x0b, 0x58, 0xbe, 0x15, 0xa8, 0x8e, 0xfd,
	0x10, 0xf5, 0xbd, 0xa5, 0x15, 0x6f, 0x05, 0x2b, 0xd5, 0xcb, 0xeb, 0xf7, 0xce, 0x5a, 0x5a, 0x7a,
	0xc0, 0xd8, 0x26, 0xb6, 0x6b, 0xc7, 0x4d, 0x54, 0x4d, 0x0f, 0x88, 0xeb, 0x54, 0xe0, 0x4f, 0x15,
	0xfc, 0x8f, 0x09, 0x36, 0x30, 0x6f, 0x34, 0x4c, 0xab, 0x81, 0x8c, 0xcb, 0x65, 0xc6, 0xa4, 0xcc,
	0xd7, 0xea, 0x3f, 0x2b, 0xf3, 0x4b, 0x4b, 0x1b, 0x1b, 0x99, 0xd0, 0xe5, 0x0a, 0xc1, 0x95, 0x8e,
	0x98, 0x0b, 0xdd, 0x2e, 0xe3, 0x9e, 0x4d, 0x17, 0x69, 0xfe, 0xaa, 0x82, 0x25, 0xc4, 0x3d, 0x7a,
	0x89, 0xda, 0x78, 0x2e, 0xb6, 0x32, 0x59, 0x3c, 0x11, 0x6a, 0xdf, 0x2b, 0xe0, 0xd9, 0x88, 0xd4,
	0x6e, 0x06, 0x81, 0x79, 0x0c, 0x2b, 0x23, 0xf3, 0x2b, 0x79, 0x23, 0x13, 0x9c, 0x29, 0xd5, 0xcb,
	0x74, 0xb8, 0xd6, 0x8c, 0xd0, 0x35, 0x5c, 0xe6, 0x06, 0x48, 0x74, 0x27, 0x51, 0x7a, 0x3c, 0xa7,
	0xac, 0xa4, 0xd6, 0x32, 0x7a, 0x7b, 0x14, 0xe9, 0xdd, 0x51, 0xa4, 0x3f, 0xed, 0x6c, 0x28, 0xc6,
	0xdf, 0xfd, 0xd0, 0x94, 0x6a, 0xd8, 0x00, 0xf7, 0xc1, 0x5c, 0xdf, 0x7d, 0x0f, 0x49, 0x13, 0x57,
	0x91, 0x12, 0xe2, 0x6e, 0x49, 0xda, 0x4c, 0xcf, 0x27, 0xef, 0x96, 0xe1, 0x23, 0x90, 0x0c, 0xe7,
	0x5f, 0xfa, 0x3f, 0x09, 0xcb, 0x0e, 0xc0, 0x6a, 0xdd, 0x1d, 0xc5, 0xf8, 0x1b, 0x41, 0x8a, 0x5a,
	0xe0, 0x01, 0x98, 0xef, 0x0b, 0x16, 0xc1, 0x12, 0x57, 0xc2, 0x64, 0x34, 0x09, 0x9c, 0xed, 0x89,
	0x16, 0xd6, 0xf3, 0x79, 0x00, 0xf6, 0xb8, 0xd5, 0x19, 0xc5, 0x70, 0x16, 0x8c, 0x1f, 0x79, 0xa8,
	0x61, 0xcb, 0x61, 0x9b, 0xac, 0xb6, 0x1f, 0xf2, 0x3b, 0x60, 0x21, 0xda, 0x23, 0xc4, 0x13, 0xce,
	0x3a, 0xf7, 0x08, 0x05, 0x14, 0x6a, 0x20, 0x45, 0x98, 0x8b, 0x02, 0xa3, 0xb7, 0x15, 0xc8, 0xa5,
	0x2d, 0xb1, 0xf2, 0x30, 0x71, 0x7e, 0x9a, 0x89, 0x27, 0xc6, 0xa6, 0xc6, 0xf2, 0x9f, 0x54, 0x30,
	0xd3, 0xc3, 0xe9, 0xc6, 0x81, 0xf7, 0x41, 0x8c, 0x72, 0x4b, 0xb6, 0xa6, 0xd6, 0x34, 0x7d, 0xc8,
	0x9f, 0x94, 0x1e, 0x25, 0x28, 0xc6, 0xc5, 0x99, 0xaa, 0xa2, 0x03, 0x3e, 0x00, 0x71, 0xca, 0x2d,
	0x9a, 0x56, 0x73, 0xb1, 0xbf, 0xef, 0x94, 0x2d, 0x70, 0x0f, 0x24, 0xdb, 0xb1, 0xc5, 0x9b, 0x63,
	0xf2, 0xcd, 0x77, 0xae, 0xe8, 0x1f, 0x38, 0x7b, 0x07, 0x98, 0x90, 0xa0, 0x3d, 0x6e, 0xc1, 0x3a,
	0x00, 0x21, 0x94, 0xca, 0x1f, 0xee, 0xcd, 0xa9, 0xc9, 0x2e, 0x95, 0xe6, 0x3f, 0x28, 0x7d, 0xde,
	0x4a, 0xbe, 0x85, 0x6c, 0x1b, 0xd9, 0xd7, 0xf5, 0xf6, 0xad, 0xa5, 0x29, 0x6d, 0x6f, 0x7d, 0x87,
	0x57, 0x6f, 0x7a, 0x78, 0xc9, 0x0b, 0x0f, 0x9f, 0xff, 0xa8, 0x80, 0x85, 0x21, 0x5f, 0x37, 0x4c,
	0xbb, 0x71, 0xad, 0xb4, 0xf2, 0xe6, 0x46, 0x89, 0xf7, 0x47, 0x91, 0x38, 0x62, 0x86, 0xa9, 0x8b,
	0x8f, 0xe5, 0xed, 0x54, 0xa6, 0x94, 0x83, 0xf5, 0x9b, 0x4d, 0x2e, 0x6b, 0x42, 0x16, 0xee, 0xfe,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x09, 0xd2, 0xc6, 0xb8, 0x69, 0x09, 0x00, 0x00,
}
