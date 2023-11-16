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

	assert.Equal(t, aNil, CastOrNil[*string](aNil))
}

type ISample interface {
	GetValue() string
}

type sample struct {
	aValue string
}

func (s *sample) GetValue() string {
	return s.aValue
}

func Test_CastOrNil_returns_casted_value_for_not_nil(t *testing.T) {
	var aValue sample = sample{aString}

	assert.Equal(t, &aValue, CastOrNil[*sample](&aValue))
	assert.IsType(t, &sample{}, CastOrNil[*sample](&aValue))
}

func Test_CastOrNil_returns_casted_value_for_interfaces(t *testing.T) {
	var aValue sample = sample{aString}

	assert.Equal(t, &aValue, CastOrNil[ISample](&aValue))
	assert.IsType(t, &sample{}, CastOrNil[ISample](&aValue))
}

func TestOrDefaultF(t *testing.T) {
	type args[T any] struct {
		value        *T
		defaultValue func() T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[string]{
		{
			name: "value is not nil",
			args: args[string]{
				value:        &aString,
				defaultValue: func() string { return "bbb" },
			},
			want: aString,
		},
		{
			name: "value is nil",
			args: args[string]{
				value:        nil,
				defaultValue: func() string { return "bbb" },
			},
			want: "bbb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, OrDefaultF(tt.args.value, tt.args.defaultValue), "OrDefaultF(%v, %v)", tt.args.value, tt.args.defaultValue)
		})
	}
}

func TestCoalesce(t *testing.T) {
	type args[T any] struct {
		values []*T
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		want      T
		wantPanic bool
	}
	tests := []testCase[string]{
		{
			name: "first value is not nil",
			args: args[string]{
				values: []*string{&aString, nil},
			},
			want: aString,
		},
		{
			name: "first value is nil",
			args: args[string]{
				values: []*string{nil, &aString},
			},
			want: aString,
		},
		{
			name: "all values are nil",
			args: args[string]{
				values: []*string{nil, nil, nil, nil},
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panicsf(t, func() { Coalesce(tt.args.values...) }, "Coalesce(%v)", tt.args.values)
			} else {
				assert.Equalf(t, tt.want, Coalesce(tt.args.values...), "Coalesce(%v)", tt.args.values)
			}
		})
	}
}
