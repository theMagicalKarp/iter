package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

// Set returns a map containing unique elements from the given iterable.
// The elements in the iterable must be of a comparable type.
// The returned map contains each unique element as a key, with a value of true.
func Set[K comparable](items iter.Iterable[K]) map[K]bool {
	seen := make(map[K]bool)

	for item, more := items.Next(); more; item, more = items.Next() {
		seen[item] = true
	}

	return seen
}
