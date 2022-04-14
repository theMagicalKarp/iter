package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestTakeWhileBasic(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 6)
	items := itertools.TakeWhile(target, func(i int) bool {
		return i < 4
	})

	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())

	assert.True(t, target.HasNext())
	assert.Equal(t, 5, target.Next())
	assert.True(t, target.HasNext())
	assert.Equal(t, 6, target.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestTakeWhileOverflow(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 6)
	items := itertools.TakeWhile(target, func(i int) bool {
		return true
	})

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

func TestTakeWhileNever(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 6)
	items := itertools.TakeWhile(target, func(i int) bool {
		return false
	})

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())

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

func TestTakeWhileEmtpy(t *testing.T) {
	t.Parallel()

	target := iter.NewIter[int]()
	items := itertools.TakeWhile(target, func(i int) bool {
		return i < 4
	})

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())

	assert.False(t, target.HasNext())
	assert.Equal(t, 0, target.Next())
}
