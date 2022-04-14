package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestAllTrue(t *testing.T) {
	t.Parallel()

	result := itertools.All(iter.NewIter(1, 2, 3, 4, 5, 6, 7), func(n int) bool {
		return n > 0
	})

	assert.True(t, result)
}

func TestAllSomeTrue(t *testing.T) {
	t.Parallel()

	result := itertools.All(iter.NewIter(1, 2, 3, 4, -5, 6, 7), func(n int) bool {
		return n > 0
	})

	assert.False(t, result)
}

func TestAllNoTrue(t *testing.T) {
	t.Parallel()

	result := itertools.All(iter.NewIter(-1, -2, -3, -4, -5, -6, -7), func(n int) bool {
		return n > 0
	})

	assert.False(t, result)
}

func TestAllEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.All(iter.NewIter[int](), func(n int) bool {
		return n > 0
	})

	assert.True(t, result)
}
