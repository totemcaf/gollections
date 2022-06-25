package lists

import (
	"github.com/totemcaf/gollections/types"
)

// Map computes a new list with the result of applying function mapper to the source list
// This is implemented as a method in List[T] because GO generics does not support (yet) the
// use of type parameters in methods signatures.
func Map[S any, T any](src List[S], m types.Mapper[S, T]) List[T] {
	values := src.Values()

	target := make([]T, len(values))

	for idx, s := range values {
		target[idx] = m(s)
	}

	return Of(target...)
}

// Reduce convert this list in a single value of the same type. Starts with the zero value of the type
// This is implemented as a method in List[T] because GO generics does not support (yet) the
// use of type parameters in methods signatures.
func Reduce[S any, T any](src List[S], reducer func(accum T, element S) T) T {
	var initial T
	return Fold(src, initial, reducer)
}

// Fold convert this list in a single value of the same type
// This is implemented as a method in List[T] because GO generics does not support (yet) the
// use of type parameters in methods signatures.
func Fold[S any, T any](src List[S], initial T, reducer func(accum T, element S) T) T {
	values := src.Values()

	result := initial

	for _, s := range values {
		result = reducer(result, s)
	}

	return result
}
