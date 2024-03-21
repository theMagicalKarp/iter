package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func Partition[T any](input iter.Iterable[T], function predicates.Predicate[T]) (iter.Iterable[T], iter.Iterable[T]) {
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
