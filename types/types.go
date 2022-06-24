package types

type Predicate[T any] func(T) bool
type Mapper[S any, T any] func(S) T

type Function1[T any] func(T)

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Iterable[T any] interface {
	ForEach(Function1[T])
	Iterator() Iterator[T]
}
