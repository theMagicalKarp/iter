package itertools_test

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleLog() {
	logger := log.New(os.Stdout, "", 0)

	items := itertools.Log(iter.New(1, 2, 3), logger, func(item int) string {
		return fmt.Sprintf("Item: %d", item)
	})

	itertools.Drain(items)
	// Output:
	// Item: 1
	// Item: 2
	// Item: 3
}

func TestLogBasic(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	logger := log.New(&buf, "~", 0)

	items := itertools.Log(iter.New(1, 2, 3), logger, func(item int) string {
		return fmt.Sprintf("Item: %d", item)
	})

	v, more := items.Next()
	assert.True(t, more)
	assert.Equal(t, 1, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 2, v)

	v, more = items.Next()
	assert.True(t, more)
	assert.Equal(t, 3, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	assert.Equal(t, "~Item: 1\n~Item: 2\n~Item: 3\n", buf.String())
}

func TestLogEmpty(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	logger := log.New(&buf, "~", 0)

	items := itertools.Log(iter.New[int](), logger, func(item int) string {
		return fmt.Sprintf("Item: %d", item)
	})

	v, more := items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	v, more = items.Next()
	assert.False(t, more)
	assert.Equal(t, 0, v)

	assert.Empty(t, buf.String())
}
