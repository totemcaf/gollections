package syncs

import (
	"context"
	"sync"
	"time"
)

// Waitable is a function that can be waited on.
type Waitable[T any] func(ctx context.Context) (T, error)

// WaitAll waits for all Waitable to complete successfully or to fail.
// Cancelling the context will cause all Waitable to fail.
//
// The results and errors are returned in the same order as the Waitable.
// If a Waitable fails, the error is returned in the errs slice.
// If a Waitable succeeds, the result is returned in the results slice.
// If a Waitable is cancelled, the result is returned in the results slice and the error is returned in the errs slice.
//
// ctx (context) can be used to cancel the WaitAll. If nil is provided, a context with a deadline of 1 second is used.
// Remember to cancel the context when you're done with it.
//
// Example:
//  ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 3))
//  defer cancel()
//  results, errs := WaitAll(ctx, funcs...)
//
func WaitAll[T any](ctx context.Context, waitables ...Waitable[T]) ([]T, []error) {
	var wg sync.WaitGroup

	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second))
		defer cancel()
	}

	results := make([]T, len(waitables))
	errs := make([]error, len(waitables))

	for i, waitable := range waitables {
		wg.Add(1)
		go func(i int, waitable Waitable[T]) {
			defer wg.Done()
			results[i], errs[i] = waitable(ctx)
		}(i, waitable)
	}
	wg.Wait()
	return results, errs
}
