package math_utils

func Pow(base, exp int) int {
	switch {
	case exp == 0:
		return 1
	case exp%2 == 0:
		return Pow(base*base, exp/2)
	}

	return base * Pow(base, exp-1)
}
