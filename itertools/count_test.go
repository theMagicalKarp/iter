package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
	"github.com/theMagicalKarp/iter/predicates"
)

func TestCountBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Count(iter.NewIter(1, 2, 3, 4, 5, 6, 7), func(n int) bool {
		return n > 3
	})

	assert.Equal(t, 4, result)
}

func TestCountAllFalse(t *testing.T) {
	t.Parallel()

	result := itertools.Count(iter.NewIter(1, 2, 3, 4, 5, 6, 7), predicates.False[int])

	assert.Equal(t, 0, result)
}

func TestCountEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.Count(iter.NewIter[int](), predicates.True[int])

	assert.Equal(t, 0, result)
}
