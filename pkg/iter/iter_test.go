package iter_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
)

func Example() {
	items := iter.New("hello", "world", "!!!")

	for item, ok := items.Next(); ok; item, ok = items.Next() {
		fmt.Println(item)
	}
	// Output:
	// hello
	// world
	// !!!
}

func TestNewIter(t *testing.T) {
	t.Parallel()

	items := iter.New(1, 2, 3)

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestNewIterEmpty(t *testing.T) {
	t.Parallel()

	items := iter.New[int]()
	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
