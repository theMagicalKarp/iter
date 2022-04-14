package iter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
)

func TestSetBasic(t *testing.T) {
	t.Parallel()

	s := iter.NewSet[string]()
	s.Add("foo")
	s.Add("bar")
	s.Add("qux")
	s.Add("bax")

	assert.True(t, s.Has("foo"))
	assert.True(t, s.Has("bar"))
	assert.True(t, s.Has("qux"))
	assert.True(t, s.Has("bax"))
	assert.Equal(t, 4, s.Len())
}

func TestSetRemove(t *testing.T) {
	t.Parallel()

	s := iter.NewSet[string]()
	s.Add("foo")
	s.Add("bar")
	s.Add("qux")
	s.Add("bax")

	s.Remove("qux")
	s.Remove("qux")

	assert.True(t, s.Has("foo"))
	assert.True(t, s.Has("bar"))
	assert.False(t, s.Has("qux"))
	assert.True(t, s.Has("bax"))
	assert.Equal(t, 3, s.Len())
}

func TestSetSub(t *testing.T) {
	t.Parallel()

	s1 := iter.NewSet[string]()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("qux")
	s1.Add("bax")

	s2 := iter.NewSet[string]()
	s2.Add("qux")

	s3 := s1.Sub(s2)

	assert.True(t, s3.Has("foo"))
	assert.True(t, s3.Has("bar"))
	assert.False(t, s3.Has("qux"))
	assert.True(t, s3.Has("bax"))
	assert.Equal(t, 3, s3.Len())
}

func TestSetUnion(t *testing.T) {
	t.Parallel()

	s1 := iter.NewSet[string]()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("qux")
	s1.Add("bax")

	s2 := iter.NewSet[string]()
	s2.Add("purple")

	s3 := s1.Union(s2)

	assert.True(t, s3.Has("foo"))
	assert.True(t, s3.Has("bar"))
	assert.True(t, s3.Has("qux"))
	assert.True(t, s3.Has("bax"))
	assert.True(t, s3.Has("purple"))
	assert.Equal(t, 5, s3.Len())
}

func TestSetIntersection(t *testing.T) {
	t.Parallel()

	s1 := iter.NewSet[string]()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("qux")
	s1.Add("bax")

	s2 := iter.NewSet[string]()
	s2.Add("bar")

	s3 := s1.Intersection(s2)

	assert.False(t, s3.Has("foo"))
	assert.True(t, s3.Has("bar"))
	assert.False(t, s3.Has("qux"))
	assert.False(t, s3.Has("bax"))
	assert.Equal(t, 1, s3.Len())
}

func TestSetIter(t *testing.T) {
	t.Parallel()

	s := iter.NewSet[string]()
	s.Add("foo")
	s.Add("bar")
	s.Add("qux")
	s.Add("bax")

	itemIter := s.Iter()
	items := make([]string, 0, 4)

	for itemIter.HasNext() {
		items = append(items, itemIter.Next())
	}

	assert.ElementsMatch(t, []string{"foo", "bar", "qux", "bax"}, items)
}

func TestSetFromIter(t *testing.T) {
	t.Parallel()

	s := iter.NewSet(
		iter.NewIter("foo", "bar", "qux"),
		iter.NewIter("foo", "bax"),
	)

	assert.True(t, s.Has("foo"))
	assert.True(t, s.Has("bar"))
	assert.True(t, s.Has("qux"))
	assert.True(t, s.Has("bax"))
	assert.Equal(t, 4, s.Len())
}
