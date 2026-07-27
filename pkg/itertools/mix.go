package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type mixIter[T any] struct {
	mix   []iter.Iterable[T]
	index int
}

func (m *mixIter[T]) Next() (T, bool) {
	for range len(m.mix) {
		value, more := m.mix[m.index].Next()
		m.index = (m.index + 1) % len(m.mix)

		if more {
			return value, true
		}
	}

	var value T

	return value, false
}

// Mix takes multiple iterables and returns a new iterable that iterates over the elements
// of each input iterable in a round-robin fashion.
// The returned iterable will continue until all input iterables are exhausted.
// The type parameter T represents the type of elements in the input iterables.
func Mix[T any](mix ...iter.Iterable[T]) iter.Iterable[T] {
	return &mixIter[T]{
		mix:   mix,
		index: 0,
	}
}
