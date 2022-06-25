package slices_test

import (
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
