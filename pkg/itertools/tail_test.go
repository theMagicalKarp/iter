package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleTail() {
	items := itertools.Tail(iter.New(1, 2, 3, 4, 5), 2)

	itertools.Print(items)
	// Output:
	// 4, 5
}

func TestTailBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Tail(iter.New(1, 2, 3, 4, 5), 2)
	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 4, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 5, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestTailEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Tail(iter.New[int](), 2)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestTailOverflow(t *testing.T) {
	t.Parallel()

	items := itertools.Tail(iter.New(1, 2, 3), 5)

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
}

func TestTailExact(t *testing.T) {
	t.Parallel()

	items := itertools.Tail(iter.New(1, 2, 3), 3)

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
}

func TestTailZero(t *testing.T) {
	t.Parallel()

	items := itertools.Tail(iter.New(1, 2, 3, 4, 5), 0)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
