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

func ExampleReader() {
	iter := iter.New([]byte("Hello"), []byte(" "), []byte("World!"))

	reader := itertools.Reader(iter)

	buffer := new(strings.Builder)

	_, err := io.Copy(buffer, reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(buffer.String())

	// Output:
	// Hello World!
}

func TestBasicReader(t *testing.T) {
	t.Parallel()

	// Create a test iterable
	testData := [][]byte{
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
	}
	iter := iter.New(testData...)

	// Create a reader using the Reader function
	reader := itertools.Reader(iter)

	// Read from the reader and verify the output
	expectedOutput := []byte("HelloWorldAbcDefSharks!")
	output := make([]byte, len(expectedOutput))
	n, err := reader.Read(output)
	require.NoError(t, err)
	assert.Equal(t, len(expectedOutput), n)
	assert.Equal(t, expectedOutput, output)

	n, err = reader.Read(output)
	assert.Equal(t, err, io.EOF)
	assert.Equal(t, 0, n)
}

func TestTinyReads(t *testing.T) {
	t.Parallel()

	// Create a test iterable
	testData := [][]byte{
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
	}
	iter := iter.New(testData...)

	// Create a reader using the Reader function
	reader := itertools.Reader(iter)

	// Read from the reader and verify the output
	expectedOutput := []byte("HelloWorldAbcDefSharks!")

	for i := range expectedOutput {
		output := make([]byte, 1)
		n, err := reader.Read(output)
		require.NoError(t, err)
		assert.Equal(t, 1, n)
		assert.Equal(t, expectedOutput[i:i+1], output)
	}

	n, err := reader.Read([]byte("foo"))
	assert.Equal(t, err, io.EOF)
	assert.Equal(t, 0, n)
}

func TestSmallReads(t *testing.T) {
	t.Parallel()

	// Create a test iterable
	testData := [][]byte{
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
		[]byte("foo"),
	}
	iter := iter.New(testData...)

	// Create a reader using the Reader function
	reader := itertools.Reader(iter)

	// Read from the reader and verify the output
	expectedOutput := []byte("HelloWorldAbcDefSharks!foo")

	for i := range len(expectedOutput) / 2 {
		output := make([]byte, 2)
		n, err := reader.Read(output)
		require.NoError(t, err)
		assert.Equal(t, 2, n)
		assert.Equal(t, expectedOutput[i*2:i*2+2], output)
	}

	n, err := reader.Read([]byte("foo"))
	assert.Equal(t, err, io.EOF)
	assert.Equal(t, 0, n)
}

func TestLittleReads(t *testing.T) {
	t.Parallel()

	// Create a test iterable
	testData := [][]byte{
		[]byte("!"),
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
		[]byte("foo"),
	}
	iter := iter.New(testData...)

	// Create a reader using the Reader function
	reader := itertools.Reader(iter)

	// Read from the reader and verify the output
	expectedOutput := []byte("!HelloWorldAbcDefSharks!foo")

	for i := range len(expectedOutput) / 3 {
		output := make([]byte, 3)
		n, err := reader.Read(output)
		require.NoError(t, err)
		assert.Equal(t, 3, n)
		assert.Equal(t, expectedOutput[i*3:i*3+3], output)
	}

	n, err := reader.Read([]byte("foo"))
	assert.Equal(t, err, io.EOF)
	assert.Equal(t, 0, n)
}

func TestReaderBigBuffer(t *testing.T) {
	t.Parallel()

	// Create a test iterable
	testData := [][]byte{
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
	}
	iter := iter.New(testData...)

	// Create a reader using the Reader function
	reader := itertools.Reader(iter)

	// Read from the reader and verify the output
	expectedOutput := []byte("HelloWorldAbcDefSharks!")
	output := make([]byte, 256)
	n, err := reader.Read(output)
	require.NoError(t, err)
	assert.Equal(t, len(expectedOutput), n)
	assert.Equal(t, append(expectedOutput, make([]byte, 233)...), output)

	n, err = reader.Read([]byte("foo"))
	assert.Equal(t, err, io.EOF)
	assert.Equal(t, 0, n)
}

func TestEmptyReader(t *testing.T) {
	t.Parallel()

	iter := iter.New[[]byte]()
	reader := itertools.Reader(iter)

	output := make([]byte, 256)
	n, err := reader.Read(output)
	assert.Equal(t, err, io.EOF)
	assert.Equal(t, 0, n)

	for i := range 256 {
		assert.Equal(t, byte(0), output[i])
	}
}

func TestReaderCopy(t *testing.T) {
	t.Parallel()

	testData := [][]byte{
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
	}
	iter := iter.New(testData...)

	var b bytes.Buffer

	written, err := io.Copy(&b, itertools.Reader(iter))
	require.NoError(t, err)
	assert.Equal(t, int64(23), written)
	assert.Equal(t, "HelloWorldAbcDefSharks!", b.String())
}

func TestReaderCopySmallBuffer(t *testing.T) {
	t.Parallel()

	testData := [][]byte{
		[]byte("Hello"),
		[]byte("World"),
		[]byte("AbcDef"),
		[]byte("Sharks!"),
	}
	iter := iter.New(testData...)

	var b bytes.Buffer

	copyBuffer := make([]byte, 3)
	written, err := io.CopyBuffer(&b, itertools.Reader(iter), copyBuffer)
	require.NoError(t, err)
	assert.Equal(t, int64(23), written)
	assert.Equal(t, "HelloWorldAbcDefSharks!", b.String())
}
