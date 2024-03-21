package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleRange() {
	items := itertools.Range(0, 5)

	itertools.Print(items)
	// Output: 0, 1, 2, 3, 4
}

func ExampleRangeInc() {
	items := itertools.RangeInc(0, 5, 2)

	itertools.Print(items)
	// Output: 0, 2, 4
}

func ExampleRangeInc_reverse() {
	items := itertools.RangeInc(5, 0, -1)

	itertools.Print(items)
	// Output: 5, 4, 3, 2, 1
}

func TestRangeBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Range(1, 10)

	for i := 1; i < 10; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeNeg(t *testing.T) {
	t.Parallel()

	items := itertools.Range(-10, 10)

	for i := -10; i < 10; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeWeird(t *testing.T) {
	t.Parallel()

	items := itertools.Range(10, -10)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeBad(t *testing.T) {
	t.Parallel()

	items := itertools.Range(10, 10)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeAdvanced(t *testing.T) {
	t.Parallel()

	items := itertools.Range(10, 20)

	for i := 10; i < 20; i++ {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeIncBasic(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(1, 10, 3)

	for i := 1; i < 10; i += 3 {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeIncNeg(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(-10, 10, 2)

	for i := -10; i < 10; i += 2 {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeIncWeird(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(10, -10, 2)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeIncBad(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(10, 10, 3)

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeIncAdvanced(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(10, 20, 4)

	for i := 10; i < 20; i += 4 {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeIncOverflow(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(1, 10, 10)

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestRangeIncNegReverse(t *testing.T) {
	t.Parallel()

	items := itertools.RangeInc(10, -10, -2)

	for i := 10; i > -10; i -= 2 {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
