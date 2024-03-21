package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleSubtract() {
	first := iter.New(1, 2, 3, 4, 5)
	second := iter.New(3, 4, 5, 6, 7)
	items := itertools.Subtract(first, second)

	itertools.Println(items)
	// Unordered output:
	// 1
	// 2
}

func TestBasicSubtract(t *testing.T) {
	t.Parallel()

	// Create the first iterable
	first := iter.New(1, 2, 3, 4, 5)

	// Create the second iterable
	second := iter.New(3, 4, 5, 6, 7)

	// Call the Subtract function
	result := itertools.Slice(itertools.Subtract(first, second))

	expected := []int{1, 2}

	// Assert that the result set is equal to the expected set
	assert.ElementsMatch(t, expected, result)
}

func TestEmptyFirstSubtract(t *testing.T) {
	t.Parallel()

	// Create the first iterable
	first := iter.New[int]()

	// Create the second iterable
	second := iter.New(3, 4, 5, 6, 7)

	// Call the Subtract function
	result := itertools.Slice(itertools.Subtract(first, second))

	expected := []int{}

	// Assert that the result set is equal to the expected set
	assert.ElementsMatch(t, expected, result)
}

func TestEmptySecondSubtract(t *testing.T) {
	t.Parallel()

	// Create the first iterable
	first := iter.New(1, 2, 3, 4, 5)

	// Create the second iterable
	second := iter.New[int]()

	// Call the Subtract function
	result := itertools.Slice(itertools.Subtract(first, second))

	expected := []int{1, 2, 3, 4, 5}

	// Assert that the result set is equal to the expected set
	assert.ElementsMatch(t, expected, result)
}

func TestBothEmptySubtract(t *testing.T) {
	t.Parallel()

	// Create the first iterable
	first := iter.New[int]()

	// Create the second iterable
	second := iter.New[int]()

	// Call the Subtract function
	result := itertools.Slice(itertools.Subtract(first, second))

	expected := []int{}

	// Assert that the result set is equal to the expected set
	assert.ElementsMatch(t, expected, result)
}
