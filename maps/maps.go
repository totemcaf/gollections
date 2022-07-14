package maps

import "github.com/totemcaf/gollections/types"

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))

	for k, v := range m {
		entries = append(entries, Entry[K, V]{k, v})
	}

	return entries
}

func Keys[K comparable, V any](m map[K]V) []K {
	values := make([]K, 0, len(m))

	for k := range m {
		values = append(values, k)
	}

	return values
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

// Clone copy a map entries into a new map
func Clone[K comparable, V any](source map[K]V) map[K]V {
	target := make(map[K]V, len(source))

	for k, v := range source {
		target[k] = v
	}

	return target
}

// ReduceKeys applies a reducer function to the keys of the map
func ReduceKeys[Value any, Key comparable, Element any](
	initialValue Value,
	elements map[Key]Element,
	reducer func(Value, Key) Value,
) Value {
	accum := initialValue

	for key := range elements {
		accum = reducer(accum, key)
	}

	return accum
}

// ReduceValues applies a reducer function to the values of the map
func ReduceValues[Value any, Key comparable, Element any](
	initialValue Value,
	elements map[Key]Element,
	reducer func(Value, Element) Value,
) Value {
	accum := initialValue

	for _, element := range elements {
		accum = reducer(accum, element)
	}

	return accum
}

// ReduceEntries applies a reducer function to the entries of the map
func ReduceEntries[Value any, Key comparable, Element any](
	initialValue Value,
	elements map[Key]Element,
	reducer func(Value, Key, Element) Value,
) Value {
	accum := initialValue

	for key, element := range elements {
		accum = reducer(accum, key, element)
	}

	return accum
}

// Map maps values. Returns a new Map with same keys and values transformed by map function
func Map[K comparable, V any, W any](m map[K]V, mapper types.Mapper[V, W]) map[K]W {
	result := make(map[K]W, len(m))

	for k, v := range m {
		result[k] = mapper(v)
	}

	return result
}
