package itertools

import "github.com/theMagicalKarp/iter/pkg/iter"

// Logger is an interface for logging messages.
type Logger interface {
	Println(v ...any)
}

type logIter[T any] struct {
	logger Logger
	iter   iter.Iterable[T]
	format func(item T) string
}

func (l *logIter[T]) Next() (T, bool) {
	value, more := l.iter.Next()

	if !more {
		return value, more
	}

	l.logger.Println(l.format(value))

	return value, more
}

// Log is a function that wraps an iterable with logging capabilities.
// It takes an iterable, a logger, and a format function as input,
// and returns a new iterable that logs each item before yielding it.
// The logger is used to log each item, and the format function is used
// to format the item before logging.
func Log[T any](iter iter.Iterable[T], logger Logger, format func(item T) string) iter.Iterable[T] {
	return &logIter[T]{
		logger: logger,
		iter:   iter,
		format: format,
	}
}
