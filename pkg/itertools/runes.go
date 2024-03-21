package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

// Runes returns an iterable of runes from the given string value.
func Runes(value string) iter.Iterable[rune] {
	return iter.New([]rune(value)...)
}
