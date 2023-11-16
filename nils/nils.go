package nils

import "reflect"

func IsNil[T any](t *T) bool {
	return t == nil
}

func IsNotNil[T any](t *T) bool {
	return !IsNil(t)
}

// OrDefault Return value if no nil of default value in the another case
func OrDefault[T any](value *T, defaultValue T) T {
	if value != nil {
		return *value
	}
	return defaultValue
}

// OrDefaultF Return value if no nil of default value from function in the another case
func OrDefaultF[T any](value *T, defaultValue func() T) T {
	if value != nil {
		return *value
	}
	return defaultValue()
}

// Coalesce Returns first no-nil value or panic if none
func Coalesce[T any](values ...*T) T {
	for _, v := range values {
		if v != nil {
			return *v
		}
	}

	panic("All values are nil, provide at least one no nil value")
}

func Copy[T any](value *T) *T {
	if value == nil {
		return nil
	}

	t := *value

	return &t
}

func CastOrNil[T any](value any) T {
	if value == nil {
		var t T
		return t
	}

	t, ok := value.(T)

	if !ok {
		panic("cannot cast value to type " + reflect.TypeOf(value).String())
	}
	return t
}
