package itertools

import (
	"sync"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

type teeParent[T any] struct {
	items iter.Iterable[T]
	slice []T
	mutex *sync.Mutex
}

type teeIter[T any] struct {
	parent *teeParent[T]
	cursor int
}

func (t *teeIter[T]) Next() (T, bool) {
	t.parent.mutex.Lock()
	defer t.parent.mutex.Unlock()

	if t.cursor >= len(t.parent.slice) {
		value, more := t.parent.items.Next()
		if !more {
			return value, false
		}

		t.parent.slice = append(t.parent.slice, value)
		t.cursor++

		return value, true
	}

	value := t.parent.slice[t.cursor]
	t.cursor++

	return value, true
}

// Tee splits an iterable into two separate iterables.
// Both iterables will produce the same elements as the original iterable.
// The elements are buffered, so they can be consumed independently.
// Any changes made to one iterable will not affect the other.
func Tee[T any](iter iter.Iterable[T]) (iter.Iterable[T], iter.Iterable[T]) {
	var mutex sync.Mutex

	slice := make([]T, 0)

	parent := &teeParent[T]{
		items: iter,
		slice: slice,
		mutex: &mutex,
	}

	first := &teeIter[T]{
		cursor: 0,
		parent: parent,
	}

	second := &teeIter[T]{
		cursor: 0,
		parent: parent,
	}

	return first, second
}
