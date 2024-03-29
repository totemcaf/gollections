package cond

// Condition is a conditional statement that can be chained with else-if statements
type Condition[T any] interface {
	// Else returns the value if the condition is true, ifFalse otherwise
	Else(ifFalse T) T
	// ElseF returns the value if the condition is true, ifFalse() otherwise
	ElseF(ifFalse func() T) T
	// ElseIf returns the same condition if the condition is true, ifFalse otherwise
	ElseIf(cond bool, ifFalse T) Condition[T]
	// ElseIfF returns the same condition if the condition is true, ifFalse() otherwise
	ElseIfF(cond bool, ifFalse func() T) Condition[T]
}

// If returns a Condition that can be used to chain else-if statements
func If[T any](condition bool, value T) Condition[T] {
	if condition {
		return trueCond[T]{value: value}
	}
	return falseCond[T]{}
}

// IfF returns a Condition that can be used to chain else-if statements
// If the condition is true, the value is evaluated and returned
func IfF[T any](condition bool, value func() T) Condition[T] {
	if condition {
		return trueCond[T]{value: value()}
	}
	return falseCond[T]{}
}

// trueCond is a Condition that is true, so it will return the value
type trueCond[T any] struct {
	value T
}

func (t trueCond[T]) Else(T) T {
	return t.value
}

func (t trueCond[T]) ElseF(func() T) T {
	return t.value
}

func (t trueCond[T]) ElseIf(bool, T) Condition[T] {
	return t
}

func (t trueCond[T]) ElseIfF(bool, func() T) Condition[T] {
	return t
}

// falseCond is a Condition that is false, so it will return the ifFalse value
type falseCond[T any] struct {
}

func (f falseCond[T]) Else(ifFalse T) T {
	return ifFalse
}

func (f falseCond[T]) ElseF(ifFalse func() T) T {
	return ifFalse()
}

func (f falseCond[T]) ElseIf(cond bool, ifFalse T) Condition[T] {
	return If(cond, ifFalse)
}

func (f falseCond[T]) ElseIfF(cond bool, ifFalse func() T) Condition[T] {
	return If(cond, ifFalse())
}
