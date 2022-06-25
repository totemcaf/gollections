package lists

import (
	"reflect"
	"strings"
	"testing"
)

func Test_sliceList_AppendAll(t *testing.T) {
	tests := []struct {
		name string
		s    List[string]
		args []string
		want List[string]
	}{
		{
			name: "Empty and empty",
			s:    Empty[string](),
			args: []string{},
			want: Empty[string](),
		},
		{
			name: "Empty and element",
			s:    Empty[string](),
			args: []string{"hello"},
			want: Of("hello"),
		},
		{
			name: "Element and empty",
			s:    Of("hello"),
			args: []string{},
			want: Of("hello"),
		},
		{
			name: "Element and element",
			s:    Of("hello"),
			args: []string{"world"},
			want: Of("hello", "world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AppendAll(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceList_Count(t *testing.T) {
	tests := []struct {
		name string
		s    List[string]
		want int
	}{
		{
			name: "empty list",
			s:    Empty[string](),
			want: 0,
		},
		{
			name: "one element",
			s:    Of("hello"),
			want: 1,
		},
		{
			name: "several elements",
			s:    Of("nice", "example", "to", "test"),
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceList_At2(t *testing.T) {
	tests := []struct {
		name  string
		s     List[string]
		idx   int
		want  string
		found bool
	}{
		{
			name:  "no elements",
			s:     Empty[string](),
			idx:   0,
			want:  "",
			found: false,
		},
		{
			name:  "found",
			s:     Of("hello", "nice", "world"),
			idx:   0,
			want:  "hello",
			found: true,
		},
		{
			name:  "found 1",
			s:     Of("hello", "nice", "world"),
			idx:   1,
			want:  "nice",
			found: true,
		},
		{
			name:  "found 2",
			s:     Of("hello", "nice", "world"),
			idx:   2,
			want:  "world",
			found: true,
		},
		{
			name:  "after end",
			s:     Of("hello", "nice", "world"),
			idx:   3,
			want:  "",
			found: false,
		},
		{
			name:  "before start",
			s:     Of("hello", "nice", "world"),
			idx:   -1,
			want:  "",
			found: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.At2(tt.idx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("At2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.found {
				t.Errorf("At2() got1 = %v, want %v", got1, tt.found)
			}
		})
	}
}

func Test_sliceList_Map(t *testing.T) {
	capitalize := func(s string) string { return strings.ToUpper(s) }

	tests := []struct {
		name string
		s    List[string]
		want List[string]
	}{
		{
			name: "no elements",
			s:    Empty[string](),
			want: Of[string](),
		},
		{
			name: "Some elements",
			s:    Of("hello", "nice", "world"),
			want: Of("HELLO", "NICE", "WORLD"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Map(capitalize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceList_FilterBy(t *testing.T) {
	isLenGT4 := func(s string) bool { return len(s) > 4 }

	tests := []struct {
		name string
		s    List[string]
		idx  int
		want List[string]
	}{
		{
			name: "no elements",
			s:    Empty[string](),
			want: Of[string](),
		},
		{
			name: "Some elements match",
			s:    Of("hello", "nice", "world"),
			want: Of("hello", "world"),
		},
		{
			name: "No elements match",
			s:    Of("no", "one", "is", "good"),
			want: Of[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.FilterBy(isLenGT4)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}
