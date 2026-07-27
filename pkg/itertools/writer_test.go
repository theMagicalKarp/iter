package itertools_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleWriter() {
	writer := itertools.Writer()

	_, err := io.Copy(writer, strings.NewReader("Hello, World!"))
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	out, _ := writer.Next()
	fmt.Println(string(out))

	// Output:
	// Hello, World!
}

func TestWriterBasic(t *testing.T) {
	t.Parallel()

	writer := itertools.Writer()

	n, err := writer.Write([]byte("Hello, "))
	require.NoError(t, err)
	assert.Equal(t, 7, n)

	n, err = writer.Write([]byte("World!"))
	require.NoError(t, err)
	assert.Equal(t, 6, n)

	v, more := writer.Next()
	assert.True(t, more)
	assert.Equal(t, []byte("Hello, World!"), v)

	n, err = writer.Write([]byte("weeeeee"))
	require.NoError(t, err)
	assert.Equal(t, 7, n)

	v, more = writer.Next()
	assert.True(t, more)
	assert.Equal(t, []byte("weeeeee"), v)

	v, more = writer.Next()
	assert.True(t, more)
	assert.Equal(t, []byte{}, v)

	err = writer.Close()
	require.NoError(t, err)

	v, more = writer.Next()
	assert.False(t, more)
	assert.Equal(t, []byte{}, v)

	v, more = writer.Next()
	assert.False(t, more)
	assert.Equal(t, []byte{}, v)

	n, err = writer.Write([]byte("foooooo"))
	assert.Equal(t, itertools.ErrClosedIter, err)
	assert.Equal(t, 0, n)
}

func TestWriterCopy(t *testing.T) {
	t.Parallel()

	writer := itertools.Writer()

	var expected bytes.Buffer

	placeholder := "Hello, World!\n"

	for range 100 {
		n, err := io.Copy(writer, strings.NewReader(placeholder))
		require.NoError(t, err)
		assert.Equal(t, int64(14), n)

		n, err = io.Copy(&expected, strings.NewReader(placeholder))

		require.NoError(t, err)
		assert.Equal(t, int64(14), n)
	}

	v, more := writer.Next()
	assert.True(t, more)
	assert.Equal(t, expected.Bytes(), v)
}

func TestWriterCopySmallBuffer(t *testing.T) {
	t.Parallel()

	writer := itertools.Writer()

	var expected bytes.Buffer

	placeholder := "Hello, World!\n"
	buffer := make([]byte, 3)

	for range 100 {
		n, err := io.CopyBuffer(writer, strings.NewReader(placeholder), buffer)
		require.NoError(t, err)
		assert.Equal(t, int64(14), n)

		n, err = io.CopyBuffer(&expected, strings.NewReader(placeholder), buffer)
		require.NoError(t, err)
		assert.Equal(t, int64(14), n)
	}

	v, more := writer.Next()
	assert.True(t, more)
	assert.Equal(t, expected.Bytes(), v)
}

func TestWriterAndReader(t *testing.T) {
	t.Parallel()

	writer := itertools.Writer()
	reader := itertools.Reader(iter.New([][]byte{
		[]byte("!"),
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
		[]byte("foo"),
	}...))

	n, err := io.Copy(writer, reader)
	require.NoError(t, err)
	assert.Equal(t, int64(27), n)

	v, more := writer.Next()
	assert.True(t, more)
	assert.Equal(t, []byte("!HelloWorldAbcDefSharks!foo"), v)

	v, more = writer.Next()
	assert.True(t, more)
	assert.Equal(t, []byte{}, v)

	err = writer.Close()
	require.NoError(t, err)

	v, more = writer.Next()
	assert.False(t, more)
	assert.Equal(t, []byte{}, v)

	v, more = writer.Next()
	assert.False(t, more)
	assert.Equal(t, []byte{}, v)
}

func TestWriterAndReaderSmallBuff(t *testing.T) {
	t.Parallel()

	writer := itertools.Writer()
	reader := itertools.Reader(iter.New([][]byte{
		[]byte("!"),
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
		[]byte("foo"),
	}...))

	buffer := make([]byte, 3)
	n, err := io.CopyBuffer(writer, reader, buffer)
	require.NoError(t, err)
	assert.Equal(t, int64(27), n)

	v, more := writer.Next()
	assert.True(t, more)
	assert.Equal(t, []byte("!HelloWorldAbcDefSharks!foo"), v)

	v, more = writer.Next()
	assert.True(t, more)
	assert.Equal(t, []byte{}, v)

	err = writer.Close()
	require.NoError(t, err)

	v, more = writer.Next()
	assert.False(t, more)
	assert.Equal(t, []byte{}, v)

	v, more = writer.Next()
	assert.False(t, more)
	assert.Equal(t, []byte{}, v)
}
