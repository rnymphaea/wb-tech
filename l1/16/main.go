package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
)

var (
	debug  bool
	random bool
)

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.BoolVar(&random, "random", false, "use the random array to check the correctness of algorithm")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Random array: %t\n", random)
	}

	var (
		arr     []int
		size    int
		reverse bool
	)

	if random {
		size = rand.Intn(20)
		arr = make([]int, 0, size)

		for i := 0; i < size; i++ {
			arr = append(arr, rand.Intn(50))
		}
	} else {
		fmt.Printf("Enter the size of array: ")
		fmt.Scan(&size)

		arr = make([]int, size)

		fmt.Printf("Enter %d elements of array: ", size)
		for i := 0; i < size; i++ {
			fmt.Scan(&arr[i])
		}
	}

	if debug {
		log.Printf("[DEBUG] initial array: %v\n", arr)
	}

	var reverseInput string
	fmt.Printf("Reverse order? [y/n] ")
	fmt.Scan(&reverseInput)

	if reverseInput != "y" && reverseInput != "n" {
		fmt.Println("Invalid option")
		return
	}

	reverse = reverseInput == "y"

	res := QuickSort(arr, reverse)
	fmt.Println("Result: ", res)
}

func QuickSort(arr []int, reverse bool) []int {
	if debug {
		log.Printf("\n[DEBUG] current array: %v", arr)
	}

	if len(arr) < 2 {
		return arr
	}

	left := make([]int, 0)
	right := make([]int, 0)
	base := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] <= base {
			if !reverse {
				left = append(left, arr[i])
			} else {
				right = append(right, arr[i])
			}
		} else {
			if !reverse {
				right = append(right, arr[i])
			} else {
				left = append(left, arr[i])
			}
		}
	}

	if debug {
		log.Printf("[DEBUG] base = %d\n", base)
		log.Printf("[DEBUG] left array: %v\n", left)
		log.Printf("[DEBUG] right array: %v\n", right)
	}

	res := make([]int, 0)
	res = append(res, QuickSort(left, reverse)...)
	res = append(res, base)
	res = append(res, QuickSort(right, reverse)...)

	return res
}
