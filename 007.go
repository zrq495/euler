package main

import (
	"fmt"
	"math"
)

func isPrime(a int) bool {
	sq := math.Sqrt(float64(a))
	for i := 2; i <= int(sq); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	cnt := 0
	prime := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			cnt++
			prime = i
		}
		if cnt == 10001 {
			break
		}
	}
	fmt.Println(cnt, prime)

}
