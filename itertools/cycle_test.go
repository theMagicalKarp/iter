package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestCycleBasic(t *testing.T) {
	t.Parallel()

	cycleIter := itertools.Cycle(iter.NewIter(1, 2, 3), 3)

	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			assert.True(t, cycleIter.HasNext())
			assert.Equal(t, j, cycleIter.Next())
		}
	}
	assert.False(t, cycleIter.HasNext())
	assert.Equal(t, 0, cycleIter.Next())
}

func TestCycleOnce(t *testing.T) {
	t.Parallel()

	cycleIter := itertools.Cycle(iter.NewIter(1, 2, 3), 1)

	for j := 1; j < 4; j++ {
		assert.True(t, cycleIter.HasNext())
		assert.Equal(t, j, cycleIter.Next())
	}

	assert.False(t, cycleIter.HasNext())
	assert.Equal(t, 0, cycleIter.Next())
}

func TestCycleEmpty(t *testing.T) {
	t.Parallel()

	cycleIter := itertools.Cycle(iter.NewIter[int](), 3)

	assert.False(t, cycleIter.HasNext())
	assert.Equal(t, 0, cycleIter.Next())
}

func TestCycleZero(t *testing.T) {
	t.Parallel()

	cycleIter := itertools.Cycle(iter.NewIter(1), 0)

	assert.False(t, cycleIter.HasNext())
	assert.Equal(t, 0, cycleIter.Next())
}

func TestCycleEmptyWithZero(t *testing.T) {
	t.Parallel()

	cycleIter := itertools.Cycle(iter.NewIter[int](), 0)

	assert.False(t, cycleIter.HasNext())
	assert.Equal(t, 0, cycleIter.Next())
}

func TestCycleInfinite(t *testing.T) {
	t.Parallel()

	cycleIter := itertools.Cycle(iter.NewIter(1, 2, 3), -1)

	for i := 0; i < 100; i++ {
		for j := 1; j < 4; j++ {
			assert.True(t, cycleIter.HasNext())
			assert.Equal(t, j, cycleIter.Next())
		}
	}

	assert.True(t, cycleIter.HasNext())
	assert.Equal(t, 1, cycleIter.Next())
	assert.Equal(t, 2, cycleIter.Next())
	assert.Equal(t, 3, cycleIter.Next())
}
