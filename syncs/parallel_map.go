package syncs

import (
	"context"
	"time"

	"github.com/totemcaf/gollections/slices"
)

// ParallelMap applies the given mapper to each element of the given slice in parallel.
// The results and errors are returned in the same order as the slice.
// If the context expires before all the mappers are finished, the remaining mappers are cancelled.
//  If a mapper fails, the error is returned in the errs slice.
//  If a mapper succeeds, the result is returned in the results slice.
//
// Example:
//  values := []int{1, 2, 3, 4, 5}
//  results, errs := ParallelMap(values, time.Second, func(ctx context.Context, v V) (T, error) {
//	    return v * v, nil
//  })
func ParallelMap[V, T any](values []V, maxWait time.Duration, mapper func(context.Context, V) (T, error)) ([]T, []error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(maxWait))
	defer cancel()

	return WaitAll(ctx, slices.Map(values, func(v V) Waitable[T] {
		return func(ctx context.Context) (T, error) {
			return mapper(ctx, v)
		}
	})...)
}
