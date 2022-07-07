package lib

import (
	"context"
	"errors"
	"time"
)

type RetryOptions struct {
	retriesCount    int
	retriesInterval time.Duration
}

func WithRetries(ctx context.Context, retryFn func(ctx context.Context) error, options RetryOptions) error {
	for i := 0; i < options.retriesCount; i++ {
		if err := retryFn(ctx); err == nil {
			return nil
		}
		time.Sleep(options.retriesInterval)
	}
	return errors.New("maximum number of retries exceeded")
}
