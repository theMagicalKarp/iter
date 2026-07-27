package itertools_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleLines() {
	items := itertools.Lines(strings.NewReader("Hello\nWorld\n!"))

	itertools.Print(items)
	// Output:
	// Hello, World, !
}

func TestLinesBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Lines(strings.NewReader("foo\nbar\nhello\nworld\n\nsup\n"))

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, "foo", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "bar", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "hello", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "world", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Empty(t, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "sup", v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)
}

func TestLinesNoNewLines(t *testing.T) {
	t.Parallel()

	items := itertools.Lines(strings.NewReader("hello"))

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, "hello", v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)
}

func TestLinesEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Lines(strings.NewReader(""))

	v, more := items.Next()
	assert.False(t, more)
	assert.Empty(t, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)
}

func TestWithReader(t *testing.T) {
	t.Parallel()

	reader := itertools.Reader(iter.New([][]byte{
		[]byte("!"),
		[]byte("Hello\n"),
		[]byte("World"),
		[]byte("Abc\nD\nef"),
		[]byte("Sharks!"),
		[]byte("foo\n"),
	}...))

	items := itertools.Lines(reader)

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, "!Hello", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "WorldAbc", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "D", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "efSharks!foo", v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)
}

func TestWithReaderWriter(t *testing.T) {
	t.Parallel()

	writer := itertools.Writer()
	items := itertools.Lines(itertools.Reader(writer))

	toWrite := [][]byte{
		[]byte("!"),
		[]byte("Hello\n"),
		[]byte("World"),
		[]byte("Abc\nD\nef"),
		[]byte("Sharks!"),
		[]byte("foo\n"),
	}

	for _, b := range toWrite {
		n, err := io.Copy(writer, bytes.NewReader(b))
		require.NoError(t, err)
		assert.Equal(t, int64(len(b)), n)
	}

	require.NoError(t, writer.Close())

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, "!Hello", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "WorldAbc", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "D", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "efSharks!foo", v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)
}

func TestWithReaderWriterSmallBuffer(t *testing.T) {
	t.Parallel()

	writer := itertools.Writer()
	items := itertools.Lines(itertools.Reader(writer))

	toWrite := [][]byte{
		[]byte("!"),
		[]byte("Hello\n"),
		[]byte("World"),
		[]byte("Abc\nD\nef"),
		[]byte("Sharks!"),
		[]byte("foo\n"),
	}

	buffer := make([]byte, 3)
	for _, b := range toWrite {
		n, err := io.CopyBuffer(writer, bytes.NewReader(b), buffer)
		require.NoError(t, err)
		assert.Equal(t, int64(len(b)), n)
	}

	require.NoError(t, writer.Close())

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, "!Hello", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "WorldAbc", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "D", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "efSharks!foo", v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Empty(t, v)
}
