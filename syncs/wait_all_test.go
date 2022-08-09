package syncs

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/totemcaf/gollections/slices"
)

func Test_WaitAll_waits_for_all_the_waitables(t *testing.T) {
	funcs := slices.Map([]int{1, 2, 3, 4, 5}, func(i int) Waitable[int] {
		return func(ctx context.Context) (int, error) {
			fmt.Println("waiting for", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			fmt.Println("finished for", i)
			return i, nil
		}
	})

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	results, errs := WaitAll(ctx, funcs...)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, results)
	assert.Equal(t, []error{nil, nil, nil, nil, nil}, errs)
}

func Test_WaitAll_report_failing_waitables(t *testing.T) {
	funcs := slices.Map([]int{1, 2, 3, 4, 5}, func(i int) Waitable[int] {
		return func(ctx context.Context) (int, error) {
			fmt.Println("waiting for", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			fmt.Println("finished for", i)

			if i == 3 || i == 4 {
				return 0, fmt.Errorf("error for %d", i)
			}

			return i, nil
		}
	})

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	results, errs := WaitAll(ctx, funcs...)

	assert.Equal(t, []int{1, 2, 0, 0, 5}, results)
	assert.Equal(t, []error{nil, nil, fmt.Errorf("error for 3"), fmt.Errorf("error for 4"), nil}, errs)
}
