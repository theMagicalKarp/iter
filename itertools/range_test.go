package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestRangeBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Range(1, 10)

	for i := 1; i < 10; i++ {
		assert.True(t, items.HasNext())
		assert.Equal(t, i, items.Next())
	}

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeNeg(t *testing.T) {
	t.Parallel()

	items := itertools.Range(-10, 10)

	for i := -10; i < 10; i++ {
		assert.True(t, items.HasNext())
		assert.Equal(t, i, items.Next())
	}

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeWeird(t *testing.T) {
	t.Parallel()

	items := itertools.Range(10, -10)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeBad(t *testing.T) {
	t.Parallel()

	items := itertools.Range(10, 10)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeAdvanced(t *testing.T) {
	t.Parallel()

	items := itertools.Range(10, 20)

	for i := 10; i < 20; i++ {
		assert.True(t, items.HasNext())
		assert.Equal(t, i, items.Next())
	}

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeIncBasic(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(1, 10, 3)

	for i := 1; i < 10; i += 3 {
		assert.True(t, items.HasNext())
		assert.Equal(t, i, items.Next())
	}

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeIncNeg(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(-10, 10, 2)

	for i := -10; i < 10; i += 2 {
		assert.True(t, items.HasNext())
		assert.Equal(t, i, items.Next())
	}

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeIncWeird(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(10, -10, 2)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeIncBad(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(10, 10, 3)

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeIncAdvanced(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(10, 20, 4)

	for i := 10; i < 20; i += 4 {
		assert.True(t, items.HasNext())
		assert.Equal(t, i, items.Next())
	}

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestRangeIncOverflow(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(1, 10, 10)

	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}
