package lists

// Generate use provided generator to generate array of N elements
func Generate[T any](times int, generator func(i int) T) []T {
	if times <= 0 {
		return []T{}
	}

	elements := make([]T, times)

	for idx := range elements {
		elements[idx] = generator(idx)
	}

	return elements
}
