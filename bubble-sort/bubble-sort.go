package main

import "fmt"

func swap(sli []int, i, j int) {
	temp := sli[i]
	sli[i] = sli[j]
	sli[j] = temp
}

func bubbleSort(sli []int) {
	for i := len(sli); i > 0; i-- {
		for j := 1; j < i; j++ {
			if sli[j-1] > sli[j] {
				swap(sli, j, j-1)
			}
		}
	}
}

func main() {
	var num int
	fmt.Printf("Please enter a the number of numbers: ")
	fmt.Scan(&num)
	sli := make([]int, 0, 10)
	fmt.Printf("Please enter the numbers one by one: ")
	for i := 0; i < num; i++ {
		var x int
		fmt.Scan(&x)
		sli = append(sli, x)
	}
	fmt.Println(sli)
	bubbleSort(sli)
	fmt.Println(sli)
}
