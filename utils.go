package randomkit

// Kahan implements the Kahan Summation Algorithm: https://en.wikipedia.org/wiki/Kahan_summation_algorithm
func Kahan(a []float64) float64 {
	if len(a) == 0 {
		return 0 // panic?
	}
	var c float64
	sum := a[0]
	for i := 1; i < len(a); i++ {
		y := a[i] - c
		t := sum + y
		c = (t - sum) - y
		sum = t
	}
	return sum
}
