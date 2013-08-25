package string_utils

import "testing"

func TestPalindrome(t *testing.T) {
	const in, out = "", true
	var inOutPairs = map[string]bool{
		"":             true,
		"a":            true,
		"aa":           true,
		"ab":           false,
		"aba":          true,
		"Aa":           false,
		"abcdcba":      true,
		"98":           false,
		"89":           false,
		"88":           true,
		"87":           false,
		"adblaskjdasd": false,
	}
	for in, out := range inOutPairs {
		if x := IsPalindrome(in); x != out {
			t.Errorf("IsPalindrome(%s) = %t, want %t", in, x, out)
		}
	}
}
