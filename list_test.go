package iter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
)

func TestListBasic(t *testing.T) {
	t.Parallel()

	l := iter.NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushFront(4)
	l.PushBack(3)

	items := l.Iter()
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestListBasicReverse(t *testing.T) {
	t.Parallel()

	l := iter.NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushFront(4)
	l.PushBack(3)

	items := l.Reverse()
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestListMoveFrontNode(t *testing.T) {
	t.Parallel()

	l := iter.NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	moveMe := l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushNodeFront(moveMe)

	items := l.Iter()
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 5, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestListRemoves(t *testing.T) {
	t.Parallel()

	l := iter.NewList[int]()
	first := l.PushBack(1)
	l.PushBack(2)
	mid := l.PushBack(3)
	l.PushBack(4)
	last := l.PushBack(5)

	l.Remove(first)
	l.Remove(mid)
	l.Remove(last)

	items := l.Iter()
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestListPop(t *testing.T) {
	t.Parallel()

	l := iter.NewList[int]()
	assert.Equal(t, 0, l.PopFront())
	assert.Equal(t, 0, l.PopBack())

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	assert.Equal(t, 1, l.PopFront())
	assert.Equal(t, 5, l.PopBack())

	items := l.Iter()
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestListComplicated(t *testing.T) {
	t.Parallel()

	l := iter.NewList[int]()
	l.PushFront(1)
	toMove := l.PushFront(55)
	l.PushBack(4)
	toRemove := l.PushBack(64)
	l.PushBack(5)
	l.PushFront(42)
	l.PushFront(5)
	l.PushNodeFront(toMove)
	l.Remove(toRemove)
	l.Remove(toRemove)

	items := l.Iter()
	assert.True(t, items.HasNext())
	assert.Equal(t, 55, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 5, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 42, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 5, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}

func TestListFromiter(t *testing.T) {
	t.Parallel()

	l := iter.NewList(
		iter.NewIter(1, 2, 3),
		iter.NewIter(4, 5, 6),
	)

	items := l.Iter()
	assert.True(t, items.HasNext())
	assert.Equal(t, 1, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 2, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 3, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 4, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 5, items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, 6, items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, 0, items.Next())
}
