package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleStartsWith() {
	items := iter.New(1, 2, 3, 4, 5)
	pre_items := iter.New(1, 2)

	fmt.Println(itertools.StartsWith(items, pre_items))
}

func TestStartsWithBasic(t *testing.T) {
	t.Parallel()

	a := iter.New(1, 2, 3, 4, 5)
	b := iter.New(1, 2)

	assert.True(t, itertools.StartsWith(a, b))
}

func TestStartsWithLong(t *testing.T) {
	t.Parallel()

	a := iter.New(1, 2)
	b := iter.New(1, 2, 3, 4, 5)

	assert.False(t, itertools.StartsWith(a, b))
}

func TestStartsWithEmpty(t *testing.T) {
	t.Parallel()

	a := iter.New(1, 2, 3, 4, 5)
	b := iter.New[int]()

	assert.True(t, itertools.StartsWith(a, b))
}

func TestStartsWithBothEmpty(t *testing.T) {
	t.Parallel()

	a := iter.New[int]()
	b := iter.New[int]()

	assert.True(t, itertools.StartsWith(a, b))
}

func TestStartsWithEqual(t *testing.T) {
	t.Parallel()

	a := iter.New(1, 2, 3, 4, 5)
	b := iter.New(1, 2, 3, 4, 5)

	assert.True(t, itertools.StartsWith(a, b))
}

func TestStartsWithTricky(t *testing.T) {
	t.Parallel()

	a := iter.New(0, 0, 0, 0)
	b := iter.New[int]()

	assert.True(t, itertools.StartsWith(a, b))

	a = iter.New[int]()
	b = iter.New(0, 0, 0, 0)

	assert.False(t, itertools.StartsWith(a, b))
}

func TestStartsWithNotEqual(t *testing.T) {
	t.Parallel()

	a := iter.New(1, 2, 3, 4)
	b := iter.New(1, 1, 1, 1)

	assert.False(t, itertools.StartsWith(a, b))
}
