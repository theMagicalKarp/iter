package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type takeWhileIter[T any] struct {
	fn   func(item T) bool
	done bool
	tmp  *T
	iter iter.Iter[T]
}

func (t *takeWhileIter[T]) Next() T {
	if !t.HasNext() {
		var t T

		return t
	}

	next := *t.tmp
	t.tmp = nil

	return next
}

func (t *takeWhileIter[T]) HasNext() bool {
	if t.done {
		return false
	}

	if t.tmp != nil {
		return true
	}

	if !t.iter.HasNext() {
		t.done = true

		return false
	}

	next := t.iter.Next()

	if !t.fn(next) {
		t.done = true

		return false
	}

	t.tmp = &next

	return true
}

func TakeWhile[T any](iter iter.Iter[T], fn func(item T) bool) iter.Iter[T] {
	return &takeWhileIter[T]{
		fn:   fn,
		iter: iter,
	}
}
