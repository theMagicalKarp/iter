package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestSumBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Sum(iter.NewIter(1, 2, 3, 4, 5))

	assert.Equal(t, 15, result)
}

func TestSumNegatives(t *testing.T) {
	t.Parallel()

	result := itertools.Sum(iter.NewIter(1, 2, -3, -4, 5))

	assert.Equal(t, 1, result)
}

func TestSumEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.Sum(iter.NewIter[int]())

	assert.Equal(t, 0, result)
}

func TestSumStrings(t *testing.T) {
	t.Parallel()

	result := itertools.Sum(iter.NewIter("a", "b", "c", "d", "hello"))

	assert.Equal(t, "abcdhello", result)
}
