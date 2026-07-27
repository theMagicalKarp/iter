package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleSlice() {
	items := itertools.Slice(iter.New(1, 2, 3, 4, 5))

	fmt.Println(items)
	// Output: [1 2 3 4 5]
}

func ExampleSliceN() {
	items := itertools.SliceN(iter.New(1, 2, 3, 4, 5), 1, 4)

	fmt.Println(items)
	// Output: [2 3 4]
}

func TestBasicSlice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{10, 20, 30}, []int{10, 20, 30}},
		{[]int{100}, []int{100}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		result := itertools.Slice(iter.New(tc.input...))
		assert.ElementsMatch(t, tc.expected, result)
	}
}

func TestBasicSliceN(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input    []int
		start    int
		stop     int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, 4, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 2, 3, []int{3}},
		{[]int{1, 2, 3, 4, 5}, 5, 10, []int{}},
		{[]int{}, 5, 10, []int{}},
		{[]int{}, 0, 5, []int{}},
	}

	for _, tc := range testCases {
		result := itertools.SliceN(iter.New(tc.input...), tc.start, tc.stop)
		assert.ElementsMatch(t, tc.expected, result)
	}
}
