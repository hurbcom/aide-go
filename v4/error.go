package v4

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// ErrorGroup resolves all go routines. It fails at first error encountered
// warning: Creating goroutines in the args will cause them to run in background
func ErrorGroup(ctx context.Context, args ...func() error) error {
	g, _ := errgroup.WithContext(ctx)

	for _, fn := range args {
		g.Go(fn)
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
