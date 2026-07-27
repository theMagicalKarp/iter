package itertools

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

// IterWriter is an io.WriteCloser which buffers everything written to it and
// hands it back out as an iterable of byte slices. Use Writer to construct one.
type IterWriter struct {
	closed bool
	buffer bytes.Buffer
}

// ErrClosedIter is returned by Write when the writer has already been closed.
var ErrClosedIter = errors.New("io: write on closed iter")

// Next returns the next batch of bytes from the iterator's buffer.
// If the buffer is empty and the iterator is closed, it returns an empty byte slice and false.
// Otherwise, it returns a copy of the buffer's contents and true.
func (i *IterWriter) Next() ([]byte, bool) {
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
func (i *IterWriter) Write(p []byte) (int, error) {
	if i.closed {
		return 0, ErrClosedIter
	}

	n, err := i.buffer.Write(p)
	if err != nil {
		return n, fmt.Errorf("iter write error: %w", err)
	}

	return n, nil
}

// Close closes the IterWriter and marks it as closed.
// It returns nil error.
func (i *IterWriter) Close() error {
	i.closed = true

	return nil
}

// IterWriteCloser is the behaviour implemented by IterWriter: an io.WriteCloser
// whose contents can also be consumed as an iterable of byte slices.
type IterWriteCloser interface {
	io.WriteCloser
	iter.Iterable[[]byte]
}

// Writer returns a new IterWriter, which implements both io.WriteCloser and iter.Iterable.
// Data written to this writer is stored to a buffer and then read it back as an iterable of byte slices.
// The writer can be closed to signal the end of the data, otherwise it will infinity return empty slices.
// After closing, the writer will continue to return results from the buffer until it is empty,
// and will not accept any more writes.
func Writer() *IterWriter {
	return &IterWriter{closed: false, buffer: bytes.Buffer{}}
}
