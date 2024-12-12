package rbmv

func Variance(vals []float64, avgVal float64) float64 {
	sumVal := 0.0
	for _, val := range vals {
		diff := val - avgVal
		sumVal += diff * diff
	}
	return sumVal / float64(len(vals))
}
