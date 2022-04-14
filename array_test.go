package iter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
)

func TestArrayAtSet(t *testing.T) {
	t.Parallel()

	a := iter.NewArray(10, 10, iter.NewIter(1, 2, 3))
	assert.Equal(t, 13, a.Len())

	for i := 0; i < 10; i++ {
		assert.Equal(t, 0, a.At(i))
	}

	assert.Equal(t, 1, a.At(10))
	assert.Equal(t, 2, a.At(11))
	assert.Equal(t, 3, a.At(12))

	a.Set(3, 42)
	assert.Equal(t, 42, a.At(3))
}

func TestArrayIter(t *testing.T) {
	t.Parallel()

	a := iter.NewArray(0, 10, iter.NewIter(11, 22, 33, 44, 55))

	itemsIter := a.Iter()
	items := make([]int, 0, 5)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.Equal(t, []int{11, 22, 33, 44, 55}, items)
}

func TestArrayReverseIter(t *testing.T) {
	t.Parallel()

	a := iter.NewArray(0, 10, iter.NewIter(11, 22, 33, 44, 55))

	itemsIter := a.Reverse()
	items := make([]int, 0, 5)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.Equal(t, []int{55, 44, 33, 22, 11}, items)
}

func TestArrayRangeIter(t *testing.T) {
	t.Parallel()

	a := iter.NewArray(0, 10, iter.NewIter(11, 22, 33, 44, 55))

	itemsIter := a.Range(1, 4)
	items := make([]int, 0, 5)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.Equal(t, []int{22, 33, 44}, items)
}

func TestArrayTailIter(t *testing.T) {
	t.Parallel()

	a := iter.NewArray(0, 10, iter.NewIter(11, 22, 33, 44, 55))

	itemsIter := a.Tail(2)
	items := make([]int, 0, 5)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.Equal(t, []int{33, 44, 55}, items)
}

func TestArrayHeadIter(t *testing.T) {
	t.Parallel()

	a := iter.NewArray(0, 10, iter.NewIter(11, 22, 33, 44, 55))

	itemsIter := a.Head(2)
	items := make([]int, 0, 5)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.Equal(t, []int{11, 22}, items)
}

func TestArrayFromIters(t *testing.T) {
	t.Parallel()

	a := iter.NewArray(
		0, 5,
		iter.NewIter(11, 22, 33, 44, 55),
		iter.NewIter(66, 77, 88, 99, 110),
	)

	assert.Equal(t, 10, a.Len())
	assert.Equal(t, 11, a.At(0))
	assert.Equal(t, 22, a.At(1))
	assert.Equal(t, 33, a.At(2))
	assert.Equal(t, 44, a.At(3))
	assert.Equal(t, 55, a.At(4))
	assert.Equal(t, 66, a.At(5))
	assert.Equal(t, 77, a.At(6))
	assert.Equal(t, 88, a.At(7))
	assert.Equal(t, 99, a.At(8))
	assert.Equal(t, 110, a.At(9))
}

func TestArrayAppend(t *testing.T) {
	t.Parallel()

	a := iter.NewArray[int](0, 0)
	assert.Equal(t, 0, a.Len())
	a.Append(11)
	assert.Equal(t, 1, a.Len())
	assert.Equal(t, 11, a.At(0))

	a.Append(22)
	assert.Equal(t, 2, a.Len())
	assert.Equal(t, 22, a.At(1))
}

func TestArrayPop(t *testing.T) {
	t.Parallel()

	a := iter.NewArray(0, 10, iter.NewIter(11, 22, 33, 44, 55))

	assert.Equal(t, 5, a.Len())
	assert.Equal(t, 55, a.Pop())
	assert.Equal(t, 4, a.Len())

	assert.Equal(t, 4, a.Len())
	assert.Equal(t, 44, a.Pop())
	assert.Equal(t, 3, a.Len())
}
