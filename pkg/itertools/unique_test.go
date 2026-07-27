package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleUnique() {
	items := itertools.Unique(iter.New(1, 2, 2, 3, 3, 3))

	itertools.Print(items)
	// Unordered output: 1, 2, 3
}

func TestUniqueBasic(t *testing.T) {
	t.Parallel()

	input2 := iter.New(1, 2, 2, 3, 3, 4, 5)
	expected2 := []int{1, 2, 3, 4, 5}
	result2 := itertools.Unique(input2)
	assert.ElementsMatch(t, expected2, itertools.Slice(result2))
}

func TestUniqueString(t *testing.T) {
	t.Parallel()

	input2 := iter.New("a", "aa", "b", "c", "b")
	expected2 := []string{"a", "aa", "b", "c"}
	result2 := itertools.Unique(input2)
	assert.ElementsMatch(t, expected2, itertools.Slice(result2))
}

func TestUniqueEmpty(t *testing.T) {
	t.Parallel()

	input3 := iter.New[int]()
	expected3 := []int{}
	result3 := itertools.Unique(input3)
	assert.ElementsMatch(t, expected3, itertools.Slice(result3))
}
