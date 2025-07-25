package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const numbers = "0123456789"

var debug bool

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	var s string
	fmt.Printf("Enter your string: ")

	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		fmt.Println("error scanning string:", err)
	}

	res, err := unpack(s)
	if err != nil {
		fmt.Println("error unpacking string: ", err)
		return
	}

	fmt.Println(res)
}

func unpack(s string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}

	if strings.Contains(numbers, string(s[0])) {
		return "", fmt.Errorf("first char must not be numeric")
	}

	mul := -1 // current multiplier
	result := ""

	for _, v := range s {
		char := string(v)

		if debug {
			log.Printf("[DEBUG] current char = %s, mul = %d, result = %s\n", char, mul, result)
		}

		num, err := strconv.Atoi(char)
		if err == nil {
			// escape sequence
			if tmp := []rune(result); tmp[len(tmp)-1] == rune('\\') {
				tmp = tmp[:len(tmp)-1]
				tmp = append(tmp, rune(v))
				result = string(tmp)
				continue
			}

			if mul == -1 {
				mul = num
			} else {
				mul = mul*10 + num
			}
		} else {
			if mul != -1 {
				// we use []rune because of incorrect addressing with unicode characters
				tmp := []rune(result)

				if mul == 0 {
					result = string(tmp[:len(tmp)-1])
				} else {
					last := string(tmp[len(tmp)-1])
					result += strings.Repeat(last, mul-1)
				}
			}

			result += char
			mul = -1
		}

		if debug {
			log.Printf("[DEBUG] after processing got: mul = %d, result = %s\n", mul, result)
		}
	}

	if mul != -1 {
		tmp := []rune(result)

		if mul == 0 {
			result = string(tmp[:len(tmp)-1])
		} else {
			last := string(tmp[len(tmp)-1])
			result += strings.Repeat(last, mul-1)
		}
	}

	return result, nil
}
