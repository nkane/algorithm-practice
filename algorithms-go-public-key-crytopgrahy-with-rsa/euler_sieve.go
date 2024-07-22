package cryptography

func EulerSieve(max int) []bool {
	output := make([]bool, max+1)
	output[2] = true
	for i := 3; i < len(output); i += 2 {
		output[i] = true
	}
	for i := 3; i <= max; i += 2 {
		if output[i] {
			maxQ := max / i
			if maxQ%2 == 0 {
				maxQ--
			}
			for j := maxQ; j >= i; j -= 2 {
				if output[j] {
					output[i*j] = false
				}
			}
		}
	}
	return output
}
