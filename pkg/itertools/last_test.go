package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleLast() {
	result, _ := itertools.Last(iter.New(1, 2, 3))

	fmt.Println(result)
	// Output: 3
}

func TestLastIfBasic(t *testing.T) {
	t.Parallel()

	result, found := itertools.LastIf(iter.New(1, 2, 3, 4, 5), func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, 4, result)
	assert.True(t, found)
}

func TestLastIFBasicDNE(t *testing.T) {
	t.Parallel()

	result, found := itertools.LastIf(iter.New(1, 2, 3, 4, 5), func(_ int) bool {
		return false
	})

	assert.Equal(t, 0, result)
	assert.False(t, found)
}

func TestLastIfBasicEmpty(t *testing.T) {
	t.Parallel()

	result, found := itertools.LastIf(iter.New[int](), func(_ int) bool {
		return true
	})

	assert.Equal(t, 0, result)
	assert.False(t, found)
}

func TestLastBasic(t *testing.T) {
	t.Parallel()

	result, found := itertools.Last(iter.New(1, 2, 3, 4, 5))

	assert.Equal(t, 5, result)
	assert.True(t, found)
}

func TestLastEmpty(t *testing.T) {
	t.Parallel()

	result, found := itertools.Last(iter.New[int]())

	assert.Equal(t, 0, result)
	assert.False(t, found)
}
