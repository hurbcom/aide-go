package lib

import (
	"context"
	"errors"
	"testing"
)

func TestErrorGroup(t *testing.T) {
	type args struct {
		ctx  context.Context
		args []func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should run successfully when all go routines complete successfully",
			args: args{
				ctx: context.Background(),
				args: []func() error{
					func() error {
						return nil
					},
					func() error {
						return nil
					},
				},
			},
		},
		{
			name: "Should fail when first fo routine fails",
			args: args{
				ctx: context.Background(),
				args: []func() error{
					func() error {
						return errors.New("Oops, you received an error")
					},
					func() error {
						return nil
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ErrorGroup(tt.args.ctx, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("ErrorGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
