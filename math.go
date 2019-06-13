package goutils

// AbsInt returns the absolute value of x.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// PCTChange computes percent change
func PCTChange(before float64, after float64) float64 {
	diff := after - before
	realDiff := diff / before
	return realDiff
}

// Max greater of two values
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
