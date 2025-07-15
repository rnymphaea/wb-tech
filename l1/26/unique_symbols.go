package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var debug bool

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	var s string
	fmt.Printf("Enter string: ")
	fmt.Scan(&s)

	unique := UniqueSymbols(s)

	if unique {
		fmt.Printf("String '%s' contains unique symbols\n", s)
	} else {
		fmt.Printf("String '%s' does not contain unique symbols\n", s)
	}
}

func UniqueSymbols(s string) bool {
	lower := []rune(strings.ToLower(s))
	m := make(map[rune]struct{})

	var excess []string
	if debug {
		excess = make([]string, 0, len(lower))
	}

	for i := 0; i < len(lower); i++ {
		if _, exists := m[lower[i]]; exists {
			if debug {
				excess = append(excess, string(lower[i]))
			} else {
				return false
			}
		} else {
			m[lower[i]] = struct{}{}
		}
	}

	if debug {
		if len(excess) != 0 {
			log.Printf("[DEBUG] excess symbols: %v", excess)
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}
