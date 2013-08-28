package math_utils

import "testing"

type BaseExp struct{ base, exp int }

func TestPow(t *testing.T) {
	expectedIOPairs := map[BaseExp]int{
		BaseExp{base: 2, exp: 1}: 2,
	}

	for in, out := range expectedIOPairs {
		if x := Pow(in.base, in.exp); x != out {
			t.Errorf("Pow(%v, %v) = %v, wanted %v", in.base, in.exp, x, out)
		}

	}

}
