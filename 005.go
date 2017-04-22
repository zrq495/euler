package main

import "fmt"

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	a, b := 10, 8
	fmt.Println(gcd(a, b))
	fmt.Println(lcm(a, b))
	n := 20
	m := 1
	for i := 1; i < n; i++ {
		m = lcm(m, i)
	}
	fmt.Println(m)
}
