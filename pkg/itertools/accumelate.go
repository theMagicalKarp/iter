package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"golang.org/x/exp/constraints"
)

type accumulateIter[T constraints.Ordered] struct {
	iter  iter.Iterable[T]
	value T
}

func (a *accumulateIter[T]) Next() (T, bool) {
	value, more := a.iter.Next()

	if !more {
		return value, false
	}

	a.value += value

	return a.value, true
}

// Accumulate returns an iterable that yields accumulated values from the input iterable.
// The accumulated value at each position is the sum of all previous values in the input iterable.
// The input iterable must contain elements of a type that satisfies the constraints.Ordered interface.
func Accumulate[T constraints.Ordered](iter iter.Iterable[T]) iter.Iterable[T] {
	var value T

	return &accumulateIter[T]{
		iter:  iter,
		value: value,
	}
}
