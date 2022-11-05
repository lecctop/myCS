package utils

import "math"

func MaxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func AbsInt(n int) int {
	return int(math.Abs(float64(n)))
}
