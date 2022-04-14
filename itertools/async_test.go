package itertools_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestAsyncBasic(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.NewIter(1, 2, 3, 4, 5, 6), 3,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.ElementsMatch(t, []string{"1", "2", "3", "4", "5", "6"}, items)
}

func TestAsyncEmpty(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.NewIter[int](), 3,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.ElementsMatch(t, []string{}, items)
}

func TestAsyncNegativeWorker(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.NewIter(1, 2, 3, 4, 5, 6), -3,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.ElementsMatch(t, []string{}, items)
}

func TestAsyncSingleWorker(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.NewIter(1, 2, 3, 4, 5, 6), 1,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.ElementsMatch(t, []string{"1", "2", "3", "4", "5", "6"}, items)
}

func TestAsyncDeadContext(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	itemsIter := itertools.Async(ctx, iter.NewIter(1, 2, 3, 4, 5, 6), 1,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	assert.ElementsMatch(t, []string{}, items)
}

func TestAsyncDryPull(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.NewIter[int](), 3,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	assert.Equal(t, "", itemsIter.Next())
}

func TestAsyncStressTest(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), itertools.Range(0, 10000), 16,
		func(ctx context.Context, i int) int {
			return -i
		})

	items := make([]int, 0, 10000)

	for itemsIter.HasNext() {
		items = append(items, itemsIter.Next())
	}

	expected := make([]int, 0, 10000)

	for i := 0; i < 10000; i++ {
		expected = append(expected, -i)
	}

	assert.ElementsMatch(t, expected, items)
}
