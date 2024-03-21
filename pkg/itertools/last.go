package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

// Last returns the last element of the given iterable and a boolean value indicating if an element was found.
// If the iterable is empty, the returned boolean value will be false.
func Last[T any](iter iter.Iterable[T]) (T, bool) {
	var (
		last  T
		found bool
	)

	for {
		value, more := iter.Next()

		if !more {
			return last, found
		}

		last = value
		found = true
	}
}

// LastIf returns the last element in the iterable that satisfies the given predicate function.
// It iterates over the elements of the iterable and returns the last element that satisfies the predicate,
// along with a boolean value indicating whether such an element was found.
// If no element satisfies the predicate, the zero value of the element type is returned along with false.
func LastIf[T any](iter iter.Iterable[T], function predicates.Predicate[T]) (T, bool) {
	var (
		last  T
		found bool
	)

	for {
		value, more := iter.Next()

		if !more {
			return last, found
		}

		if function(value) {
			last = value
			found = true
		}
	}
}
