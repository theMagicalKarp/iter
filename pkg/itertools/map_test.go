package itertools_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleMap() {
	items := itertools.Map(iter.New(1, 2, 3),
		func(i int) int {
			return i * 2
		})

	itertools.Print(items)
	// Output: 2, 4, 6
}

func TestMapBasic(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Map(iter.New(1, 2, 3, 4, 5, 6), strconv.Itoa)

	items := make([]string, 0, 6)

	for {
		item, more := itemsIter.Next()
		if !more {
			break
		}

		items = append(items, item)
	}

	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, items)

	item, more := itemsIter.Next()
	assert.False(t, more)
	assert.Empty(t, item)
}

func TestMapBasicEmpty(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Map(iter.New[int](), strconv.Itoa)

	items := make([]string, 0, 6)

	for {
		item, more := itemsIter.Next()
		if !more {
			break
		}

		items = append(items, item)
	}

	assert.Equal(t, []string{}, items)

	item, more := itemsIter.Next()
	assert.False(t, more)
	assert.Empty(t, item)
}
