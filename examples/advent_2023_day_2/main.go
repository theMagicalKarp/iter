package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

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

func UnmarshalRecord(line string) Record {
	record := Record{}
	for _, entry := range strings.Split(line, ",") {
		if strings.Contains(entry, "red") {
			fmt.Sscanf(entry, "%d", &record.red)
			continue
		}

		if strings.Contains(entry, "green") {
			fmt.Sscanf(entry, "%d", &record.green)
			continue
		}

		if strings.Contains(entry, "blue") {
			fmt.Sscanf(entry, "%d", &record.blue)
			continue
		}
	}

	return record
}

func UnmarshalGame(line string) Game {
	game := Game{}
	fmt.Sscanf(line, "Game %d", &game.id)

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
	resp, err := http.Get("https://gist.githubusercontent.com/theMagicalKarp/089e97377f559b65503d17f8dddea5f4/raw/18d74d27fee9d727d4dbc50ce85569eb2ced8a75/advent_2023_day2.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	lines := itertools.Lines(resp.Body)
	games := itertools.Map(lines, UnmarshalGame)

	minProducts := itertools.Map(games, LowestAllowedProduct)
	fmt.Println(itertools.Sum(minProducts))
}
