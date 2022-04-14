package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type cycleIter[T any] struct {
	iter    iter.Iter[T]
	current *iter.Node[T]
	start   *iter.Node[T]
	limit   int
	loop    int
}

func (c *cycleIter[T]) Next() T {
	if !c.HasNext() {
		var t T

		return t
	}

	if c.current.Next == nil && c.iter.HasNext() {
		c.current.Value = c.iter.Next()

		if !c.iter.HasNext() {
			c.current.Next = c.start
		} else {
			c.current.Next = &iter.Node[T]{}
		}
	}

	if c.current.Next == c.start {
		c.loop++
	}

	nextNode := c.current
	c.current = nextNode.Next

	return nextNode.Value
}

func (c *cycleIter[T]) HasNext() bool {
	if c.current.Next == nil && !c.iter.HasNext() {
		return false
	}

	return c.limit < 0 || c.loop < c.limit
}

func Cycle[T any](chamgeme iter.Iter[T], limit int) iter.Iter[T] {
	node := &iter.Node[T]{}

	return &cycleIter[T]{
		iter:    chamgeme,
		current: node,
		start:   node,
		limit:   limit,
		loop:    0,
	}
}
