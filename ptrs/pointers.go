package ptrs

// SetIfPresent sets target to value if value is not nil
func SetIfPresent[T any](target *T, value *T) {
	if value != nil {
		*target = *value
	}
}
