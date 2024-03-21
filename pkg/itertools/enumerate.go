package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Enumerate returns an iterable that pairs each element of the input iterable
// with its corresponding index. The index starts from 0.
func Enumerate[T any](items iter.Iterable[T]) iter.Iterable[tuple.Tuple[int, T]] {
	return Zip(Incrementer(0), items)
}
