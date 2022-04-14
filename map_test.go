package iter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
)

func TestMapGetSet(t *testing.T) {
	t.Parallel()

	m := iter.NewMap[string, int]()
	m.Set("foo", 1)
	m.Set("bar", 2)
	m.Set("qux", 3)
	m.Set("bar", 7)
	m.Set("bax", 4)

	assert.Equal(t, 1, m.Get("foo"))
	assert.Equal(t, 7, m.Get("bar"))
	assert.Equal(t, 3, m.Get("qux"))
	assert.Equal(t, 4, m.Get("bax"))
	assert.Equal(t, 4, m.Len())
}

func TestMapRemove(t *testing.T) {
	t.Parallel()

	m := iter.NewMap[string, int]()
	m.Set("foo", 1)
	m.Set("bar", 2)
	m.Set("qux", 3)
	m.Set("bax", 4)

	assert.Equal(t, 4, m.Len())
	assert.True(t, m.Has("qux"))
	m.Remove("qux")
	assert.Equal(t, 3, m.Len())
	assert.False(t, m.Has("qux"))
	m.Remove("qux")
	assert.Equal(t, 3, m.Len())
	assert.False(t, m.Has("qux"))

	assert.Equal(t, 1, m.Get("foo"))
	assert.Equal(t, 2, m.Get("bar"))
	assert.Equal(t, 0, m.Get("qux"))
	assert.Equal(t, 4, m.Get("bax"))

	assert.Equal(t, 0, m.Get("dne"))
}

func TestMapFromIter(t *testing.T) {
	t.Parallel()

	m := iter.NewMap(
		iter.NewIter(
			iter.Tuple[string, string]{"foo", "bar"},
			iter.Tuple[string, string]{"qux", "bax"},
		),
		iter.NewIter(
			iter.Tuple[string, string]{"purple", "monkey"},
		),
	)

	assert.Equal(t, "bar", m.Get("foo"))
	assert.Equal(t, "bax", m.Get("qux"))
	assert.Equal(t, "monkey", m.Get("purple"))

	assert.Equal(t, 3, m.Len())
}

func TestMapValues(t *testing.T) {
	t.Parallel()

	m := iter.NewMap[string, int]()
	m.Set("foo", 10)
	m.Set("bar", 20)
	m.Set("qux", 30)
	m.Set("delete me", 50)
	m.Set("bax", 40)
	m.Remove("delete me")

	valuesIter := m.Values()
	values := make([]int, 0, 4)

	for valuesIter.HasNext() {
		values = append(values, valuesIter.Next())
	}

	assert.ElementsMatch(t, []int{10, 20, 30, 40}, values)
}

func TestMapKeys(t *testing.T) {
	t.Parallel()

	m := iter.NewMap[string, int]()
	m.Set("foo", 10)
	m.Set("delete me", 50)
	m.Set("bar", 20)
	m.Set("qux", 30)
	m.Set("bax", 40)
	m.Remove("delete me")

	keysIter := m.Keys()
	keys := make([]string, 0, 4)

	for keysIter.HasNext() {
		keys = append(keys, keysIter.Next())
	}

	assert.ElementsMatch(t, []string{"foo", "bar", "qux", "bax"}, keys)
}

func TestMapItems(t *testing.T) {
	t.Parallel()

	m := iter.NewMap[string, int]()
	m.Set("delete me", 50)
	m.Set("foo", 10)
	m.Set("bar", 20)
	m.Set("qux", 30)
	m.Set("bax", 40)
	m.Remove("delete me")

	itemsIter := m.Items()
	items := map[string]int{}

	for itemsIter.HasNext() {
		next := itemsIter.Next()
		items[next.First] = next.Second
	}

	assert.Equal(t, map[string]int{
		"foo": 10,
		"bar": 20,
		"qux": 30,
		"bax": 40,
	}, items)
}
