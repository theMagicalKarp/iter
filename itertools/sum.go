package itertools

import (
	"github.com/theMagicalKarp/iter"
	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Ordered](iter iter.Iter[T]) T {
	var total T

	for iter.HasNext() {
		total += iter.Next()
	}

	return total
}
