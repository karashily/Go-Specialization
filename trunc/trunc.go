package main

import "fmt"

func main() {
	var num float32
	fmt.Printf("Please enter a floating point number: ")
	fmt.Scan(&num)
	fmt.Println("Your truncated number is: ", int(num))
}
