package itertools

import (
	"context"
	"sync"

	"github.com/theMagicalKarp/iter"
)

type asyncIter[T any, V any] struct {
	next      *V
	workerOut chan V
}

func (a *asyncIter[T, V]) Next() V {
	if !a.HasNext() {
		var v V

		return v
	}

	defer a.clear()

	return *a.next
}

func (a *asyncIter[T, V]) clear() {
	a.next = nil
}

func (a *asyncIter[T, V]) HasNext() bool {
	if a.next != nil {
		return true
	}

	next, ok := <-a.workerOut

	if ok {
		a.next = &next
	}

	return a.next != nil
}

func worker[T any, V any](ctx context.Context, workerIn chan T, workerOut chan V,
	fn func(ctx context.Context, item T) V,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case next, ok := <-workerIn:
			if !ok {
				return
			}

			workerOut <- fn(ctx, next)
		}
	}
}

func Async[T any, V any](ctx context.Context, in iter.Iter[T], workers int,
	fn func(ctx context.Context, item T) V,
) iter.Iter[V] {
	if workers <= 0 {
		return iter.NewIter[V]()
	}

	var mu sync.Mutex

	aliveWorkers := workers
	workerIn := make(chan T, workers)
	workerOut := make(chan V, workers)

	async := &asyncIter[T, V]{
		workerOut: workerOut,
	}

	go func() {
		defer close(workerIn)

		for in.HasNext() {
			select {
			case <-ctx.Done():
				return
			case workerIn <- in.Next():
			}
		}
	}()

	for i := 0; i < workers; i++ {
		go func() {
			worker(ctx, workerIn, workerOut, fn)
			mu.Lock()
			defer mu.Unlock()
			aliveWorkers--

			if aliveWorkers == 0 {
				close(workerOut)
			}
		}()
	}

	return async
}
