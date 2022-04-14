package itertools

import (
	"github.com/theMagicalKarp/iter"
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](iter iter.Iter[T]) T {
	var max T

	if iter.HasNext() {
		max = iter.Next()
	}

	for iter.HasNext() {
		current := iter.Next()
		if current > max {
			max = current
		}
	}

	return max
}
