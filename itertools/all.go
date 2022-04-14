package itertools

import (
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/predicates"
)

func All[T any](iter iter.Iter[T], fn predicates.Predicate[T]) bool {
	for iter.HasNext() {
		if !fn(iter.Next()) {
			return false
		}
	}

	return true
}
