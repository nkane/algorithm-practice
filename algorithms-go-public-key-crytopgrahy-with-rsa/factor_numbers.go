package cryptography

var primes []int

func FindFactors(num int) []int {
	factors := []int{}
	if num < 0 {
		factors = append(factors, -1)
		num = -num
	}
	for num%2 == 0 {
		factors = append(factors, 2)
		num /= 2
	}
	factor := 3
	for factor*factor <= num {
		if num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		} else {
			factor += 2
		}
	}
	if num > 1 {
		factors = append(factors, num)
	}
	return factors
}

func MultipleSlice(arr []int) int {
	result := 1
	for _, i := range arr {
		result *= i
	}
	return result
}

func FindFactorsSieve(num int) []int {
	factors := []int{}
	if num < 0 {
		factors = append(factors, -1)
		num = -num
	}
	for _, factor := range primes {
		for num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
			if num == 1 {
				break
			}
		}
		if factor*factor > num {
			factors = append(factors, num)
			break
		}
	}
	return factors
}
