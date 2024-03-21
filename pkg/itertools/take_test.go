package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleTake() {
	items := itertools.Take(iter.New(1, 2, 3, 4, 5, 6), 3)

	itertools.Print(items)
	// Output:
	// 1, 2, 3
}

func TestTakeBasic(t *testing.T) {
	t.Parallel()

	target := iter.New(1, 2, 3, 4, 5, 6)
	items := itertools.Take(target, 3)

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)
	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)
	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 4, v)
	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 5, v)
	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 6, v)
	v, more = target.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = target.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestTakeOverflow(t *testing.T) {
	t.Parallel()

	target := iter.New(1, 2, 3, 4, 5, 6)
	items := itertools.Take(target, 100)

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)
	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)
	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)
	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 4, v)
	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 5, v)
	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 6, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = target.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = target.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestTakeZero(t *testing.T) {
	t.Parallel()

	target := iter.New(1, 2, 3, 4, 5, 6)
	items := itertools.Take(target, 0)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)
	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)
	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)
	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 4, v)
	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 5, v)
	v, more = target.Next()
	assert.True(t, more)
	assert.Equal(t, 6, v)
	v, more = target.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = target.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestTakeEmtpy(t *testing.T) {
	t.Parallel()

	target := iter.New[int]()
	items := itertools.Take(target, 3)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = target.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = target.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
