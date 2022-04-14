package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestMixBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Mix(
		iter.NewIter(1, 2, 3, 4, 5),
		iter.NewIter(11, 22, 33, 44, 55, 66, 77),
		iter.NewIter[int](),
		iter.NewIter(111, 222),
	)

	expected := []int{
		1, 11, 111,
		2, 22, 222,
		3, 33,
		4, 44,
		5, 55,
		66,
		77,
	}

	for _, exp := range expected {
		assert.True(t, items.HasNext())
		assert.Equal(t, exp, items.Next())
	}

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestMixEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Mix[int]()

	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}
