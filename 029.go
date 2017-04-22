package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64 = 100
	m := map[interface{}]bool{}
	for i := int64(2); i <= n; i++ {
		bigI := big.NewInt(i)
		for j := int64(2); j <= n; j++ {
			bigJ := big.NewInt(j)
			var r big.Int
			r.Exp(bigI, bigJ, nil)
			m[r.String()] = true
		}
	}
	fmt.Println(len(m))

}
