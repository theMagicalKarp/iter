package itertools_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleProduct() {
	result := itertools.Product(iter.New(1, 2, 3, 4))

	fmt.Println(result)
	// Output:
	// 24
}

func TestProductBasic(t *testing.T) {
	t.Parallel()

	result := itertools.Product(iter.New(1, 2, 3, 6, 4, 5))

	assert.Equal(t, 720, result)
}

func TestProeuctEmpty(t *testing.T) {
	t.Parallel()

	result := itertools.Product(iter.New[int]())

	assert.Equal(t, 0, result)
}
