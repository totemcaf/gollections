package ord

import (
	"golang.org/x/exp/constraints"
)

// Min returns the minimum of a and all bs.
func Min[N constraints.Ordered](a N, b N) N {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of a and all bs.
func Max[N constraints.Ordered](a N, b N) N {
	if a > b {
		return a
	}
	return b
}

// NMin returns the minimum of a and all bs.
func NMin[N constraints.Ordered](a N, b ...N) N {
	min := a
	for _, v := range b {
		if v < min {
			min = v
		}
	}
	return min
}

// NMax returns the maximum of a and all bs.
func NMax[N constraints.Ordered](a N, b ...N) N {
	max := a
	for _, v := range b {
		if v > max {
			max = v
		}
	}
	return max
}

// Clamp returns a value clamped between min and max.
func Clamp[N constraints.Ordered](value, min, max N) N {
	return Max(Min(value, max), min)
}
