package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type batchIter[T any] struct {
	iter iter.Iter[T]
	size int
}

func (b *batchIter[T]) Next() iter.Iter[T] {
	buffer := make([]T, b.size)
	i := 0

	for i = 0; i < b.size && b.iter.HasNext(); i++ {
		buffer[i] = b.iter.Next()
	}

	return iter.NewIter(buffer[:i]...)
}

func (b *batchIter[T]) HasNext() bool {
	return b.iter.HasNext()
}

func Batch[T any](iter iter.Iter[T], size int) iter.Iter[iter.Iter[T]] {
	return &batchIter[T]{
		iter: iter,
		size: size,
	}
}
