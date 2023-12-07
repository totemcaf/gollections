package cond

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {

	assert.Equal(t, 1, If(true, 1).Else(2))
	assert.Equal(t, 2, If(false, 1).Else(2))
	assert.Equal(t, 1, If(true, 1).ElseF(func() int { return 2 }))
	assert.Equal(t, 2, If(false, 1).ElseF(func() int { return 2 }))

	assert.Equal(t, 1, If(true, 1).ElseIf(true, 2).Else(3))
	assert.Equal(t, 2, If(false, 1).ElseIf(true, 2).Else(3))
	assert.Equal(t, 3, If(false, 1).ElseIf(false, 2).Else(3))
	assert.Equal(t, 1, If(true, 1).ElseIf(true, 2).ElseF(func() int { return 3 }))
	assert.Equal(t, 2, If(false, 1).ElseIf(true, 2).ElseF(func() int { return 3 }))
	assert.Equal(t, 3, If(false, 1).ElseIf(false, 2).ElseF(func() int { return 3 }))
}
