package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestDropWhileBasic(t *testing.T) {
	t.Parallel()

	items := itertools.DropWhile(iter.NewIter(1, 2, 3, 4, 5), func(i int) bool {
		return i < 3
	})

	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.Equal(t, 4, items.Next())
	assert.Equal(t, 5, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestDropWhileOverflow(t *testing.T) {
	t.Parallel()

	items := itertools.DropWhile(iter.NewIter(1, 2, 3, 4, 5), func(i int) bool {
		return i < 100
	})

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestDropEmtpy(t *testing.T) {
	t.Parallel()

	items := itertools.DropWhile(iter.NewIter[int](), func(i int) bool {
		return true
	})

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestDropEmtpyFalse(t *testing.T) {
	t.Parallel()

	items := itertools.DropWhile(iter.NewIter[int](), func(i int) bool {
		return false
	})

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestDropWhileBasicTrue(t *testing.T) {
	t.Parallel()

	items := itertools.DropWhile(iter.NewIter(1, 2, 3, 4, 5), func(i int) bool {
		return true
	})

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestDropWhileBasicFalse(t *testing.T) {
	t.Parallel()

	items := itertools.DropWhile(iter.NewIter(1, 2, 3, 4, 5), func(i int) bool {
		return false
	})

	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.Equal(t, 2, items.Next())
	assert.Equal(t, 3, items.Next())
	assert.Equal(t, 4, items.Next())
	assert.Equal(t, 5, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}
