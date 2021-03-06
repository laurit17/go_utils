package math_utils

import (
	"github.com/laurit17/go_utils/iterator_utils"
)

func SieveOfEratosthenes() func() int {
	numberStream := iterator_utils.IterStartingFrom(2)
	return func() int {
		nextNumber := numberStream()
		filter_func := func(n int) bool { return n%nextNumber != 0 }
		numberStream = iterator_utils.FilterIter(filter_func, numberStream)
		return nextNumber
	}
}

func PrimesUpTo(n int) []bool {
	sieve := make([]bool, n+1)

	for i := 2; i < len(sieve); i++ {
		sieve[i] = true
	}

	for i := 2; i*i < len(sieve); i++ {
		if sieve[i] {
			for j := 2; i*j < len(sieve); j++ {
				sieve[i*j] = false
			}
		}
	}

	return sieve
}

func NthPrime(n int) int {
	iter := SieveOfEratosthenes()
	for i := 0; i < n-1; i++ {
		iter()
	}
	return iter()
}
