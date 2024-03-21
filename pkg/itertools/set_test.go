package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleSet() {
	items := itertools.Set(iter.New(1, 2, 3, 1, 2, 3, 4, 5))

	for k := range items {
		fmt.Println(k)
	}
	// Unordered output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func TestBasicSet(t *testing.T) {
	t.Parallel()

	items := itertools.Set(iter.New(1, 2, 3, 1, 2, 3, 4, 5))
	expected := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
	}

	assert.Equal(t, expected, items)
}

func TestBasicSetEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Set(iter.New[int]())
	expected := map[int]bool{}

	assert.Equal(t, expected, items)
}
