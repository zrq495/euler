package main

import "fmt"

const (
	n = 1000
)

func main() {
	max := 0
	maxI := 0
	for i := n; i >= 1; i-- {
		var a [n]int
		s := 1
		cnt := 0
		for a[s] == 0 && s != 0 {
			a[s] = cnt
			s = s * 10
			s = s % i
			cnt++
		}
		// fmt.Println(i, cnt)
		if cnt-1 > max {
			max = cnt - 1
			maxI = i
		}
	}
	fmt.Println(maxI, max)
}
