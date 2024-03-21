package predicates

import (
	"errors"

	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
)

// IsError returns a function that checks if the given error matches a specific error.
// It uses the errors.Is function to perform the comparison.
func IsError(this error) func(error) bool {
	return func(that error) bool {
		return errors.Is(that, this)
	}
}

// HasError returns a predicate function that checks if the given error matches the error in a tuple.
// It takes an error as input and returns a function that takes a tuple of type `containers.Tuple[T, error]`
// as input and returns a boolean value. The returned function checks if the error in the tuple matches the input
// error using the `errors.Is` function.
func HasError[T any](this error) func(tuple.Tuple[T, error]) bool {
	return func(that tuple.Tuple[T, error]) bool {
		return errors.Is(that.Second(), this)
	}
}
