package ptrs

import (
	"github.com/stretchr/testify/require"
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
