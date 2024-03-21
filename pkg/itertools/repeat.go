package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type repeatIter[T any] struct {
	value T
}

func (r *repeatIter[T]) Next() (T, bool) {
	return r.value, true
}

// Repeat returns an iterable that repeatedly yields the specified value.
// The iterable will continue indefinitely, producing the same value each time.
func Repeat[T any](value T) iter.Iterable[T] {
	return &repeatIter[T]{value: value}
}
