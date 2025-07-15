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
	flag.BoolVar(&random, "random", false, "use the random array to check the correctness")
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
		size = rand.Intn(20) + 1
		arr = make([]int, 0, size)

		for i := 0; i < size; i++ {
			arr = append(arr, rand.Intn(50))
		}

		target = rand.Intn(size)
	} else {
		fmt.Printf("Enter the size of array: ")
		fmt.Scan(&size)

		arr = make([]int, size)

		fmt.Printf("Enter %d elements of array: ", size)
		for i := 0; i < size; i++ {
			fmt.Scan(&arr[i])
		}

		fmt.Printf("Enter the target index to delete: ")
		fmt.Scan(&target)

		if target < 0 || target > size-1 {
			fmt.Println("Invalid index")
			return
		}
	}

	if debug {
		log.Printf("[DEBUG] initial array: %v", arr)
		log.Printf("[DEBUG] index to delete: %d, arr[%d] = %d", target, target, arr[target])
	} else {
		fmt.Println("Initial array: ", arr)
		fmt.Println("Index to delete: ", target)
	}

	copy(arr[target:], arr[target+1:])
	arr = arr[:len(arr)-1]

	fmt.Println("Result: ", arr)
}
