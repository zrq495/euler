package main

import (
	"fmt"
	"math"
)

const n int = 10000

var a [n * 10]int

func getAmicable(x int) int {
	if a[x] != 0 {
		return a[x]
	}
	sum := 1
	sq := int(math.Sqrt(float64(x)))
	for i := 2; i <= sq; i++ {
		if x%i == 0 {
			sum += i
			if i != sq {
				sum += x / i
			}
		}
	}
	a[x] = sum
	return sum
}

func main() {
	// fmt.Println(getAmicable(220))
	sum := 0
	for i := 1; i < n; i++ {
		am := getAmicable(i)
		pAm := getAmicable(am)
		if i == pAm && i != am {
			sum += i
			fmt.Println(i, am)
		}
	}
	fmt.Println(sum)
}
