package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type cycleIter[T any] struct {
	slice []T
	iter  iter.Iterable[T]
	index int
}

func (c *cycleIter[T]) Next() (T, bool) {
	value, more := c.iter.Next()

	if !more {
		if len(c.slice) == 0 {
			return value, false
		}

		toReturn := c.slice[c.index]
		c.index = (c.index + 1) % len(c.slice)

		return toReturn, true
	}

	c.slice = append(c.slice, value)

	return value, more
}

// Cycle returns an iterable that cycles through the elements of the input iterable indefinitely.
// The input iterable must support iteration.
// The returned iterable will keep track of the elements it has already iterated through,
// allowing it to cycle through the elements without consuming additional memory.
func Cycle[T any](iter iter.Iterable[T]) iter.Iterable[T] {
	return &cycleIter[T]{
		iter:  iter,
		index: 0,
		slice: make([]T, 0),
	}
}
