package itertools

import (
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/predicates"
)

func Last[T any](iter iter.Iter[T], fn predicates.Predicate[T]) T {
	var toReturn T

	for iter.HasNext() {
		next := iter.Next()
		if fn(next) {
			toReturn = next
		}
	}

	return toReturn
}
