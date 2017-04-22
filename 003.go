package main

import "fmt"

func max_prime(n int) int {
	// t := math.Sqrt(float64(n))
	new := 1
	max := 1
	for i := 2; new < n; i++ {
		if n%i == 0 {
			new *= i
			max = i
			fmt.Println(new, max, i)
		}
	}
	return max
}

func main() {
	n := 600851475143
	fmt.Println(max_prime(n))
}
