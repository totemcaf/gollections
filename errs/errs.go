package errs

// OrPanic panics if err is not nil, otherwise returns value
func OrPanic[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

// Must panics if err is not nil, otherwise returns value
// It is an alias for OrPanic
func Must[T any](value T, err error) T {
	return OrPanic(value, err)
}
