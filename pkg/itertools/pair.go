package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type pairIter[T any] struct {
	items iter.Iterable[T]
}

func (p *pairIter[T]) Next() (tuple.Tuple[T, T], bool) {
	var empty T

	return tuple.New(empty, empty), false
}

// Pair groups the elements of an iterable into two element tuples.
//
// Pair is not implemented yet. The returned iterable is always empty,
// regardless of the input.
func Pair[T any](items iter.Iterable[T]) iter.Iterable[tuple.Tuple[T, T]] {
	return &pairIter[T]{
		items: items,
	}
}
