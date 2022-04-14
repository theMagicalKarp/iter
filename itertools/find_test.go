package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestFindBasic(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 6, 3)

	resp := itertools.Find(target, func(i int) bool {
		return i == 3
	})

	assert.Equal(t, 3, resp)
	assert.True(t, target.HasNext())
}

func TestFindDNE(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 3, 6, 3)

	resp := itertools.Find(target, func(i int) bool {
		return i == 42
	})

	assert.Equal(t, 0, resp)
	assert.False(t, target.HasNext())
}

func TestFindEmpty(t *testing.T) {
	t.Parallel()

	target := iter.NewIter[int]()

	resp := itertools.Find(target, func(i int) bool {
		return true
	})

	assert.Equal(t, 0, resp)
	assert.False(t, target.HasNext())
}
