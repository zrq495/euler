package main

import (
	"fmt"
	"math"
)

func get_factors(index int) int {
	cnt := 0
	sqrt := int(math.Sqrt(float64(index)))
	for i := 1; i <= sqrt; i++ {
		if index%i == 0 {
			cnt += 2
		}
	}
	if sqrt*sqrt == index {
		cnt--
	}
	return cnt
}

func main() {
	total_cnt := 500
	cnt := 0
	add_sum := 0
	index := 0
	for cnt <= total_cnt {
		index++
		add_sum += index
		cnt = get_factors(add_sum)
		fmt.Println(index, add_sum, cnt)
	}
	fmt.Println(index, add_sum, cnt)
}
