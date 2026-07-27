package itertools

import (
	"cmp"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Max returns the maximum element from the given iterable.
// The iterable must contain elements of a type that satisfies the Ordered constraint.
// If the iterable is empty, it returns the zero value of the element type.
func Max[T cmp.Ordered](iter iter.Iterable[T]) T {
	largest, more := iter.Next()

	for more {
		value, more := iter.Next()
		if !more {
			break
		}

		if value > largest {
			largest = value
		}
	}

	return largest
}
