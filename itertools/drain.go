package itertools

import (
	"github.com/theMagicalKarp/iter"
)

func Drain[T any](iter iter.Iter[T]) {
	for iter.HasNext() {
		iter.Next()
	}
}
