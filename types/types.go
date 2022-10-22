package types

type Predicate[T any] func(T) bool
type Mapper[S any, T any] func(S) T
type MapperWithError[S any, T any] func(S) (T, error)

type Function1[T any] func(T)

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Iterable[T any] interface {
	ForEach(Function1[T])
	Iterator() Iterator[T]
}

type Cloneable[T any] interface {
	Clone() T
}

type Comparable[T any] interface {
	// Compare returns -1, 0, 1 if this is less than, equal to, to greater than other
	Compare(other T) int
}
