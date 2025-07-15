package main

import (
	"flag"
	"fmt"
	"log"
)

var debug bool

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	var num1, num2 int

	fmt.Printf("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Printf("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)

	fmt.Println("First way (arithmetic operations):")
	if debug {
		log.Println("[DEBUG] changing values by using addition and subtraction")
		log.Printf("[DEBUG] step 1: assign to num2 sum of num1 and num2: num1 = %d, num2 = %d\n", num1, num1+num2)
		log.Printf("[DEBUG] step 2: assign to num1 sub of num1 and num2 (num1 + num2): num1 = %d, num2 = %d\n", num2, num1+num2)
		log.Printf("[DEBUG] step 3: assign to num2 sub of num1 (num2) and num2 (num1 + num2): num1 = %d, num2 = %d\n", num2, num1)
	}

	num2 = num1 + num2
	num1 = num2 - num1
	num2 = num2 - num1

	fmt.Printf("num1 = %d, num = %d\n", num1, num2)

	fmt.Println("Second way (logical operations):")
	if debug {
		log.Println("[DEBUG] changing values by using XOR")
		log.Printf("[DEBUG] step 1: assign to num2 XOR of num1 and num2: num1 = %b (%d), num2 = %b (%d)\n", num1, num1, num1^num2, num1^num2)
		log.Printf("[DEBUG] step 2: assign to num1 XOR of num1 and num2 (num1 XOR num2): num1 = %b (%d), num2 = %b (%d)\n", num2, num2, num1^num2, num1^num2)
		log.Printf("[DEBUG] step 3: assign to num2 XOR of num1 (num2) and num2 (num1 XOR num2): num1 = %b (%d), num2 = %b (%d)\n", num2, num2, num1, num1)
	}

	num2 = num1 ^ num2
	num1 = num2 ^ num1
	num2 = num2 ^ num1

	fmt.Printf("num1 = %d, num = %d\n", num1, num2)
}
