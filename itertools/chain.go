package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type chainIter[T any] struct {
	chain []iter.Iter[T]
	index int
}

func (c *chainIter[T]) Next() T {
	if !c.HasNext() {
		var t T

		return t
	}

	return c.chain[c.index].Next()
}

func (c *chainIter[T]) HasNext() bool {
	c.prime()

	return c.index < len(c.chain) && c.chain[c.index].HasNext()
}

func (c *chainIter[T]) prime() {
	for c.index < len(c.chain) {
		if c.chain[c.index].HasNext() {
			return
		}

		c.index++
	}
}

func Chain[T any](chain ...iter.Iter[T]) iter.Iter[T] {
	return &chainIter[T]{
		chain: chain,
		index: 0,
	}
}
