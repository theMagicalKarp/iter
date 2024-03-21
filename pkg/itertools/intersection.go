package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

type intersectionIter[K comparable] struct {
	seen   map[K]bool
	found  map[K]bool
	first  iter.Iterable[K]
	second iter.Iterable[K]
}

func (s *intersectionIter[K]) Next() (K, bool) {
	if s.seen == nil {
		s.seen = make(map[K]bool)
		for item, more := s.first.Next(); more; item, more = s.first.Next() {
			s.seen[item] = true
		}
	}

	for item, more := s.second.Next(); more; item, more = s.second.Next() {
		if s.seen[item] && !s.found[item] {
			s.found[item] = true

			return item, true
		}
	}

	var empty K

	return empty, false
}

// Intersection returns an iterable that represents the intersection of two iterables.
// The returned iterable contains only the elements that are present in both input iterables.
// The elements are compared using the `==` operator.
// The input iterables must be of the same type.
// The order of elements in the returned iterable is not guaranteed.
func Intersection[K comparable](first, second iter.Iterable[K]) iter.Iterable[K] {
	return &intersectionIter[K]{
		first:  first,
		second: second,
		found:  make(map[K]bool),
		seen:   nil,
	}
}
