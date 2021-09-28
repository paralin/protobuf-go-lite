// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: enums.proto

package test

import (
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegularEnum int32

const (
	RegularEnum_REGULAR_UNKNOWN RegularEnum = 0
	RegularEnum_REGULAR_A       RegularEnum = 1
	RegularEnum_REGULAR_B       RegularEnum = 2
)

// Enum value maps for RegularEnum.
var (
	RegularEnum_name = map[int32]string{
		0: "REGULAR_UNKNOWN",
		1: "REGULAR_A",
		2: "REGULAR_B",
	}
	RegularEnum_value = map[string]int32{
		"REGULAR_UNKNOWN": 0,
		"REGULAR_A":       1,
		"REGULAR_B":       2,
	}
)

func (x RegularEnum) Enum() *RegularEnum {
	p := new(RegularEnum)
	*p = x
	return p
}

func (x RegularEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegularEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[0].Descriptor()
}

func (RegularEnum) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[0]
}

func (x RegularEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegularEnum.Descriptor instead.
func (RegularEnum) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

type CustomEnum int32

const (
	CustomEnum_CUSTOM_UNKNOWN CustomEnum = 0
	CustomEnum_CUSTOM_V1_0    CustomEnum = 1
	CustomEnum_CUSTOM_V1_0_1  CustomEnum = 2
)

// Enum value maps for CustomEnum.
var (
	CustomEnum_name = map[int32]string{
		0: "CUSTOM_UNKNOWN",
		1: "CUSTOM_V1_0",
		2: "CUSTOM_V1_0_1",
	}
	CustomEnum_value = map[string]int32{
		"CUSTOM_UNKNOWN": 0,
		"CUSTOM_V1_0":    1,
		"CUSTOM_V1_0_1":  2,
	}
)

func (x CustomEnum) Enum() *CustomEnum {
	p := new(CustomEnum)
	*p = x
	return p
}

func (x CustomEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[1].Descriptor()
}

func (CustomEnum) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[1]
}

func (x CustomEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomEnum.Descriptor instead.
func (CustomEnum) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{1}
}

type CustomEnumValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value CustomEnum `protobuf:"varint,1,opt,name=value,proto3,enum=thethings.json.test.CustomEnum" json:"value,omitempty"`
}

func (x *CustomEnumValue) Reset() {
	*x = CustomEnumValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEnumValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEnumValue) ProtoMessage() {}

func (x *CustomEnumValue) ProtoReflect() protoreflect.Message {
	mi := &file_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEnumValue.ProtoReflect.Descriptor instead.
func (*CustomEnumValue) Descriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

func (x *CustomEnumValue) GetValue() CustomEnum {
	if x != nil {
		return x.Value
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

type MessageWithEnums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regular        RegularEnum        `protobuf:"varint,1,opt,name=regular,proto3,enum=thethings.json.test.RegularEnum" json:"regular,omitempty"`
	Regulars       []RegularEnum      `protobuf:"varint,2,rep,packed,name=regulars,proto3,enum=thethings.json.test.RegularEnum" json:"regulars,omitempty"`
	Custom         CustomEnum         `protobuf:"varint,3,opt,name=custom,proto3,enum=thethings.json.test.CustomEnum" json:"custom,omitempty"`
	Customs        []CustomEnum       `protobuf:"varint,4,rep,packed,name=customs,proto3,enum=thethings.json.test.CustomEnum" json:"customs,omitempty"`
	WrappedCustom  *CustomEnumValue   `protobuf:"bytes,5,opt,name=wrapped_custom,json=wrappedCustom,proto3" json:"wrapped_custom,omitempty"`
	WrappedCustoms []*CustomEnumValue `protobuf:"bytes,6,rep,name=wrapped_customs,json=wrappedCustoms,proto3" json:"wrapped_customs,omitempty"`
}

func (x *MessageWithEnums) Reset() {
	*x = MessageWithEnums{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithEnums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithEnums) ProtoMessage() {}

func (x *MessageWithEnums) ProtoReflect() protoreflect.Message {
	mi := &file_enums_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithEnums.ProtoReflect.Descriptor instead.
func (*MessageWithEnums) Descriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{1}
}

func (x *MessageWithEnums) GetRegular() RegularEnum {
	if x != nil {
		return x.Regular
	}
	return RegularEnum_REGULAR_UNKNOWN
}

func (x *MessageWithEnums) GetRegulars() []RegularEnum {
	if x != nil {
		return x.Regulars
	}
	return nil
}

func (x *MessageWithEnums) GetCustom() CustomEnum {
	if x != nil {
		return x.Custom
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

func (x *MessageWithEnums) GetCustoms() []CustomEnum {
	if x != nil {
		return x.Customs
	}
	return nil
}

func (x *MessageWithEnums) GetWrappedCustom() *CustomEnumValue {
	if x != nil {
		return x.WrappedCustom
	}
	return nil
}

func (x *MessageWithEnums) GetWrappedCustoms() []*CustomEnumValue {
	if x != nil {
		return x.WrappedCustoms
	}
	return nil
}

type MessageWithOneofEnums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*MessageWithOneofEnums_Regular
	//	*MessageWithOneofEnums_Custom
	//	*MessageWithOneofEnums_WrappedCustom
	Value isMessageWithOneofEnums_Value `protobuf_oneof:"value"`
}

func (x *MessageWithOneofEnums) Reset() {
	*x = MessageWithOneofEnums{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithOneofEnums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithOneofEnums) ProtoMessage() {}

func (x *MessageWithOneofEnums) ProtoReflect() protoreflect.Message {
	mi := &file_enums_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithOneofEnums.ProtoReflect.Descriptor instead.
func (*MessageWithOneofEnums) Descriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{2}
}

func (m *MessageWithOneofEnums) GetValue() isMessageWithOneofEnums_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *MessageWithOneofEnums) GetRegular() RegularEnum {
	if x, ok := x.GetValue().(*MessageWithOneofEnums_Regular); ok {
		return x.Regular
	}
	return RegularEnum_REGULAR_UNKNOWN
}

func (x *MessageWithOneofEnums) GetCustom() CustomEnum {
	if x, ok := x.GetValue().(*MessageWithOneofEnums_Custom); ok {
		return x.Custom
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

func (x *MessageWithOneofEnums) GetWrappedCustom() *CustomEnumValue {
	if x, ok := x.GetValue().(*MessageWithOneofEnums_WrappedCustom); ok {
		return x.WrappedCustom
	}
	return nil
}

type isMessageWithOneofEnums_Value interface {
	isMessageWithOneofEnums_Value()
}

type MessageWithOneofEnums_Regular struct {
	Regular RegularEnum `protobuf:"varint,1,opt,name=regular,proto3,enum=thethings.json.test.RegularEnum,oneof"`
}

type MessageWithOneofEnums_Custom struct {
	Custom CustomEnum `protobuf:"varint,2,opt,name=custom,proto3,enum=thethings.json.test.CustomEnum,oneof"`
}

type MessageWithOneofEnums_WrappedCustom struct {
	WrappedCustom *CustomEnumValue `protobuf:"bytes,3,opt,name=wrapped_custom,json=wrappedCustom,proto3,oneof"`
}

func (*MessageWithOneofEnums_Regular) isMessageWithOneofEnums_Value() {}

func (*MessageWithOneofEnums_Custom) isMessageWithOneofEnums_Value() {}

func (*MessageWithOneofEnums_WrappedCustom) isMessageWithOneofEnums_Value() {}

var File_enums_proto protoreflect.FileDescriptor

var file_enums_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74,
	0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x06,
	0xea, 0xaa, 0x19, 0x02, 0x18, 0x01, 0x22, 0x9c, 0x03, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x72,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x07,
	0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x72, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x39,
	0x0a, 0x07, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x75, 0x6d,
	0x52, 0x07, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x12, 0x4b, 0x0a, 0x0e, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x4d, 0x0a, 0x0f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x75, 0x6d,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x12,
	0x3c, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f,
	0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x6e,
	0x75, 0x6d, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x39, 0x0a,
	0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x4d, 0x0a, 0x0e, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6a, 0x73, 0x6f,
	0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x75,
	0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2a, 0x52, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x13, 0x0a, 0x0f, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x5f,
	0x41, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x5f, 0x42,
	0x10, 0x02, 0x1a, 0x10, 0xea, 0xaa, 0x19, 0x04, 0x08, 0x00, 0x20, 0x00, 0x88, 0xa3, 0x1e, 0x00,
	0xb0, 0xa4, 0x1e, 0x00, 0x2a, 0x73, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x0b, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x5f, 0x56, 0x31, 0x5f, 0x30, 0x10, 0x01, 0x1a, 0x10, 0xea, 0xaa, 0x19, 0x0c, 0x0a, 0x03, 0x31,
	0x2e, 0x30, 0x12, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x12, 0x1e, 0x0a, 0x0d, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x5f, 0x56, 0x31, 0x5f, 0x30, 0x5f, 0x31, 0x10, 0x02, 0x1a, 0x0b, 0xea, 0xaa,
	0x19, 0x07, 0x0a, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x31, 0x1a, 0x0e, 0xea, 0xaa, 0x19, 0x0a, 0x18,
	0x01, 0x2a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x42, 0x48, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0xf0, 0xe2, 0x1e, 0x01, 0xe8, 0xe2, 0x1e, 0x00, 0xea, 0xaa, 0x19, 0x04, 0x08,
	0x01, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_proto_rawDescOnce sync.Once
	file_enums_proto_rawDescData = file_enums_proto_rawDesc
)

func file_enums_proto_rawDescGZIP() []byte {
	file_enums_proto_rawDescOnce.Do(func() {
		file_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_proto_rawDescData)
	})
	return file_enums_proto_rawDescData
}

var file_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_enums_proto_goTypes = []interface{}{
	(RegularEnum)(0),              // 0: thethings.json.test.RegularEnum
	(CustomEnum)(0),               // 1: thethings.json.test.CustomEnum
	(*CustomEnumValue)(nil),       // 2: thethings.json.test.CustomEnumValue
	(*MessageWithEnums)(nil),      // 3: thethings.json.test.MessageWithEnums
	(*MessageWithOneofEnums)(nil), // 4: thethings.json.test.MessageWithOneofEnums
}
var file_enums_proto_depIdxs = []int32{
	1,  // 0: thethings.json.test.CustomEnumValue.value:type_name -> thethings.json.test.CustomEnum
	0,  // 1: thethings.json.test.MessageWithEnums.regular:type_name -> thethings.json.test.RegularEnum
	0,  // 2: thethings.json.test.MessageWithEnums.regulars:type_name -> thethings.json.test.RegularEnum
	1,  // 3: thethings.json.test.MessageWithEnums.custom:type_name -> thethings.json.test.CustomEnum
	1,  // 4: thethings.json.test.MessageWithEnums.customs:type_name -> thethings.json.test.CustomEnum
	2,  // 5: thethings.json.test.MessageWithEnums.wrapped_custom:type_name -> thethings.json.test.CustomEnumValue
	2,  // 6: thethings.json.test.MessageWithEnums.wrapped_customs:type_name -> thethings.json.test.CustomEnumValue
	0,  // 7: thethings.json.test.MessageWithOneofEnums.regular:type_name -> thethings.json.test.RegularEnum
	1,  // 8: thethings.json.test.MessageWithOneofEnums.custom:type_name -> thethings.json.test.CustomEnum
	2,  // 9: thethings.json.test.MessageWithOneofEnums.wrapped_custom:type_name -> thethings.json.test.CustomEnumValue
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_enums_proto_init() }
func file_enums_proto_init() {
	if File_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEnumValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithEnums); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enums_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithOneofEnums); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_enums_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MessageWithOneofEnums_Regular)(nil),
		(*MessageWithOneofEnums_Custom)(nil),
		(*MessageWithOneofEnums_WrappedCustom)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_proto_goTypes,
		DependencyIndexes: file_enums_proto_depIdxs,
		EnumInfos:         file_enums_proto_enumTypes,
		MessageInfos:      file_enums_proto_msgTypes,
	}.Build()
	File_enums_proto = out.File
	file_enums_proto_rawDesc = nil
	file_enums_proto_goTypes = nil
	file_enums_proto_depIdxs = nil
}
