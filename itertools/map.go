package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type mapIter[T any, V any] struct {
	iter iter.Iter[T]
	fn   func(item T) V
}

func (m *mapIter[T, V]) Next() V {
	return m.fn(m.iter.Next())
}

func (m *mapIter[T, V]) HasNext() bool {
	return m.iter.HasNext()
}

func Map[T any, V any](iter iter.Iter[T], fn func(item T) V) iter.Iter[V] {
	return &mapIter[T, V]{
		iter: iter,
		fn:   fn,
	}
}
