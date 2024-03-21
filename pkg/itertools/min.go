package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"golang.org/x/exp/constraints"
)

// Min returns the minimum element from the given iterable.
// The iterable must contain elements of a type that satisfies the Ordered constraint.
// If the iterable is empty, the behavior is undefined.
func Min[T constraints.Ordered](iter iter.Iterable[T]) T {
	min, more := iter.Next()

	for more {
		value, more := iter.Next()
		if !more {
			break
		}

		if value < min {
			min = value
		}
	}

	return min
}
