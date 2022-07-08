package lib

import (
	"context"
	"errors"
	"time"
)

type RetryOptions struct {
	RetriesCount    int
	RetriesInterval time.Duration
}

func WithRetries(ctx context.Context, retryFn func(ctx context.Context) error, options RetryOptions) error {
	for i := 0; i < options.RetriesCount; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := retryFn(ctx); err == nil {
			return nil
		}
		time.Sleep(options.RetriesInterval)
	}
	return errors.New("maximum number of retries exceeded")
}
