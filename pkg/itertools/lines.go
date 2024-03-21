package itertools

import (
	"bufio"
	"io"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

type linesIter struct {
	next *string
	src  *bufio.Scanner
}

func (l *linesIter) Next() (string, bool) {
	if l.src.Scan() {
		value := l.src.Text()

		return value, true
	}

	return "", false
}

// Lines returns an iterable that produces each line of text from the given io.Reader.
// Each line is represented as a string.
func Lines(src io.Reader) iter.Iterable[string] {
	return &linesIter{next: nil, src: bufio.NewScanner(src)}
}
