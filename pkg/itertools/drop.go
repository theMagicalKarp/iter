package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type dropIter[T any] struct {
	count int
	index int
	iter  iter.Iterable[T]
}

func (d *dropIter[T]) Next() (T, bool) {
	var (
		value T
		more  bool
	)

	for d.index < d.count {
		value, more = d.iter.Next()
		if !more {
			return value, false
		}

		d.index++
	}

	return d.iter.Next()
}

// Drop returns an iterable that skips the first `count` elements from the input iterable.
// The returned iterable will yield the remaining elements from the input iterable.
// If `count` is greater than the length of the input iterable, an empty iterable is returned.
func Drop[T any](iter iter.Iterable[T], count int) iter.Iterable[T] {
	return &dropIter[T]{
		count: count,
		index: 0,
		iter:  iter,
	}
}
