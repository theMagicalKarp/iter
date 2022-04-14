package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestBatchBasic(t *testing.T) {
	t.Parallel()

	batchIter := itertools.Batch(iter.NewIter(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), 3)

	assert.True(t, batchIter.HasNext())
	batch1 := batchIter.Next()
	assert.True(t, batch1.HasNext())
	assert.Equal(t, 1, batch1.Next())
	assert.Equal(t, 2, batch1.Next())
	assert.Equal(t, 3, batch1.Next())
	assert.Equal(t, 0, batch1.Next())
	assert.False(t, batch1.HasNext())

	assert.True(t, batchIter.HasNext())
	batch2 := batchIter.Next()
	assert.True(t, batch2.HasNext())
	assert.Equal(t, 4, batch2.Next())
	assert.Equal(t, 5, batch2.Next())
	assert.Equal(t, 6, batch2.Next())
	assert.Equal(t, 0, batch2.Next())
	assert.False(t, batch1.HasNext())

	assert.True(t, batchIter.HasNext())
	batch3 := batchIter.Next()
	assert.True(t, batch3.HasNext())
	assert.Equal(t, 7, batch3.Next())
	assert.Equal(t, 8, batch3.Next())
	assert.Equal(t, 9, batch3.Next())
	assert.Equal(t, 0, batch3.Next())
	assert.False(t, batch1.HasNext())

	batch4 := batchIter.Next()
	assert.True(t, batch4.HasNext())
	assert.Equal(t, 10, batch4.Next())
	assert.Equal(t, 0, batch4.Next())
	assert.Equal(t, 0, batch4.Next())
	assert.Equal(t, 0, batch4.Next())
	assert.False(t, batch4.HasNext())

	assert.False(t, batchIter.HasNext())
	batch5 := batchIter.Next()
	assert.False(t, batch5.HasNext())
	assert.Equal(t, 0, batch5.Next())
	assert.Equal(t, 0, batch5.Next())
	assert.Equal(t, 0, batch5.Next())
	assert.Equal(t, 0, batch5.Next())
}

func TestBatchEmpty(t *testing.T) {
	t.Parallel()

	batchIter := itertools.Batch(iter.NewIter[int](), 3)

	assert.False(t, batchIter.HasNext())
	batch1 := batchIter.Next()
	assert.False(t, batch1.HasNext())
	assert.Equal(t, 0, batch1.Next())
}
