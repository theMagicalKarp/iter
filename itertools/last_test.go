package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestLastBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Last(iter.NewIter(1, 2, 3, 4, 5), func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, 4, result)
}

func TestLastBasicDNE(t *testing.T) {
	t.Parallel()

	result := itertools.Last(iter.NewIter(1, 2, 3, 4, 5), func(n int) bool {
		return false
	})

	assert.Equal(t, 0, result)
}

func TestLastBasicEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.Last(iter.NewIter[int](), func(n int) bool {
		return true
	})

	assert.Equal(t, 0, result)
}
