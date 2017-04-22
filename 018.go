package main

import "fmt"

var a = [15][15]int{
	{75},
	{95, 64},
	{17, 47, 82},
	{18, 35, 87, 10},
	{20, 4, 82, 47, 65},
	{19, 1, 23, 75, 3, 34},
	{88, 2, 77, 73, 7, 63, 67},
	{99, 65, 4, 28, 6, 16, 70, 92},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
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
	fmt.Println(a)
	s := new([15][15]int)
	for i := 0; i < 15; i++ {
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
	m := getMaxNum(s[14][:])
	fmt.Println(m)
}
