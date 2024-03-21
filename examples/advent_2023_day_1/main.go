package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

var WORD_TO_DIGIT = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func ToDigit(s string) int {
	if unicode.IsDigit(rune(s[0])) {
		return int(s[0] - '0')
	}

	for k, v := range WORD_TO_DIGIT {
		if strings.HasPrefix(s, k) {
			return v
		}
	}

	return -1
}

func ChopString(s string) func(int) string {
	return func(i int) string {
		return s[i:]
	}
}

func ProcessLine(line string) int {
	irange := itertools.Range(0, len(line))
	subStrings := itertools.Map(irange, ChopString(line))
	allDigits := itertools.Map(subStrings, ToDigit)
	positiveDigits := itertools.Filter(allDigits, predicates.Positive)

	first, more := positiveDigits.Next()
	if !more {
		return 0
	}

	last, found := itertools.Last(positiveDigits)
	if !found {
		last = first
	}

	return first*10 + last
}

func main() {
	f, err := os.Open("examples/example1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := itertools.Lines(f)
	readings := itertools.Map(lines, ProcessLine)
	fmt.Println(itertools.Sum(readings))
}
