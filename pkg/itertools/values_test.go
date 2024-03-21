package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleValues() {
	input := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}
	items := itertools.Values(input)

	itertools.Println(items)
	// Unordered output:
	// 2
	// 3
	// 5
}

func TestBasicValues(t *testing.T) {
	t.Parallel()

	expected := []int{5, 3, 2}
	actual := itertools.Slice(itertools.Values(map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}))

	assert.ElementsMatch(t, expected, actual)
}

func TestEmptyValues(t *testing.T) {
	t.Parallel()

	expected := []int{}
	actual := itertools.Slice(itertools.Values(map[string]int{}))

	assert.ElementsMatch(t, expected, actual)
}
