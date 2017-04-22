package main

import "fmt"

// (a+b+c)**2 = a**2+b**2+c**2+2ab+2ac+2bc
func main() {
	n := 100
	x := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			x += 2 * i * j
		}
	}
	fmt.Println(x)
}
