package itertools

import (
	"github.com/theMagicalKarp/iter"
)

func At[T any](iter iter.Iter[T], index int) T {
	for i := 0; i < index; i++ {
		if !iter.HasNext() {
			break
		}

		iter.Next()
	}

	return iter.Next()
}
