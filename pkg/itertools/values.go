package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

// Values returns an iterable containing the values of the given map.
// The keys of the map are ignored.
// The order of the values in the returned iterable is not guaranteed to be the same as the original map.
func Values[K comparable, V any](items map[K]V) iter.Iterable[V] {
	toReturn := make([]V, 0)

	for _, item := range items {
		toReturn = append(toReturn, item)
	}

	return iter.New(toReturn...)
}
