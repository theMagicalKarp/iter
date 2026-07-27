// Package main solves Advent of Code 2023 day 2 as a demonstration of the
// iter library.
package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

const inputURL = "https://gist.githubusercontent.com/theMagicalKarp/089e97377f559b65503d17f8dddea5f4/raw/18d74d27fee9d727d4dbc50ce85569eb2ced8a75/advent_2023_day2.txt"

type Record struct {
	red   int
	green int
	blue  int
}

func Red(record Record) int {
	return record.red
}

func Green(record Record) int {
	return record.green
}

func Blue(record Record) int {
	return record.blue
}

type Game struct {
	id      int
	records []Record
}

// Scanf reads values out of text using the given format, panicking on
// malformed input. The puzzle input is well formed, so a failure here means
// the parser itself is wrong.
func Scanf(text string, format string, targets ...any) {
	_, err := fmt.Sscanf(text, format, targets...)
	if err != nil {
		panic(err)
	}
}

func UnmarshalRecord(line string) Record {
	var record Record

	for entry := range strings.SplitSeq(line, ",") {
		entry = strings.TrimSpace(entry)

		switch {
		case strings.Contains(entry, "red"):
			Scanf(entry, "%d", &record.red)
		case strings.Contains(entry, "green"):
			Scanf(entry, "%d", &record.green)
		case strings.Contains(entry, "blue"):
			Scanf(entry, "%d", &record.blue)
		}
	}

	return record
}

func UnmarshalGame(line string) Game {
	var game Game

	Scanf(line, "Game %d", &game.id)

	rawRecords := strings.TrimPrefix(line, fmt.Sprintf("Game %d: ", game.id))
	iRawRecords := iter.New(strings.Split(rawRecords, ";")...)
	game.records = itertools.Slice(itertools.Map(iRawRecords, UnmarshalRecord))

	return game
}

func LowestAllowedProduct(game Game) int {
	red := itertools.Max(itertools.Map(iter.New(game.records...), Red))
	green := itertools.Max(itertools.Map(iter.New(game.records...), Green))
	blue := itertools.Max(itertools.Map(iter.New(game.records...), Blue))

	return itertools.Product(iter.New(red, green, blue))
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
	games := itertools.Map(lines, UnmarshalGame)
	minProducts := itertools.Map(games, LowestAllowedProduct)

	fmt.Println(itertools.Sum(minProducts))
}
