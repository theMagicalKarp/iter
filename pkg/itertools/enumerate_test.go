package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleEnumerate() {
	items := iter.New("hello", "world", "!")

	itertools.Print(itertools.Enumerate(items))
	// Output: {0 hello}, {1 world}, {2 !}
}

func TestEnumerate(t *testing.T) {
	t.Parallel()

	// Test case 2: Non-empty iterable
	items := iter.New[int](1, 2, 3, 4, 5)
	result := itertools.Slice(itertools.Enumerate(items))

	expected := []tuple.Tuple[int, int]{
		tuple.New(0, 1),
		tuple.New(1, 2),
		tuple.New(2, 3),
		tuple.New(3, 4),
		tuple.New(4, 5),
	}

	assert.Equal(t, expected, result)
}

func TestEnumerateEmpty(t *testing.T) {
	t.Parallel()

	// Test case 2: Non-empty iterable
	items := iter.New[int]()
	result := itertools.Slice(itertools.Enumerate(items))

	expected := []tuple.Tuple[int, int]{}

	assert.Equal(t, expected, result)
}
