package sets

import "fmt"

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) AddAll(v ...T) {
	for _, v := range v {
		s.Add(v)
	}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for v := range s {
		values = append(values, v)
	}
	return values
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := New[T]()
	for v := range s {
		union.Add(v)
	}
	for v := range other {
		union.Add(v)
	}
	return union
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := New[T]()
	for v := range s {
		if other.Contains(v) {
			intersection.Add(v)
		}
	}
	return intersection
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := New[T]()
	for v := range s {
		if !other.Contains(v) {
			difference.Add(v)
		}
	}
	return difference
}

func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	symmetricDifference := New[T]()
	for v := range s {
		if !other.Contains(v) {
			symmetricDifference.Add(v)
		}
	}
	for v := range other {
		if !s.Contains(v) {
			symmetricDifference.Add(v)
		}
	}
	return symmetricDifference
}

func (s Set[T]) IsSubset(other Set[T]) bool {
	for v := range s {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsSuperset(other Set[T]) bool {
	for v := range other {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s Set[T]) Clear() {
	for v := range s {
		delete(s, v)
	}
}

func (s Set[T]) Copy() Set[T] {
	setCopy := New[T]()
	for v := range s {
		setCopy.Add(v)
	}
	return setCopy
}

func (s Set[T]) Equal(other Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for v := range s {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

func (s Set[T]) String() string {
	str := "{"
	for v := range s {
		str += fmt.Sprintf("%v ", v)
	}
	str += "}"
	return str
}

func (s Set[T]) GoString() string {
	return s.String()
}
