package itertools

import (
	"strings"

	"github.com/theMagicalKarp/iter"
)

func String(iter iter.Iter[string]) string {
	var builder strings.Builder

	for iter.HasNext() {
		builder.WriteString(iter.Next())
	}

	return builder.String()
}
