package itertools_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleAsync() {
	items := itertools.Async(context.Background(), iter.New(1, 2, 3), 2,
		func(ctx context.Context, x int) int {
			return x * 2
		})

	itertools.Println(items)
	// Unordered output:
	// 2
	// 4
	// 6
}

func TestAsyncBasic(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.New(1, 2, 3, 4, 5, 6), 3,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for {
		item, more := itemsIter.Next()
		if !more {
			break
		}
		items = append(items, item)
	}

	assert.ElementsMatch(t, []string{"1", "2", "3", "4", "5", "6"}, items)
}

func TestAsyncEmpty(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.New[int](), 3,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for {
		item, more := itemsIter.Next()
		if !more {
			break
		}
		items = append(items, item)
	}

	assert.ElementsMatch(t, []string{}, items)
}

func TestAsyncNegativeWorker(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.New(1, 2, 3, 4, 5, 6), -3,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for {
		item, more := itemsIter.Next()
		if !more {
			break
		}
		items = append(items, item)
	}

	assert.ElementsMatch(t, []string{}, items)
}

func TestAsyncSingleWorker(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.New(1, 2, 3, 4, 5, 6), 1,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for {
		item, more := itemsIter.Next()
		if !more {
			break
		}
		items = append(items, item)
	}

	assert.ElementsMatch(t, []string{"1", "2", "3", "4", "5", "6"}, items)
}

func TestAsyncDeadContext(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	itemsIter := itertools.Async(ctx, iter.New(1, 2, 3, 4, 5, 6), 6,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	items := make([]string, 0, 6)

	for {
		item, more := itemsIter.Next()
		if !more {
			break
		}
		items = append(items, item)
	}

	assert.ElementsMatch(t, []string{}, items, fmt.Sprintf("%v", items))
}

func TestAsyncDryPull(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), iter.New[int](), 3,
		func(ctx context.Context, i int) string {
			return fmt.Sprintf("%d", i)
		})

	v, more := itemsIter.Next()
	assert.Equal(t, v, "")
	assert.Equal(t, more, false)
}

func TestAsyncStressTest(t *testing.T) {
	t.Parallel()

	itemsIter := itertools.Async(context.TODO(), itertools.Range(0, 10000), 16,
		func(ctx context.Context, i int) int {
			return -i
		})

	items := make([]int, 0, 10000)

	for {
		item, more := itemsIter.Next()
		if !more {
			break
		}
		items = append(items, item)
	}

	expected := make([]int, 0, 10000)

	for i := 0; i < 10000; i++ {
		expected = append(expected, -i)
	}

	assert.ElementsMatch(t, expected, items)
}
