package main

import "fmt"

var a = [1001]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func fill() {
	for i, v := range a {
		if v == "" {
			fmt.Println(i)
			if i > 20 && i < 100 {
				a[i] = a[i/10*10] + a[i%10]
			} else if i >= 100 && i < 1000 {
				a[i] = a[i/100] + "hundred"
				if i%100 != 0 {
					a[i] += "and" + a[i%100]
				}
			} else if i >= 1000 && i < 10000 {
				a[i] = a[i/1000] + "thousand" + a[i%1000]
			}
		}
	}
}

func count() int {
	cnt := 0
	for _, s := range a {
		cnt += len(s)
	}
	return cnt
}

func main() {
	fill()
	fmt.Println(a)
	fmt.Println(count())

}
