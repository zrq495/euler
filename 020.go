package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func factorial(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}
	t := new(big.Int)
	t.Set(n)
	return n.Mul(n, factorial(t.Sub(t, big.NewInt(1))))

}

func main() {
	n := big.NewInt(100)
	res := factorial(n)
	sum := 0
	for _, c := range res.String() {
		t, _ := strconv.Atoi(string(c))
		sum += t
	}
	fmt.Println(sum)
}
