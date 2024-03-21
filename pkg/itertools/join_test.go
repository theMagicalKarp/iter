package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleJoin() {
	items := itertools.Join(iter.New(1, 2, 3), 0)

	itertools.Print(items)
	// Output: 1, 0, 2, 0, 3
}

func TestBasicJoin(t *testing.T) {
	t.Parallel()

	items := itertools.Join(iter.New(1, 2, 3), 42)

	for i := 1; i < 3; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)

		v, more = items.Next()
		assert.True(t, more)
		assert.Equal(t, v, 42)
	}

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)
}

func TestJoinEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Join(iter.New[int](), 42)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
