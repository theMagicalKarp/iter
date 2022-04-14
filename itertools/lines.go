package itertools

import (
	"bufio"
	"io"

	"github.com/theMagicalKarp/iter"
)

type linesIter struct {
	next *string
	src  *bufio.Scanner
}

func (l *linesIter) Next() string {
	if !l.HasNext() {
		return ""
	}

	next := *l.next
	l.next = nil

	return next
}

func (l *linesIter) HasNext() bool {
	if l.next != nil {
		return true
	}

	if l.src.Scan() {
		s := l.src.Text()
		l.next = &s

		return true
	}

	return false
}

func Lines(src io.Reader) iter.Iter[string] {
	return &linesIter{next: nil, src: bufio.NewScanner(src)}
}
