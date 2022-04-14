package itertools

import (
	"github.com/theMagicalKarp/iter"
	"golang.org/x/exp/constraints"
)

type accumulateIter[T constraints.Ordered] struct {
	iter  iter.Iter[T]
	value T
}

func (a *accumulateIter[T]) Next() T {
	a.value += a.iter.Next()

	return a.value
}

func (a *accumulateIter[T]) HasNext() bool {
	return a.iter.HasNext()
}

func Accumulate[T constraints.Ordered](iter iter.Iter[T]) iter.Iter[T] {
	return &accumulateIter[T]{
		iter: iter,
	}
}
