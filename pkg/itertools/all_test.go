package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleAll_true() {
	result := itertools.All(iter.New(1, 2, 3), func(n int) bool {
		return n != 0
	})

	fmt.Println(result)
	// Output: true
}

func ExampleAll_false() {
	result := itertools.All(iter.New(1, 2, 3), func(n int) bool {
		return n == 1
	})

	fmt.Println(result)
	// Output: false
}

func TestAllTrue(t *testing.T) {
	t.Parallel()

	result := itertools.All(iter.New(1, 2, 3, 4, 5, 6, 7), func(n int) bool {
		return n > 0
	})

	assert.True(t, result)
}

func TestAllSomeTrue(t *testing.T) {
	t.Parallel()

	result := itertools.All(iter.New(1, 2, 3, 4, -5, 6, 7), func(n int) bool {
		return n > 0
	})

	assert.False(t, result)
}

func TestAllNoTrue(t *testing.T) {
	t.Parallel()

	result := itertools.All(iter.New(-1, -2, -3, -4, -5, -6, -7), func(n int) bool {
		return n > 0
	})

	assert.False(t, result)
}

func TestAllEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.All(iter.New[int](), func(n int) bool {
		return n > 0
	})

	assert.True(t, result)
}
