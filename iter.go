package iter

// Represents a stateful structure which can be used for iterating over sets
// of elements. Next should return the current element while advancing the
// iterator to the next element. HasNext should determine if there are any
// elements left to retrieve from the iterator.
type Iter[T any] interface {
	Next() T
	HasNext() bool
}

type iter[T any] struct {
	items []T
	index int
}

func (i *iter[T]) Next() T {
	if i.index >= len(i.items) {
		var t T

		return t
	}

	value := i.items[i.index]
	i.index++

	return value
}

func (i *iter[T]) HasNext() bool {
	return i.index < len(i.items)
}

// Creates a basic iterator which iterates over the set of elements provided
// in the parameters of the constructor.
func NewIter[T any](items ...T) Iter[T] {
	return &iter[T]{
		items: items,
		index: 0,
	}
}
