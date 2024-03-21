package predicates_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func ExamplePositive() {
	items := itertools.Filter(iter.New(1, -1, 0, -4, 5), predicates.Positive[int])

	itertools.Print(items)
	// Output: 1, 5
}

func ExampleNegative() {
	items := itertools.Filter(iter.New(1, -1, 0, -4, 5), predicates.Negative[int])

	itertools.Print(items)
	// Output: -1, -4
}

func ExampleEven() {
	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), predicates.Even[int])

	itertools.Print(items)
	// Output: 2, 4
}

func ExampleOdd() {
	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), predicates.Odd[int])

	itertools.Print(items)
	// Output: 1, 3, 5
}

func ExampleDivisible() {
	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), predicates.Divisible(2))

	itertools.Print(items)
	// Output: 2, 4
}

func TestDivisible(t *testing.T) {
	t.Parallel()

	// Test cases
	testCases := []struct {
		divisor int
		number  int
		result  bool
	}{
		{2, 4, true},
		{3, 9, true},
		{5, 7, false},
	}

	// Run tests
	for _, tc := range testCases {
		predicate := predicates.Divisible(tc.divisor)
		assert.Equal(t, tc.result, predicate(tc.number))
	}
}

func TestEven(t *testing.T) {
	t.Parallel()

	// Test cases
	testCases := []struct {
		input  int
		result bool
	}{
		{4, true},
		{7, false},
		{10, true},
	}

	// Run tests
	for _, tc := range testCases {
		assert.Equal(t, tc.result, predicates.Even(tc.input))
	}
}

func TestPositive(t *testing.T) {
	t.Parallel()

	// Test cases
	testCases := []struct {
		input  int
		result bool
	}{
		{4, true},
		{-7, false},
		{0, false},
	}

	// Run tests
	for _, tc := range testCases {
		assert.Equal(t, tc.result, predicates.Positive(tc.input))
	}
}

func TestNegative(t *testing.T) {
	t.Parallel()

	// Test cases
	testCases := []struct {
		input  int
		result bool
	}{
		{-5, true},
		{0, false},
		{10, false},
	}

	// Run tests
	for _, tc := range testCases {
		assert.Equal(t, tc.result, predicates.Negative(tc.input))
	}
}

func TestOdd(t *testing.T) {
	t.Parallel()

	// Test cases
	testCases := []struct {
		input  int
		result bool
	}{
		{3, true},
		{6, false},
		{9, true},
	}

	// Run tests
	for _, tc := range testCases {
		assert.Equal(t, tc.result, predicates.Odd(tc.input))
	}
}
