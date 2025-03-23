package ppopt

import (
	"github.com/bitcrshr/protoproxy/types/protoproxypb"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Get(d protoreflect.Descriptor, o Opt) {
	switch desc := d.(type) {
	case protoreflect.FileDescriptor:
		file := &File{}
		o = file

		fd, ok := desc.Options().(*descriptorpb.FileOptions)
		if !ok {
			return
		}

		file.Proto = fd

		if ppfd, ok := proto.GetExtension(fd, protoproxypb.E_File).(*protoproxypb.FileOptions); ok {
			file.PP = ppfd
		}
	case protoreflect.ServiceDescriptor:
		svc := &Service{}
		o = svc

		sd, ok := desc.Options().(*descriptorpb.ServiceOptions)
		if !ok {
			return
		}

		svc.Proto = sd

		if ppsd, ok := proto.GetExtension(sd, protoproxypb.E_Service).(*protoproxypb.ServiceOptions); ok {
			svc.PP = ppsd
		}
	case protoreflect.MethodDescriptor:
		method := &Method{}
		o = method

		md, ok := desc.Options().(*descriptorpb.MethodOptions)
		if !ok {
			return
		}

		method.Proto = md

		if ppmd, ok := proto.GetExtension(md, protoproxypb.E_Method).(*protoproxypb.MethodOptions); ok {
			method.PP = ppmd
		}

		http, ok := proto.GetExtension(md, annotations.E_Http).(*annotations.HttpRule)
		if ok {
			method.Http = http
		}
	case protoreflect.MessageDescriptor:
		msg := &Message{}
		o = msg

		md, ok := desc.Options().(*descriptorpb.MessageOptions)
		if !ok {
			return
		}

		msg.Proto = md

		if ppmd, ok := proto.GetExtension(md, protoproxypb.E_Message).(*protoproxypb.MessageOptions); ok {
			msg.PP = ppmd
		}
	case protoreflect.FieldDescriptor:
		field := &Field{}
		o = field

		fd, ok := desc.Options().(*descriptorpb.FieldOptions)
		if !ok {
			return
		}

		field.Proto = fd

		if ppfd, ok := proto.GetExtension(fd, protoproxypb.E_Field).(*protoproxypb.FieldOptions); ok {
			field.PP = ppfd
		}
	case protoreflect.EnumDescriptor:
		enum := &Enum{}
		o = enum

		ed, ok := desc.Options().(*descriptorpb.EnumOptions)
		if !ok {
			return
		}

		enum.Proto = ed

		if pped, ok := proto.GetExtension(ed, protoproxypb.E_Enum).(*protoproxypb.EnumOptions); ok {
			enum.PP = pped
		}
	case protoreflect.EnumValueDescriptor:
		enumValue := &EnumValue{}
		o = enumValue

		evd, ok := desc.Options().(*descriptorpb.EnumValueOptions)
		if !ok {
			return
		}

		if ppevd, ok := proto.GetExtension(evd, protoproxypb.E_EnumValue).(*protoproxypb.EnumValueOptions); ok {
			enumValue.PP = ppevd
		}

	case protoreflect.OneofDescriptor:
		oneof := &Oneof{}
		o = oneof

		od, ok := desc.Options().(*descriptorpb.OneofOptions)
		if !ok {
			return
		}

		oneof.Proto = od

		if ppod, ok := proto.GetExtension(od, protoproxypb.E_Oneof).(*protoproxypb.OneofOptions); ok {
			oneof.PP = ppod
		}
	}
}

type Opt interface {
	isOpt()
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

type Oneof struct {
	PP    *protoproxypb.OneofOptions
	Proto *descriptorpb.OneofOptions
}

func (f *File) isOpt()       {}
func (s *Service) isOpt()    {}
func (m *Method) isOpt()     {}
func (m *Message) isOpt()    {}
func (f *Field) isOpt()      {}
func (e *Enum) isOpt()       {}
func (ev *EnumValue) isOpt() {}
func (o *Oneof) isOpt()      {}
