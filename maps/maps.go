package maps

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

	for k, _ := range m {
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
