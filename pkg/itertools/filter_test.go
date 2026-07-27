package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func ExampleFilter() {
	result := itertools.Filter(iter.New(1, 2, 3, 4, 5, 6), predicates.Even[int])

	itertools.Print(result)
	// Output: 2, 4, 6
}

func TestFilterBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), func(i int) bool {
		return i%2 == 0
	})

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 4, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestFilterEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Filter(iter.New[int](), func(_ int) bool {
		return true
	})

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}

func TestFilterNoMatch(t *testing.T) {
	t.Parallel()

	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), func(_ int) bool {
		return false
	})

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)
}
