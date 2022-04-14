package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestTakeBasic(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 6)
	items := itertools.Take(target, 3)

	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())

	assert.True(t, target.HasNext())
	assert.Equal(t, 4, target.Next())
	assert.True(t, target.HasNext())
	assert.Equal(t, 5, target.Next())
	assert.True(t, target.HasNext())
	assert.Equal(t, 6, target.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestTakeOverflow(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 6)
	items := itertools.Take(target, 100)

	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 5, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 6, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())

	assert.False(t, target.HasNext())
	assert.Equal(t, 0, target.Next())
}

func TestTakeZero(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 6)
	items := itertools.Take(target, 0)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())

	assert.True(t, target.HasNext())
	assert.Equal(t, 1, target.Next())
	assert.True(t, target.HasNext())
	assert.Equal(t, 2, target.Next())
	assert.True(t, target.HasNext())
	assert.Equal(t, 3, target.Next())
	assert.True(t, target.HasNext())
	assert.Equal(t, 4, target.Next())
	assert.True(t, target.HasNext())
	assert.Equal(t, 5, target.Next())
	assert.True(t, target.HasNext())
	assert.Equal(t, 6, target.Next())
	assert.False(t, target.HasNext())
	assert.Equal(t, 0, target.Next())
}

func TestTakeEmtpy(t *testing.T) {
	t.Parallel()

	target := iter.NewIter[int]()
	items := itertools.Take(target, 3)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())

	assert.False(t, target.HasNext())
	assert.Equal(t, 0, target.Next())
}
