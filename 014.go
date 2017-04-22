package main

import "fmt"

func get_next(i int) int {
	if i%2 == 0 {
		i /= 2
	} else {
		i = 3*i + 1
	}
	return i
}

const flag = 1000000

func main() {
	max_len := 0
	term := flag
	for i := flag; i > 1; i-- {
		t := i
		c_len := 0
		for t != 1 {
			t = get_next(t)
			c_len += 1
		}
		if c_len > max_len {
			max_len = c_len
			term = i
		}
	}
	fmt.Println(term, max_len)
}
