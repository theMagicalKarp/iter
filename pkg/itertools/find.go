package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

// Find returns the first element in the iterable that satisfies the given predicate function.
// If no element satisfies the predicate, it returns the zero value of the element type and false.
// The iterable must implement the `Next` method to retrieve the next element.
// The predicate function takes an element of the iterable as input and returns a
// boolean value indicating whether the element satisfies the condition.
func Find[T any](iter iter.Iterable[T], function predicates.Predicate[T]) (T, bool) {
	for {
		value, more := iter.Next()

		if !more {
			return value, false
		}

		if function(value) {
			return value, true
		}
	}
}
