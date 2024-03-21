package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleMin() {
	result := itertools.Min(iter.New(2, 1, 5, 4, 3))

	fmt.Println(result)
	// Output: 1
}

func TestMinBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Min(iter.New(1, 2, 3, 6, 4, 5))

	assert.Equal(t, 1, result)
}

func TestMinEmtpy(t *testing.T) {
	t.Parallel()

	result := itertools.Min(iter.New[int]())

	assert.Equal(t, 0, result)
}

func TestMinNegatives(t *testing.T) {
	t.Parallel()

	result := itertools.Min(iter.New(-5, -4, -3, -2, -1))

	assert.Equal(t, -5, result)
}

func TestMinMixed(t *testing.T) {
	t.Parallel()

	result := itertools.Min(iter.New(3, 4, 1, 2, 5, -5))

	assert.Equal(t, -5, result)
}
