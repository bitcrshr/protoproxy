package desckind

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DescriptorKind int

const (
	File DescriptorKind = iota
	Service
	Method
	Message
	Field
	Enum
	EnumValue
	Oneof
)

func Of(d protoreflect.Descriptor) DescriptorKind {
	switch d.(type) {
	case protoreflect.FileDescriptor:
		return File
	case protoreflect.ServiceDescriptor:
		return Service
	case protoreflect.MethodDescriptor:
		return Method
	case protoreflect.MessageDescriptor:
		return Message
	case protoreflect.FieldDescriptor:
		return Field
	case protoreflect.EnumDescriptor:
		return Enum
	case protoreflect.EnumValueDescriptor:
		return EnumValue
	case protoreflect.OneofDescriptor:
		return Oneof
	default:
		log.Fatal().Msgf("desckind.From: unhandled descriptor of type %T", d)
		return -1
	}
}
