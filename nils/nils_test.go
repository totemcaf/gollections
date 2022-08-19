package nils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var aString = "aaa"
var aNumber = 10

func Test_IsNil_returns_true_for_nil(t *testing.T) {
	var aNil *struct{} = nil
	assert.True(t, IsNil(aNil))
}

func Test_IsNil_returns_false_for_not_nil(t *testing.T) {
	assert.False(t, IsNil(&aString))
	assert.False(t, IsNil(&aNumber))
}

func Test_IsNotNil_returns_false_for_nil(t *testing.T) {
	var aNil *struct{} = nil
	assert.False(t, IsNotNil(aNil))
}

func Test_IsNotNil_returns_true_for_not_nil(t *testing.T) {
	assert.True(t, IsNotNil(&aString))
	assert.True(t, IsNotNil(&aNumber))
}

func Test_OrDefault_returns_value_if_not_nil(t *testing.T) {
	assert.Equal(t, aString, OrDefault(&aString, "bbb"))
	assert.Equal(t, aNumber, OrDefault(&aNumber, 20))
}

func Test_OrDefault_returns_default_value_if_nil(t *testing.T) {
	assert.Equal(t, "bbb", OrDefault(nil, "bbb"))
	assert.Equal(t, 20, OrDefault(nil, 20))
}

func Test_Copy_returns_nil_for_nil(t *testing.T) {
	var aNil *struct{} = nil
	assert.Equal(t, aNil, Copy(aNil))
}

func Test_Copy_returns_copied_value_for_not_nil(t *testing.T) {
	assert.Equal(t, aString, *Copy(&aString))
	assert.Equal(t, aNumber, *Copy(&aNumber))
	assert.NotSame(t, aString, *Copy(&aString))
	assert.NotSame(t, aNumber, *Copy(&aNumber))
}

func Test_CastOrNil_returns_nil_for_nil(t *testing.T) {
	var aNil *string = nil

	assert.Equal(t, aNil, CastOrNil[string](aNil))
}

type sample struct {
	aValue string
}

func Test_CastOrNil_returns_casted_value_for_not_nil(t *testing.T) {
	var aValue sample = sample{aString}

	assert.Equal(t, &aValue, CastOrNil[sample](&aValue))
	assert.IsType(t, &sample{}, CastOrNil[sample](&aValue))
}
