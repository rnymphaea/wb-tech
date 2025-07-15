package main

import "fmt"

func main() {
	var s string
	fmt.Printf("Enter the string: ")
	fmt.Scan(&s)

	result := ReverseString(s)
	fmt.Println("Reversed string: ", result)

}

func ReverseString(s string) string {
	symbols := []rune(s)
	l := len(symbols)

	for i := 0; i < l/2; i++ {
		symbols[i], symbols[l-i-1] = symbols[l-i-1], symbols[i]
	}

	return string(symbols)
}
