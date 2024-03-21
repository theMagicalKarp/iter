package predicates

import "golang.org/x/exp/constraints"

// Positive checks if the given value is positive.
// It returns true if the value is greater than zero, otherwise false.
func Positive[T constraints.Signed](t T) bool {
	return t > 0
}

// Negative returns true if the given value is negative, otherwise false.
func Negative[T constraints.Signed](t T) bool {
	return t < 0
}

// Even checks if the given integer is even.
// It returns true if the integer is even, and false otherwise.
func Even[T constraints.Integer](t T) bool {
	return t%2 == 0
}

// Odd returns true if the given integer is odd, false otherwise.
func Odd[T constraints.Integer](t T) bool {
	return t%2 != 0
}

// Divisible returns a predicate function that checks if a given number is divisible by the specified divisor.
// The divisor must be an integer type.
func Divisible[T constraints.Integer](divisor T) Predicate[T] {
	return func(t T) bool {
		return t%divisor == 0
	}
}
