package main

import "fmt"

func main() {
	a, b := 1, 2
	max := 4000000
	sum := 0
	for b < max {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	fmt.Println(sum)
}
