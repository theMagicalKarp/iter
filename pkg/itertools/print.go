package itertools

import (
	"fmt"
	"os"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

// Println prints the elements of an iterable to the standard output.
// It iterates over the elements of the iterable and prints each element on a new line.
// The type parameter T represents the type of elements in the iterable.
func Println[T any](iter iter.Iterable[T]) {
	for {
		value, more := iter.Next()
		if !more {
			return
		}

		// Errors writing to stdout are deliberately dropped, matching the
		// behaviour of fmt.Println.
		_, _ = fmt.Fprintln(os.Stdout, value)
	}
}

func toString[T any](value T) string {
	return fmt.Sprintf("%v", value)
}

// Print prints the elements of the given iterable to the standard output.
// It converts each element to a string and joins them with a comma.
// The elements are printed in the order they are returned by the iterable.
// After printing all the elements, a newline character is printed.
func Print[T any](iter iter.Iterable[T]) {
	items := Join(Map(iter, toString[T]), ", ")

	for {
		value, more := items.Next()
		if !more {
			_, _ = fmt.Fprint(os.Stdout, "\n")

			return
		}

		_, _ = fmt.Fprint(os.Stdout, value)
	}
}
