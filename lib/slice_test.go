package lib

import (
	"reflect"
	"testing"
)

type TestUser struct {
	Name string
	Age int
}

func TestSliceFilter(t *testing.T) {
	type args struct {
		users Slice[TestUser]
		predicate Predicate[TestUser]
	}
	tests := []struct {
		name string
		args args
		want Slice[TestUser]
	}{
		{
			name: "should return [Jonah]",
			args: args{
				users: Slice[TestUser]{
					{Name: "Jonah", Age: 18},
					{Name: "Camile", Age: 22},
				},
				predicate: func(u TestUser) bool {
					return u.Name == "Jonah"
				},
			},
			want: Slice[TestUser]{
				{Name: "Jonah", Age: 18},
			},
		},
		{
			name: "should return []",
			args: args{
				users: Slice[TestUser]{
					{Name: "Jonah", Age: 18},
					{Name: "Camile", Age: 22},
				},
				predicate: func(u TestUser) bool {
					return u.Name == "Math"
				},
			},
			want: Slice[TestUser]{},
		},
		{
			name: "should return all if predicate is nil",
			args: args{
				users: Slice[TestUser]{
					{Name: "Jonah", Age: 18},
					{Name: "Camile", Age: 22},
				},
				predicate: nil,
			},
			want: Slice[TestUser]{
				{Name: "Jonah", Age: 18},
				{Name: "Camile", Age: 22},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := tt.args.users.Filter(tt.args.predicate); !reflect.DeepEqual(tt.want, actual) {
				t.Fail()
			}
		})
	}
}
