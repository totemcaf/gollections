package cond

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	f1 := func() int { return 1 }
	f2 := func() int { return 2 }
	f3 := func() int { return 3 }

	assert.Equal(t, 1, If(true, 1).Else(2))
	assert.Equal(t, 2, If(false, 1).Else(2))
	assert.Equal(t, 1, If(true, 1).ElseF(f2))
	assert.Equal(t, 2, If(false, 1).ElseF(f2))

	assert.Equal(t, 1, If(true, 1).ElseIf(true, 2).Else(3))
	assert.Equal(t, 2, If(false, 1).ElseIf(true, 2).Else(3))
	assert.Equal(t, 3, If(false, 1).ElseIf(false, 2).Else(3))
	assert.Equal(t, 1, If(true, 1).ElseIf(true, 2).ElseF(f3))
	assert.Equal(t, 2, If(false, 1).ElseIf(true, 2).ElseF(f3))
	assert.Equal(t, 3, If(false, 1).ElseIf(false, 2).ElseF(f3))

	assert.Equal(t, 1, IfF(true, f1).Else(2))
	assert.Equal(t, 2, IfF(false, f1).Else(2))
	assert.Equal(t, 1, IfF(true, f1).ElseF(f2))
	assert.Equal(t, 3, IfF(false, f1).ElseIf(false, 2).ElseF(f3))
}
