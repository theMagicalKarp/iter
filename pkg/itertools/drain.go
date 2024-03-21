package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Drain consumes all elements from the given iterable until it is exhausted.
// It discards the values returned by the iterable's Next method.
func Drain[T any](iter iter.Iterable[T]) {
	more := true
	for more {
		_, more = iter.Next()
	}
}
