package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

type uniqueIterable[K comparable] struct {
	seen map[K]bool
	iter iter.Iterable[K]
}

func (u *uniqueIterable[K]) Next() (K, bool) {
	item, more := u.iter.Next()
	if !more {
		return item, more
	}

	for more {
		if !u.seen[item] {
			u.seen[item] = true

			return item, true
		}

		item, more = u.iter.Next()
	}

	var empty K

	return empty, false
}

// Unique returns an iterable that contains only the unique elements from the input iterable.
// The input iterable can contain elements of any type that is comparable.
// The returned iterable will preserve the order of the unique elements.
func Unique[K comparable](input iter.Iterable[K]) iter.Iterable[K] {
	return &uniqueIterable[K]{
		seen: make(map[K]bool),
		iter: input,
	}
}
