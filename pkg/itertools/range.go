package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

type rangeIter struct {
	index int
	stop  int
	inc   int
}

func (r *rangeIter) Next() (int, bool) {
	if r.inc > 0 && r.index >= r.stop {
		return 0, false
	}

	if r.inc < 0 && r.index <= r.stop {
		return 0, false
	}

	resp := r.index
	r.index += r.inc

	return resp, true
}

// Range returns an iterable that generates a sequence of integers from start to stop (exclusive).
// The start parameter specifies the starting value of the sequence.
// The stop parameter specifies the ending value of the sequence.
// The returned iterable can be used in a for loop or with other iterator functions.
func Range(start, stop int) iter.Iterable[int] {
	return &rangeIter{
		index: start,
		stop:  stop,
		inc:   1,
	}
}

// RangeInc returns an iterable that generates a sequence of integers starting from 'start' up to 'stop' (inclusive),
// with each integer incremented by 'inc' amount. The 'inc' parameter can be negative to generate a decreasing sequence.
func RangeInc(start, stop, inc int) iter.Iterable[int] {
	return &rangeIter{
		index: start,
		stop:  stop,
		inc:   inc,
	}
}
