package main

import (
	"fmt"
)

const n, m int = 21, 21

func main() {
	var a [n][m]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a[i][j] = 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				a[i][j] = a[i][j-1] + a[i-1][j]
			}
		}
	}
	fmt.Println(a)
	fmt.Println(a[n-1][m-1])
}
