package itertools

import (
	"fmt"

	"github.com/theMagicalKarp/iter"
)

func Print[T any](iter iter.Iter[T]) {
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
