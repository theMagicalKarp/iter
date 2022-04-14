package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestSplitBasic(t *testing.T) {
	t.Parallel()

	items1, items2 := itertools.Split(iter.NewIter(1, 2, 3, 4, 5, 6), func(i int) bool {
		return i%2 == 0
	})

	assert.True(t, items1.HasNext())
	assert.Equal(t, 2, items1.Next())
	assert.True(t, items1.HasNext())
	assert.Equal(t, 4, items1.Next())
	assert.True(t, items1.HasNext())
	assert.Equal(t, 6, items1.Next())
	assert.False(t, items1.HasNext())
	assert.Equal(t, 0, items1.Next())

	assert.True(t, items2.HasNext())
	assert.Equal(t, 1, items2.Next())
	assert.True(t, items2.HasNext())
	assert.Equal(t, 3, items2.Next())
	assert.True(t, items2.HasNext())
	assert.Equal(t, 5, items2.Next())
	assert.False(t, items2.HasNext())
	assert.Equal(t, 0, items2.Next())
}

func TestSplitLopsided(t *testing.T) {
	t.Parallel()

	items1, items2 := itertools.Split(iter.NewIter(1, 2, 3, 4, 5, 6), func(i int) bool {
		return true
	})

	assert.True(t, items1.HasNext())
	assert.Equal(t, 1, items1.Next())
	assert.True(t, items1.HasNext())
	assert.Equal(t, 2, items1.Next())
	assert.True(t, items1.HasNext())
	assert.Equal(t, 3, items1.Next())
	assert.True(t, items1.HasNext())
	assert.Equal(t, 4, items1.Next())
	assert.True(t, items1.HasNext())
	assert.Equal(t, 5, items1.Next())
	assert.True(t, items1.HasNext())
	assert.Equal(t, 6, items1.Next())
	assert.False(t, items1.HasNext())
	assert.Equal(t, 0, items1.Next())

	assert.False(t, items2.HasNext())
	assert.Equal(t, 0, items2.Next())
}

func TestSplitEmpty(t *testing.T) {
	t.Parallel()

	items1, items2 := itertools.Split(iter.NewIter[int](), func(i int) bool {
		return i%2 == 0
	})

	assert.False(t, items1.HasNext())
	assert.Equal(t, 0, items1.Next())

	assert.False(t, items2.HasNext())
	assert.Equal(t, 0, items2.Next())
}
