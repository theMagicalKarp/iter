package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func TestBasicEndsWith(t *testing.T) {
	t.Parallel()

	result := itertools.EndsWith(iter.New(1, 2, 3, 4, 5), iter.New(4, 5))

	assert.True(t, result)
}

func TestBasicEndsWithEqual(t *testing.T) {
	t.Parallel()

	result := itertools.EndsWith(iter.New(1, 2, 3), iter.New(1, 2, 3))

	assert.True(t, result)
}

func TestBasicEndsWithSuffixLonger(t *testing.T) {
	t.Parallel()

	result := itertools.EndsWith(iter.New(1, 2, 3), iter.New(1, 2, 3, 4, 5))

	assert.False(t, result)
}

func TestBasicEndsWithEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.EndsWith(iter.New[int](), iter.New[int]())

	assert.True(t, result)
}

func TestBasicEndsWithSuffixEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.EndsWith(iter.New[int](1, 2, 3), iter.New[int]())

	assert.True(t, result)
}

func TestBasicEndsWithInputEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.EndsWith(iter.New[int](), iter.New[int](1, 2, 3))

	assert.False(t, result)
}
