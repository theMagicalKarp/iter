package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestAtBasic(t *testing.T) {
	t.Parallel()

	result := itertools.At(iter.NewIter(1, 2, 3, 4, 5, 6), 3)

	assert.Equal(t, 4, result)
}

func TestAtOverflow(t *testing.T) {
	t.Parallel()

	result := itertools.At(iter.NewIter(1, 2, 3, 4, 5, 6), 100)

	assert.Equal(t, 0, result)
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.At(iter.NewIter[int](), 3)

	assert.Equal(t, 0, result)
}

func TestAtZero(t *testing.T) {
	t.Parallel()

	result := itertools.At(iter.NewIter(1, 2, 3, 4, 5, 6), 0)

	assert.Equal(t, 1, result)
}
