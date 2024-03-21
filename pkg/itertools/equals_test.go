package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleEquals() {
	a := iter.New(1, 2, 3)
	b := iter.New(1, 2, 3)

	fmt.Println(itertools.Equals(a, b))
	// Output: true
}

func TestBasicEquals(t *testing.T) {
	t.Parallel()

	a := iter.New(1, 2, 3)
	b := iter.New(1, 2, 3)

	assert.True(t, itertools.Equals(a, b))
}

func TestBasicNotEquals(t *testing.T) {
	t.Parallel()

	a := iter.New(1, 2, 3)
	b := iter.New(1, 3, 2)

	assert.False(t, itertools.Equals(a, b))
}

func TestBasicLopsidedEquals(t *testing.T) {
	t.Parallel()

	a := iter.New(1, 2, 3)
	b := iter.New(1, 2, 3, 4)

	assert.False(t, itertools.Equals(a, b))
}

func TestBasicEmptyEquals(t *testing.T) {
	t.Parallel()

	a := iter.New[int]()
	b := iter.New(1, 2, 3)

	assert.False(t, itertools.Equals(a, b))

	a = iter.New(1, 2, 3)
	b = iter.New[int]()

	assert.False(t, itertools.Equals(a, b))
}

func TestBasicBothEmptyEquals(t *testing.T) {
	t.Parallel()

	a := iter.New[int]()
	b := iter.New[int]()

	assert.True(t, itertools.Equals(a, b))
}

func TestTrickyEquals(t *testing.T) {
	t.Parallel()

	a := iter.New[int]()
	b := iter.New[int](0, 0, 0)

	assert.False(t, itertools.Equals(a, b))
}
