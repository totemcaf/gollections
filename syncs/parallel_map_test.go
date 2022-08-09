package syncs

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ParallelMap_maps_values(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	results, errs := ParallelMap(values, time.Second, func(_ context.Context, v int) (int, error) {
		return v * v, nil
	})

	assert.Equal(t, []int{1, 4, 9, 16, 25}, results)
	assert.Equal(t, []error{nil, nil, nil, nil, nil}, errs)

}

func Test_ParallelMap_cancels_remaining_waitables_if_context_expires(t *testing.T) {
	values := []int{1, 2, 3}

	results, errs := ParallelMap(values, time.Millisecond*100, func(ctx context.Context, v int) (int, error) {
		if Sleep(ctx, time.Second) {
			return 0, ctx.Err()
		}
		return v * v, nil
	})

	assert.Equal(t, []int{0, 0, 0}, results)
	assert.Equal(t, []error{context.DeadlineExceeded, context.DeadlineExceeded, context.DeadlineExceeded}, errs)
}
