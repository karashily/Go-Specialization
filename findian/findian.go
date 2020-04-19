package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	lower := strings.ToLower(s[0 : len(s)-2]) // a slice to eliminate \n at the end
	if lower[0] == 'i' && lower[len(lower)-1] == 'n' && strings.Contains(lower, "a") {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
