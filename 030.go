package main

import (
	"fmt"
	"math"
	"strconv"
)

var max = 400000

func power(s string) int {
	var sum float64
	for _, c := range s {
		intC, _ := strconv.Atoi(string(c))
		sum += math.Pow(float64(intC), 5)
	}
	return int(sum)
}

func main() {
	sum := 0
	for i := 10; i < max; i++ {
		p := power(strconv.Itoa(i))
		if p == i {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("sum:", sum)
}
