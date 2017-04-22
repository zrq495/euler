package main

import "fmt"

func main() {
	n := 1001
	sum := 1
	for i := 3; i <= n; i += 2 {
		a := i * i
		b := a - i + 1
		c := b - i + 1
		d := c - i + 1
		sum += a + b + c + d
	}
	fmt.Println(sum)
}
