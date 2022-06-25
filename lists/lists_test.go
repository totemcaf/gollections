package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	reducer := func(sum int, s string) int { return sum + len(s) }

	tests := []struct {
		name string
		args List[string]
		want int
	}{
		{
			name: "empty",
			args: Of[string](),
			want: 0,
		},
		{
			name: "not empty",
			args: Of("hello", "world"),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Reduce(tt.args, reducer), "Reduce(%v)", tt.args)
		})
	}
}
