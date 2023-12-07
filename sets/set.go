package sets

import "fmt"

type Set[T comparable] map[T]struct{}

// New creates a new empty set.
func New[T comparable]() Set[T] {
	return make(Set[T])
}

// Of creates a new set with the given elements.
func Of[T comparable](ts ...T) Set[T] {
	s := make(Set[T])
	s.AddAll(ts...)
	return s
}

// Add adds the given element to the set.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// AddAll adds the given elements to the set.
func (s Set[T]) AddAll(v ...T) {
	for _, v := range v {
		s.Add(v)
	}
}

// Remove removes the given element from the set.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// Contains returns true if the set contains the given element.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Size returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s)
}

// Values returns the elements of the set as a slice.
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for v := range s {
		values = append(values, v)
	}
	return values
}

// Union returns a new set with all the elements of the set and the given set.
// Common elements are only added once.
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

// Intersection returns a new set with the elements that are in both this set and the other set.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := New[T]()
	for v := range s {
		if other.Contains(v) {
			intersection.Add(v)
		}
	}
	return intersection
}

// Difference returns a new set with the elements that are in this set but not in the other.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := New[T]()
	for v := range s {
		if !other.Contains(v) {
			difference.Add(v)
		}
	}
	return difference
}

// SymmetricDifference returns a new set with the elements that are in this set or the other set but not in both.
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

// IsSubset returns true if this set is a subset of the other set. A set is a subset if all elements of this set are
// also in the other set.
func (s Set[T]) IsSubset(other Set[T]) bool {
	for v := range s {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if this set is a superset of the other set. A set is a superset if all elements of the
// other set are also in this set.
func (s Set[T]) IsSuperset(other Set[T]) bool {
	for v := range other {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the set is empty. It has no elements.
func (s Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Clear removes all elements from the set.
func (s Set[T]) Clear() {
	for v := range s {
		delete(s, v)
	}
}

// Copy returns a copy of the set.
func (s Set[T]) Copy() Set[T] {
	setCopy := New[T]()
	for v := range s {
		setCopy.Add(v)
	}
	return setCopy
}

// Equal returns true if the set is equal to the other set. Two sets are equal if they have the same elements.
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

// String returns a string representation of the set.
func (s Set[T]) String() string {
	str := "{"
	for v := range s {
		str += fmt.Sprintf("%v ", v)
	}
	str += "}"
	return str
}

// GoString returns a Go string representation of the set.
func (s Set[T]) GoString() string {
	return s.String()
}
