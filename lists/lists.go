package lists

import (
	"github.com/totemcaf/gollections/types"
)

func Of[S any](ss ...S) []S {
	return ss
}

func Map[S any, T any](ss []S, m types.Mapper[S, T]) []T {
	tt := make([]T, len(ss))
	for idx, s := range ss {
		tt[idx] = m(s)
	}
	return tt
}

func FlatMap[S any, T any](ss []S, m types.Mapper[S, []T]) []T {
	tt := make([]T, 0, len(ss))
	for _, s := range ss {
		tt = append(tt, m(s)...)
	}
	return tt
}

func Filter[S any](ss []S, p types.Predicate[S]) []S {
	var tt []S
	for _, s := range ss {
		if p(s) {
			tt = append(tt, s)
		}
	}
	return tt
}

func Count[S any](ss []S, p types.Predicate[S]) int {
	var count = 0
	for _, s := range ss {
		if p(s) {
			count++
		}
	}
	return count
}

func Find[S any](ss []S, p types.Predicate[S]) (S, bool) {
	for _, s := range ss {
		if p(s) {
			return s, true
		}
	}
	var notFound S
	return notFound, false
}

func Index[S comparable](ss []S, toFind S) int {
	for idx, s := range ss {
		if s == toFind {
			return idx
		}
	}
	return -1
}

func Index2[S comparable](ss []S, toFind S) (int, bool) {
	idx := Index(ss, toFind)
	return idx, idx >= 0
}

func IndexBy[S any](ss []S, p types.Predicate[S]) int {
	for idx, s := range ss {
		if p(s) {
			return idx
		}
	}
	return -1
}

func IndexBy2[S any](ss []S, p types.Predicate[S]) (int, bool) {
	idx := IndexBy(ss, p)
	return idx, idx >= 0
}
