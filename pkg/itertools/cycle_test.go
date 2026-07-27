package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleCycle() {
	items := itertools.Cycle(iter.New("a", "b", "c"))
	itertools.Print(itertools.Take(items, 7))
	// Output: a, b, c, a, b, c, a
}

func TestCycleBasic(t *testing.T) {
	t.Parallel()

	cycleIter := itertools.Cycle(iter.New(1, 2, 3))

	for range 10 {
		for j := 1; j < 4; j++ {
			v, more := cycleIter.Next()
			assert.True(t, more)
			assert.Equal(t, j, v)
		}
	}
}

func TestCycleEmpty(t *testing.T) {
	t.Parallel()

	cycleIter := itertools.Cycle(iter.New[int]())

	v, more := cycleIter.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = cycleIter.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
