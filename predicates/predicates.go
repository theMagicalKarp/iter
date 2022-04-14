package predicates

import (
	"strings"

	"golang.org/x/exp/constraints"
)

type Predicate[T any] func(t T) bool

func True[T any](t T) bool {
	return true
}

func False[T any](t T) bool {
	return false
}

func Positive[T constraints.Signed](t T) bool {
	return t > 0
}

func Negative[T constraints.Signed](t T) bool {
	return t < 0
}

func Even[T constraints.Integer](t T) bool {
	return t%2 == 0
}

func Odd[T constraints.Integer](t T) bool {
	return !Even(t)
}

func Zero[T constraints.Integer](t T) bool {
	return t == 0
}

func NotZero[T constraints.Integer](t T) bool {
	return !Zero(t)
}

func Divisible[T constraints.Integer](divisor T) Predicate[T] {
	return func(t T) bool {
		return t%divisor == 0
	}
}

func Empty(s string) bool {
	return s == ""
}

func NotEmpty(s string) bool {
	return !Empty(s)
}

func Contains(substr string) Predicate[string] {
	return func(s string) bool {
		return strings.Contains(s, substr)
	}
}

func HasPrefix(prefix string) Predicate[string] {
	return func(s string) bool {
		return strings.HasPrefix(s, prefix)
	}
}

func HasSuffix(suffix string) Predicate[string] {
	return func(s string) bool {
		return strings.HasSuffix(s, suffix)
	}
}
