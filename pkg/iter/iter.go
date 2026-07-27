// Package iter provides an interface for efficiently working with streams of data
// across different data structures and algorithms, enabling a more ergonomic coding.
package iter

// Iterable represents a generic iterable collection.
type Iterable[T any] interface {
	// Next returns the next element in the collection and a boolean value indicating if the value was found.
	Next() (T, bool)
}

type iter[T any] struct {
	items []T
	index int
}

// Next returns the next item in the iterator and a boolean value indicating
// whether the next value was found.
func (i *iter[T]) Next() (T, bool) {
	if i.index >= len(i.items) {
		var t T

		return t, false
	}

	value := i.items[i.index]
	i.index++

	return value, true
}

// New creates a new iterable from the given items.
// It takes a variadic parameter `items` of type `T` and returns an `Iterable` of type `T`.
// The `Iterable` allows iterating over the items using the `Next` method.
func New[T any](items ...T) Iterable[T] {
	return &iter[T]{
		items: items,
		index: 0,
	}
}
