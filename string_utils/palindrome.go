package string_utils

func IsPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}
	if str[0] != str[len(str)-1] {
		return false
	}
	return IsPalindrome(str[1 : len(str)-1])
}
