package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestZipBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Zip(
		iter.NewIter(1, 2, 3, 4, 5, 6),
		iter.NewIter("a", "b", "c", "d", "e", "f"),
	)

	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 1, Second: "a"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 2, Second: "b"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 3, Second: "c"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 4, Second: "d"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 5, Second: "e"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 6, Second: "f"}, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 0, Second: ""}, items.Next())
}

func TestZipUneven(t *testing.T) {
	t.Parallel()

	items := itertools.Zip(
		iter.NewIter(1, 2, 3),
		iter.NewIter("a", "b", "c", "d", "e", "f"),
	)

	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 1, Second: "a"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 2, Second: "b"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 3, Second: "c"}, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 0, Second: ""}, items.Next())
}

func TestZipOtherUneven(t *testing.T) {
	t.Parallel()

	items := itertools.Zip(
		iter.NewIter(1, 2, 3, 4, 5, 6),
		iter.NewIter("a", "b", "c"),
	)

	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 1, Second: "a"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 2, Second: "b"}, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 3, Second: "c"}, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 0, Second: ""}, items.Next())
}

func TestZipEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Zip(
		iter.NewIter[int](),
		iter.NewIter[string](),
	)

	assert.False(t, items.HasNext())
	assert.Equal(t, iter.Tuple[int, string]{First: 0, Second: ""}, items.Next())
}
