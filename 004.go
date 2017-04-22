package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("aba"))
	max := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= i; j-- {
			p := i * j
			if isPalindrome(strconv.Itoa(p)) {
				if p > max {
					max = p
				}
			}
		}
	}
	fmt.Println(max)
}
