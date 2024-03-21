package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"golang.org/x/exp/constraints"
)

// Max returns the maximum element from the given iterable.
// The iterable must contain elements of a type that satisfies the Ordered constraint.
// If the iterable is empty, it returns the zero value of the element type.
func Max[T constraints.Ordered](iter iter.Iterable[T]) T {
	min, more := iter.Next()

	for more {
		value, more := iter.Next()
		if !more {
			break
		}

		if value > min {
			min = value
		}
	}

	return min
}
