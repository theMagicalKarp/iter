package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type zipIter[T any, V any] struct {
	first  iter.Iter[T]
	second iter.Iter[V]
}

func (z *zipIter[T, V]) Next() iter.Tuple[T, V] {
	if !z.HasNext() {
		var resp iter.Tuple[T, V]

		return resp
	}

	return iter.Tuple[T, V]{
		First:  z.first.Next(),
		Second: z.second.Next(),
	}
}

func (z *zipIter[T, V]) HasNext() bool {
	return z.first.HasNext() && z.second.HasNext()
}

func Zip[T any, V any](first iter.Iter[T], second iter.Iter[V]) iter.Iter[iter.Tuple[T, V]] {
	return &zipIter[T, V]{
		first:  first,
		second: second,
	}
}
