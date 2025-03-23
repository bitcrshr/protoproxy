package pathmap

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/bitcrshr/protoproxy/opts"
	"github.com/bitcrshr/protoproxy/types/protoproxypb"
	"github.com/bitcrshr/protoproxy/util"
	"github.com/iancoleman/strcase"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type PathMap interface {
	SetValueOf(msgName protoreflect.FullName, fieldName protoreflect.FullName, value any) error
	ParseRequest(method string) ([]byte, error)
}

type pathMap struct {
	msgMap   map[protoreflect.FullName]*entry
	fieldMap map[protoreflect.FullName]*entry
	reqMap   map[string]protoreflect.MessageDescriptor
	resMap   map[string]protoreflect.MessageDescriptor
}

func (pm *pathMap) SetValueOf(msgName protoreflect.FullName, fieldName protoreflect.FullName, value any) error {
	ent, ok := pm.fieldMap[fieldName]
	if !ok || ent == nil {
		err := eris.Errorf("entry for field `%s` not found", string(fieldName))
		log.Error().Err(err).Msg("PathMap.SetValueOf failed")
		return err
	}

	ptr, ok := ent.valueMap[fieldName]
	if !ok || ptr == nil {
		err := eris.Errorf("value map entry for field `%s` not found", string(fieldName))
		log.Error().Err(err).Msg("PathMap.SetValueOf failed")
		return err
	}

	*ptr = value

	return nil
}

func (pm *pathMap) ParseRequest(method string) ([]byte, error) {
	msg, ok := pm.reqMap[method]
	if !ok || msg == nil {
		err := status.Errorf(codes.NotFound, "method %s not found", method)
		log.Error().Err(err).Msg("PathMap.ParseRequest failed")
		return nil, err
	}

	msgEntry, ok := pm.msgMap[msg.FullName()]
	if !ok || msgEntry == nil {
		err := status.Errorf(codes.NotFound, "entry for msg `%s` not found", string(msg.FullName()))
		log.Error().Err(err).Msg("PathMap.MarshalJSON failed")
		return nil, err
	}

	jsonBytes, err := json.Marshal(msgEntry.jsonMap)
	if err != nil {
		err = eris.Wrapf(err, "failed to marshal message `%s`", string(msg.FullName()))
		log.Error().Err(err).Msg("PathMap.MarshalJSON failed")
		return nil, err
	}

	return jsonBytes, nil
}

// represents the input or output of a method
type entry struct {
	jsonMap  map[string]*any
	valueMap map[protoreflect.FullName]*any
	pathMap  map[protoreflect.FullName][]string

	rootMessage protoreflect.MessageDescriptor

	method protoreflect.MethodDescriptor
}

func New(fds *descriptorpb.FileDescriptorSet) (PathMap, error) {
	files, err := protodesc.NewFiles(fds)
	if err != nil {
		err = eris.Wrap(err, "failed to parse file descriptor set")
		log.Error().Err(err).Msg("pathmap.New failed")
		return nil, err
	}

	pm := &pathMap{
		msgMap:   map[protoreflect.FullName]*entry{},
		fieldMap: map[protoreflect.FullName]*entry{},
		reqMap:   map[string]protoreflect.MessageDescriptor{},
		resMap:   map[string]protoreflect.MessageDescriptor{},
	}

	for file := range files.RangeFiles {
		for svc := range util.DescIter(file.Services()) {
			for method := range util.DescIter(svc.Methods()) {
				input := method.Input()
				output := method.Output()

				pm.reqMap[fqMethodName(method)] = input
				pm.resMap[fqMethodName(method)] = output

				inputEntry := &entry{
					jsonMap:     map[string]*any{},
					valueMap:    map[protoreflect.FullName]*any{},
					pathMap:     map[protoreflect.FullName][]string{},
					rootMessage: input,
					method:      method,
				}

				outputEntry := &entry{
					jsonMap:     map[string]*any{},
					valueMap:    map[protoreflect.FullName]*any{},
					pathMap:     map[protoreflect.FullName][]string{},
					rootMessage: output,
					method:      method,
				}

				for field := range util.DescIter(input.Fields()) {
					traverse(field, []string{getJsonFieldName(field)}, inputEntry.jsonMap, inputEntry.valueMap, inputEntry.pathMap)
					pm.fieldMap[field.FullName()] = inputEntry
				}

				for field := range util.DescIter(output.Fields()) {
					traverse(field, []string{getJsonFieldName(field)}, outputEntry.jsonMap, outputEntry.valueMap, outputEntry.pathMap)
					pm.fieldMap[field.FullName()] = outputEntry
				}

				pm.msgMap[input.FullName()] = inputEntry
				pm.msgMap[output.FullName()] = outputEntry
			}
		}
	}

	reqKeys := []string{}
	for k, _ := range pm.reqMap {
		reqKeys = append(reqKeys, k)
	}

	resKeys := []string{}
	for k := range pm.resMap {
		resKeys = append(resKeys, k)
	}

	log.Debug().Any("reqKeys", reqKeys).Any("resKeys", resKeys).Send()

	return pm, nil
}

func traverse(
	fd protoreflect.FieldDescriptor,
	currPath []string,
	currJsonMap map[string]*any,
	valueMap map[protoreflect.FullName]*any,
	pathMap map[protoreflect.FullName][]string,
) {
	value := zeroValueFor(fd)
	currJsonMap[getJsonFieldName(fd)] = &value
	valueMap[fd.FullName()] = &value
	pathMap[fd.FullName()] = currPath

	if msg := fd.Message(); msg != nil {
		for field := range util.DescIter(msg.Fields()) {
			traverse(field, append(currPath, getJsonFieldName(fd)), map[string]*any{}, valueMap, pathMap)
		}
	}
}

func zeroValueFor(fd protoreflect.FieldDescriptor) any {
	if fd.IsList() {
		return reflect.MakeSlice(
			reflect.SliceOf(reflect.TypeOf(scalarZeroValue(fd.Kind()))),
			0,
			0,
		).Interface()
	}

	if fd.IsMap() {
		key := fd.MapKey()
		value := fd.MapValue()

		return reflect.MakeMap(
			reflect.MapOf(
				reflect.TypeOf(scalarZeroValue(key.Kind())),
				reflect.TypeOf(scalarZeroValue(value.Kind())),
			),
		).Interface()
	}

	return scalarZeroValue(fd.Kind())
}

func scalarZeroValue(k protoreflect.Kind) any {
	switch k {
	case protoreflect.BoolKind:
		return false
	case protoreflect.BytesKind:
		return []byte{}
	case protoreflect.DoubleKind:
		return float64(0.0)
	case protoreflect.EnumKind:
		return protoreflect.EnumNumber(0)
	case protoreflect.Fixed32Kind:
		return uint32(0)
	case protoreflect.Fixed64Kind:
		return uint64(0)
	case protoreflect.FloatKind:
		return float32(0.0)
	case protoreflect.Int32Kind:
		return int32(0)
	case protoreflect.Int64Kind:
		return int64(0)
	case protoreflect.MessageKind:
		return map[string]*any{}
	case protoreflect.Sfixed32Kind:
		return int32(0)
	case protoreflect.Sfixed64Kind:
		return int64(0)
	case protoreflect.Sint32Kind:
		return int32(0)
	case protoreflect.Sint64Kind:
		return int64(0)
	case protoreflect.StringKind:
		return ""
	case protoreflect.Uint32Kind:
		return uint32(0)
	case protoreflect.Uint64Kind:
		return uint64(0)
	default:
		log.Fatal().Msgf("pathmap.scalarZeroValue: unhandled protoreflect.Kind %s", k.String())
		return nil
	}
}

const emptyTc = protoproxypb.TextCase_TEXT_CASE_UNSPECIFIED

func getJsonFieldName(fd protoreflect.FieldDescriptor) string {
	protoName := string(fd.Name())

	fieldOpts := opts.Get[*opts.Field](fd)

	if pp := fieldOpts.PP; pp != nil {
		if name := pp.Name; name != nil {
			return *name
		}

		if tc := pp.TextCase; tc != emptyTc {
			return renderTextCase(tc, protoName)
		}
	}

	// TODO! need to actually traverse up through all nested msgs to get to the root msg (method input or output)
	// to look for the lowest-scoped set TextCase before moving on to file.
	msg := fd.Parent().(protoreflect.MessageDescriptor)
	msgOpts := opts.Get[*opts.Message](msg)

	if pp := msgOpts.PP; pp != nil {
		if tc := pp.TextCase; tc != emptyTc {
			return renderTextCase(tc, protoName)
		}
	}

	file := fd.ParentFile()
	fileOpts := opts.Get[*opts.File](file)

	if pp := fileOpts.PP; pp != nil {
		if tc := pp.TextCase; tc != emptyTc {
			return renderTextCase(tc, protoName)
		}
	}

	return protoName
}

func renderTextCase(tc protoproxypb.TextCase, s string) string {
	switch tc {
	case protoproxypb.TextCase_CAMEL:
		return strcase.ToLowerCamel(s)
	case protoproxypb.TextCase_KEBAB:
		return strcase.ToKebab(s)
	case protoproxypb.TextCase_LOWER:
		return strings.ToLower(stripNonAlphanumeric(s))
	case protoproxypb.TextCase_PASCAL:
		return strcase.ToCamel(s)
	case protoproxypb.TextCase_SCREAMING_KEBAB:
		return strcase.ToScreamingKebab(s)
	case protoproxypb.TextCase_SCREAMING_SNAKE:
		return strcase.ToScreamingSnake(s)
	case protoproxypb.TextCase_SNAKE:
		return strcase.ToSnake(s)
	case protoproxypb.TextCase_UPPER:
		return strings.ToUpper(stripNonAlphanumeric(s))
	default:
		return s
	}
}

func stripNonAlphanumeric(s string) string {
	stripped := []rune{}
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' || c <= 'Z') || (c >= '0' && c <= '9') {
			stripped = append(stripped, c)
		}
	}

	return string(stripped)
}

func fqMethodName(method protoreflect.MethodDescriptor) string {
	svc := method.Parent().(protoreflect.ServiceDescriptor)

	return fmt.Sprintf("/%s/%s", string(svc.FullName()), string(method.Name()))
}
