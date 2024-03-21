package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// MapFn is a function type that defines a mapping operation.
// It takes an input item of type T and returns a transformed value of type V.
type MapFn[T any, V any] func(item T) V

type mapIter[T any, V any] struct {
	iter     iter.Iterable[T]
	function MapFn[T, V]
}

func (m *mapIter[T, V]) Next() (V, bool) {
	input, more := m.iter.Next()

	if !more {
		var value V

		return value, false
	}

	return m.function(input), true
}

// Map applies a function to each item in the input iterable and returns an iterable of the results.
// The input iterable can be of any type, and the function can transform each item to a different type.
// The returned iterable will yield the transformed values.
func Map[T any, V any](iter iter.Iterable[T], function MapFn[T, V]) iter.Iterable[V] {
	return &mapIter[T, V]{
		iter:     iter,
		function: function,
	}
}
