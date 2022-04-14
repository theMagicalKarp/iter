package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestBasicJoin(t *testing.T) {
	t.Parallel()

	items := itertools.Join(iter.NewIter(1, 2, 3), 42)

	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.Equal(t, 42, items.Next())
	assert.Equal(t, 2, items.Next())
	assert.Equal(t, 42, items.Next())
	assert.Equal(t, 3, items.Next())
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}

func TestJoinEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Join(iter.NewIter[int](), 42)
	assert.Equal(t, 0, items.Next())
	assert.False(t, items.HasNext())
}
