package itertools

import (
	"github.com/theMagicalKarp/iter"
)

func Each[T any](iter iter.Iter[T], fn func(item T)) {
	for iter.HasNext() {
		fn(iter.Next())
	}
}

func EachUntil[T any](iter iter.Iter[T], fn func(item T) bool) {
	for iter.HasNext() {
		if fn(iter.Next()) {
			return
		}
	}
}
