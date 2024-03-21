package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

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
	y := line.First()
	runes := itertools.Enumerate(itertools.Runes(line.Second()))
	numbers := make([]Number, 0)

	value := 0
	neighbors := make([]Coordinate, 0)

	itertools.Each(runes, func(item tuple.Tuple[int, rune]) {
		x, r := item.Unpack()

		if !unicode.IsDigit(r) && value > 0 {
			numbers = append(numbers, Number{
				value:     value,
				neighbors: neighbors,
			})

			value = 0
			neighbors = make([]Coordinate, 0)
		}

		if unicode.IsDigit(r) {
			value = value*10 + int(r-'0')
			neighbors = append(neighbors, Coordinate{x - 1, y - 1})
			neighbors = append(neighbors, Coordinate{x, y - 1})
			neighbors = append(neighbors, Coordinate{x + 1, y - 1})

			neighbors = append(neighbors, Coordinate{x - 1, y})
			neighbors = append(neighbors, Coordinate{x + 1, y})

			neighbors = append(neighbors, Coordinate{x - 1, y + 1})
			neighbors = append(neighbors, Coordinate{x, y + 1})
			neighbors = append(neighbors, Coordinate{x + 1, y + 1})
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

func RuneToCoordinate(y int) func(tuple.Tuple[int, rune]) tuple.Tuple[Coordinate, rune] {
	return func(item tuple.Tuple[int, rune]) tuple.Tuple[Coordinate, rune] {
		x, r := item.Unpack()
		return tuple.New(Coordinate{x, y}, r)
	}
}

func RuneIsSymbol(item tuple.Tuple[Coordinate, rune]) bool {
	r := item.Second()
	return !(unicode.IsDigit(r) || r == '.')
}

func SymbolCoordinates(line tuple.Tuple[int, string]) iter.Iterable[Coordinate] {
	y := line.First()

	runes := itertools.Enumerate(itertools.Runes(line.Second()))
	runeCordinates := itertools.Map(runes, RuneToCoordinate(y))
	runeCordinates = itertools.Filter(runeCordinates, RuneIsSymbol)
	return itertools.Map(runeCordinates, tuple.First)
}

func IsAdjacent(coordinates map[Coordinate]bool) func(Number) bool {
	return func(number Number) bool {
		return itertools.Any(iter.New(number.neighbors...), func(c Coordinate) bool {
			return coordinates[c]
		})
	}
}

func main() {
	f, err := os.Open("examples/example3/input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	first, second := itertools.Tee(itertools.Enumerate(itertools.Lines(f)))

	numbers := itertools.Flatten(itertools.Map(first, ProcessLine))
	symbolCoordinates := itertools.Set(itertools.Flatten(
		itertools.Map(second, SymbolCoordinates),
	))

	numbers = itertools.Filter(numbers, IsAdjacent(symbolCoordinates))

	fmt.Println(itertools.Sum(itertools.Map(numbers, Value)))
}
