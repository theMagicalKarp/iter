package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

type filterIter[T any] struct {
	next *T
	iter iter.Iterable[T]
	fn   predicates.Predicate[T]
}

func (f *filterIter[T]) Next() (T, bool) {
	for {
		value, more := f.iter.Next()
		if !more {
			return value, more
		}

		if f.fn(value) {
			return value, true
		}
	}
}

// Filter returns an iterable that only yields elements from the input iterable
// for which the given predicate function returns true.
// The predicate function `fn` is called with each element from the input iterable,
// and only elements for which the predicate returns true are included in the output iterable.
func Filter[T any](iter iter.Iterable[T], fn predicates.Predicate[T]) iter.Iterable[T] {
	return &filterIter[T]{
		next: nil,
		iter: iter,
		fn:   fn,
	}
}
