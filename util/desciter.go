package util

import (
	"iter"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type IterableDesc[T protoreflect.Descriptor] interface {
	Len() int
	Get(int) T
}

func DescIter[T protoreflect.Descriptor](id IterableDesc[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range id.Len() {
			if !yield(id.Get(i)) {
				return
			}
		}
	}
}
