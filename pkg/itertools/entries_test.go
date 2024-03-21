package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleEntries() {
	input := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}
	items := itertools.Entries(input)

	itertools.Println(items)
	// Unordered output:
	// {apple 5}
	// {banana 3}
	// {orange 2}
}

func TestBasicEntries(t *testing.T) {
	t.Parallel()

	expected := []tuple.Tuple[string, int]{
		tuple.New("apple", 5),
		tuple.New("banana", 3),
		tuple.New("orange", 2),
	}

	actual := itertools.Slice(itertools.Entries(map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}))

	assert.ElementsMatch(t, expected, actual)
}

func TestEmptyEntries(t *testing.T) {
	t.Parallel()

	expected := []tuple.Tuple[string, int]{}
	actual := itertools.Slice(itertools.Entries(map[string]int{}))

	assert.ElementsMatch(t, expected, actual)
}
