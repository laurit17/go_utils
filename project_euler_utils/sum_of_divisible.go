package project_euler_utils

func TriangleSum(num int) int {
	return num * (num + 1) / 2
}

func DivisibleSumLessThan(max int, divisor int) int {
	return TriangleSum((max-1)/divisor) * divisor
}
