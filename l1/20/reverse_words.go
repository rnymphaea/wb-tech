package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		size int
	)

	fmt.Printf("Enter number of words: ")
	fmt.Scan(&size)

	s := make([]string, size)
	fmt.Printf("Enter %d words: ", size)

	for i := 0; i < size; i++ {
		fmt.Scan(&s[i])
	}

	result := ReverseWordsOrder(s)
	fmt.Println("Result: ", result)
}

func ReverseWordsOrder(s []string) string {
	l := len(s)

	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}

	return strings.Join(s, " ")
}
