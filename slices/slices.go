package slices

import (
	"github.com/totemcaf/gollections/types"
)

func Of[S any](ss ...S) []S {
	return ss
}

func Map[S any, T any](ss []S, m types.Mapper[S, T]) []T {
	tt := make([]T, len(ss))
	for idx, s := range ss {
		tt[idx] = m(s)
	}
	return tt
}

// MapWithError applies mapper to all elements of 'ss' and returns slice of results.
// If mapper returns error, then it is returned immediately.
func MapWithError[S any, T any](ss []S, m types.MapperWithError[S, T]) ([]T, error) {
	tt := make([]T, len(ss))
	var err error
	for idx, s := range ss {
		tt[idx], err = m(s)
		if err != nil {
			return nil, err
		}
	}
	return tt, nil
}

// MapNonNil apply mapper to all elements of 'ss' and return slice of all non-nil results
func MapNonNil[S any, T any](ss []S, mapper types.Mapper[S, *T]) []T {
	us := make([]T, 0, len(ss))

	for _, s := range ss {
		if t := mapper(s); t != nil {
			us = append(us, *t)
		}
	}

	return us
}

func FlatMap[S any, T any](ss []S, m types.Mapper[S, []T]) []T {
	tt := make([]T, 0, len(ss))
	for _, s := range ss {
		tt = append(tt, m(s)...)
	}
	return tt
}

func Filter[S any](ss []S, p types.Predicate[S]) []S {
	var tt []S
	for _, s := range ss {
		if p(s) {
			tt = append(tt, s)
		}
	}
	return tt
}

// FilterNot removes (filter out) all elements the satisfies predicate
func FilterNot[S any](ss []S, p types.Predicate[S]) []S {
	var tt []S
	for _, s := range ss {
		if !p(s) {
			tt = append(tt, s)
		}
	}
	return tt
}

// Remove all occurrences of the given element from the slice
func Remove[T types.Comparable[T]](ts []T, toRemove T) []T {
	return FilterNot(ts, func(t T) bool { return toRemove.Compare(t) == 0 })
}

// Any returns true if at least one element satisfies predicate
func Any[T any](ts []T, predicate func(t T) bool) bool {
	for _, t := range ts {
		if predicate(t) {
			return true
		}
	}
	return false
}

// FilterNonNil returns all the elements of array that are non nil
func FilterNonNil[T any](ts []*T) []*T {
	return Filter(ts, func(t *T) bool { return t != nil })
}

func Count[S any](ss []S, p types.Predicate[S]) int {
	var count = 0
	for _, s := range ss {
		if p(s) {
			count++
		}
	}
	return count
}

func Find[S any](ss []S, p types.Predicate[S]) (S, bool) {
	for _, s := range ss {
		if p(s) {
			return s, true
		}
	}
	var notFound S
	return notFound, false
}

func Index[S comparable](ss []S, toFind S) int {
	for idx, s := range ss {
		if s == toFind {
			return idx
		}
	}
	return -1
}

func Index2[S comparable](ss []S, toFind S) (int, bool) {
	idx := Index(ss, toFind)
	return idx, idx >= 0
}

func IndexBy[S any](ss []S, p types.Predicate[S]) int {
	for idx, s := range ss {
		if p(s) {
			return idx
		}
	}
	return -1
}

func IndexBy2[S any](ss []S, p types.Predicate[S]) (int, bool) {
	idx := IndexBy(ss, p)
	return idx, idx >= 0
}

func Has[T types.Comparable[T]](ts []T, other T) bool {
	for _, t := range ts {
		if t.Compare(other) == 0 {
			return true
		}
	}
	return false
}

func Has2[T comparable](ts []T, other T) bool {
	for _, t := range ts {
		if t == other {
			return true
		}
	}
	return false
}

func DeepClone[T types.Cloneable[T]](source []T) []T {
	return Map(source, func(t T) T { return t.Clone() })
}

func Clone[T any](source []T) []T {
	target := make([]T, len(source))

	copy(target, source)

	return target
}

func Reduce[Value any, Element any](
	initialValue Value,
	elements []Element,
	reducer func(Value, Element) Value,
) Value {

	accum := initialValue

	for _, element := range elements {
		accum = reducer(accum, element)
	}

	return accum
}
