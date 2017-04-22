package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxN int = 100

var a [maxN][maxN]int

func getData() {
	fname := "p067_triangle.txt"
	fl, err := os.Open(fname)
	if err != nil {
		fmt.Println(fname, err)
		return
	}
	defer fl.Close()
	scanner := bufio.NewScanner(fl)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineList := strings.Split(strings.TrimSpace(line), " ")
		for j, t := range lineList {
			tt, _ := strconv.Atoi(string(t))
			a[i][j] = tt
		}
		i++
	}
	fmt.Println(a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxNum(a []int) int {
	t := a[0]
	for _, d := range a {
		if d > t {
			t = d
		}
	}
	return t
}

func main() {
	getData()
	fmt.Println(a)
	var s [maxN][maxN]int
	for i := 0; i < maxN; i++ {
		for j := 0; j <= i; j++ {
			if i-1 < 0 {
				s[i][j] = a[i][j]
				continue
			}
			if j-1 < 0 {
				s[i][j] = s[i-1][j] + a[i][j]
				continue
			}
			s[i][j] = max(s[i-1][j], s[i-1][j-1]) + a[i][j]
		}
	}
	fmt.Println(s)
	m := getMaxNum(s[maxN-1][:])
	fmt.Println(m)
}
