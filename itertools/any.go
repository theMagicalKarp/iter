package itertools

import (
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/predicates"
)

func Any[T any](iter iter.Iter[T], fn predicates.Predicate[T]) bool {
	for iter.HasNext() {
		if fn(iter.Next()) {
			return true
		}
	}

	return false
}
