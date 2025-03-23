// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: test/test.proto

package testpb

import (
	_ "github.com/bitcrshr/protoproxy/types/protoproxypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Foobar int32

const (
	Foobar_FOOBAR_UNSPECIFIED Foobar = 0
	Foobar_FOO                Foobar = 1
	Foobar_BAR                Foobar = 2
)

// Enum value maps for Foobar.
var (
	Foobar_name = map[int32]string{
		0: "FOOBAR_UNSPECIFIED",
		1: "FOO",
		2: "BAR",
	}
	Foobar_value = map[string]int32{
		"FOOBAR_UNSPECIFIED": 0,
		"FOO":                1,
		"BAR":                2,
	}
)

func (x Foobar) Enum() *Foobar {
	p := new(Foobar)
	*p = x
	return p
}

func (x Foobar) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Foobar) Descriptor() protoreflect.EnumDescriptor {
	return file_test_test_proto_enumTypes[0].Descriptor()
}

func (Foobar) Type() protoreflect.EnumType {
	return &file_test_test_proto_enumTypes[0]
}

func (x Foobar) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Foobar.Descriptor instead.
func (Foobar) EnumDescriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{0}
}

type TestRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldUno      string                 `protobuf:"bytes,1,opt,name=field_uno,json=fieldUno,proto3" json:"field_uno,omitempty"`
	FieldDos      string                 `protobuf:"bytes,2,opt,name=field_dos,json=fieldDos,proto3" json:"field_dos,omitempty"`
	FoobarEnumYay Foobar                 `protobuf:"varint,3,opt,name=foobar_enum_yay,json=foobarEnumYay,proto3,enum=test.v1.Foobar" json:"foobar_enum_yay,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	mi := &file_test_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequest) GetFieldUno() string {
	if x != nil {
		return x.FieldUno
	}
	return ""
}

func (x *TestRequest) GetFieldDos() string {
	if x != nil {
		return x.FieldDos
	}
	return ""
}

func (x *TestRequest) GetFoobarEnumYay() Foobar {
	if x != nil {
		return x.FoobarEnumYay
	}
	return Foobar_FOOBAR_UNSPECIFIED
}

type AnothaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnothaRequest) Reset() {
	*x = AnothaRequest{}
	mi := &file_test_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnothaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnothaRequest) ProtoMessage() {}

func (x *AnothaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnothaRequest.ProtoReflect.Descriptor instead.
func (*AnothaRequest) Descriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{1}
}

type AnothaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Woohoos       []string               `protobuf:"bytes,1,rep,name=woohoos,proto3" json:"woohoos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnothaResponse) Reset() {
	*x = AnothaResponse{}
	mi := &file_test_test_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnothaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnothaResponse) ProtoMessage() {}

func (x *AnothaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_test_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnothaResponse.ProtoReflect.Descriptor instead.
func (*AnothaResponse) Descriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{2}
}

func (x *AnothaResponse) GetWoohoos() []string {
	if x != nil {
		return x.Woohoos
	}
	return nil
}

type TestResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bar           string                 `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	mi := &file_test_test_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_test_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{3}
}

func (x *TestResponse) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

var File_test_test_proto protoreflect.FileDescriptor

var file_test_test_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x70, 0x62, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x75, 0x6e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x82, 0xb5, 0x18, 0x02, 0x20, 0x01, 0x52,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x6e, 0x6f, 0x12, 0x23, 0x0a, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x64, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x82, 0xb5,
	0x18, 0x02, 0x18, 0x03, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x6f, 0x73, 0x12, 0x37,
	0x0a, 0x0f, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x79, 0x61,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x52, 0x0d, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x59, 0x61, 0x79, 0x3a, 0x06, 0x82, 0xb5, 0x18, 0x02, 0x08, 0x01, 0x22,
	0x0f, 0x0a, 0x0d, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x32, 0x0a, 0x0e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x77, 0x6f, 0x6f, 0x68, 0x6f, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x06, 0x82, 0xb5, 0x18, 0x02, 0x08, 0x01, 0x52, 0x07, 0x77, 0x6f, 0x6f,
	0x68, 0x6f, 0x6f, 0x73, 0x22, 0x20, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x61, 0x72, 0x2a, 0x32, 0x0a, 0x06, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72,
	0x12, 0x16, 0x0a, 0x12, 0x46, 0x4f, 0x4f, 0x42, 0x41, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x4f, 0x4f, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41, 0x52, 0x10, 0x02, 0x32, 0xac, 0x01, 0x0a, 0x0b, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0a, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a,
	0x22, 0x0a, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4b, 0x0a, 0x09,
	0x41, 0x6e, 0x6f, 0x74, 0x68, 0x61, 0x4f, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6f, 0x74,
	0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x07, 0x2a, 0x05, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x42, 0x8d, 0x01, 0x82, 0xb5, 0x18, 0x02,
	0x08, 0x05, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x72, 0x73, 0x68,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0xa2, 0x02,
	0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x07, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x54, 0x65, 0x73, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_test_test_proto_rawDescOnce sync.Once
	file_test_test_proto_rawDescData []byte
)

func file_test_test_proto_rawDescGZIP() []byte {
	file_test_test_proto_rawDescOnce.Do(func() {
		file_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_test_test_proto_rawDesc), len(file_test_test_proto_rawDesc)))
	})
	return file_test_test_proto_rawDescData
}

var file_test_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_test_proto_goTypes = []any{
	(Foobar)(0),            // 0: test.v1.Foobar
	(*TestRequest)(nil),    // 1: test.v1.TestRequest
	(*AnothaRequest)(nil),  // 2: test.v1.AnothaRequest
	(*AnothaResponse)(nil), // 3: test.v1.AnothaResponse
	(*TestResponse)(nil),   // 4: test.v1.TestResponse
}
var file_test_test_proto_depIdxs = []int32{
	0, // 0: test.v1.TestRequest.foobar_enum_yay:type_name -> test.v1.Foobar
	1, // 1: test.v1.TestService.TestMethod:input_type -> test.v1.TestRequest
	2, // 2: test.v1.TestService.AnothaOne:input_type -> test.v1.AnothaRequest
	4, // 3: test.v1.TestService.TestMethod:output_type -> test.v1.TestResponse
	3, // 4: test.v1.TestService.AnothaOne:output_type -> test.v1.AnothaResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_test_proto_init() }
func file_test_test_proto_init() {
	if File_test_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_test_test_proto_rawDesc), len(file_test_test_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_test_proto_goTypes,
		DependencyIndexes: file_test_test_proto_depIdxs,
		EnumInfos:         file_test_test_proto_enumTypes,
		MessageInfos:      file_test_test_proto_msgTypes,
	}.Build()
	File_test_test_proto = out.File
	file_test_test_proto_goTypes = nil
	file_test_test_proto_depIdxs = nil
}
