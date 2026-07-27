package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type jointIter[T any] struct {
	iter      iter.Iterable[T]
	separator T
	last      *T
	firstPull bool
}

func (j *jointIter[T]) Next() (T, bool) {
	if j.last != nil {
		value := *j.last
		j.last = nil

		return value, true
	}

	value, more := j.iter.Next()

	if j.firstPull {
		j.firstPull = false

		return value, more
	}

	if more {
		j.last = &value

		return j.separator, true
	}

	return value, more
}

// Join concatenates the elements of an iterable into a single iterable,
// separating each element with a specified separator.
// The separator can be of any type.
// The returned iterable will yield the elements of the original iterable,
// with the separator inserted between each pair of elements.
func Join[T any](iter iter.Iterable[T], separator T) iter.Iterable[T] {
	return &jointIter[T]{
		iter:      iter,
		separator: separator,
		last:      nil,
		firstPull: true,
	}
}
