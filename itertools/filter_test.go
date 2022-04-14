package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestFilterBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Filter(iter.NewIter(1, 2, 3, 4, 5), func(i int) bool {
		return i%2 == 0
	})

	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestFilterEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Filter(iter.NewIter[int](), func(i int) bool {
		return true
	})

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestFilterNoMatch(t *testing.T) {
	t.Parallel()

	items := itertools.Filter(iter.NewIter(1, 2, 3, 4, 5), func(i int) bool {
		return false
	})

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}
