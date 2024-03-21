package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Equals checks if two iterables are equal by comparing their elements.
// It returns true if the iterables have the same elements in the same order,
// and false otherwise.
func Equals[T comparable](a, b iter.Iterable[T]) bool {
	for {
		first, firstOk := a.Next()
		second, secondOk := b.Next()

		if firstOk != secondOk {
			return false
		}

		if first != second {
			return false
		}

		if !firstOk {
			return true
		}
	}
}
