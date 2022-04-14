package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type tailIter[T any] struct {
	iter   iter.Iter[T]
	amount int
	list   *iter.List[T]
}

func (t *tailIter[T]) Next() T {
	if !t.HasNext() {
		var t T

		return t
	}

	return t.list.PopFront()
}

func (t *tailIter[T]) HasNext() bool {
	t.prime()

	return t.list.Front() != nil
}

func (t *tailIter[T]) prime() {
	for t.iter.HasNext() {
		t.list.PushBack(t.iter.Next())

		if t.amount <= 0 {
			t.list.PopFront()
		} else {
			t.amount--
		}
	}
}

func Tail[T any](changme iter.Iter[T], amount int) iter.Iter[T] {
	return &tailIter[T]{
		iter:   changme,
		amount: amount,
		list:   iter.NewList[T](),
	}
}
