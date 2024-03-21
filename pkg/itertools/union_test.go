package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleUnion() {
	items := itertools.Union(
		iter.New(1, 2, 3, 4),
		iter.New(3, 4, 5, 6),
	)

	itertools.Println(items)
	// Unordered output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func TestBasicUnion(t *testing.T) {
	t.Parallel()

	// Create the first iterable
	first := iter.New(1, 2, 3, 4)

	// Create the second iterable
	second := iter.New(3, 4, 5, 6)

	// Call the Union function
	result := itertools.Slice(itertools.Union(first, second))

	expected := []int{1, 2, 3, 4, 5, 6}
	assert.ElementsMatch(t, expected, result)
}

func TestFirstEmpty(t *testing.T) {
	t.Parallel()

	// Create the first iterable
	first := iter.New[int]()

	// Create the second iterable
	second := iter.New(3, 4, 5, 6)

	// Call the Union function
	result := itertools.Slice(itertools.Union(first, second))

	expected := []int{3, 4, 5, 6}
	assert.ElementsMatch(t, expected, result)
}

func TestSecondEmpty(t *testing.T) {
	t.Parallel()

	// Create the first iterable
	first := iter.New(1, 2, 3, 4)

	// Create the second iterable
	second := iter.New[int]()

	// Call the Union function
	result := itertools.Slice(itertools.Union(first, second))

	expected := []int{1, 2, 3, 4}
	assert.ElementsMatch(t, expected, result)
}

func TestBothEmpty(t *testing.T) {
	t.Parallel()

	// Create the first iterable
	first := iter.New[int]()

	// Create the second iterable
	second := iter.New[int]()

	// Call the Union function
	result := itertools.Slice(itertools.Union(first, second))

	expected := []int{}
	assert.ElementsMatch(t, expected, result)
}
