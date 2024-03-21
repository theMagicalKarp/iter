package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

// Keys returns an iterable containing all the keys from the given map.
// The keys are returned in an arbitrary order.
// The type parameters K and V specify the key and value types of the map.
func Keys[K comparable, V any](items map[K]V) iter.Iterable[K] {
	toReturn := make([]K, 0)

	for key := range items {
		toReturn = append(toReturn, key)
	}

	return iter.New(toReturn...)
}
