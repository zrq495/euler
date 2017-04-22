package main

import "fmt"

func main() {
	max := 1000
	for i := 1; i < max; i++ {
		for j := 1; j < max; j++ {
			for k := 1; k < max; k++ {
				if i+j+k == max && i*i+j*j == k*k {
					fmt.Println(i, j, k, i*j*k)
					break
				}
			}
		}
	}
}
