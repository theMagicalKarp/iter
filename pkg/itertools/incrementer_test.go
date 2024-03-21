package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleIncrementer() {
	items := itertools.Incrementer(0)

	itertools.Print(itertools.Take(items, 3))
	// Output: 0, 1, 2
}

func TestIncrementer(t *testing.T) {
	t.Parallel()

	repeat := itertools.Incrementer(5)

	for i := 5; i < 10; i++ {
		value, more := repeat.Next()
		assert.True(t, more)
		assert.Equal(t, i, value)
	}
}

func TestDecrementer(t *testing.T) {
	t.Parallel()

	repeat := itertools.Decrementer(5)

	for i := 5; i >= 0; i-- {
		value, more := repeat.Next()
		assert.True(t, more)
		assert.Equal(t, i, value)
	}
}
