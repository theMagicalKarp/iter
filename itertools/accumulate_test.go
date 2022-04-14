package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestAccumulateInts(t *testing.T) {
	t.Parallel()

	result := iter.NewArray(0, 8, itertools.Accumulate(iter.NewIter(1, -2, 3, 4, -5, 6, 7)))

	assert.Equal(t, 1, result.At(0))
	assert.Equal(t, -1, result.At(1))
	assert.Equal(t, 2, result.At(2))
	assert.Equal(t, 6, result.At(3))
	assert.Equal(t, 1, result.At(4))
	assert.Equal(t, 7, result.At(5))
	assert.Equal(t, 14, result.At(6))
	assert.Equal(t, 7, result.Len())
}

func TestAccumulateStrings(t *testing.T) {
	t.Parallel()

	result := iter.NewArray(0, 8, itertools.Accumulate(iter.NewIter("hello", " world", "!")))

	assert.Equal(t, "hello", result.At(0))
	assert.Equal(t, "hello world", result.At(1))
	assert.Equal(t, "hello world!", result.At(2))
	assert.Equal(t, 3, result.Len())
}

func TestAccumulateEmpty(t *testing.T) {
	t.Parallel()

	result := iter.NewArray(0, 8, itertools.Accumulate(iter.NewIter[string]()))

	assert.Equal(t, 0, result.Len())
}
