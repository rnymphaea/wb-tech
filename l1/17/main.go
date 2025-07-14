package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"slices"

	"golang.org/x/exp/constraints"
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
		log.Printf("Random: %t\n", random)
	}

	var (
		arr    []int
		size   int
		target int
	)

	if random {
		size = rand.Intn(20)
		arr = make([]int, 0, size)

		for i := 0; i < size; i++ {
			arr = append(arr, rand.Intn(50))
		}

		target = rand.Intn(50)
	} else {
		fmt.Printf("Enter the size of array: ")
		fmt.Scan(&size)

		arr = make([]int, size)

		fmt.Printf("Enter %d elements of array: ", size)
		for i := 0; i < size; i++ {
			fmt.Scan(&arr[i])
		}

		fmt.Printf("Enter the target to find: ")
		fmt.Scan(&target)
	}

	slices.Sort(arr)

	if debug {
		log.Printf("[DEBUG] initial array: %v\n", arr)
		log.Printf("[DEBUG] target: %d\n", target)
	} else {
		fmt.Println("Sorted arr: ", arr)
	}

	result := BinarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Index of %d in array is: %d\n", target, result)
	} else {
		fmt.Printf("Element %d is not in array!\n", target)
	}
}

func BinarySearch[T constraints.Ordered](arr []T, target T) int {
	left := 0
	right := len(arr)
	iteration := 1

	for left < right {
		mid := (left + right) / 2

		if debug {
			log.Printf("[DEBUG] iteration #%d\n", iteration)
			log.Printf("[DEBUG] left = %d, right = %d, mid = %d\n", left, right, mid)
			log.Printf("[DEBUG] arr[mid] = %v\n", arr[mid])
		}

		if target < arr[mid] {
			right = mid

			if debug {
				log.Printf("[DEBUG] target < arr[mid]. Finding in %v\n", arr[left:right])
			}
		} else if target == arr[mid] {
			return mid
		} else {
			left = mid + 1

			if debug {
				log.Printf("[DEBUG] target > arr[mid]. Finding in %v\n", arr[left:right])
			}
		}

		iteration++
	}
	return -1
}
