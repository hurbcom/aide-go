package lib

import (
	"context"
	"errors"
	"time"
)

func WithRetries(ctx context.Context, retryFn func() error, options struct {
	retriesCount    int
	retriesInterval int
}) error {
	for i := 0; i < options.retriesCount; i++ {
		if err := retryFn(); err == nil {
			return nil
		}
		time.Sleep(time.Duration(options.retriesInterval) * time.Second)
	}
	return errors.New("maximum number of retries exceeded")
}
