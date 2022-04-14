package itertools

import (
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/predicates"
)

func Find[T any](iter iter.Iter[T], fn predicates.Predicate[T]) (resp T) {
	for iter.HasNext() {
		next := iter.Next()
		if fn(next) {
			resp = next

			return
		}
	}

	return
}
