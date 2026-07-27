package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type zipIter[T any, V any] struct {
	first  iter.Iterable[T]
	second iter.Iterable[V]
}

func (z *zipIter[T, V]) Next() (tuple.Tuple[T, V], bool) {
	first, moreFirst := z.first.Next()
	second, moreSecond := z.second.Next()

	if !moreFirst || !moreSecond {
		var resp tuple.Tuple[T, V]

		return resp, false
	}

	return tuple.New[T, V](first, second), true
}

// Zip returns an iterable that combines elements from two iterables into tuples.
// The resulting iterable yields tuples where the i-th tuple contains the i-th element from the first iterable
// and the i-th element from the second iterable.
// The iterator stops when the shortest input iterable is exhausted.
func Zip[T any, V any](
	first iter.Iterable[T],
	second iter.Iterable[V],
) iter.Iterable[tuple.Tuple[T, V]] {
	return &zipIter[T, V]{
		first:  first,
		second: second,
	}
}
