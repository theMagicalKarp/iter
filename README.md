## Iter

Iter is a _simple_ and _ergonomic_ package for implementing and utilinzing
iterators in Golang.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/stretchr/testify)](https://pkg.go.dev/github.com/stretchr/testify)

### Features

- Easy to implement Iterable interface
- Iterable data structures
- Standardized toolset for utilizing iterables

### Getting Started

```Go
package main

import (
	"fmt"

	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func main() {
	// create iterator from static values
	numbers := iter.NewIter(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// filter to collect even numbers
	evenNumbers := itertools.Filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	// square numbers and convert to strings
	squaredNumbers := itertools.Map(evenNumbers, func(n int) string {
		return fmt.Sprintf("%d", n*n)
	})

	fmt.Println(
		// combine into all values into single string
		itertools.String(
			// join values by commas
			itertools.Join(squaredNumbers, ", "),
		),
	)
	// 0, 4, 16, 36, 64
}
```

### How

Everything in the iter packages utilizes a single interface for reading data.

```Go
type Iter[T any] interface {
	Next() T
	HasNext() bool
}
```

### Why

Universal iteratorable interfaces are foundational to powerful problem sovling.
When using iterables as the standard for functional communciation you can create
and utilize standard sets of libraries and structs seamlessly. This results in
clever and easy to understand solutions.

For example if you wanted find the max integer within a file.

```Go
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/theMagicalKarp/iter/itertools"
)

func main() {
	f, err := os.Open("nums.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	lines := itertools.Lines(f)
	nums := itertools.Map(lines, func(line string) int {
		n, _ := strconv.Atoi(line)
		return n
	})

	fmt.Println(itertools.Max(nums))
}
```

Since many of the itertools functions accept and return iterables, you can
plugin and play various functions to achieve a desired result.
