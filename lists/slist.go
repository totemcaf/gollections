package lists

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/totemcaf/gollections/slices"
	"github.com/totemcaf/gollections/types"
)

type sliceList[T any] struct {
	es []T
}

func Empty[T any]() List[T] {
	return &sliceList[T]{nil}
}

func Of[T any](e ...T) List[T] {
	return Empty[T]().AppendAll(e...)
}

func (s *sliceList[T]) Values() []T {
	slice := make([]T, len(s.es))
	copy(slice, s.es)

	return slice
}

func (s *sliceList[T]) Append(t T) List[T] {
	result := make([]T, len(s.es), len(s.es)+1)
	copy(result, s.es)

	ts := append(result, t)
	return &sliceList[T]{ts}
}

func (s *sliceList[T]) AppendAll(t ...T) List[T] {
	if len(t) == 0 {
		return s
	}

	result := make([]T, len(s.es), len(s.es)+len(t))
	copy(result, s.es)

	ts := append(result, t...)
	return &sliceList[T]{ts}
}

func (s *sliceList[T]) Concat(second List[T]) List[T] {
	return s.AppendAll(second.Values()...)
}

func (s *sliceList[T]) Count() int {
	return len(s.es)
}

func (s *sliceList[T]) CountBy(predicate types.Predicate[T]) int {
	count := 0
	for _, e := range s.es {
		if predicate(e) {
			count++
		}
	}
	return count
}

func (s *sliceList[T]) At2(idx int) (T, bool) {
	if idx < 0 || idx >= len(s.es) {
		var empty T
		return empty, false
	}
	return s.es[idx], true
}

func (s *sliceList[T]) At(idx int) T {
	e, _ := s.At2(idx)
	return e
}

func (s *sliceList[T]) Map(mapper func(T) T) List[T] {
	if len(s.es) == 0 {
		return s
	}

	result := make([]T, len(s.es))

	for idx, e := range s.es {
		result[idx] = mapper(e)
	}

	return &sliceList[T]{result}
}

// Reduce convert this list in a single value of the same type
func (s *sliceList[T]) Reduce(reducer func(accum T, element T) T) T {
	var result T
	return s.Fold(result, reducer)
}

// Fold convert this list in a single value of the same type
func (s *sliceList[T]) Fold(initial T, reducer func(accum T, element T) T) T {
	result := initial

	for _, e := range s.es {
		result = reducer(result, e)
	}

	return result
}

func (s *sliceList[T]) FilterBy(predicate types.Predicate[T]) List[T] {
	if len(s.es) == 0 {
		return s
	}

	var result []T

	for _, e := range s.es {
		if predicate(e) {
			result = append(result, e)
		}
	}

	return &sliceList[T]{result}
}

func (s *sliceList[T]) Any(predicate types.Predicate[T]) bool {
	for _, e := range s.es {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (s *sliceList[T]) areEqual(a interface{}, b T) bool {
	if ac, ok := a.(types.Comparable[T]); ok {
		return ac.Compare(b) == 0
	}
	return reflect.DeepEqual(a, b)
}

func (s *sliceList[T]) All(predicate types.Predicate[T]) bool {
	for _, e := range s.es {
		if !predicate(e) {
			return false
		}
	}
	return true
}

func (s *sliceList[T]) Index(t T) int {
	for idx, e := range s.es {
		if s.areEqual(e, t) {
			return idx
		}
	}
	return -1
}

func (s *sliceList[T]) Index2(t T) (int, bool) {
	idx := s.Index(t)
	return idx, idx >= 0
}

func (s *sliceList[T]) IndexBy(predicate types.Predicate[T]) int {
	for idx, e := range s.es {
		if predicate(e) {
			return idx
		}
	}
	return -1
}

func (s *sliceList[T]) IndexBy2(t types.Predicate[T]) (int, bool) {
	idx := s.IndexBy(t)
	return idx, idx >= 0
}

func (s *sliceList[T]) toString(x T) string {
	return fmt.Sprintf("%v", x)
}

func (s *sliceList[T]) Join(separator string) string {
	return strings.Join(slices.Map(s.es, s.toString), separator)
}
