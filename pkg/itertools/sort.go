package itertools

import (
	"slices"

	"github.com/theMagicalKarp/iter/pkg/iter"
	"golang.org/x/exp/constraints"
)

// Number is an interface that represents a numeric value.
// It can be either an integer or a float.
type Number interface {
	constraints.Integer | constraints.Float
}

// Compare is a function type that compares two values of type T and returns an integer.
// It should return a negative value if a < b, 0 if a == b, and a positive value if a > b.
type Compare[T any] func(a, b T) int

// Asc returns the difference between `a` and `b`.
// It is used to define the ascending order for sorting.
func Asc[T Number](this, that T) T {
	return this - that
}

// Desc returns the difference between the second argument and the first argument.
// It is used as a comparison function for sorting in descending order.
func Desc[T Number](this, that T) T {
	return that - this
}

type sortIter[T any] struct {
	iter iter.Iterable[T]
	out  iter.Iterable[T]

	cmp Compare[T]
}

func (s *sortIter[T]) Next() (T, bool) {
	if s.out == nil {
		slice := Slice(s.iter)
		slices.SortFunc(slice, s.cmp)
		s.out = iter.New(slice...)
	}

	return s.out.Next()
}

// Sort returns an iterable that produces the elements of the input iterable in sorted order.
// The sorting is performed using the provided comparison function cmp.
// The input iterable must support iteration and the comparison function must return an integer
// less than, equal to, or greater than zero if the first argument is considered to be respectively
// less than, equal to, or greater than the second argument.
func Sort[T any](iter iter.Iterable[T], cmp Compare[T]) iter.Iterable[T] {
	return &sortIter[T]{
		iter: iter,
		out:  nil,
		cmp:  cmp,
	}
}
