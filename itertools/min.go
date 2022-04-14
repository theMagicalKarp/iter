package itertools

import (
	"github.com/theMagicalKarp/iter"
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](iter iter.Iter[T]) T {
	var min T

	if iter.HasNext() {
		min = iter.Next()
	}

	for iter.HasNext() {
		current := iter.Next()
		if current < min {
			min = current
		}
	}

	return min
}
