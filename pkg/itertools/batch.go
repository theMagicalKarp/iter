package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type batchIter[T any] struct {
	iter iter.Iterable[T]
	size int
}

func (b *batchIter[T]) Next() (iter.Iterable[T], bool) {
	buffer := make([]T, 0, b.size)
	index := 0

	for index < b.size {
		value, more := b.iter.Next()

		if !more {
			break
		}

		buffer = append(buffer, value)
		index++
	}

	if index == 0 {
		return nil, false
	}

	return iter.New(buffer...), true
}

// Batch returns an iterable that groups elements from the input iterable into batches of the specified size.
// Each batch is itself an iterable of elements from the input iterable.
// The last batch may contain fewer elements if the input iterable's length is not divisible by the batch size.
func Batch[T any](iter iter.Iterable[T], size int) iter.Iterable[iter.Iterable[T]] {
	return &batchIter[T]{
		iter: iter,
		size: size,
	}
}
