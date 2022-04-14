package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestRepeatBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Repeat(42, 5)

	for i := 0; i < 5; i++ {
		assert.True(t, items.HasNext())
		assert.Equal(t, 42, items.Next())
	}

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRepeatZero(t *testing.T) {
	t.Parallel()

	items := itertools.Repeat(42, 0)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRepeatInf(t *testing.T) {
	t.Parallel()

	items := itertools.Repeat(42, -1)

	for i := 0; i < 1000; i++ {
		assert.True(t, items.HasNext())
		assert.Equal(t, 42, items.Next())
	}

	assert.True(t, items.HasNext())
	assert.Equal(t, 42, items.Next())
}
