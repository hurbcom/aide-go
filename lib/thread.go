package lib

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// ErrorGroup resolve all go routine, if some one fail will return the first error
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
