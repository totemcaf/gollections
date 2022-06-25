package lists

import "github.com/totemcaf/gollections/types"

type List[T any] interface {
	Unmarshal(data []byte) error
	// Append returns a list with all the elements of this and a new one to the end of the list
	Append(T) List[T]
	// AppendAll returns a list with all the elements of this and a all new ones to the end of the list
	AppendAll(...T) List[T]
	// Count returns number of elements in list
	Count() int
	// CountBy returns number of elements in list that satisfies predicate
	CountBy(predicate types.Predicate[T]) int
	// At2 returns element at idx, or report empty element if idx <0 or >= List.Count
	At2(idx int) (T, bool)
	// At returns element at idx, or empty if idx <0 or >= List.Count
	At(idx int) T
	// Map converts a list of one type to another list of same type with possible different values
	Map(mapper func(T) T) List[T]

	// FilterBy returns list with all elements except the ones that satisfies predicate
	FilterBy(types.Predicate[T]) List[T]

	// Any returns true if at least one element satisfies predicate
	Any(types.Predicate[T]) bool
	// All returns true if all the elements satisfies predicate
	All(types.Predicate[T]) bool

	// Index returns the first position of value in list or -1
	Index(T) int

	// Index2 returns the first position of value or reports not found
	Index2(T) (int, bool)

	// IndexBy returns the first position of value that satisfies predicate or -1 if not found
	IndexBy(types.Predicate[T]) int

	// IndexBy2 returns the first position of value that satisfies predicate
	IndexBy2(types.Predicate[T]) (int, bool)
}
