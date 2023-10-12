package cond

// IIf returns ifTrue if condition is true, ifFalse otherwise
func IIf[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// IIfFunc returns ifTrue() if condition is true, ifFalse() otherwise
func IIfFunc[T any](condition bool, ifTrue func() T, ifFalse func() T) T {
	if condition {
		return ifTrue()
	}
	return ifFalse()
}
