package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestChainBasic(t *testing.T) {
	t.Parallel()

	chainIter := itertools.Chain(
		iter.NewIter(1, 2, 3),
		iter.NewIter(4, 5, 6),
		iter.NewIter(7, 8, 9),
	)

	for i := 1; i < 10; i++ {
		assert.True(t, chainIter.HasNext())
		assert.Equal(t, i, chainIter.Next())
	}

	assert.False(t, chainIter.HasNext())
	assert.Equal(t, 0, chainIter.Next())
}

func TestChainEmpty(t *testing.T) {
	t.Parallel()

	chainIter := itertools.Chain[int]()

	assert.False(t, chainIter.HasNext())
	assert.Equal(t, 0, chainIter.Next())
}
