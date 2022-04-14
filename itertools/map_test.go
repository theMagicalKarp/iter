package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestMapBasic(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Map(iter.NewIter(1, 2, 3, 4, 5, 6),
		func(i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, items)
	assert.Equal(t, "0", itemsIter.Next())
}

func TestMapBasicEmpty(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Map(iter.NewIter[int](),
		func(i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.Equal(t, []string{}, items)
	assert.Equal(t, "0", itemsIter.Next())
}
