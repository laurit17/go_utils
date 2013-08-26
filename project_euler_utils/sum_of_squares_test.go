package project_euler_utils

import "testing"

func TestSumOfSquares(t *testing.T) {
	inputsAndOutputs := map[int]int{
		0: 0,
		1: 1,
		2: 5,
		3: 14,
	}
	for in, out := range inputsAndOutputs {
		if x := SumOfSquares(in); x != out {
			t.Errorf("SumOfSquares(%d) = %d, want %d", in, x, out)
		}
	}
}
