package main

import (
	"fmt"
	"math"
)

func main() {
	n := 2000000
	// n := 10
	primes := make([]int, 1)
	primes[0] = 2
	sum := 2
	for i := 3; i <= n; i++ {
		is_p := true
		sq := int(math.Sqrt(float64(i)))
		for _, p := range primes {
			if p <= sq && i%p == 0 {
				is_p = false
				break
			}
		}
		if is_p {
			primes = append(primes, i)
			sum += i
		}
	}
	fmt.Println(sum)
}
