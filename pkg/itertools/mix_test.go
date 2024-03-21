package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleMix() {
	items := itertools.Mix(
		iter.New(1, 2, 3),
		iter.New(11, 22, 33, 44),
		iter.New(111, 222),
	)

	itertools.Print(items)
	// Output:
	// 1, 11, 111, 2, 22, 222, 3, 33, 44
}

func TestMixBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Mix(
		iter.New(1, 2, 3, 4, 5),
		iter.New(11, 22, 33, 44, 55, 66, 77),
		iter.New[int](),
		iter.New(111, 222),
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
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, exp, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestMixEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Mix[int]()

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
