package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleDrop() {
	items := itertools.Drop(iter.New("a", "b", "c", "d", "e"), 2)
	itertools.Print(items)
	// Output: c, d, e
}

func TestDropBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Drop(iter.New(1, 2, 3, 4, 5), 3)

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

func TestDropOverflow(t *testing.T) {
	t.Parallel()

	items := itertools.Drop(iter.New(1, 2, 3, 4, 5), 6)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestDropEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Drop(iter.New[int](), 3)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestDropZero(t *testing.T) {
	t.Parallel()

	items := itertools.Drop(iter.New(1, 2, 3, 4, 5), 0)

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
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
