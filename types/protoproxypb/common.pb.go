// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: protoproxypb/common.proto

package protoproxypb

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
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

type TextCase int32

const (
	TextCase_TEXT_CASE_UNSPECIFIED TextCase = 0
	// camelCase
	TextCase_CAMEL TextCase = 1
	// PascalCase
	TextCase_PASCAL TextCase = 2
	// kebab-case
	TextCase_KEBAB TextCase = 3
	// snake_case
	TextCase_SNAKE TextCase = 4
	// SCREAMING_SNAKE_CASE
	TextCase_SCREAMING_SNAKE TextCase = 5
	// SCREAMING-KEBAB-CASE
	TextCase_SCREAMING_KEBAB TextCase = 6
	// lowercase
	TextCase_LOWER TextCase = 7
	// UPPERCASE
	TextCase_UPPER TextCase = 8
)

// Enum value maps for TextCase.
var (
	TextCase_name = map[int32]string{
		0: "TEXT_CASE_UNSPECIFIED",
		1: "CAMEL",
		2: "PASCAL",
		3: "KEBAB",
		4: "SNAKE",
		5: "SCREAMING_SNAKE",
		6: "SCREAMING_KEBAB",
		7: "LOWER",
		8: "UPPER",
	}
	TextCase_value = map[string]int32{
		"TEXT_CASE_UNSPECIFIED": 0,
		"CAMEL":                 1,
		"PASCAL":                2,
		"KEBAB":                 3,
		"SNAKE":                 4,
		"SCREAMING_SNAKE":       5,
		"SCREAMING_KEBAB":       6,
		"LOWER":                 7,
		"UPPER":                 8,
	}
)

func (x TextCase) Enum() *TextCase {
	p := new(TextCase)
	*p = x
	return p
}

func (x TextCase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TextCase) Descriptor() protoreflect.EnumDescriptor {
	return file_protoproxypb_common_proto_enumTypes[0].Descriptor()
}

func (TextCase) Type() protoreflect.EnumType {
	return &file_protoproxypb_common_proto_enumTypes[0]
}

func (x TextCase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TextCase.Descriptor instead.
func (TextCase) EnumDescriptor() ([]byte, []int) {
	return file_protoproxypb_common_proto_rawDescGZIP(), []int{0}
}

type Encoding int32

const (
	Encoding_ENCODING_UNSPECIFIED    Encoding = 0
	Encoding_JSON                    Encoding = 1
	Encoding_URL_ENCODED_FORM_VALUES Encoding = 2
)

// Enum value maps for Encoding.
var (
	Encoding_name = map[int32]string{
		0: "ENCODING_UNSPECIFIED",
		1: "JSON",
		2: "URL_ENCODED_FORM_VALUES",
	}
	Encoding_value = map[string]int32{
		"ENCODING_UNSPECIFIED":    0,
		"JSON":                    1,
		"URL_ENCODED_FORM_VALUES": 2,
	}
)

func (x Encoding) Enum() *Encoding {
	p := new(Encoding)
	*p = x
	return p
}

func (x Encoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Encoding) Descriptor() protoreflect.EnumDescriptor {
	return file_protoproxypb_common_proto_enumTypes[1].Descriptor()
}

func (Encoding) Type() protoreflect.EnumType {
	return &file_protoproxypb_common_proto_enumTypes[1]
}

func (x Encoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Encoding.Descriptor instead.
func (Encoding) EnumDescriptor() ([]byte, []int) {
	return file_protoproxypb_common_proto_rawDescGZIP(), []int{1}
}

type Scheme int32

const (
	Scheme_SCHEME_UNSPECIFIED Scheme = 0
	Scheme_HTTP               Scheme = 1
	Scheme_HTTPS              Scheme = 2
)

// Enum value maps for Scheme.
var (
	Scheme_name = map[int32]string{
		0: "SCHEME_UNSPECIFIED",
		1: "HTTP",
		2: "HTTPS",
	}
	Scheme_value = map[string]int32{
		"SCHEME_UNSPECIFIED": 0,
		"HTTP":               1,
		"HTTPS":              2,
	}
)

func (x Scheme) Enum() *Scheme {
	p := new(Scheme)
	*p = x
	return p
}

func (x Scheme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scheme) Descriptor() protoreflect.EnumDescriptor {
	return file_protoproxypb_common_proto_enumTypes[2].Descriptor()
}

func (Scheme) Type() protoreflect.EnumType {
	return &file_protoproxypb_common_proto_enumTypes[2]
}

func (x Scheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scheme.Descriptor instead.
func (Scheme) EnumDescriptor() ([]byte, []int) {
	return file_protoproxypb_common_proto_rawDescGZIP(), []int{2}
}

type StatusMapping struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	HttpStatusCode int32                  `protobuf:"varint,1,opt,name=http_status_code,json=httpStatusCode,proto3" json:"http_status_code,omitempty"`
	GrpcStatusCode code.Code              `protobuf:"varint,2,opt,name=grpc_status_code,json=grpcStatusCode,proto3,enum=google.rpc.Code" json:"grpc_status_code,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *StatusMapping) Reset() {
	*x = StatusMapping{}
	mi := &file_protoproxypb_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusMapping) ProtoMessage() {}

func (x *StatusMapping) ProtoReflect() protoreflect.Message {
	mi := &file_protoproxypb_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusMapping.ProtoReflect.Descriptor instead.
func (*StatusMapping) Descriptor() ([]byte, []int) {
	return file_protoproxypb_common_proto_rawDescGZIP(), []int{0}
}

func (x *StatusMapping) GetHttpStatusCode() int32 {
	if x != nil {
		return x.HttpStatusCode
	}
	return 0
}

func (x *StatusMapping) GetGrpcStatusCode() code.Code {
	if x != nil {
		return x.GrpcStatusCode
	}
	return code.Code(0)
}

var File_protoproxypb_common_proto protoreflect.FileDescriptor

var file_protoproxypb_common_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x70, 0x62, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75,
	0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x28, 0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x10, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x92, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x43, 0x41, 0x53, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x43, 0x41, 0x4d, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x53, 0x43,
	0x41, 0x4c, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x45, 0x42, 0x41, 0x42, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x4e, 0x41, 0x4b, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x43,
	0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4e, 0x41, 0x4b, 0x45, 0x10, 0x05, 0x12,
	0x13, 0x0a, 0x0f, 0x53, 0x43, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x42,
	0x41, 0x42, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x07, 0x12,
	0x09, 0x0a, 0x05, 0x55, 0x50, 0x50, 0x45, 0x52, 0x10, 0x08, 0x2a, 0x4b, 0x0a, 0x08, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x52,
	0x4c, 0x5f, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x53, 0x10, 0x02, 0x2a, 0x35, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54,
	0x50, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x53, 0x10, 0x02, 0x42, 0x98,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x74,
	0x63, 0x72, 0x73, 0x68, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0xca, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0xe2, 0x02, 0x16, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_protoproxypb_common_proto_rawDescOnce sync.Once
	file_protoproxypb_common_proto_rawDescData []byte
)

func file_protoproxypb_common_proto_rawDescGZIP() []byte {
	file_protoproxypb_common_proto_rawDescOnce.Do(func() {
		file_protoproxypb_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protoproxypb_common_proto_rawDesc), len(file_protoproxypb_common_proto_rawDesc)))
	})
	return file_protoproxypb_common_proto_rawDescData
}

var file_protoproxypb_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protoproxypb_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protoproxypb_common_proto_goTypes = []any{
	(TextCase)(0),         // 0: protoproxy.TextCase
	(Encoding)(0),         // 1: protoproxy.Encoding
	(Scheme)(0),           // 2: protoproxy.Scheme
	(*StatusMapping)(nil), // 3: protoproxy.StatusMapping
	(code.Code)(0),        // 4: google.rpc.Code
}
var file_protoproxypb_common_proto_depIdxs = []int32{
	4, // 0: protoproxy.StatusMapping.grpc_status_code:type_name -> google.rpc.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protoproxypb_common_proto_init() }
func file_protoproxypb_common_proto_init() {
	if File_protoproxypb_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protoproxypb_common_proto_rawDesc), len(file_protoproxypb_common_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoproxypb_common_proto_goTypes,
		DependencyIndexes: file_protoproxypb_common_proto_depIdxs,
		EnumInfos:         file_protoproxypb_common_proto_enumTypes,
		MessageInfos:      file_protoproxypb_common_proto_msgTypes,
	}.Build()
	File_protoproxypb_common_proto = out.File
	file_protoproxypb_common_proto_goTypes = nil
	file_protoproxypb_common_proto_depIdxs = nil
}
