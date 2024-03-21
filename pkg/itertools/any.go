package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

// Any returns true if at least one element in the iterable satisfies the given predicate function.
// It iterates over the elements of the iterable until it finds an element that satisfies the predicate,
// and then it returns true. If no such element is found, it returns false.
// The predicate function takes an element of type T as input and returns a boolean value.
// The iterable must implement the Iterable interface, and the predicate function must implement the Predicate
// interface.
func Any[T any](iter iter.Iterable[T], function predicates.Predicate[T]) bool {
	for {
		value, more := iter.Next()
		if !more {
			return false
		}

		if function(value) {
			return true
		}
	}
}
