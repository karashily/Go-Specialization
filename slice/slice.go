package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Printf("enter a number: ")
	var s string
	fmt.Scan(&s)

	sli := make([]int, 0, 3)

	for s != "X" {
		num, _ := strconv.Atoi(s)
		sli = append(sli, num)
		sort.Ints(sli)
		fmt.Println(sli)
		fmt.Printf("enter a number (X to end): ")
		fmt.Scan(&s)
	}
}
