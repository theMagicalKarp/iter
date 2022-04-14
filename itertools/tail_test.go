package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestTailBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Tail(iter.NewIter(1, 2, 3, 4, 5), 2)

	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.Equal(t, 5, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestTailEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Tail(iter.NewIter[int](), 2)

	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestTailOverflow(t *testing.T) {
	t.Parallel()

	items := itertools.Tail(iter.NewIter(1, 2, 3), 5)

	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.Equal(t, 2, items.Next())
	assert.Equal(t, 3, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}
