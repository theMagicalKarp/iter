package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type tailIter[T any] struct {
	iter   iter.Iterable[T]
	out    iter.Iterable[T]
	amount int
}

func (t *tailIter[T]) Next() (T, bool) {
	if t.out != nil {
		return t.out.Next()
	}

	index := 0
	slice := make([]T, 0, t.amount)

	for item, more := t.iter.Next(); more; item, more = t.iter.Next() {
		if len(slice) < t.amount {
			slice = append(slice, item)
		} else {
			slice[index] = item
			index = (index + 1) % t.amount
		}
	}

	t.out = Chain(
		iter.New(slice[index:]...),
		iter.New(slice[:index]...),
	)

	return t.out.Next()
}

// Tail returns an iterable that yields the last `amount` elements from the given iterable.
// The `iter` parameter is the iterable to extract elements from.
// The `amount` parameter specifies the number of elements to extract.
// The returned iterable will yield the elements in the same order as the original iterable.
func Tail[T any](input iter.Iterable[T], amount int) iter.Iterable[T] {
	if amount <= 0 {
		return iter.New[T]()
	}

	return &tailIter[T]{
		iter:   input,
		amount: amount,
		out:    nil,
	}
}
