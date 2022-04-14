package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type takeIter[T any] struct {
	iter   iter.Iter[T]
	amount int
	index  int
}

func (t *takeIter[T]) Next() T {
	if !t.HasNext() {
		var t T

		return t
	}

	t.index++

	return t.iter.Next()
}

func (t *takeIter[T]) HasNext() bool {
	return t.index < t.amount && t.iter.HasNext()
}

func Take[T any](iter iter.Iter[T], amount int) iter.Iter[T] {
	return &takeIter[T]{
		iter:   iter,
		amount: amount,
		index:  0,
	}
}
