package itertools

import (
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/predicates"
)

type filterIter[T any] struct {
	next *T
	iter iter.Iter[T]
	fn   predicates.Predicate[T]
}

func (f *filterIter[T]) Next() T {
	if !f.HasNext() {
		var t T

		return t
	}
	defer f.clear()

	return *f.next
}

func (f *filterIter[T]) HasNext() bool {
	f.prime()

	return f.next != nil
}

func (f *filterIter[T]) prime() {
	if f.next != nil {
		return
	}

	for f.iter.HasNext() {
		next := f.iter.Next()
		if f.fn(next) {
			f.next = &next

			return
		}
	}
}

func (f *filterIter[T]) clear() {
	f.next = nil
}

func Filter[T any](iter iter.Iter[T], fn predicates.Predicate[T]) iter.Iter[T] {
	return &filterIter[T]{
		next: nil,
		iter: iter,
		fn:   fn,
	}
}
