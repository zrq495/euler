package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := big.NewInt(0)
	fmt.Println(sum)
	var n int64 = 1000
	for i := int64(1); i <= n; i++ {
		bigI := big.NewInt(i)
		sum.Add(sum, bigI.Exp(bigI, bigI, nil))
	}
	sumStr := sum.String()
	fmt.Println(sumStr[len(sumStr)-10:])
}
