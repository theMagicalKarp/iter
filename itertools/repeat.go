package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type repeatIter[T any] struct {
	value T
	limit int
	index int
}

func (r *repeatIter[T]) Next() T {
	if !r.HasNext() {
		var t T

		return t
	}

	r.index++

	return r.value
}

func (r *repeatIter[T]) HasNext() bool {
	return r.limit < 0 || r.index < r.limit
}

func Repeat[T any](value T, limit int) iter.Iter[T] {
	return &repeatIter[T]{
		value: value,
		limit: limit,
		index: 0,
	}
}
