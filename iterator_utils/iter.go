package iterator_utils

func Iter(n int) func() int {
	i := 0
	return func() int {
		if i >= n {
			return -1
		}
		defer func() { i += 1 }()
		return i
	}
}

func IterStartingFrom(n int) func() int {
	return func() int {
		next := n
		n += 1
		return next
	}
}
