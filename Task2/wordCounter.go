package main

import (
	"strings"
	"fmt"
)

func main() {
	s := "Hello World"
	words := strings.Fields(s)
	count := make(map[string]int)

	for _, v := range words {
		count[v] += 1
	}

	fmt.Println(count)
}