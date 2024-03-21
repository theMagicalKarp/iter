package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleDrain() {
	items := iter.New("a", "b", "c")
	itertools.Drain(items)

	itertools.Print(items)
	// Output:
}

func TestDrainBasic(t *testing.T) {
	t.Parallel()

	items := iter.New(1, 2, 3, 4, 5, 6, 7)
	itertools.Drain(items)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestDrainEmpty(t *testing.T) {
	t.Parallel()

	items := iter.New[int]()
	itertools.Drain(items)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
