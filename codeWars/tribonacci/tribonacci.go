package tribonacci

// Function returns a tribonacci slice with n elements
func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}
	var result []float64 = make([]float64, 3, 6)
	copy(result, signature[:])
	if n < 3 {
		return signature[:n]
	} else {
		for i := 0; i < n-3; i++ {
			plus := result[i] + result[i+1] + result[i+2]
			result = append(result, plus)
		}
	}
	return result
}
