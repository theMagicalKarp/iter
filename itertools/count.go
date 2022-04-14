package itertools

import (
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/predicates"
)

func Count[T any](iter iter.Iter[T], fn predicates.Predicate[T]) int {
	total := 0

	for iter.HasNext() {
		if fn(iter.Next()) {
			total++
		}
	}

	return total
}
