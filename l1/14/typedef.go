package main

import "fmt"

func main() {
	fmt.Println("Trying int")
	PrintType(123)

	fmt.Println("Trying string")
	PrintType("simple")

	fmt.Println("Trying bool")
	PrintType(true)

	fmt.Println("Trying chan")
	PrintType(make(chan interface{}))

	fmt.Println("Trying other type")
	PrintType('a')
}

func PrintType(i interface{}) {
	fmt.Printf("The type of [%v] is ", i)
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan")
		break
	default:
		fmt.Println("undefined")
	}
}
