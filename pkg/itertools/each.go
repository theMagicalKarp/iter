package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Each applies a function to each item in the iterable.
// It iterates over the iterable and calls the provided function for each item.
// The function should take an item of type T as its argument.
// The iterable should implement the `iter.Iterable` interface.
// The function does not return any value.
func Each[T any](iter iter.Iterable[T], function func(item T)) {
	for {
		value, more := iter.Next()
		if !more {
			break
		}

		function(value)
	}
}

// EachUntil applies the given function to each item in the iterable until the function returns true.
// It stops iterating and returns immediately when the function returns true.
// The function takes an item of type T as input and returns a boolean value.
// If the iterable is empty, the function does nothing.
func EachUntil[T any](iter iter.Iterable[T], function func(item T) bool) {
	for {
		value, more := iter.Next()
		if !more {
			break
		}

		if function(value) {
			return
		}
	}
}
