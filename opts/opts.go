package opts

import (
	"github.com/bitcrshr/protoproxy/types/protoproxypb"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Opts interface {
	isOpt()
}

func Get[T Opts](desc protoreflect.Descriptor) T {
	switch t := desc.(type) {
	case protoreflect.FileDescriptor:
		file := &File{}

		if fileOpts, ok := t.Options().(*descriptorpb.FileOptions); ok {
			file.Proto = fileOpts
		}

		if ppFileOpts, ok := proto.GetExtension(t.Options(), protoproxypb.E_File).(*protoproxypb.FileOptions); ok {
			file.PP = ppFileOpts
		}

		return Opts(file).(T)
	case protoreflect.ServiceDescriptor:
		svc := &Service{}

		if svcOpts, ok := t.Options().(*descriptorpb.ServiceOptions); ok {
			svc.Proto = svcOpts
		}

		if ppSvcOpts, ok := proto.GetExtension(t.Options(), protoproxypb.E_Service).(*protoproxypb.ServiceOptions); ok {
			svc.PP = ppSvcOpts
		}

		return Opts(svc).(T)
	case protoreflect.MethodDescriptor:
		method := &Method{}

		if methodOpts, ok := t.Options().(*descriptorpb.MethodOptions); ok {
			method.Proto = methodOpts
		}

		if ppMethodOpts, ok := proto.GetExtension(t.Options(), protoproxypb.E_Method).(*protoproxypb.MethodOptions); ok {
			method.PP = ppMethodOpts
		}

		if httpRule, ok := proto.GetExtension(t.Options(), annotations.E_Http).(*annotations.HttpRule); ok {
			method.Http = httpRule
		}

		return Opts(method).(T)
	case protoreflect.MessageDescriptor:
		msg := &Message{}

		if msgOpts, ok := t.Options().(*descriptorpb.MessageOptions); ok {
			msg.Proto = msgOpts
		}

		if ppMsgOpts, ok := proto.GetExtension(t.Options(), protoproxypb.E_Message).(*protoproxypb.MessageOptions); ok {
			msg.PP = ppMsgOpts
		}

		return Opts(msg).(T)
	case protoreflect.OneofDescriptor:
		oo := &Oneof{}

		if ooOpts, ok := t.Options().(*descriptorpb.OneofOptions); ok {
			oo.Proto = ooOpts
		}

		if ppOoOpts, ok := proto.GetExtension(t.Options(), protoproxypb.E_Oneof).(*protoproxypb.OneofOptions); ok {
			oo.PP = ppOoOpts
		}

		return Opts(oo).(T)
	case protoreflect.FieldDescriptor:
		field := &Field{}

		if fieldOpts, ok := t.Options().(*descriptorpb.FieldOptions); ok {
			field.Proto = fieldOpts
		}

		if ppFieldOpts, ok := proto.GetExtension(t.Options(), protoproxypb.E_Field).(*protoproxypb.FieldOptions); ok {
			field.PP = ppFieldOpts
		}

		return Opts(field).(T)
	case protoreflect.EnumDescriptor:
		enum := &Enum{}

		if enumOpts, ok := t.Options().(*descriptorpb.EnumOptions); ok {
			enum.Proto = enumOpts
		}

		if ppEnumOpts, ok := proto.GetExtension(t.Options(), protoproxypb.E_Enum).(*protoproxypb.EnumOptions); ok {
			enum.PP = ppEnumOpts
		}

		return Opts(enum).(T)
	case protoreflect.EnumValueDescriptor:
		enumValue := &EnumValue{}

		if enumValueOpts, ok := t.Options().(*descriptorpb.EnumValueOptions); ok {
			enumValue.Proto = enumValueOpts
		}

		if ppEnumValueOpts, ok := proto.GetExtension(t.Options(), protoproxypb.E_EnumValue).(*protoproxypb.EnumValueOptions); ok {
			enumValue.PP = ppEnumValueOpts
		}

		return Opts(enumValue).(T)
	default:
		log.Fatal().Msgf("opts.Get: unhandled protoreflect.Descriptor type: %v", t)
		return Opts(&File{}).(T) // stupid compiler
	}
}

type File struct {
	PP    *protoproxypb.FileOptions
	Proto *descriptorpb.FileOptions
}

type Service struct {
	PP    *protoproxypb.ServiceOptions
	Proto *descriptorpb.ServiceOptions
}

type Method struct {
	PP    *protoproxypb.MethodOptions
	Proto *descriptorpb.MethodOptions
	Http  *annotations.HttpRule
}

type Message struct {
	PP    *protoproxypb.MessageOptions
	Proto *descriptorpb.MessageOptions
}

type Oneof struct {
	PP    *protoproxypb.OneofOptions
	Proto *descriptorpb.OneofOptions
}

type Field struct {
	PP    *protoproxypb.FieldOptions
	Proto *descriptorpb.FieldOptions
}

type Enum struct {
	PP    *protoproxypb.EnumOptions
	Proto *descriptorpb.EnumOptions
}

type EnumValue struct {
	PP    *protoproxypb.EnumValueOptions
	Proto *descriptorpb.EnumValueOptions
}

func (o *File) isOpt()      {}
func (o *Service) isOpt()   {}
func (o *Method) isOpt()    {}
func (o *Message) isOpt()   {}
func (o *Oneof) isOpt()     {}
func (o *Field) isOpt()     {}
func (o *Enum) isOpt()      {}
func (o *EnumValue) isOpt() {}
