package main

import "fmt"

var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// var a = []int{0, 1, 2, 3}

var cnt = 0

func main() {
	perm(0, len(a)-1)
	fmt.Println(cnt)
}

func perm(s, m int) {
	if s > m {
		cnt++
		// fmt.Println(a)
		if cnt == 1000000 {
			for _, e := range a {
				fmt.Print(e)
			}
			fmt.Println()
		}
		return
	}
	for i := s; i <= m; i++ {
		a[s], a[i] = a[i], a[s]
		t := i
		for ; t-1 >= 0 && t-1 > s && a[t] < a[t-1]; t-- {
			a[t], a[t-1] = a[t-1], a[t]
		}
		perm(s+1, m)
		tt := t
		for ; tt < i; tt++ {
			a[tt], a[tt+1] = a[tt+1], a[tt]
		}
		a[s], a[i] = a[i], a[s]
	}
}
