package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleChain() {
	items := itertools.Chain(
		iter.New(1, 2, 3),
		iter.New(4, 5, 6),
		iter.New(7, 8, 9),
	)
	itertools.Print(items)
	// Output:
	// 1, 2, 3, 4, 5, 6, 7, 8, 9
}

func TestChainBasic(t *testing.T) {
	t.Parallel()

	chainIter := itertools.Chain(
		iter.New(1, 2, 3),
		iter.New(4, 5, 6),
		iter.New(7, 8, 9),
	)

	for i := 1; i < 10; i++ {
		v, more := chainIter.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}
	v, more := chainIter.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestChainEmpty(t *testing.T) {
	t.Parallel()

	chainIter := itertools.Chain[int]()

	v, more := chainIter.Next()

	assert.False(t, more)
	assert.Equal(t, 0, v)
}
