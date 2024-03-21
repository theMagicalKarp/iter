package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Empty returns an empty iterable of type T.
func Empty[T any]() iter.Iterable[T] {
	return iter.New[T]()
}
