package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type flattenIter[T any] struct {
	chain   iter.Iterable[iter.Iterable[T]]
	current iter.Iterable[T]
}

func (f *flattenIter[T]) Next() (T, bool) {
	if f.current != nil {
		value, more := f.current.Next()

		if more {
			return value, true
		}

		f.current = nil
	}

	for item, more := f.chain.Next(); more; item, more = f.chain.Next() {
		f.current = item
		value, more := f.current.Next()

		if more {
			return value, true
		}
	}

	var empty T

	return empty, false
}

func Flatten[T any](chain iter.Iterable[iter.Iterable[T]]) iter.Iterable[T] {
	return &flattenIter[T]{
		chain:   chain,
		current: nil,
	}
}
