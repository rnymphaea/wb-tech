package main

import (
	"flag"
	"fmt"
	"log"
)

var debug bool

const debugPrefix = "[DEBUG]"

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	var (
		num   int64
		shift uint8
		bit   int64
	)

	fmt.Printf("Enter the number: ")
	fmt.Scan(&num)
	if debug {
		log.Printf("%s bit representation: %b\n", debugPrefix, num)
	}

	fmt.Printf("Enter the number of the bit to be replaced: ")
	fmt.Scan(&shift)
	if shift < 1 || shift > 64 {
		fmt.Println("Invalid number: the number must be greater than 0 and less than 65")
		return
	}

	fmt.Printf("Replace bit #%d with (1 or 0): ", shift)
	fmt.Scan(&bit)

	if bit != 1 && bit != 0 {
		fmt.Println("Bit must be 1 or 0")
		return
	}

	result := ReplaceBit(num, bit, shift)

	fmt.Printf("Result: %d\n", result)
	if debug {
		log.Printf("%s bit representation: %b\n", debugPrefix, result)
	}
}

func ReplaceBit(num, bit int64, shift uint8) int64 {
	mask := int64(1) << (shift - 1)
	result := (num & ^mask) | ((bit & 1) << (shift - 1))

	if debug {
		log.Printf("%s bit mask: %b\n", debugPrefix, mask)
		log.Printf("%s set zero value to %d bit: %b\n", debugPrefix, shift, num & ^mask)
		log.Printf("%s replace %d bit with new value: %b\n", debugPrefix, shift, result)
	}

	return result
}
