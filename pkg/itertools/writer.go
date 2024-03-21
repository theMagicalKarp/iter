package itertools

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

type iterWriter struct {
	closed bool
	buffer bytes.Buffer
}

var ErrClosedIter = errors.New("io: write on closed iter")

// Next returns the next batch of bytes from the iterator's buffer.
// If the buffer is empty and the iterator is closed, it returns an empty byte slice and false.
// Otherwise, it returns a copy of the buffer's contents and true.
func (i *iterWriter) Next() ([]byte, bool) {
	if i.buffer.Len() == 0 && i.closed {
		return []byte{}, false
	}

	toReturn := make([]byte, i.buffer.Len())
	copy(toReturn, i.buffer.Bytes())
	i.buffer.Reset()

	return toReturn, true
}

// Write writes the contents of the given byte slice to the iterator writer's buffer.
// It returns the number of bytes written and any error encountered.
// If the iterator writer is closed, it returns an error indicating that the writer is closed.
func (i *iterWriter) Write(p []byte) (int, error) {
	if i.closed {
		return 0, ErrClosedIter
	}

	n, err := i.buffer.Write(p)
	if err != nil {
		return n, fmt.Errorf("iter write error: %w", err)
	}

	return n, nil
}

// Close closes the iterWriter and marks it as closed.
// It returns nil error.
func (i *iterWriter) Close() error {
	i.closed = true

	return nil
}

type IterWriteCloser interface {
	io.WriteCloser
	iter.Iterable[[]byte]
}

// Writer returns a new instance of io.WriteCloser and iter.Iterable.
// Data written to this writer is stored to a buffer and then read it back as an iterable of byte slices.
// The writer can be closed to signal the end of the data, otherwise it will infinity return empty slices.
// After closing, the writer will continue to return results from the buffer until it is empty,
// and will not accept any more writes.
func Writer() IterWriteCloser {
	return &iterWriter{closed: false}
}
