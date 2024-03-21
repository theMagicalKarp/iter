package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleKeys() {
	input := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}
	items := itertools.Keys(input)

	itertools.Println(items)
	// Unordered output:
	// apple
	// banana
	// orange
}

func TestBasicKeys(t *testing.T) {
	t.Parallel()

	expected := []string{"apple", "banana", "orange"}
	actual := itertools.Slice(itertools.Keys(map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}))

	assert.ElementsMatch(t, expected, actual)
}

func TestEmptyKeys(t *testing.T) {
	t.Parallel()

	expected := []string{}
	actual := itertools.Slice(itertools.Keys(map[string]int{}))

	assert.ElementsMatch(t, expected, actual)
}
