package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		counter[word]++
	}

	return counter
}

const test1 = "I am learning Go!"
const test2 = "The quick brown fox jumped over the lazy dog."
const test3 = "I ate a donut. Then I ate another donut."
const test4 = "A man a plan a canal panama."

func main() {
	fmt.Println(WordCount(test1))
	fmt.Println(WordCount(test2))
	fmt.Println(WordCount(test3))
	fmt.Println(WordCount(test4))
}
