package cond

import (
	"reflect"
	"testing"
)

func TestIIf(t *testing.T) {
	type args[T any] struct {
		condition bool
		ifTrue    T
		ifFalse   T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[string]{
		{
			name: "true",
			args: args[string]{
				condition: true,
				ifTrue:    "true",
				ifFalse:   "false",
			},
			want: "true",
		},
		{
			name: "false",
			args: args[string]{
				condition: false,
				ifTrue:    "true",
				ifFalse:   "false",
			},
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IIf(tt.args.condition, tt.args.ifTrue, tt.args.ifFalse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIIfFunc(t *testing.T) {
	type args[T any] struct {
		condition bool
		ifTrue    func() T
		ifFalse   func() T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "true",
			args: args[int]{
				condition: true,
				ifTrue:    func() int { return 1 },
				ifFalse:   func() int { return 2 },
			},
			want: 1,
		},
		{
			name: "false",
			args: args[int]{
				condition: false,
				ifTrue:    func() int { return 1 },
				ifFalse:   func() int { return 2 },
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IIfFunc(tt.args.condition, tt.args.ifTrue, tt.args.ifFalse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IIfFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
