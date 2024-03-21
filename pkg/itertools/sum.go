package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"golang.org/x/exp/constraints"
)

// Sum returns the sum of all elements in the given iterable.
// The elements must be of a type that satisfies the Ordered constraint.
func Sum[T constraints.Ordered](iter iter.Iterable[T]) T {
	sum, more := iter.Next()

	for more {
		value, more := iter.Next()
		if !more {
			break
		}

		sum += value
	}

	return sum
}
