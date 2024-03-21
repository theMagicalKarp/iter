package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleRepeat() {
	items := itertools.Repeat("hello world")

	itertools.Print(itertools.Take(items, 3))
	// Output: hello world, hello world, hello world
}

func TestRepeat(t *testing.T) {
	t.Parallel()

	value := "hello"
	repeat := itertools.Repeat(value)

	for i := 0; i < 5; i++ {
		v, more := repeat.Next()
		assert.True(t, more)
		assert.Equal(t, v, value)
	}

	v, more := repeat.Next()
	assert.True(t, more)
	assert.Equal(t, v, "hello")
}
