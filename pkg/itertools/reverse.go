package itertools

import (
	"slices"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

type reverseIter[T any] struct {
	in  iter.Iterable[T]
	out iter.Iterable[T]
}

func (t *reverseIter[T]) Next() (T, bool) {
	if t.out == nil {
		out := Slice(t.in)
		slices.Reverse(out)
		t.out = iter.New(out...)
	}

	return t.out.Next()
}

// Reverse returns an iterable that iterates over the elements of the input iterable in reverse order.
// The input iterable can be of any type.
func Reverse[T any](iter iter.Iterable[T]) iter.Iterable[T] {
	return &reverseIter[T]{
		in:  iter,
		out: nil,
	}
}
