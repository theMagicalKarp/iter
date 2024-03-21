package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

type subtractIterable[K comparable] struct {
	seen   map[K]bool
	first  iter.Iterable[K]
	second iter.Iterable[K]
}

func (s *subtractIterable[K]) Next() (K, bool) {
	if s.seen == nil {
		s.seen = make(map[K]bool)
		for item, more := s.second.Next(); more; item, more = s.second.Next() {
			s.seen[item] = true
		}
	}

	for item, more := s.first.Next(); more; item, more = s.first.Next() {
		if s.seen[item] {
			continue
		}

		return item, true
	}

	var empty K

	return empty, false
}

// Subtract returns an iterable that contains the elements from the first iterable
// that are not present in the second iterable.
// The type of elements in the iterables must be comparable.
func Subtract[K comparable](first, second iter.Iterable[K]) iter.Iterable[K] {
	return &subtractIterable[K]{
		first:  first,
		second: second,
		seen:   nil,
	}
}
