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
	return after / before - 1
}

// Scales `val` using `mean` and `std` of training dataset. When scaling data 
// via `preprocessing.Scale` `mean` and `std` are taken from entire data set.
// When using model for prediciton `mean` and `std` woudl be computed only 
// from given sample. This could result in inacurate predictions. This functions 
// allows to pass mean and std of entire dataset 
func Scale(val, mean, std float64) float64 {
	return (val - mean) / std
}

// Max greater of two values
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min lesser of two values
func Min(x int, y int) int {
	if y > x {
		return x
	}
	return y
}
