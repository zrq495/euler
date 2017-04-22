package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var letters = "0ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getNameScore(name string, index int) int {
	sum := 0
	for _, c := range name {
		i := strings.Index(letters, string(c))
		if i > 0 {
			sum += i
		} else {
			fmt.Println(string(c), c, name)
		}
	}
	return index * sum
}

func getData() []string {
	fname := "p022_names.txt"
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(fname, err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	data, err := reader.ReadString('\n')
	data = strings.Replace(data, "\n", "", -1)
	data = strings.Replace(data, "\"", "", -1)
	dataSlice := strings.Split(data, ",")
	sort.Strings(dataSlice)
	return dataSlice
}

func main() {
	// fmt.Println(getNameScore("COLIN", 938))
	// return
	data := getData()
	fmt.Println(data)
	sum := 0
	for i, d := range data {
		s := getNameScore(d, i+1)
		sum += s
	}
	fmt.Println(sum)
}
