package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// StartsWith checks if the given iterable starts with the specified prefix.
// It compares the elements of the iterable and the prefix in order, and returns true if they match.
// If the prefix is longer than the iterable, it returns false.
// The elements are compared using the equality operator (==) for the given type T.
// The iterable and the prefix must both implement the Iterable interface.
// The type T must satisfy the Ordered constraint, meaning it must be comparable using the comparison
// operators (<, >, <=, >=).
func StartsWith[T comparable](iter, prefix iter.Iterable[T]) bool {
	for {
		current, currentOk := iter.Next()
		pre, preOk := prefix.Next()

		if !preOk {
			return true
		}

		if currentOk != preOk {
			return false
		}

		if current != pre {
			return false
		}
	}
}
