package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type chainIter[T any] struct {
	chain []iter.Iterable[T]
	index int
}

func (c *chainIter[T]) Next() (T, bool) {
	for c.index < len(c.chain) {
		value, more := c.chain[c.index].Next()

		if more {
			return value, true
		}

		c.index++
	}

	var value T

	return value, false
}

// Chain returns an iterable that combines multiple iterables into a single iterable.
// The elements from each iterable are returned in the order they appear.
// The type parameter T represents the type of elements in the iterables.
func Chain[T any](chain ...iter.Iterable[T]) iter.Iterable[T] {
	return &chainIter[T]{
		chain: chain,
		index: 0,
	}
}
