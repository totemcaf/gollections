package slices

import (
	"github.com/totemcaf/gollections/sets"
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

func FlatMap[SS ~[]S, S any, TS ~[]T, T any](ss SS, m types.Mapper[S, TS]) TS {
	tt := make(TS, 0, len(ss))
	for _, s := range ss {
		tt = append(tt, m(s)...)
	}
	return tt
}

func Filter[SS ~[]S, S any](ss SS, p types.Predicate[S]) SS {
	var tt SS
	for _, s := range ss {
		if p(s) {
			tt = append(tt, s)
		}
	}
	return tt
}

// FilterNot removes (filter out) all elements the satisfies predicate
func FilterNot[SS ~[]S, S any](ss SS, p types.Predicate[S]) SS {
	var tt SS
	for _, s := range ss {
		if !p(s) {
			tt = append(tt, s)
		}
	}
	return tt
}

// Remove all occurrences of the given element from the slice
func Remove[TS ~[]T, T types.Comparable[T]](ts TS, toRemove T) TS {
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
func FilterNonNil[TS ~[]*T, T any](ts TS) TS {
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

// Has2 returns true if the slice contains the given element
// Deprecated: use Contains instead
func Has2[T comparable](ts []T, other T) bool {
	return Contains(ts, other)
}

// Contains returns true if the slice contains the given element
func Contains[T comparable](ts []T, other T) bool {
	return Index(ts, other) >= 0
}

func DeepClone[SS ~[]S, S types.Cloneable[S]](source SS) SS {
	return Map(source, func(t S) S { return t.Clone() })
}

func Clone[TS ~[]T, T any](source TS) TS {
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

func JoinDistinct[TS ~[]T, T comparable](es ...TS) TS {
	result := sets.New[T]()

	for _, e := range es {
		result.AddAll(e...)
	}

	return result.Values()
}

// CastAll converts the type of all elements of the slice
// If an element cannot be converted, it will panic
func CastAll[SS ~[]S, S any, TS ~[]T, T any](ss SS) TS {
	tt := make(TS, len(ss))
	for idx, s := range ss {
		var a any = s
		tt[idx] = a.(T)
	}
	return tt
}

// HasDuplicates returns true if the slice contains at least one element duplicated
func HasDuplicates[TS ~[]T, T comparable](ts TS) bool {
	return len(ts) != sets.Of[T](ts...).Size()
}

// RemoveAt removes the element at the given index by shifting all elements after it to the left
func RemoveAt[T any](elements []T, idx int) []T { //
	if idx < 0 || idx >= len(elements) {
		return elements
	}
	newElements := make([]T, len(elements)-1)
	copy(newElements, elements[:idx])
	copy(newElements[idx:], elements[idx+1:])
	return newElements
}

// InsertAt inserts the given element at the given index by shifting all elements after it to the right
func InsertAt[T any](elements []T, idx int, element T) []T {
	if idx < 0 || idx > len(elements) {
		panic("index out of bounds")
	}
	newElements := make([]T, len(elements)+1)
	copy(newElements, elements[:idx])
	newElements[idx] = element
	copy(newElements[idx+1:], elements[idx:])
	return newElements
}
