package project_euler_utils

func TriangleSum(low int, high int) int {
	return high*(high+1)/2 - low*(low+1)/2
}

func DivisibleSumLessThan(max int, divisor int) int {
	return TriangleSum(0, (max-1)/divisor) * divisor
}
