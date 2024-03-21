package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleAccumulate() {
	items := itertools.Accumulate(iter.New("a", "b", "c"))
	itertools.Print(items)
	// Output: a, ab, abc
}

func TestAccumulateInts(t *testing.T) {
	t.Parallel()

	items := itertools.Accumulate(iter.New(0, 1, 2, 3, 4, 5))

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 6, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 10, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 15, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestAccumulateStrings(t *testing.T) {
	t.Parallel()

	items := itertools.Accumulate(iter.New("a", "b", "c"))

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, "a", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "ab", v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, "abc", v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, "", v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, "", v)
}

func TestAccumulateEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Accumulate(iter.New[string]())

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, "", v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, "", v)
}
