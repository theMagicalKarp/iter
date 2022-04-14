package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestEachBasic(t *testing.T) {
	t.Parallel()

	target := iter.NewIter(1, 2, 3, 4, 5, 6)
	expected := []int{1, 2, 3, 4, 5, 6}
	expectedIndex := 0

	itertools.Each(target, func(i int) {
		assert.Equal(t, expected[expectedIndex], i)
		expectedIndex++
	})

	assert.Equal(t, 6, expectedIndex)
	assert.False(t, target.HasNext())
}

func TestEachEmpty(t *testing.T) {
	t.Parallel()

	itertools.Each(iter.NewIter[int](), func(i int) {
		assert.Fail(t, "should not have been called")
	})
}

func TestEachUntilBasic(t *testing.T) {
	t.Parallel()

	expected := []int{1, 2, 3, 4, 5, 6}
	expectedIndex := 0
	target := iter.NewIter(1, 2, 3, 4, 5, 6)

	itertools.EachUntil(target, func(i int) bool {
		assert.Equal(t, expected[expectedIndex], i)
		expectedIndex++

		return i >= 3
	})

	assert.Equal(t, 3, expectedIndex)
	assert.True(t, target.HasNext())
}

func TestEachUntilEmpty(t *testing.T) {
	t.Parallel()

	itertools.EachUntil(iter.NewIter[int](), func(i int) bool {
		assert.Fail(t, "should not have been called")

		return i >= 3
	})
}

func TestEachUntilNever(t *testing.T) {
	t.Parallel()

	expected := []int{1, 2, 3, 4, 5, 6}
	expectedIndex := 0
	target := iter.NewIter(1, 2, 3, 4, 5, 6)

	itertools.EachUntil(target, func(i int) bool {
		assert.Equal(t, expected[expectedIndex], i)
		expectedIndex++

		return false
	})

	assert.Equal(t, 6, expectedIndex)
	assert.False(t, target.HasNext())
}

func TestEachUntilAlways(t *testing.T) {
	t.Parallel()

	expected := []int{1}
	expectedIndex := 0
	target := iter.NewIter(1, 2, 3, 4, 5, 6)

	itertools.EachUntil(target, func(i int) bool {
		assert.Equal(t, expected[expectedIndex], i)
		expectedIndex++

		return true
	})

	assert.Equal(t, 1, expectedIndex)
	assert.True(t, target.HasNext())
}
