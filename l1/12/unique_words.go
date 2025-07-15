package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	debug      bool
	defaultarr bool
)

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.BoolVar(&defaultarr, "default", false, "use default array of strings ('cat', 'cat', 'dog', 'cat', 'tree')")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	var (
		s          []string
		ignoreCase bool
	)

	if defaultarr {
		s = []string{"cat", "cat", "dog", "cat", "tree"}
		ignoreCase = false
	} else {
		var N int
		fmt.Printf("Enter number of words: ")
		fmt.Scan(&N)

		fmt.Printf("Enter %d words: ", N)
		s = make([]string, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&s[i])
		}

		var ignoreCaseInput string
		fmt.Printf("Ignore case? [y/n] ")
		fmt.Scan(&ignoreCaseInput)

		if ignoreCaseInput != "y" && ignoreCaseInput != "n" {
			fmt.Println("Invalid option")
			return
		}

		ignoreCase = strings.EqualFold(ignoreCaseInput, "y")
	}

	if debug {
		log.Printf("[DEBUG] initial array: %v\n", s)
		if !defaultarr {
			log.Printf("[DEBUG] ingore case: %t\n", ignoreCase)
		}
	}

	result := GetUniqueWords(s, ignoreCase)
	fmt.Println("Result: ", result)
}

func GetUniqueWords(s []string, ignoreCase bool) []string {
	m := make(map[string]struct{})
	result := make([]string, 0, len(s))

	for i := 0; i < len(s); i++ {
		word := s[i]
		if ignoreCase {
			word = strings.ToLower(word)
		}

		if _, exists := m[word]; !exists {
			result = append(result, word)
			m[word] = struct{}{}
		}
	}

	if debug {
		log.Printf("[DEBUG] map after iterating over array: %v\n", m)
	}

	return result
}
