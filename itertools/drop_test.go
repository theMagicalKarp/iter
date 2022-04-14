package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestDropBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Drop(iter.NewIter(1, 2, 3, 4, 5), 3)

	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.Equal(t, 5, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestDropOverflow(t *testing.T) {
	t.Parallel()

	items := itertools.Drop(iter.NewIter(1, 2, 3, 4, 5), 6)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestDropEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Drop(iter.NewIter[int](), 3)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestDropZero(t *testing.T) {
	t.Parallel()

	items := itertools.Drop(iter.NewIter(1, 2, 3, 4, 5), 0)

	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.Equal(t, 4, items.Next())
	assert.Equal(t, 5, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}
