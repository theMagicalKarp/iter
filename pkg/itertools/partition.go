package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

// Partition splits an iterable in two by applying the given predicate to every
// element. The first returned iterable holds the elements the predicate
// accepted, the second the elements it rejected, each in their original order.
// The input iterable is fully drained before either result is returned.
func Partition[T any](
	input iter.Iterable[T],
	function predicates.Predicate[T],
) (iter.Iterable[T], iter.Iterable[T]) {
	trueIter := make([]T, 0)
	falseIter := make([]T, 0)

	value, more := input.Next()

	for more {
		if function(value) {
			trueIter = append(trueIter, value)
		} else {
			falseIter = append(falseIter, value)
		}

		value, more = input.Next()
	}

	return iter.New(trueIter...), iter.New(falseIter...)
}
