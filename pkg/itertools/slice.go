package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Slice returns a new slice containing all the elements from the given iterable.
// The elements are retrieved by calling the Next method on the iterable until it returns false.
// The type parameter T specifies the type of elements in the iterable.
func Slice[T any](items iter.Iterable[T]) []T {
	toReturn := make([]T, 0)

	for item, more := items.Next(); more; item, more = items.Next() {
		toReturn = append(toReturn, item)
	}

	return toReturn
}

// SliceN returns a slice of elements from the given iterable, starting from
// the specified start index and ending at the specified stop index.
// The function iterates through the iterable until it reaches the start index,
// and then appends each element to the returned slice until it reaches the stop index.
// If the iterable is exhausted before reaching the stop index, the function
// returns the slice with the available elements.
// The type parameter T represents the type of elements in the iterable.
func SliceN[T any](items iter.Iterable[T], start, stop int) []T {
	toReturn := make([]T, 0)

	var index int
	for index = 0; index < start; index++ {
		_, more := items.Next()
		if !more {
			return toReturn
		}
	}

	for item, more := items.Next(); more; item, more = items.Next() {
		if index >= stop {
			break
		}

		toReturn = append(toReturn, item)
		index++
	}

	return toReturn
}
