package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func ExamplePartition() {
	evenItems, oddItems := itertools.Partition(iter.New(1, 2, 3, 4, 5, 6), predicates.Even[int])

	itertools.Print(evenItems)
	itertools.Print(oddItems)
	// Output:
	// 2, 4, 6
	// 1, 3, 5
}

func TestPartitionBasic(t *testing.T) {
	t.Parallel()

	items1, items2 := itertools.Partition(iter.New(1, 2, 3, 4, 5, 6), func(i int) bool {
		return i%2 == 0
	})

	v, more := items1.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)

	v, more = items1.Next()
	assert.True(t, more)
	assert.Equal(t, 4, v)

	v, more = items1.Next()
	assert.True(t, more)
	assert.Equal(t, 6, v)

	v, more = items1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items2.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)

	v, more = items2.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)

	v, more = items2.Next()
	assert.True(t, more)
	assert.Equal(t, 5, v)

	v, more = items2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestPartitionLopsided(t *testing.T) {
	t.Parallel()

	items1, items2 := itertools.Partition(iter.New(1, 2, 3, 4, 5, 6), func(_ int) bool {
		return true
	})

	for i := 1; i <= 6; i++ {
		v, more := items1.Next()
		assert.True(t, more)
		assert.Equal(t, i, v)
	}

	v, more := items1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestPartitionEmpty(t *testing.T) {
	t.Parallel()

	items1, items2 := itertools.Partition(iter.New[int](), func(i int) bool {
		return i%2 == 0
	})
	v, more := items1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
