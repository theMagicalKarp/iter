package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func TestEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Empty[int]()
	assert.NotNil(t, items)

	value, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, value)

	value, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, value)
}
