package cryptography

func GDC(a int, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for {
		remainder := a % b
		if remainder == 0 {
			return b
		}
		a = b
		b = remainder
	}
}

func LCM(a int, b int) int {
	result := a * (b / GDC(a, b))
	return result
}
