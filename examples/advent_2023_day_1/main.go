// Package main solves Advent of Code 2023 day 1 as a demonstration of the
// iter library.
package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

const inputURL = "https://gist.githubusercontent.com/theMagicalKarp/089e97377f559b65503d17f8dddea5f4/raw/18d74d27fee9d727d4dbc50ce85569eb2ced8a75/advent_2023_day1.txt"

// ToDigit reads the digit at the start of text, accepting either a numeral or
// an English digit name, and returns -1 when text starts with neither.
func ToDigit(text string) int {
	if unicode.IsDigit(rune(text[0])) {
		return int(text[0] - '0')
	}

	// Ordered by value, so a word's index is one less than the digit it names.
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for index, word := range words {
		if strings.HasPrefix(text, word) {
			return index + 1
		}
	}

	return -1
}

func ChopString(text string) func(int) string {
	return func(i int) string {
		return text[i:]
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
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, inputURL, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer func() { _ = resp.Body.Close() }()

	lines := itertools.Lines(resp.Body)
	readings := itertools.Map(lines, ProcessLine)

	fmt.Println(itertools.Sum(readings))
}
