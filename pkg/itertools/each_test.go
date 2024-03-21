package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleEach() {
	itertools.Each(iter.New("a", "b", "c"), func(item string) {
		fmt.Println(item)
	})
	// Output:
	// a
	// b
	// c
}

func ExampleEachUntil() {
	itertools.EachUntil(iter.New(1, 2, 3, 4, 5), func(item int) bool {
		fmt.Println(item)
		return item == 3
	})
	// Output:
	// 1
	// 2
	// 3
}

func TestEachBasic(t *testing.T) {
	t.Parallel()

	target := iter.New(1, 2, 3, 4, 5, 6)
	expected := []int{1, 2, 3, 4, 5, 6}
	expectedIndex := 0

	itertools.Each(target, func(i int) {
		assert.Equal(t, expected[expectedIndex], i)
		expectedIndex++
	})

	assert.Equal(t, 6, expectedIndex)
}

func TestEachEmpty(t *testing.T) {
	t.Parallel()

	itertools.Each(iter.New[int](), func(i int) {
		assert.Fail(t, "should not have been called")
	})
}

func TestEachUntilBasic(t *testing.T) {
	t.Parallel()

	expected := []int{1, 2, 3, 4, 5, 6}
	expectedIndex := 0
	target := iter.New(1, 2, 3, 4, 5, 6)

	itertools.EachUntil(target, func(i int) bool {
		assert.Equal(t, expected[expectedIndex], i)
		expectedIndex++

		return i >= 3
	})

	assert.Equal(t, 3, expectedIndex)
}

func TestEachUntilEmpty(t *testing.T) {
	t.Parallel()

	itertools.EachUntil(iter.New[int](), func(i int) bool {
		assert.Fail(t, "should not have been called")

		return i >= 3
	})
}

func TestEachUntilNever(t *testing.T) {
	t.Parallel()

	expected := []int{1, 2, 3, 4, 5, 6}
	expectedIndex := 0
	target := iter.New(1, 2, 3, 4, 5, 6)

	itertools.EachUntil(target, func(i int) bool {
		assert.Equal(t, expected[expectedIndex], i)
		expectedIndex++

		return false
	})

	assert.Equal(t, 6, expectedIndex)
}

func TestEachUntilAlways(t *testing.T) {
	t.Parallel()

	expected := []int{1}
	expectedIndex := 0
	target := iter.New(1, 2, 3, 4, 5, 6)

	itertools.EachUntil(target, func(i int) bool {
		assert.Equal(t, expected[expectedIndex], i)
		expectedIndex++

		return true
	})

	assert.Equal(t, 1, expectedIndex)
}
