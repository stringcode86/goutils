package goutils

func intfArrToFloat64Arr(intfArr []interface{}) []float64 {
	floats := make([]float64, 0, len(intfArr))
	for idx, val := range intfArr {
		floats[idx] = val.(float64)
	}
	return floats
}

func intfArrToFloat32Arr(intfArr []interface{}) []float32 {
	floats := make([]float32, 0, len(intfArr))
	for idx, val := range intfArr {
		floats[idx] = val.(float32)
	}
	return floats
}