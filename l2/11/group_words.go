package main

import (
	"flag"
	"fmt"
	"log"
	"slices"
	"sort"
	"strings"
)

var (
	debug      bool
	defaultarr bool
)

const debugPrefix = "[DEBUG]"

type Group struct {
	first string
	set   map[string]struct{}
}

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.BoolVar(&defaultarr, "default", false, "use default array of words (пятак, пятка, тяпка, листок, слиток, столик, стол)")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	var words []string

	if defaultarr {
		words = []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	} else {
		var size int
		fmt.Printf("Enter the number of words: ")

		_, err := fmt.Scan(&size)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		words = make([]string, size)

		fmt.Printf("Enter %d words: ", size)

		for i := 0; i < size; i++ {
			_, err = fmt.Scan(&words[i])
			if err != nil {
				fmt.Println("error: ", err)
				return
			}
		}
	}

	res := groupWords(words)
	fmt.Printf("Result: ")

	for k, v := range res {
		fmt.Printf("%s: %q\n", k, v)
	}

}

func groupWords(words []string) map[string][]string {
	const funcName = "groupWords"

	groups := make(map[string]*Group)

	for _, w := range words {
		lower := strings.ToLower(w)
		sorted := sortString(lower)

		if debug {
			log.Printf("%s %s: lower='%s', sorted='%s'\n", debugPrefix, funcName, lower, sorted)
		}

		if _, ok := groups[sorted]; !ok {
			newGroup := Group{
				first: w,
				set:   make(map[string]struct{}),
			}

			newGroup.set[w] = struct{}{}

			groups[sorted] = &newGroup
		} else {
			groups[sorted].set[w] = struct{}{}
		}
	}

	res := make(map[string][]string, len(words))

	for _, group := range groups {
		if len(group.set) == 1 {
			continue
		}

		wordSlice := make([]string, 0, len(group.set))

		for key := range group.set {
			wordSlice = append(wordSlice, key)
		}

		slices.SortFunc(wordSlice, func(a, b string) int {
			return strings.Compare(strings.ToLower(a), strings.ToLower(b))
		})

		res[group.first] = wordSlice
	}

	return res
}

func sortString(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	return string(s)
}
