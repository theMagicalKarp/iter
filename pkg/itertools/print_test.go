package itertools_test

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExamplePrint() {
	itertools.Print(iter.New("a", "b", "c"))
	// Output: a, b, c
}

func ExamplePrintln() {
	itertools.Println(iter.New("a", "b", "c"))
	// Output:
	// a
	// b
	// c
}
