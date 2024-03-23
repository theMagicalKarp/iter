package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleFlatten() {
	items := itertools.Flatten(
		iter.New(
			iter.New(1, 2, 3),
			iter.New(4, 5, 6),
		),
		iter.New(
			iter.New(7, 8, 9),
		),
	)
	itertools.Print(items)
	// Output:
	// 1, 2, 3, 4, 5, 6, 7, 8, 9
}

func TestFlattenBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Flatten(
		iter.New(
			iter.New(1, 2, 3),
			iter.New(4, 5, 6),
			iter.New(7, 8, 9),
		),
	)

	for i := 1; i < 10; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}
	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestFlattenBasicMulti(t *testing.T) {
	t.Parallel()

	items := itertools.Flatten(
		iter.New(
			iter.New(1, 2, 3),
			iter.New(4, 5, 6),
		),
		iter.New(
			iter.New(7, 8, 9),
		),
	)

	for i := 1; i < 10; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}
	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestFlattenWithEmpties(t *testing.T) {
	t.Parallel()

	items := itertools.Flatten(
		iter.New(
			iter.New[int](),
			iter.New(1, 2, 3),
			iter.New(4, 5, 6),
			iter.New[int](),
			iter.New(7, 8, 9),
			iter.New[int](),
		),
	)

	for i := 1; i < 10; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}
	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestFlattenWithEmptiesMulti(t *testing.T) {
	t.Parallel()

	items := itertools.Flatten(
		iter.New(
			iter.New[int](),
			iter.New(1, 2, 3),
			iter.New(4, 5, 6),
			iter.New[int](),
			iter.New(7, 8, 9),
			iter.New[int](),
		),
		iter.New[iter.Iterable[int]](),
		iter.New(
			iter.New[int](),
			iter.New(1, 2, 3),
			iter.New(4, 5, 6),
			iter.New[int](),
			iter.New(7, 8, 9),
			iter.New[int](),
		),
	)

	for i := 1; i < 10; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	for i := 1; i < 10; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestFlattenEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Flatten(iter.New[iter.Iterable[int]]())

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestFlattenEmptyMulti(t *testing.T) {
	t.Parallel()

	items := itertools.Flatten(
		iter.New[iter.Iterable[int]](),
		iter.New[iter.Iterable[int]](),
		iter.New[iter.Iterable[int]](),
	)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
