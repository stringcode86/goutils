package goutils

func IntfArrToFloat64Arr(intfArr []interface{}) []float64 {
	floats := make([]float64, 0, len(intfArr))
	for idx, val := range intfArr {
		floats[idx] = val.(float64)
	}
	return floats
}

func IntfArrToFloat32Arr(intfArr []interface{}) []float32 {
	floats := make([]float32, 0, len(intfArr))
	for idx, val := range intfArr {
		floats[idx] = val.(float32)
	}
	return floats
}

func IntfArrToIntArr(intfArr []interface{}) []int {
	ints := make([]int, 0, len(intfArr))
	for idx, val := range intfArr {
		ints[idx] = val.(int)
	}
	return ints
}

func IntfArrToInt64Arr(intfArr []interface{}) []int64 {
	ints := make([]int64, 0, len(intfArr))
	for idx, val := range intfArr {
		ints[idx] = val.(int64)
	}
	return ints
}

func IntfArrToUint64Arr(intfArr []interface{}) []uint64 {
	ints := make([]uint64, 0, len(intfArr))
	for idx, val := range intfArr {
		ints[idx] = val.(uint64)
	}
	return ints
}

func IntArrToFloat64Arr(intArr []int) []float64 {
	floats := make([]float64, 0, len(intArr))
	for idx, val := range intArr {
		floats[idx] = float64(val)
	}
	return floats
}
