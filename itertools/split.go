package itertools

import (
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/predicates"
)

func Split[T any](in iter.Iter[T], fn predicates.Predicate[T]) (iter.Iter[T], iter.Iter[T]) {
	trueIter := iter.NewList[T]()
	falseIter := iter.NewList[T]()

	for in.HasNext() {
		next := in.Next()
		if fn(next) {
			trueIter.PushBack(next)
		} else {
			falseIter.PushBack(next)
		}
	}

	return trueIter.Iter(), falseIter.Iter()
}
