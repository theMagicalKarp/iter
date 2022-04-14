package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestAnyTrue(t *testing.T) {
	t.Parallel()

	result := itertools.Any(iter.NewIter(1, 2, 3, 4, 5, 6, 7), func(n int) bool {
		return n > 0
	})

	assert.True(t, result)
}

func TestAnySomeTrue(t *testing.T) {
	t.Parallel()

	result := itertools.Any(iter.NewIter(1, 2, 3, 4, -5, 6, 7), func(n int) bool {
		return n > 0
	})

	assert.True(t, result)
}

func TestAnyNoTrue(t *testing.T) {
	t.Parallel()

	result := itertools.Any(iter.NewIter(-1, -2, -3, -4, -5, -6, -7), func(n int) bool {
		return n > 0
	})

	assert.False(t, result)
}

func TestAnyEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.Any(iter.NewIter[int](), func(n int) bool {
		return n > 0
	})

	assert.False(t, result)
}
