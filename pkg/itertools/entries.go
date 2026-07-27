package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Entries returns an iterable of tuples containing the key-value pairs from the given map.
// The keys and values can be of any type, as long as the keys are comparable.
// The returned iterable can be used to iterate over the key-value pairs in a consistent order.
func Entries[K comparable, V any](items map[K]V) iter.Iterable[tuple.Tuple[K, V]] {
	toReturn := make([]tuple.Tuple[K, V], 0, len(items))

	for key, value := range items {
		toReturn = append(toReturn, tuple.New(key, value))
	}

	return iter.New(toReturn...)
}
