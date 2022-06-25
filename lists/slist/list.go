package slist

import (
	"reflect"

	"github.com/totemcaf/gollections/lists"
	"github.com/totemcaf/gollections/types"
)

type sliceList[T any] struct {
	es []T
}

func Empty[T any]() lists.List[T] {
	return &sliceList[T]{nil}
}

func Of[T any](e ...T) lists.List[T] {
	return Empty[T]().AppendAll(e...)
}

func (s *sliceList[T]) Append(t T) lists.List[T] {
	result := make([]T, len(s.es), len(s.es)+1)
	copy(result, s.es)

	ts := append(result, t)
	return &sliceList[T]{ts}
}

func (s *sliceList[T]) AppendAll(t ...T) lists.List[T] {
	if len(t) == 0 {
		return s
	}

	result := make([]T, len(s.es), len(s.es)+len(t))
	copy(result, s.es)

	ts := append(result, t...)
	return &sliceList[T]{ts}
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

func (s *sliceList[T]) Map(mapper func(T) T) lists.List[T] {
	if len(s.es) == 0 {
		return s
	}

	result := make([]T, len(s.es))

	for idx, e := range s.es {
		result[idx] = mapper(e)
	}

	return &sliceList[T]{result}
}

func (s *sliceList[T]) FilterBy(predicate types.Predicate[T]) lists.List[T] {
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
