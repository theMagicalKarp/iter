package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

// Union returns an iterable that contains all unique elements from the two input iterables.
// The elements are returned in the order they are encountered.
// The input iterables must be of the same type.
// The type of the elements must be comparable using the == operator.
func Union[K comparable](first, second iter.Iterable[K]) iter.Iterable[K] {
	return Unique(Chain(first, second))
}
