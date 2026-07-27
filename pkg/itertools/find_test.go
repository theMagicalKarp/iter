package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleFind() {
	result, _ := itertools.Find(iter.New(1, 2, 3, 4, 5, 6), func(n int) bool {
		return n > 3
	})

	fmt.Println(result)
	// Output: 4
}

func TestFindBasic(t *testing.T) {
	t.Parallel()

	target := iter.New(1, 2, 3, 4, 5, 6, 3)

	resp, found := itertools.Find(target, func(i int) bool {
		return i == 3
	})

	assert.Equal(t, 3, resp)
	assert.True(t, found)
}

func TestFindDNE(t *testing.T) {
	t.Parallel()

	target := iter.New(1, 2, 3, 4, 5, 3, 6, 3)

	resp, found := itertools.Find(target, func(i int) bool {
		return i == 42
	})

	assert.Equal(t, 0, resp)
	assert.False(t, found)
}

func TestFindEmpty(t *testing.T) {
	t.Parallel()

	target := iter.New[int]()

	resp, found := itertools.Find(target, func(_ int) bool {
		return true
	})

	assert.Equal(t, 0, resp)
	assert.False(t, found)
}
