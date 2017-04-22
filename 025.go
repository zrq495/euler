package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)
	fmt.Println(limit)
	cnt := 2
	for b.Cmp(&limit) <= 0 {
		a.Add(a, b)
		a, b = b, a
		cnt++
	}
	fmt.Println(a, b)
	fmt.Println(cnt)
}
