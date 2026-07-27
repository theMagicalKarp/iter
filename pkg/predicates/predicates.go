// Package predicates provides a set of predicate functions that can be used in conjunction of the itertools
// package to filter and transform iterables.
package predicates

import (
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
)

// Predicate is a function type that represents a predicate.
// It takes a value of type T and returns a boolean value indicating whether the value satisfies the predicate.
type Predicate[T any] func(t T) bool

// True returns true for any input value.
func True[T any](_ T) bool {
	return true
}

// False returns false for any input value.
func False[T any](_ T) bool {
	return false
}

// Not returns a new predicate that negates the result of the given predicate function.
// It takes a predicate function `f` as input and returns a new predicate function.
// The returned predicate function returns the logical negation of the result of `f`.
func Not[T any](f Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return !f(t)
	}
}

// Is returns a Predicate function that checks if the given value is equal to the original value.
// The type of the value must be comparable.
func Is[T comparable](this T) Predicate[T] {
	return func(that T) bool {
		return this == that
	}
}

// Has returns a predicate function that checks if the second element of a tuple is equal to the given value.
// The function takes a value of type V and returns a function that takes a tuple of type Tuple[T, V] and returns
// a boolean. The function compares the second element of the tuple with the given value and returns true if they are
// equal, false otherwise.
func Has[T any, V comparable](this V) func(tuple.Tuple[T, V]) bool {
	return func(that tuple.Tuple[T, V]) bool {
		return this == that.Second()
	}
}
