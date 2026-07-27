package itertools

import (
	"cmp"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Min returns the minimum element from the given iterable.
// The iterable must contain elements of a type that satisfies the Ordered constraint.
// If the iterable is empty, the behavior is undefined.
func Min[T cmp.Ordered](iter iter.Iterable[T]) T {
	smallest, more := iter.Next()

	for more {
		value, more := iter.Next()
		if !more {
			break
		}

		if value < smallest {
			smallest = value
		}
	}

	return smallest
}
