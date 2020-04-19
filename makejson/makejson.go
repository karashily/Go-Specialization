package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Enter Address: ")
	addr, _ := reader.ReadString('\n')
	mp := map[string]string{
		"name":    strings.TrimSuffix(name, "\r\n"),
		"address": strings.TrimSuffix(addr, "\r\n"),
	}
	j, _ := json.MarshalIndent(mp, "", "    ")
	fmt.Println(string(j))
}
