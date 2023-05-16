package retry

import (
	"context"
	"fmt"
	"time"
)

type RetryAction func(ctx context.Context) error

type Retry struct {
	attempts int
	delay    time.Duration
}

func New(attemps int, delay time.Duration) *Retry {
	return &Retry{
		attempts: attemps,
		delay:    delay,
	}
}

func (r *Retry) Execute(ctx context.Context, action RetryAction) error {
	for i := 0; i < r.attempts; i++ {
		err := action(ctx)
		if err == nil {
			return nil
		}

		if err.Error() == "401 Unauthorized" || err.Error() == "404 Not Found" {
			return err
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("retry canceled: %w", ctx.Err())
		case <-time.After(r.delay):
			r.delay += 20 * time.Millisecond
		}
	}
	return fmt.Errorf("after %d attempts, last error", r.attempts)
}
