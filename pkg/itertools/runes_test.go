package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleRunes() {
	items := itertools.Runes("hello world!")

	itertools.Print(items)
	// Output: 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33
}

func TestRunesBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Runes("hello world!")

	for _, r := range "hello world!" {
		v, more := items.Next()
		assert.True(t, more)
		assert.Equal(t, r, v)
	}

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, int32(0), v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, int32(0), v)
}

func TestRunesEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Runes("")

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, int32(0), v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, int32(0), v)
}
