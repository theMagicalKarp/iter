package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleBatch() {
	result := itertools.Batch(iter.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), 3)

	itertools.Each(result, func(batch iter.Iterable[int]) {
		itertools.Print(batch)
	})

	// Output:
	// 1, 2, 3
	// 4, 5, 6
	// 7, 8, 9
	// 10
}

func TestBatchBasic(t *testing.T) {
	t.Parallel()

	batchIter := itertools.Batch(iter.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), 3)

	batch1, more := batchIter.Next()
	assert.True(t, more)

	v, more := batch1.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)
	v, more = batch1.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)
	v, more = batch1.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)
	v, more = batch1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = batch1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	batch2, more := batchIter.Next()
	assert.True(t, more)

	v, more = batch2.Next()
	assert.True(t, more)
	assert.Equal(t, 4, v)
	v, more = batch2.Next()
	assert.True(t, more)
	assert.Equal(t, 5, v)
	v, more = batch2.Next()
	assert.True(t, more)
	assert.Equal(t, 6, v)
	v, more = batch2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = batch2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	batch3, more := batchIter.Next()
	assert.True(t, more)

	v, more = batch3.Next()
	assert.True(t, more)
	assert.Equal(t, 7, v)
	v, more = batch3.Next()
	assert.True(t, more)
	assert.Equal(t, 8, v)
	v, more = batch3.Next()
	assert.True(t, more)
	assert.Equal(t, 9, v)
	v, more = batch3.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = batch3.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	batch4, more := batchIter.Next()
	assert.True(t, more)

	v, more = batch4.Next()
	assert.True(t, more)
	assert.Equal(t, 10, v)
	v, more = batch4.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = batch4.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	batch5, more := batchIter.Next()
	assert.False(t, more)
	assert.Nil(t, batch5)

	batch5, more = batchIter.Next()
	assert.False(t, more)
	assert.Nil(t, batch5)
}

func TestBatchExact(t *testing.T) {
	t.Parallel()

	batchIter := itertools.Batch(iter.New[int](1, 2, 3, 4, 5, 6), 3)

	batch1, more := batchIter.Next()
	assert.True(t, more)

	v, more := batch1.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)
	v, more = batch1.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)
	v, more = batch1.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)
	v, more = batch1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = batch1.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	batch2, more := batchIter.Next()
	assert.True(t, more)

	v, more = batch2.Next()
	assert.True(t, more)
	assert.Equal(t, 4, v)
	v, more = batch2.Next()
	assert.True(t, more)
	assert.Equal(t, 5, v)
	v, more = batch2.Next()
	assert.True(t, more)
	assert.Equal(t, 6, v)
	v, more = batch2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
	v, more = batch2.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	batch3, more := batchIter.Next()
	assert.False(t, more)
	assert.Nil(t, batch3)

	batch3, more = batchIter.Next()
	assert.False(t, more)
	assert.Nil(t, batch3)
}

func TestBatchEmpty(t *testing.T) {
	t.Parallel()

	batchIter := itertools.Batch(iter.New[int](), 3)

	batch1, more := batchIter.Next()
	assert.False(t, more)
	assert.Nil(t, batch1)

	batch1, more = batchIter.Next()
	assert.False(t, more)
	assert.Nil(t, batch1)
}

func TestBatchZeroSize(t *testing.T) {
	t.Parallel()

	batchIter := itertools.Batch(iter.New[int](), 0)

	batch1, more := batchIter.Next()
	assert.False(t, more)
	assert.Nil(t, batch1)

	batch1, more = batchIter.Next()
	assert.False(t, more)
	assert.Nil(t, batch1)
}
