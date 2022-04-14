package iter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
)

func TestNewIter(t *testing.T) {
	t.Parallel()

	items := iter.NewIter(1, 2, 3)
	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestNewIterEmpty(t *testing.T) {
	t.Parallel()

	items := iter.NewIter[int]()
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}
