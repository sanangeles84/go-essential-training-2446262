package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	fmt.Println(text)
	words := strings.Fields(text) //convert string to list of words
	fmt.Println(words)
	counts := map[string]int{} //map of words and their count
	for _, val := range words {
		counts[strings.ToLower(val)]++
	}
	fmt.Println(counts)
}
