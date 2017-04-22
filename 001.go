package main

import (
	"fmt"
)

func main() {
	n := 1000
	sum := 0
	for i := 3; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
