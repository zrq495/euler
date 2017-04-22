package main

import (
	"fmt"
	"math"
)

const n int = 28123 + 1

var a, b [n + 1]int
var c [n * 3]int

func getAmicable(x int) int {
	if a[x] != 0 {
		return a[x]
	}
	sum := 1
	sq := math.Sqrt(float64(x))
	for i := 2; i <= int(sq); i++ {
		if x%i == 0 {
			sum += i
			if float64(i) != sq {
				sum += x / i
			}
		}
	}
	a[x] = sum
	return sum
}

func main() {
	// fmt.Println(getAmicable(12))
	for i := 1; i < n; i++ {
		am := getAmicable(i)
		if i < am {
			b[i] = am
		}
	}
	for i := 1; i <= n; i++ {
		if b[i] > 0 {
			for j := i; j <= n; j++ {
				if b[j] > 0 && i+j <= n {
					c[i+j] = 1
				}
			}
		}
	}
	sum := 0
	for i := 1; i <= n; i++ {
		if c[i] == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
