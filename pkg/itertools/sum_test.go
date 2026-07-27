package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleSum() {
	result := itertools.Sum(iter.New(1, 2, 3))

	fmt.Println(result)
	// Output:
	// 6
}

func TestSumBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Sum(iter.New(1, 2, 3, 6, 4, 5))

	assert.Equal(t, 21, result)
}

func TestSumString(t *testing.T) {
	t.Parallel()

	result := itertools.Sum(iter.New("hello", " ", "world", "!"))

	assert.Equal(t, "hello world!", result)
}

func TestSumEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.Sum(iter.New[string]())

	assert.Empty(t, result)
}
