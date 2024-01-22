package ptrs

import "golang.org/x/exp/constraints"

// Number is a constraint that permits any numerical type.
type Number interface {
	constraints.Integer | constraints.Float
}

// SetIfPresent sets target to value if value is not nil
func SetIfPresent[T any](target *T, value *T) {
	if value != nil {
		*target = *value
	}
}

// Ptr returns a pointer to a value
func Ptr[T any](value T) *T {
	return &value
}

// True returns a pointer to true
func True() *bool {
	return Ptr(true)
}

// False returns a pointer to false
func False() *bool {
	return Ptr(false)
}

// Zero returns a pointer to zero
func Zero[T Number]() *T {
	t := T(0)
	return &t
}

// One returns a pointer to one
func One[T Number]() *T {
	t := T(1)
	return &t
}

// Coalesce returns the first non-nil pointer in the list
func Coalesce[T any](values ...*T) *T {
	for _, v := range values {
		if v != nil {
			return v
		}
	}
	return nil
}
