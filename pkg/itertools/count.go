package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

// CountIf counts the number of elements in the iterable that satisfy the given predicate function.
// It iterates over the elements of the iterable and applies the predicate function to each element.
// If the predicate function returns true for an element, it increments the count.
// Finally, it returns the total count of elements that satisfy the predicate.
func CountIf[T any](iter iter.Iterable[T], function predicates.Predicate[T]) int {
	total := 0

	for {
		value, more := iter.Next()
		if !more {
			break
		}

		if function(value) {
			total++
		}
	}

	return total
}

// Count returns the total number of elements in the given iterable.
// It iterates over the iterable until there are no more elements and counts each element encountered.
func Count[T any](iter iter.Iterable[T]) int {
	total := 0

	for {
		_, more := iter.Next()
		if !more {
			break
		}

		total++
	}

	return total
}
