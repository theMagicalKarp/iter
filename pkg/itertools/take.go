package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type takeIter[T any] struct {
	iter   iter.Iterable[T]
	amount int
	index  int
}

func (t *takeIter[T]) Next() (T, bool) {
	if t.index >= t.amount {
		var value T

		return value, false
	}

	value, more := t.iter.Next()
	if !more {
		var value T

		return value, false
	}

	t.index++

	return value, true
}

// Take returns a new iterable that yields the first `amount` elements from the given iterable.
// If the given iterable has fewer than `amount` elements, all elements are returned.
// The original iterable is not modified.
func Take[T any](iter iter.Iterable[T], amount int) iter.Iterable[T] {
	return &takeIter[T]{
		iter:   iter,
		amount: amount,
		index:  0,
	}
}
