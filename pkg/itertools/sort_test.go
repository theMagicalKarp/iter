package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleSort_ascending() {
	items := itertools.Sort(iter.New(2, 1, 5, 4, 3), itertools.Asc[int])

	itertools.Print(items)
	// Output: 1, 2, 3, 4, 5
}

func ExampleSort_descending() {
	items := itertools.Sort(iter.New(2, 1, 5, 4, 3), itertools.Desc[int])

	itertools.Print(items)
	// Output: 5, 4, 3, 2, 1
}

func ExampleSort() {
	items := itertools.Sort(iter.New(2, 1, 5, 4, 3), func(a, b int) int {
		return a - b
	})

	itertools.Print(items)
	// Output: 1, 2, 3, 4, 5
}

func TestSortAsc(t *testing.T) {
	t.Parallel()
	// Create a sortIter with the test list
	iter := itertools.Sort(iter.New(5, 2, 7, 1, 9), func(a, b int) int {
		return a - b
	})

	// Call Next() to get the sorted elements
	sortedList := []int{}
	for {
		elem, ok := iter.Next()
		if !ok {
			break
		}
		sortedList = append(sortedList, elem)
	}

	elemt, more := iter.Next()
	assert.False(t, more)
	assert.Equal(t, elemt, 0)

	elemt, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, elemt, 0)

	// Assert that the sorted list matches the expected result
	expected := []int{1, 2, 5, 7, 9}
	assert.Equal(t, expected, sortedList)
}

func TestSortDesc(t *testing.T) {
	t.Parallel()
	// Create a sortIter with the test list
	iter := itertools.Sort(iter.New(5, 2, 7, 1, 9), func(a, b int) int {
		return b - a
	})

	// Call Next() to get the sorted elements
	sortedList := []int{}
	for {
		elem, ok := iter.Next()
		if !ok {
			break
		}
		sortedList = append(sortedList, elem)
	}

	elemt, more := iter.Next()
	assert.False(t, more)
	assert.Equal(t, elemt, 0)

	elemt, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, elemt, 0)

	// Assert that the sorted list matches the expected result
	expected := []int{9, 7, 5, 2, 1}
	assert.Equal(t, expected, sortedList)
}

func TestSortEmpty(t *testing.T) {
	t.Parallel()
	// Create a sortIter with the test list
	iter := itertools.Sort(iter.New[int](), func(a, b int) int {
		return a - b
	})

	// Call Next() to get the sorted elements
	sortedList := []int{}
	for {
		elem, ok := iter.Next()
		if !ok {
			break
		}
		sortedList = append(sortedList, elem)
	}

	elemt, more := iter.Next()
	assert.False(t, more)
	assert.Equal(t, elemt, 0)

	elemt, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, elemt, 0)

	// Assert that the sorted list matches the expected result
	expected := []int{}
	assert.Equal(t, expected, sortedList)
}

func TestAsc(t *testing.T) {
	testCases := []struct {
		this, that, expected int
	}{
		{1, 2, -1},
		{2, 1, 1},
		{0, 0, 0},
		{-1, -1, 0},
		{-1, 1, -2},
		{100, 100, 0},
		{-100, 100, -200},
		{100, -100, 200},
	}

	for _, tc := range testCases {
		result := itertools.Asc(tc.this, tc.that)
		if result != tc.expected {
			t.Errorf("Unexpected result for Asc(%d, %d): got %d, want %d", tc.this, tc.that, result, tc.expected)
		}
	}
}

func TestDesc(t *testing.T) {
	testCases := []struct {
		this, that, expected int
	}{
		{1, 2, 1},
		{2, 1, -1},
		{0, 0, 0},
		{-1, -1, 0},
		{-1, 1, 2},
		{100, 100, 0},
		{-100, 100, 200},
		{100, -100, -200},
	}

	for _, tc := range testCases {
		result := itertools.Desc(tc.this, tc.that)
		if result != tc.expected {
			t.Errorf("Unexpected result for Desc(%d, %d): got %d, want %d", tc.this, tc.that, result, tc.expected)
		}
	}
}
