package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	var input1, input2 string
	fmt.Printf("Enter two numbers: ")
	fmt.Scan(&input1, &input2)

	a.SetString(input1, 10)
	b.SetString(input2, 10)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	product := new(big.Int).Mul(a, b)

	fmt.Printf("a = %s\nb = %s\n", a.String(), b.String())
	fmt.Println("Сложение:", sum.String())
	fmt.Println("Вычитание:", diff.String())
	fmt.Println("Умножение:", product.String())

	if b.Sign() != 0 {
		quotient := new(big.Int).Div(a, b)
		fmt.Println("Деление:", quotient.String())
	}
}
