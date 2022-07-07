package lib

import (
	"context"
	"errors"
	"testing"
)

func TestWithRetries(t *testing.T) {
	type args struct {
		ctx     context.Context
		retryFn func() error
		options struct {
			retriesCount    int
			retriesInterval int
		}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				retryFn: func() error {
					return nil
				},
				options: struct {
					retriesCount    int
					retriesInterval int
				}{
					retriesCount:    3,
					retriesInterval: 2,
				},
			},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				ctx: context.Background(),
				retryFn: func() error {
					return errors.New("maximum number of retries exceeded")
				},
				options: struct {
					retriesCount    int
					retriesInterval int
				}{
					retriesCount:    3,
					retriesInterval: 2,
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
