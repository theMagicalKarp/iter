package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type rangeIter struct {
	index int
	stop  int
	inc   int
}

func (r *rangeIter) Next() int {
	if !r.HasNext() {
		return 0
	}

	resp := r.index
	r.index += r.inc

	return resp
}

func (r *rangeIter) HasNext() bool {
	return r.index < r.stop
}

func Range(start, stop int) iter.Iter[int] {
	return &rangeIter{
		index: start,
		stop:  stop,
		inc:   1,
	}
}

func RangeInc(start, stop, inc int) iter.Iter[int] {
	return &rangeIter{
		index: start,
		stop:  stop,
		inc:   inc,
	}
}
