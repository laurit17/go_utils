package iterator_utils

func FilterIter(filter func(n int) bool, iterator func() int) func() int {
	return func() int {
		for {
			next := iterator()
			if next == -1 {
				return -1
			}
			if filter(next) {
				return next
			}
		}
	}
}
