package slices_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/totemcaf/gollections/slices"
)

func TestCount_counts_elements(t *testing.T) {
	list := slices.Of("some", "words", "to", "count")

	isLong := func(w string) bool { return len(w) > 2 }

	count := slices.Count(list, isLong)

	assert.Equal(t, 3, count)
}

func TestMap_maps_a_list(t *testing.T) {
	list := slices.Of("some", "words", "to", "count")

	wordLength := func(s string) int { return len(s) }

	mapped := slices.Map(list, wordLength)

	assert.Len(t, mapped, len(list))
	assert.Equal(t, []int{4, 5, 2, 5}, mapped)
}

func TestOf(t *testing.T) {
	type args struct {
		ss []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{},
			want: nil,
		},
		{
			name: "one element",
			args: args{[]int{42}},
			want: []int{42},
		},
		{
			name: "several elements",
			args: args{[]int{42, 74, 89}},
			want: []int{42, 74, 89},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slices.Of(tt.args.ss...), "Of(%v)", tt.args.ss)
		})
	}
}

func TestMap(t *testing.T) {
	mapper := func(s string) int { return len(s) }

	tests := []struct {
		name string
		args []string
		want []int
	}{
		{
			name: "empty",
			args: []string{},
			want: []int{},
		},
		{
			name: "one element",
			args: []string{"hello"},
			want: []int{5},
		},
		{
			name: "several elements",
			args: []string{"hello", "world", "!"},
			want: []int{5, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slices.Map(tt.args, mapper), "Map(%v, mapper)", tt.args)
		})
	}
}

func TestMapNonNil(t *testing.T) {
	mapper := func(s string) *int {
		if len(s) == 3 {
			return nil
		}
		l := len(s)
		return &l
	}

	tests := []struct {
		name string
		args []string
		want []int
	}{
		{
			name: "empty",
			args: []string{},
			want: []int{},
		},
		{
			name: "one element",
			args: []string{"hello"},
			want: []int{5},
		},
		{
			name: "one nil element",
			args: []string{"nil"},
			want: []int{},
		},
		{
			name: "several elements",
			args: []string{"hello", "world", "!"},
			want: []int{5, 5, 1},
		},
		{
			name: "some nil elements",
			args: []string{"hello", "nil", "world", "!", "nil", "nil"},
			want: []int{5, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slices.MapNonNil(tt.args, mapper), "Map(%v, mapper)", tt.args)
		})
	}
}

func TestFlatMap(t *testing.T) {
	mapper := func(s string) []string {
		return strings.Split(s, " ")
	}

	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "empty",
			args: []string{},
			want: []string{},
		},
		{
			name: "one element",
			args: []string{"hello"},
			want: []string{"hello"},
		},
		{
			name: "one elements into several",
			args: []string{"hello world"},
			want: []string{"hello", "world"},
		},
		{
			name: "several elements",
			args: []string{"hello world", "nice day", "!"},
			want: []string{"hello", "world", "nice", "day", "!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slices.FlatMap(tt.args, mapper), "FlatMap(%v, mapper)", tt.args)
		})
	}
}

type comparableInt int // implements Comparable

func (i comparableInt) Compare(other comparableInt) int {
	return int(i) - int(other)
}

func TestRemove(t *testing.T) {
	type args struct {
		ts       []comparableInt
		toRemove comparableInt
	}
	tests := []struct {
		name string
		args args
		want []comparableInt
	}{
		{
			name: "empty",
			args: args{
				ts:       []comparableInt{},
				toRemove: comparableInt(78),
			},
			want: nil,
		},
		{
			name: "not found",
			args: args{
				ts:       []comparableInt{comparableInt(1), comparableInt(2), comparableInt(3)},
				toRemove: comparableInt(4),
			},
			want: []comparableInt{comparableInt(1), comparableInt(2), comparableInt(3)},
		},
		{
			name: "found",
			args: args{
				ts:       []comparableInt{comparableInt(1), comparableInt(2), comparableInt(3)},
				toRemove: comparableInt(2),
			},
			want: []comparableInt{comparableInt(1), comparableInt(3)},
		},
		{
			name: "found at the beginning",
			args: args{
				ts:       []comparableInt{comparableInt(1), comparableInt(2), comparableInt(3)},
				toRemove: comparableInt(1),
			},
			want: []comparableInt{comparableInt(2), comparableInt(3)},
		},
		{
			name: "found at the end",
			args: args{
				ts:       []comparableInt{comparableInt(1), comparableInt(2), comparableInt(3)},
				toRemove: comparableInt(3),
			},
			want: []comparableInt{comparableInt(1), comparableInt(2)},
		},
		{
			name: "found several times",
			args: args{
				ts:       []comparableInt{comparableInt(1), comparableInt(2), comparableInt(3), comparableInt(2), comparableInt(3)},
				toRemove: comparableInt(2),
			},
			want: []comparableInt{comparableInt(1), comparableInt(3), comparableInt(3)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slices.Remove(tt.args.ts, tt.args.toRemove), "Remove(%v, %v)", tt.args.ts, tt.args.toRemove)
		})
	}
}
