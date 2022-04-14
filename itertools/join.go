package itertools

import (
	"github.com/theMagicalKarp/iter"
)

type jointIter[T any] struct {
	iter      iter.Iter[T]
	separator T
	setSep    bool
}

func (j *jointIter[T]) Next() T {
	if !j.HasNext() {
		var t T

		return t
	}

	if j.setSep {
		j.setSep = false

		return j.separator
	}

	j.setSep = true

	return j.iter.Next()
}

func (j *jointIter[T]) HasNext() bool {
	return j.iter.HasNext()
}

func Join[T any](iter iter.Iter[T], separator T) iter.Iter[T] {
	return &jointIter[T]{
		iter:      iter,
		separator: separator,
	}
}
