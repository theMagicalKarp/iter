package predicates_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func Example() {
	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), predicates.Even[int])

	itertools.Print(items)
	// Output: 2, 4
}

func ExampleTrue() {
	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), predicates.True[int])

	itertools.Print(items)
	// Output: 1, 2, 3, 4, 5
}

func ExampleFalse() {
	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), predicates.False[int])

	itertools.Print(items)
	// Output:
}

func ExampleNot() {
	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), predicates.Not(predicates.Even[int]))

	itertools.Print(items)
	// Output: 1, 3, 5
}

func ExampleIs() {
	items := itertools.Filter(iter.New(1, 2, 3, 4, 5), predicates.Is(3))

	itertools.Print(items)
	// Output: 3
}

func ExampleHas() {
	in := iter.New(
		tuple.New(1, "a"),
		tuple.New(2, "b"),
		tuple.New(3, "c"),
		tuple.New(4, "d"),
		tuple.New(5, "e"),
	)

	items := itertools.Filter(in, predicates.Has[int, string]("c"))

	itertools.Print(items)
	// Output: {3 c}
}

func TestFalse(t *testing.T) {
	t.Parallel()

	result := predicates.False("test")

	assert.False(t, result)
}

func TestTrue(t *testing.T) {
	t.Parallel()

	result := predicates.True("test")

	assert.True(t, result)
}

func TestNot(t *testing.T) {
	t.Parallel()

	// Test case 1: Not(True) should return false
	result1 := predicates.Not(predicates.True[string])("test")
	assert.False(t, result1)

	// Test case 2: Not(False) should return true
	result2 := predicates.Not(predicates.False[string])("test")
	assert.True(t, result2)
}

func TestIs(t *testing.T) {
	t.Parallel()

	// Test case 1: Is(5) should return true when compared to 5
	result1 := predicates.Is(5)(5)
	assert.True(t, result1)

	// Test case 2: Is(5) should return false when compared to 10
	result2 := predicates.Is(5)(10)
	assert.False(t, result2)

	// Test case 3: Is("hello") should return true when compared to "hello"
	result3 := predicates.Is("hello")("hello")
	assert.True(t, result3)

	// Test case 4: Is("hello") should return false when compared to "world"
	result4 := predicates.Is("hello")("world")
	assert.False(t, result4)
}

func TestHas(t *testing.T) {
	t.Parallel()

	// Test case 1: Has(5) should return true when compared to Tuple{1, 5}
	result1 := predicates.Has[int, int](5)(tuple.New(1, 5))
	assert.True(t, result1)

	// // Test case 2: Has(5) should return false when compared to Tuple{2, 10}
	result2 := predicates.Has[int, int](5)(tuple.New(5, 10))
	assert.False(t, result2)

	// Test case 3: Has("hello") should return true when compared to Tuple{5, "hello"}
	result3 := predicates.Has[int, string]("hello")(tuple.New(5, "hello"))
	assert.True(t, result3)

	// Test case 4: Has("hello") should return false when compared to Tuple{"world", "world"}
	result4 := predicates.Has[string, string]("hello")(tuple.New("world", "world"))
	assert.False(t, result4)

	// Test case 5: Has(nil) should return false when compared to Tuple{"hello", errors.New("world")}
	result5 := predicates.Has[string, error](nil)(tuple.New("hello", errors.New("world")))
	assert.False(t, result5)

	// Test case 6: Has(nil) should return true when compared to Tuple{"hello", nil}
	result6 := predicates.Has[string, error](nil)(tuple.New[string, error]("hello", nil))
	assert.True(t, result6)
}
