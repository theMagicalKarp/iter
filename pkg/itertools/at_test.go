package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleAt() {
	result, _ := itertools.At(iter.New(1, 2, 3), 1)

	fmt.Println(result)
	// Output: 2
}

func TestAtFirst(t *testing.T) {
	t.Parallel()

	result, more := itertools.At(iter.New("a", "b", "c", "d", "e"), 0)
	assert.True(t, more)
	assert.Equal(t, "a", result)
}

func TestAtLast(t *testing.T) {
	t.Parallel()

	result, more := itertools.At(iter.New("a", "b", "c", "d", "e"), 4)
	assert.True(t, more)
	assert.Equal(t, "e", result)
}

func TestAtMiddle(t *testing.T) {
	t.Parallel()

	result, more := itertools.At(iter.New("a", "b", "c", "d", "e"), 2)
	assert.True(t, more)
	assert.Equal(t, "c", result)
}

func TestAtOverflow(t *testing.T) {
	t.Parallel()

	result, more := itertools.At(iter.New("a", "b", "c", "d", "e"), 10)
	assert.False(t, more)
	assert.Equal(t, "", result)
}

func TestAtNegative(t *testing.T) {
	t.Parallel()

	result, more := itertools.At(iter.New("a", "b", "c", "d", "e"), -10)
	assert.False(t, more)
	assert.Equal(t, "", result)
}
