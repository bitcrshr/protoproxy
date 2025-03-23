package pathmap

import (
	_ "embed"
	"os"
	"testing"

	"github.com/bitcrshr/protoproxy/util"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	//go:embed test.binpb
	rawTestDescriptorSet []byte

	files *protoregistry.Files

	allFields = map[string]protoreflect.FieldDescriptor{}
)

func TestMain(m *testing.M) {
	descriptorSet := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(rawTestDescriptorSet, descriptorSet); err != nil {
		log.Fatal().Err(err).Msg("failed to unmarshal test file descriptor set")
	}

	var err error
	if files, err = protodesc.NewFiles(descriptorSet); err != nil {
		log.Fatal().Err(err).Msg("failed to parse test file descriptor set")
	}

	for file := range files.RangeFiles {
		for msg := range util.DescIter(file.Messages()) {
			for field := range util.DescIter(msg.Fields()) {
				allFields[string(field.FullName())] = field
			}
		}
	}

	code := m.Run()
	os.Exit(code)
}

func TestZeroValueFor(t *testing.T) {
	type testCase struct {
		desc     protoreflect.FieldDescriptor
		expected any
	}

	tcs := []testCase{
		{
			desc:     allFields["pathmap_test.TestMessage.double"],
			expected: float64(0.0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.doubles"],
			expected: []float64{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.float"],
			expected: float32(0.0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.floats"],
			expected: []float32{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.i32"],
			expected: int32(0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.i32s"],
			expected: []int32{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.u32"],
			expected: uint32(0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.u32s"],
			expected: []uint32{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.s32"],
			expected: int32(0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.s32s"],
			expected: []int32{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.s64"],
			expected: int64(0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.s64s"],
			expected: []int64{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.fixed32"],
			expected: uint32(0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.fixed32s"],
			expected: []uint32{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.fixed64"],
			expected: uint64(0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.fixed64s"],
			expected: []uint64{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.sf32"],
			expected: int32(0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.sf32s"],
			expected: []int32{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.bool"],
			expected: false,
		},

		{
			desc:     allFields["pathmap_test.TestMessage.bools"],
			expected: []bool{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.string"],
			expected: "",
		},

		{
			desc:     allFields["pathmap_test.TestMessage.strings"],
			expected: []string{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.bytes"],
			expected: []byte{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.bytesbytes"],
			expected: [][]byte{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.message"],
			expected: map[string]*any{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.messages"],
			expected: []map[string]*any{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.primitive_map"],
			expected: map[string]bool{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.enum"],
			expected: protoreflect.EnumNumber(0),
		},

		{
			desc:     allFields["pathmap_test.TestMessage.enums"],
			expected: []protoreflect.EnumNumber{},
		},

		{
			desc:     allFields["pathmap_test.TestMessage.msg_map"],
			expected: map[string]map[string]*any{},
		},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.expected, zeroValueFor(tc.desc))
	}
}

func TestTraverse(t *testing.T) {
	fd := allFields["pathmap_test.TestMessage.message"]

	jsonMap := make(map[string]*any)
	valueMap := make(map[protoreflect.FullName]*any)
	jsonPathMap := make(map[protoreflect.FullName][]string)

	traverse(fd, []string{string(fd.Name())}, jsonMap, valueMap, jsonPathMap)

	value := valueMap[fd.FullName()]

	expected := map[string]any{
		"foo": []string{"bar", "baz"},
	}

	*value = expected

	log.Debug().Any("jsonMap", jsonMap).Any("valueMap", valueMap).Any("pathMap", jsonPathMap).Send()

	assert.Equal(t, expected, *jsonMap["message"])
}
