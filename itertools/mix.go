package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type mixIter[T any] struct {
	next  *T
	iters []iter.Iter[T]
	index int
	done  bool
}

func (m *mixIter[T]) Next() T {
	if !m.HasNext() {
		var t T

		return t
	}

	defer m.clear()

	return *m.next
}

func (m *mixIter[T]) HasNext() bool {
	if m.done {
		return false
	}

	if m.next != nil {
		return true
	}

	for i := 0; i < len(m.iters); i++ {
		m.index = (m.index + 1) % len(m.iters)
		if m.iters[m.index].HasNext() {
			next := m.iters[m.index].Next()
			m.next = &next

			return true
		}
	}

	m.done = true

	return false
}

func (m *mixIter[T]) clear() {
	m.next = nil
}

func Mix[T any](iters ...iter.Iter[T]) iter.Iter[T] {
	return &mixIter[T]{
		iters: iters,
		index: -1,
	}
}
