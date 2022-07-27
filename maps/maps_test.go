package maps

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceEntries(t *testing.T) {
	a := map[string]int{
		"3": 10,
		"1": 20,
		"2": 30,
	}

	reducer := func(result int, key string, value int) int {
		keyNum, _ := strconv.Atoi(key)
		return result + keyNum*value
	}

	reduction := ReduceEntries(1000000, a, reducer)

	assert.Equal(t, 1000000+3*10+1*20+2*30, reduction)
}

func TestReduceKeys(t *testing.T) {
	a := map[string]int{
		"3": 10,
		"1": 20,
		"2": 30,
	}

	reducer := func(result int, key string) int {
		keyNum, _ := strconv.Atoi(key)
		return result + keyNum
	}

	reduction := ReduceKeys(100000, a, reducer)

	assert.Equal(t, 100000+3+1+2, reduction)
}

func TestReduceValues(t *testing.T) {
	a := map[string]int{
		"3": 10,
		"1": 20,
		"2": 30,
	}

	reducer := func(result int, value int) int {
		return result + value
	}

	reduction := ReduceValues(100000, a, reducer)

	assert.Equal(t, 100000+10+20+30, reduction)
}

func TestMap(t *testing.T) {
	mapper := func(value int) string {
		return strconv.Itoa(value)
	}

	tests := []struct {
		name string
		args map[string]int
		want map[string]string
	}{
		{
			name: "with value",
			args: map[string]int{"one": 1, "two": 2, "three": 3},
			want: map[string]string{"one": "1", "two": "2", "three": "3"},
		}, {
			name: "empty",
			args: map[string]int{},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Map(tt.args, mapper), "Map(%v, mapper)", tt.args)
		})
	}
}

func TestMapKeys(t *testing.T) {
	mapper := func(value int) string {
		return strconv.Itoa(value)
	}

	tests := []struct {
		name string
		args map[int]string
		want map[string]string
	}{
		{
			name: "with value",
			args: map[int]string{1: "one", 2: "two", 3: "three"},
			want: map[string]string{"1": "one", "2": "two", "3": "three"},
		}, {
			name: "empty",
			args: map[int]string{},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MapKeys(tt.args, mapper), "MapKeys(%v, mapper)", tt.args)
		})
	}
}

func TestMapEntries(t *testing.T) {
	mapper := func(key string, value int) (string, string) {
		return key + ":" + key, strconv.Itoa(value)
	}

	tests := []struct {
		name string
		args map[string]int
		want map[string]string
	}{
		{
			name: "with value",
			args: map[string]int{"one": 1, "two": 2, "three": 3},
			want: map[string]string{"one:one": "1", "two:two": "2", "three:three": "3"},
		}, {
			name: "empty",
			args: map[string]int{},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MapEntries(tt.args, mapper), "MapEntries(%v, mapper)", tt.args)
		})
	}
}
