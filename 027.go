package main

import (
	"fmt"
	"math"
)

var p = [1000000]int{}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if p[n] == 1 {
		return true
	} else if p[n] == -1 {
		return false
	}
	sq := math.Sqrt(float64(n))
	for i := 2; i <= int(sq); i++ {
		if n%i == 0 {
			p[n] = -1
			return false
		}
	}
	p[n] = 1
	return true
}

func main() {
	n := 1000
	m := 1000
	max := 0
	cnt := 0
	for i := -n; i <= n; i++ {
		for j := -m; j <= m; j++ {
			tCnt := 0
			for k := 0; ; k++ {
				v := k*k + i*k + j
				if !isPrime(v) {
					break
				}
				tCnt++
			}
			if tCnt > cnt {
				cnt = tCnt
				max = i * j
			}
		}
	}
	fmt.Println(cnt, max)

}
