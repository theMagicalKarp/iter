package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestDrainBasic(t *testing.T) {
	t.Parallel()

	items := iter.NewIter(1, 2, 3, 4, 5, 6, 7)
	itertools.Drain(items)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestDrainEmpty(t *testing.T) {
	t.Parallel()

	items := iter.NewIter[int]()
	itertools.Drain(items)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}
