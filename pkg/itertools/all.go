package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

// All returns true if all elements in the iterable satisfy the given predicate function.
// It iterates over the elements of the iterable until it finds an element that does not satisfy the predicate,
// in which case it returns false. If all elements satisfy the predicate or the iterable is empty, it returns true.
// The function takes an iterable and a predicate function as arguments.
// The predicate function should take an element of the iterable as input and return a boolean value indicating
// whether the element satisfies the condition.
func All[T any](iter iter.Iterable[T], function predicates.Predicate[T]) bool {
	for {
		value, more := iter.Next()
		if !more {
			return true
		}

		if !function(value) {
			return false
		}
	}
}
