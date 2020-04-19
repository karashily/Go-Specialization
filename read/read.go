package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Printf("Enter filename: ")
	var filename string
	_, err := fmt.Scan(&filename)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	sli := make([]Name, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		var name = Name{fname: string(words[0]), lname: string(words[1])}
		sli = append(sli, name)
	}
	for _, v := range sli {
		fmt.Printf("first name: %s, last name: %s\n", v.fname, v.lname)
	}
}
