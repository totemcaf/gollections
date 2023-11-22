package slices_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/totemcaf/gollections/slices"
	"github.com/totemcaf/gollections/types"
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

func TestClone_duplicates(t *testing.T) {
	// GIVEN a slice
	s := []int{1, 2, 3}

	// WHEN I clone it
	c := slices.Clone(s)

	// THEN the clone is equal to the original
	assert.Equal(t, s, c)

	// AND the clone is not the same slice
	assert.NotSame(t, s, c)
}

func TestFilterNonNil_returns_not_nil_elements(t *testing.T) {
	// GIVEN a slice with nil and not nil elements
	one := 1
	two := 2
	three := 3
	s := []*int{nil, &one, nil, &two, &three, nil}

	// WHEN I filter it
	f := slices.FilterNonNil(s)

	// THEN the result contains only not nil elements
	assert.Equal(t, []*int{&one, &two, &three}, f)
}

func TestIndexBy2_can_find(t *testing.T) {
	// GIVEN a slice
	s := []int{1, 2, 3}

	// WHEN I index it
	i, found := slices.IndexBy2(s, func(i int) bool {
		return i == 2
	})

	// THEN I can find the element
	assert.True(t, found)
	assert.Equal(t, 1, i)
}

func TestIndexBy2(t *testing.T) {
	equalTo := func(i int) func(int) bool {
		return func(j int) bool {
			return i == j
		}
	}

	type args struct {
		ss []int
		p  types.Predicate[int]
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantFound bool
	}{
		{
			name: "empty",
			args: args{
				ss: []int{},
				p:  equalTo(1),
			},
			want:      -1,
			wantFound: false,
		},
		{
			name: "not found",
			args: args{
				ss: []int{1, 2, 3},
				p:  equalTo(4),
			},
			want:      -1,
			wantFound: false,
		},
		{
			name: "found",
			args: args{
				ss: []int{1, 2, 3},
				p:  equalTo(2),
			},
			want:      1,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := slices.IndexBy2(tt.args.ss, tt.args.p)
			assert.Equalf(t, tt.want, got, "IndexBy2(%v, %v)", tt.args.ss, tt.args.p)
			assert.Equalf(t, tt.wantFound, got1, "IndexBy2(%v, %v)", tt.args.ss, tt.args.p)
		})
	}
}

type sampleComparable string

func (s sampleComparable) Compare(other sampleComparable) int {
	return strings.Compare(string(s), string(other))
}

func TestHas(t *testing.T) {
	type args struct {
		ts    []sampleComparable
		other sampleComparable
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				ts:    []sampleComparable{},
				other: sampleComparable("something"),
			},
			want: false,
		},
		{
			name: "not found",
			args: args{
				ts:    []sampleComparable{sampleComparable("something"), sampleComparable("else")},
				other: sampleComparable("not found"),
			},
			want: false,
		},
		{
			name: "found",
			args: args{
				ts:    []sampleComparable{sampleComparable("something"), sampleComparable("else")},
				other: sampleComparable("something"),
			},
			want: true,
		},
		{
			name: "found several times",
			args: args{
				ts:    []sampleComparable{sampleComparable("something"), sampleComparable("else"), sampleComparable("something")},
				other: sampleComparable("something"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slices.Has(tt.args.ts, tt.args.other), "Has(%v, %v)", tt.args.ts, tt.args.other)
		})
	}
}

type sampleStruct struct {
	foo string
}

func (s sampleStruct) Clone() sampleStruct {
	return sampleStruct{foo: s.foo}
}

func TestDeepClone(t *testing.T) {
	// GIVEN a slice
	s := []sampleStruct{{foo: "bar"}}

	// WHEN I clone it
	c := slices.DeepClone(s)

	// THEN the clone is equal to the original
	assert.Equal(t, s, c)

	// AND the clone is not the same slice
	assert.NotSame(t, s, c)

	// AND the clone is not the same struct
	assert.NotSame(t, s[0], c[0])
}

func TestMapWithError(t *testing.T) {
	doublePositive := func(i int) (int, error) {
		if i < 0 {
			return 0, errors.New("negative")
		}
		return i * 2, nil
	}

	tests := []struct {
		name    string
		args    []int
		want    []int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    []int{},
			want:    []int{},
			wantErr: assert.NoError,
		},
		{
			name:    "no error",
			args:    []int{1, 2, 3},
			want:    []int{2, 4, 6},
			wantErr: assert.NoError,
		},
		{
			name:    "error",
			args:    []int{1, -2, 3},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := slices.MapWithError(tt.args, doublePositive)
			if !tt.wantErr(t, err, fmt.Sprintf("MapWithError(%v, %v)", tt.args, "doublePositive")) {
				return
			}
			assert.Equalf(t, tt.want, got, "MapWithError(%v, %v)", tt.args, "doublePositive")
		})
	}
}

func TestCastAll(t *testing.T) {
	type testCase[S any, T any] struct {
		name string
		args []S
		want []T
	}
	tests := []testCase[interface{}, string]{
		{
			name: "empty",
			args: []interface{}{},
			want: []string{},
		},
		{
			name: "one element",
			args: []interface{}{"hello"},
			want: []string{"hello"},
		},
		{
			name: "several elements",
			args: []interface{}{"hello", "world", "!"},
			want: []string{"hello", "world", "!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slices.CastAll[[]any, any, []string, string](tt.args), "CastAll(%v)", tt.args)
		})
	}
}

func TestFilter(t *testing.T) {
	type args[SS ~[]S, S any] struct {
		ss SS
		p  types.Predicate[S]
	}
	type testCase[SS ~[]S, S any] struct {
		name string
		args args[SS, S]
		want SS
	}
	tests := []testCase[[]string, string]{
		{
			name: "empty",
			args: args[[]string, string]{
				ss: []string{},
				p:  func(s string) bool { return len(s) > 2 },
			},
			want: nil,
		},
		{
			name: "not found",
			args: args[[]string, string]{
				ss: []string{"a", "b", "c"},
				p:  func(s string) bool { return len(s) > 2 },
			},
			want: nil,
		},
		{
			name: "found",
			args: args[[]string, string]{
				ss: []string{"a", "bb", "ccc"},
				p:  func(s string) bool { return len(s) > 2 },
			},
			want: []string{"ccc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, slices.Filter(tt.args.ss, tt.args.p), "Filter(%v)", tt.args.ss)
		})
	}
}
