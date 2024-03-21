package itertools

import (
	"io"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

type iterReader struct {
	iter   iter.Iterable[[]byte]
	buffer []byte
}

func (r *iterReader) Read(outBytes []byte) (int, error) {
	written := 0

	if len(r.buffer) > 0 {
		toWrite := len(outBytes)
		if len(r.buffer) < toWrite {
			toWrite = len(r.buffer)
		}

		copy(outBytes, r.buffer[:toWrite])
		r.buffer = r.buffer[toWrite:]
		written += toWrite
	}

	for written < len(outBytes) {
		value, more := r.iter.Next()
		if !more {
			break
		}

		toWrite := len(outBytes) - written
		if len(value) < toWrite {
			toWrite = len(value)
		}

		copy(outBytes[written:], value[:toWrite])
		r.buffer = value[toWrite:]
		written += toWrite
	}

	if len(outBytes) > 0 && written == 0 {
		return 0, io.EOF
	}

	return written, nil
}

// Reader returns an io.Reader that reads from the given iterable of byte slices.
// Each call to Read will consume the next byte slice from the iterable.
func Reader(iter iter.Iterable[[]byte]) io.Reader {
	return &iterReader{iter: iter, buffer: make([]byte, 0)}
}
