package lib

import (
	"context"
	"errors"
	"testing"
)

func TestWithRetries(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	type args struct {
		ctx     context.Context
		retryFn func() error
		options RetryOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return nil if the retry parameter function does not return an error",
			args: args{
				ctx: context.TODO(),
				retryFn: func() error {
					return nil
				},
				options: RetryOptions{
					RetriesCount:    3,
					RetriesInterval: 2,
				},
			},
			wantErr: false,
		},
		{
			name: "should return an error if the maximum retries is exceeded",
			args: args{
				ctx: context.TODO(),
				retryFn: func() error {
					return errors.New("error")
				},
				options: RetryOptions{
					RetriesCount:    3,
					RetriesInterval: 2,
				},
			},
			wantErr: true,
		},
		{
			name: "should return an error if the context is cancelled",
			args: args{
				ctx: ctx,
				retryFn: func() error {
					return nil
				},
				options: RetryOptions{
					RetriesCount:    3,
					RetriesInterval: 2,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WithRetries(tt.args.ctx, tt.args.retryFn, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("WithRetries() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
