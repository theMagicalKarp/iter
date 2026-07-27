package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleZip() {
	items := itertools.Zip(
		iter.New(1, 2, 3, 4, 5, 6),
		iter.New("a", "b", "c", "d", "e", "f"),
	)

	itertools.Print(items)
	// Output:
	// {1 a}, {2 b}, {3 c}, {4 d}, {5 e}, {6 f}
}

func TestZipBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Zip(
		iter.New(1, 2, 3, 4, 5, 6),
		iter.New("a", "b", "c", "d", "e", "f"),
	)

	v, more := items.Next()
	assert.True(t, more)

	a, b := v.Unpack()
	assert.Equal(t, 1, a)
	assert.Equal(t, "a", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 2, a)
	assert.Equal(t, "b", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 3, a)
	assert.Equal(t, "c", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 4, a)
	assert.Equal(t, "d", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 5, a)
	assert.Equal(t, "e", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 6, a)
	assert.Equal(t, "f", b)

	v, more = items.Next()
	assert.False(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 0, a)
	assert.Empty(t, b)
}

func TestZipUneven(t *testing.T) {
	t.Parallel()

	items := itertools.Zip(
		iter.New(1, 2, 3),
		iter.New("a", "b", "c", "d", "e", "f"),
	)

	v, more := items.Next()
	assert.True(t, more)

	a, b := v.Unpack()
	assert.Equal(t, 1, a)
	assert.Equal(t, "a", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 2, a)
	assert.Equal(t, "b", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 3, a)
	assert.Equal(t, "c", b)

	v, more = items.Next()
	assert.False(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 0, a)
	assert.Empty(t, b)
}

func TestZipOtherUneven(t *testing.T) {
	t.Parallel()

	items := itertools.Zip(
		iter.New(1, 2, 3, 4, 5, 6),
		iter.New("a", "b", "c"),
	)

	v, more := items.Next()
	assert.True(t, more)

	a, b := v.Unpack()
	assert.Equal(t, 1, a)
	assert.Equal(t, "a", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 2, a)
	assert.Equal(t, "b", b)

	v, more = items.Next()
	assert.True(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 3, a)
	assert.Equal(t, "c", b)

	v, more = items.Next()
	assert.False(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 0, a)
	assert.Empty(t, b)
}

func TestZipEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Zip(
		iter.New[int](),
		iter.New[string](),
	)

	v, more := items.Next()
	assert.False(t, more)

	a, b := v.Unpack()
	assert.Equal(t, 0, a)
	assert.Empty(t, b)

	v, more = items.Next()
	assert.False(t, more)

	a, b = v.Unpack()
	assert.Equal(t, 0, a)
	assert.Empty(t, b)
}
