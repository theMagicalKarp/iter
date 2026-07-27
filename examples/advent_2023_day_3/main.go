// Package main solves Advent of Code 2023 day 3 as a demonstration of the
// iter library.
package main

import (
	"context"
	"fmt"
	"net/http"
	"unicode"

	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

const inputURL = "https://gist.githubusercontent.com/theMagicalKarp/089e97377f559b65503d17f8dddea5f4/raw/18d74d27fee9d727d4dbc50ce85569eb2ced8a75/advent_2023_day3.txt"

type Coordinate struct {
	x, y int
}

type Number struct {
	value     int
	neighbors []Coordinate
}

func Value(n Number) int {
	return n.value
}

func ProcessLine(line tuple.Tuple[int, string]) iter.Iterable[Number] {
	row := line.First()
	runes := itertools.Enumerate(itertools.Runes(line.Second()))
	numbers := make([]Number, 0)

	value := 0
	neighbors := make([]Coordinate, 0)

	itertools.Each(runes, func(item tuple.Tuple[int, rune]) {
		column, char := item.Unpack()

		if !unicode.IsDigit(char) && value > 0 {
			numbers = append(numbers, Number{
				value:     value,
				neighbors: neighbors,
			})

			value = 0
			neighbors = make([]Coordinate, 0)
		}

		if unicode.IsDigit(char) {
			value = value*10 + int(char-'0')

			neighbors = append(neighbors, Coordinate{column - 1, row - 1})
			neighbors = append(neighbors, Coordinate{column, row - 1})
			neighbors = append(neighbors, Coordinate{column + 1, row - 1})

			neighbors = append(neighbors, Coordinate{column - 1, row})
			neighbors = append(neighbors, Coordinate{column + 1, row})

			neighbors = append(neighbors, Coordinate{column - 1, row + 1})
			neighbors = append(neighbors, Coordinate{column, row + 1})
			neighbors = append(neighbors, Coordinate{column + 1, row + 1})
		}
	})

	if value > 0 {
		numbers = append(numbers, Number{
			value:     value,
			neighbors: neighbors,
		})
	}

	return iter.New(numbers...)
}

func RuneToCoordinate(row int) func(tuple.Tuple[int, rune]) tuple.Tuple[Coordinate, rune] {
	return func(item tuple.Tuple[int, rune]) tuple.Tuple[Coordinate, rune] {
		column, char := item.Unpack()

		return tuple.New(Coordinate{column, row}, char)
	}
}

func RuneIsSymbol(item tuple.Tuple[Coordinate, rune]) bool {
	char := item.Second()

	return !unicode.IsDigit(char) && char != '.'
}

func SymbolCoordinates(line tuple.Tuple[int, string]) iter.Iterable[Coordinate] {
	row := line.First()

	runes := itertools.Enumerate(itertools.Runes(line.Second()))
	runeCoordinates := itertools.Map(runes, RuneToCoordinate(row))
	runeCoordinates = itertools.Filter(runeCoordinates, RuneIsSymbol)

	return itertools.Map(runeCoordinates, tuple.First)
}

func IsAdjacent(coordinates map[Coordinate]bool) func(Number) bool {
	return func(number Number) bool {
		return itertools.Any(iter.New(number.neighbors...), func(c Coordinate) bool {
			return coordinates[c]
		})
	}
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

	first, second := itertools.Tee(itertools.Enumerate(itertools.Lines(resp.Body)))

	numbers := itertools.Flatten(itertools.Map(first, ProcessLine))
	symbolCoordinates := itertools.Set(itertools.Flatten(
		itertools.Map(second, SymbolCoordinates),
	))

	numbers = itertools.Filter(numbers, IsAdjacent(symbolCoordinates))

	fmt.Println(itertools.Sum(itertools.Map(numbers, Value)))
}
