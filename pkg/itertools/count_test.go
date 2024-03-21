package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func ExampleCount() {
	result := itertools.Count(iter.New("a", "b", "c"))

	fmt.Println(result)
	// Output: 3
}

func ExampleCountIf() {
	result := itertools.CountIf(iter.New(1, 2, 3, 4, 5), predicates.Even[int])

	fmt.Println(result)
	// Output: 2
}

func TestCountIfBasic(t *testing.T) {
	t.Parallel()

	result := itertools.CountIf(iter.New(1, 2, 3, 4, 5, 6, 7), func(n int) bool {
		return n > 3
	})

	assert.Equal(t, 4, result)
}

func TestCountIfAllFalse(t *testing.T) {
	t.Parallel()

	result := itertools.CountIf(iter.New(1, 2, 3, 4, 5, 6, 7), predicates.False[int])

	assert.Equal(t, 0, result)
}

func TestCountIfEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.CountIf(iter.New[int](), predicates.True[int])

	assert.Equal(t, 0, result)
}

func TestCountBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Count(iter.New(1, 2, 3, 4, 5, 6, 7))

	assert.Equal(t, 7, result)
}

func TestCountEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.Count(iter.New[int]())

	assert.Equal(t, 0, result)
}
