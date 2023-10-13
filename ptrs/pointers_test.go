package ptrs

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestSetIfPresent(t *testing.T) {
	value := "a value"

	type args[T any] struct {
		target T
		value  *T
	}
	type testCase[T any] struct {
		name     string
		args     args[T]
		expected T
	}
	tests := []testCase[string]{
		{
			name: "present",
			args: args[string]{
				target: "",
				value:  &value,
			},
			expected: value,
		},
		{
			name: "not present",
			args: args[string]{
				target: "original",
				value:  nil,
			},
			expected: "original",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := tt.args.target
			SetIfPresent(&target, tt.args.value)
			require.Equal(t, tt.expected, target)
		})
	}
}

func TestPtr(t *testing.T) {
	value := 42

	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *T
	}
	tests := []testCase[int]{
		{
			name: "int",
			args: args[int]{
				value: value,
			},
			want: &value,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ptr(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_True(t *testing.T) {
	require.Equal(t, Ptr(true), True())
	require.True(t, *True())
}

func Test_False(t *testing.T) {
	require.Equal(t, Ptr(false), False())
	require.False(t, *False())
}

func Test_Zero(t *testing.T) {
	require.Equal(t, Ptr(0), Zero[int]())
	require.Equal(t, 0.0, *Zero[float64]())
}

func Test_One(t *testing.T) {
	require.Equal(t, Ptr[int32](1), One[int32]())
	require.Equal(t, int32(1), *One[int32]())
}
