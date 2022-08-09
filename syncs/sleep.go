package syncs

import (
	"context"
	"time"
)

// Sleep waits for the given duration.
// Returns true if the context was canceled.
func Sleep(ctx context.Context, delay time.Duration) bool {
	select {
	case <-time.After(delay):
		return false
	case <-ctx.Done():
		return true
	}
}
