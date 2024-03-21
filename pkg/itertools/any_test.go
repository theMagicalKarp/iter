package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleAny_true() {
	result := itertools.Any(iter.New(1, 2, 3), func(n int) bool {
		return n == 1
	})

	fmt.Println(result)
	// Output: true
}

func ExampleAny_false() {
	result := itertools.All(iter.New(1, 2, 3), func(n int) bool {
		return n == 0
	})

	fmt.Println(result)
	// Output: false
}

func TestAnyTrue(t *testing.T) {
	t.Parallel()

	result := itertools.Any(iter.New(1, 2, 3, 4, 5, 6, 7), func(n int) bool {
		return n > 0
	})

	assert.True(t, result)
}

func TestAnySomeTrue(t *testing.T) {
	t.Parallel()

	result := itertools.Any(iter.New(1, 2, 3, 4, -5, 6, 7), func(n int) bool {
		return n > 0
	})

	assert.True(t, result)
}

func TestAnyNoTrue(t *testing.T) {
	t.Parallel()

	result := itertools.Any(iter.New(-1, -2, -3, -4, -5, -6, -7), func(n int) bool {
		return n > 0
	})

	assert.False(t, result)
}

func TestAnyEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.Any(iter.New[int](), func(n int) bool {
		return n > 0
	})

	assert.False(t, result)
}
