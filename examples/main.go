package main

import (
	"context"
	"fmt"
	"time"

	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
	// "github.com/theMagicalKarp/iter/predicates"
)

func main() {
	a := iter.NewIter(1, 2, 3, 4, 5, 6)

	ctx, cancel := context.WithCancel(context.Background())
	// cancel()
	stuff := itertools.Async(ctx, a, 3, func(ctx context.Context, i int) string {
		time.Sleep(time.Duration(i) * time.Second)
		cancel()
		return fmt.Sprintf("-- %d", i)
	})

	itertools.Print(stuff)

	// fmt.Println(a.Pop())
	// fmt.Println(FloatToInt(1.4))

	// x := iter.NewIter("hello", "world", "ok", "boomer", "oom", "coom")
	// x := iter.NewArray(100, itertools.Range(0, 100))
	// itertools.Print(x.Iter())
	// y := itertools.All(x, predicates.False[string])
	// fmt.Println(y)

	// z := itertools.Filter(x, predicates.HasSuffix("oom"))
	// itertools.Print(z)

	// fmt.Println("----")

	// p := itertools.Range(0, 20)
	// y := itertools.Filter(p, predicates.Divisible(5))
	// itertools.Print(y)
}

// package main

// import (

// 	"fmt"
// )

// func main() {
// 	x := make([]int, 10)
// 	x[2] = 1
// 	fmt.Println(x)
// }
