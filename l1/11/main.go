package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
)

var debug bool

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	set1 := GetRandomSet()
	set2 := GetRandomSet()

	fmt.Println("Set 1:", set1)
	fmt.Println("Set 2:", set2)

	result := Intersection(set1, set2)
	fmt.Println("Intersection of these sets: ", result)
}

func Intersection(set1, set2 []int) []int {
	result := make([]int, 0, max(len(set1), len(set2)))
	m := make(map[int]struct{})

	for _, v := range set1 {
		m[v] = struct{}{}
	}

	if debug {
		log.Printf("[DEBUG] map after writing elements of set 1: %v\n", m)
		log.Println("[DEBUG] checking set 2")
	}

	for _, v := range set2 {
		_, ok := m[v]
		if ok {
			result = append(result, v)
		}

		if debug {
			log.Printf("[DEBUG] check: element %d is in set 1: %t\n", v, ok)
		}
	}

	return result
}

func GetRandomSet() []int {
	l := rand.Intn(10)
	// this map guarantees the uniqueness of the elements
	m := make(map[int]struct{})
	for i := 0; i < l; i++ {
		key := rand.Intn(20)
		m[key] = struct{}{}
	}

	if debug {
		log.Printf("[DEBUG] map after initializing: %v\n", m)
	}

	set := make([]int, 0, l)

	for key, _ := range m {
		set = append(set, key)
	}

	return set
}
