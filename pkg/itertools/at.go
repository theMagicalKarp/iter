package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// At returns the element at the specified index in the given iterable.
// It returns the element and a boolean value indicating whether the element was found.
// If the index is out of range or negative, it returns the zero value of the element type and false.
func At[T any](iter iter.Iterable[T], index int) (T, bool) {
	if index < 0 {
		var value T

		return value, false
	}

	for range index {
		_, more := iter.Next()
		if !more {
			break
		}
	}

	return iter.Next()
}
