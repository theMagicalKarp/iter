package itertools

import (
	"context"
	"sync"

	"github.com/theMagicalKarp/iter/pkg/iter"
)

// https://pkg.go.dev/gopkg.in/go-playground/pool.v2#section-readme

// AsyncMapFn is a function type that represents an asynchronous mapping operation.
// It takes a context and an item of type T as input, and returns a value of type V.
// The context is used for cancellation and deadline propagation across function calls.
type AsyncMapFn[T any, V any] func(ctx context.Context, item T) V

type asyncIter[T any, V any] struct {
	channel chan V
	// done mirrors the cancellation signal of the context passed to Async.
	// Only the signal is retained rather than the context itself, so the
	// iterator carries no request scoped state beyond what it needs.
	done <-chan struct{}
}

func (a *asyncIter[T, V]) Next() (V, bool) {
	var (
		value V
		more  bool
	)

	select {
	case <-a.done:
		return value, false
	case value, more = <-a.channel:
		return value, more
	}
}

func pushToCannel[T any](ctx context.Context, input iter.Iterable[T], channel chan T) {
	defer close(channel)

	for value, ok := input.Next(); ok; value, ok = input.Next() {
		select {
		case <-ctx.Done():
			return
		case channel <- value:
		}
	}
}

func worker[T any, V any](
	ctx context.Context,
	input chan T,
	output chan V,
	waitGroup *sync.WaitGroup,
	function AsyncMapFn[T, V],
) {
	defer waitGroup.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case value, more := <-input:
			if !more {
				return
			}

			output <- function(ctx, value)
		}
	}
}

func waitClose[T any](wg *sync.WaitGroup, channel chan T) {
	wg.Wait()
	close(channel)
}

// Async is a function that takes a context, an input iterable, the number of workers, and an async map function.
// It returns an iterable of the mapped values.
// The function distributes the work across multiple goroutines, allowing for concurrent processing.
// The input iterable is divided into chunks and each chunk is processed by a worker goroutine.
// The number of worker goroutines is determined by the `workers` parameter.
// The `fn` parameter is the async map function that takes an input value and returns a mapped value.
// The function returns an iterable of the mapped values, which can be consumed using the `Next` method.
// If the `workers` parameter is less than 1, an empty iterable is returned.
func Async[T any, V any](
	ctx context.Context, inIter iter.Iterable[T], workers int, function AsyncMapFn[T, V],
) iter.Iterable[V] {
	if workers < 1 {
		return iter.New[V]()
	}

	input := make(chan T, workers)
	output := make(chan V, workers)

	async := &asyncIter[T, V]{done: ctx.Done(), channel: output}

	go pushToCannel(ctx, inIter, input)

	var waitGroup sync.WaitGroup

	waitGroup.Add(workers)

	for range workers {
		go worker(ctx, input, output, &waitGroup, function)
	}

	go waitClose(&waitGroup, output)

	return async
}
