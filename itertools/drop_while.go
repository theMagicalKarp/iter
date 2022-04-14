package itertools

import (
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/predicates"
)

type dropWhileIter[T any] struct {
	fn   predicates.Predicate[T]
	done bool
	tmp  *T
	iter iter.Iter[T]
}

func (d *dropWhileIter[T]) Next() T {
	d.prime()

	if d.tmp != nil {
		next := *d.tmp
		d.tmp = nil

		return next
	}

	return d.iter.Next()
}

func (d *dropWhileIter[T]) HasNext() bool {
	d.prime()

	return d.tmp != nil || d.iter.HasNext()
}

func (d *dropWhileIter[T]) prime() {
	if !d.done {
		for d.iter.HasNext() {
			next := d.iter.Next()
			if !d.fn(next) {
				d.tmp = &next

				break
			}
		}

		d.done = true
	}
}

func DropWhile[T any](iter iter.Iter[T], fn predicates.Predicate[T]) iter.Iter[T] {
	return &dropWhileIter[T]{
		fn:   fn,
		iter: iter,
	}
}
