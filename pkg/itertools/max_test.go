package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleMax() {
	result := itertools.Max(iter.New(2, 1, 5, 4, 3))

	fmt.Println(result)
	// Output: 5
}

func TestMaxBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Max(iter.New(1, 2, 3, 6, 4, 5))

	assert.Equal(t, 6, result)
}

func TestMaxEmtpy(t *testing.T) {
	t.Parallel()

	result := itertools.Max(iter.New[int]())

	assert.Equal(t, 0, result)
}

func TestMaxNegatives(t *testing.T) {
	t.Parallel()

	result := itertools.Max(iter.New(-5, -4, -3, -2, -1))

	assert.Equal(t, -1, result)
}

func TestMaxMixed(t *testing.T) {
	t.Parallel()

	result := itertools.Max(iter.New(3, 4, 1, 2, 5))

	assert.Equal(t, 5, result)
}
