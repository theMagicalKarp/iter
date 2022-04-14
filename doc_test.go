package iter_test

import (
	"fmt"

	"github.com/theMagicalKarp/iter"
)

func Example() {
	items := iter.NewIter(1, 2, 3)
	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// 1
	// 2
	// 3
}

func ExampleNewIter() {
	items := iter.NewIter(1, 2, 3)
	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// 1
	// 2
	// 3
}

func ExampleList() {
	list := iter.NewList(
		iter.NewIter(1, 2, 3),
		iter.NewIter(4, 5, 6),
	)
	list.PushBack(7)
	list.PushFront(-1)

	items := list.Iter()
	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// -1
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
}

func ExampleSet() {
	set := iter.NewSet(
		iter.NewIter("around", "the", "world"),
		iter.NewIter("around", "the", "world"),
		iter.NewIter("hello", "world"),
	)
	set.Add("foo")

	items := set.Iter()

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// around
	// the
	// world
	// hello
	// foo
}

func ExampleSet_Intersection() {
	set1 := iter.NewSet(
		iter.NewIter("a", "b", "c", "d", "e", "f"),
	)
	set2 := iter.NewSet(
		iter.NewIter("d", "e", "f", "g", "h", "i"),
	)

	newSet := set1.Intersection(set2)

	items := newSet.Iter()

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// d
	// e
	// f
}

func ExampleSet_Sub() {
	set1 := iter.NewSet(
		iter.NewIter("a", "b", "c", "d", "e", "f"),
	)
	set2 := iter.NewSet(
		iter.NewIter("d", "e", "f", "g", "h", "i"),
	)

	newSet := set1.Sub(set2)

	items := newSet.Iter()

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// a
	// b
	// c
}

func ExampleSet_Union() {
	set1 := iter.NewSet(
		iter.NewIter("a", "b", "c", "d", "e", "f"),
	)
	set2 := iter.NewSet(
		iter.NewIter("d", "e", "f", "g", "h", "i"),
	)

	newSet := set1.Union(set2)

	items := newSet.Iter()

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// a
	// b
	// c
	// d
	// e
	// f
	// g
	// h
	// i
}

func ExampleArray() {
	array := iter.NewArray[string](2, 6)
	array.Set(0, "a")
	array.Set(1, "b")
	array.Append("c")

	fmt.Println(array.At(0))
	fmt.Println(array.At(1))
	fmt.Println(array.At(2))

	// output:
	// a
	// b
	// c
}

func ExampleArray_Iter() {
	array := iter.NewArray(0, 6, iter.NewIter("a", "b", "c", "d", "e", "f"))

	items := array.Iter()

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// a
	// b
	// c
	// d
	// e
	// f
}

func ExampleArray_Range() {
	array := iter.NewArray(0, 6, iter.NewIter("a", "b", "c", "d", "e", "f"))

	items := array.Range(1, 5)

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// b
	// c
	// d
	// e
}

func ExampleArray_Reverse() {
	array := iter.NewArray(0, 6, iter.NewIter("a", "b", "c", "d", "e", "f"))

	items := array.Reverse()

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// f
	// e
	// d
	// c
	// b
	// a
}

func ExampleArray_Tail() {
	array := iter.NewArray(0, 6, iter.NewIter("a", "b", "c", "d", "e", "f"))

	items := array.Tail(2)

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// c
	// d
	// e
	// f
}

func ExampleArray_Head() {
	array := iter.NewArray(0, 6, iter.NewIter("a", "b", "c", "d", "e", "f"))

	items := array.Head(2)

	for items.HasNext() {
		fmt.Println(items.Next())
	}

	// output:
	// a
	// b
}
