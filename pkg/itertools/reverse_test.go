package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleReverse() {
	items := itertools.Reverse(iter.New(1, 2, 3))

	itertools.Print(items)
	// Output: 3, 2, 1
}

func TestReverseBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Reverse(iter.New(1, 2, 3, 4, 5))

	for i := 5; i >= 1; i-- {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestReverseEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Reverse(iter.New[int]())

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
