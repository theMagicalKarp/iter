package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestMaxBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Max(iter.NewIter(1, 2, 3, 6, 4, 5))

	assert.Equal(t, 6, result)
}

func TestMaxEmtpy(t *testing.T) {
	t.Parallel()

	result := itertools.Max(iter.NewIter[int]())

	assert.Equal(t, 0, result)
}

func TestMaxNegatives(t *testing.T) {
	t.Parallel()

	result := itertools.Max(iter.NewIter(-5, -4, -3, -2, -1))

	assert.Equal(t, -1, result)
}

func TestMaxMixed(t *testing.T) {
	t.Parallel()

	result := itertools.Max(iter.NewIter(3, 4, 1, 2, 5))

	assert.Equal(t, 5, result)
}
