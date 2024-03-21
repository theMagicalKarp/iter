package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// // Number represents a numeric value that can be either an integer or a float.
// type Number interface {
// 	constraints.Integer | constraints.Float
// }

// Product returns the product of all elements in the given iterable.
// The iterable must contain elements of a numeric type.
// If the iterable is empty, the result is 1.
func Product[T Number](iter iter.Iterable[T]) T {
	sum, more := iter.Next()

	for more {
		value, more := iter.Next()
		if !more {
			break
		}

		sum *= value
	}

	return sum
}
