package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type dropIter[T any] struct {
	count int
	index int
	iter  iter.Iter[T]
}

func (d *dropIter[T]) Next() T {
	d.prime()

	return d.iter.Next()
}

func (d *dropIter[T]) HasNext() bool {
	d.prime()

	return d.iter.HasNext()
}

func (d *dropIter[T]) prime() {
	for d.index < d.count {
		d.index++

		if !d.iter.HasNext() {
			break
		}

		d.iter.Next()
	}
}

func Drop[T any](iter iter.Iter[T], count int) iter.Iter[T] {
	return &dropIter[T]{
		count: count,
		iter:  iter,
	}
}
