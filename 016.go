package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var a big.Int
	a.Exp(big.NewInt(2), big.NewInt(1000), nil)
	fmt.Println(a)
	sum := 0
	for _, c := range a.String() {
		x, _ := strconv.Atoi(string(c))
		sum += x
	}
	fmt.Println(sum)
}
