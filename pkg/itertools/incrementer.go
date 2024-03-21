package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

type incrementIter struct {
	index int
	step  int
}

func (i *incrementIter) Next() (int, bool) {
	value := i.index
	i.index += i.step

	return value, true
}

// Incrementer returns an iterable that generates a sequence of integers starting from the given start value.
// Each subsequent value in the sequence is obtained by incrementing the previous value by 1.
func Incrementer(start int) iter.Iterable[int] {
	return &incrementIter{index: start, step: 1}
}

// Decrementer returns an iterable that generates a sequence of integers in descending order.
// The sequence starts from the given start value and decrements by 1 for each subsequent value.
func Decrementer(start int) iter.Iterable[int] {
	return &incrementIter{index: start, step: -1}
}
