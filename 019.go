package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := 0
	start := time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC)
	for start.Before(end) {
		if start.Weekday() == time.Sunday {
			cnt++
			fmt.Println(start)
		}
		start = start.AddDate(0, 1, 0)
	}
	fmt.Println(cnt)
}
