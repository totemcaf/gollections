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
