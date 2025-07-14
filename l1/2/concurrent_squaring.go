package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

var debug bool

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.Parse()

	arr := [5]int{2, 4, 6, 8, 10}
	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Initial array: %v", arr)
	}

	var wg sync.WaitGroup

	for idx, value := range arr {
		wg.Add(1)
		go func(idx, v int) {
			defer wg.Done()

			if debug {
				log.Printf("[DEBUG] Goroutine %d received value %d. After processing got: %d", idx, v, v*v)
			} else {
				fmt.Println(v * v)
			}
		}(idx, value)
	}

	wg.Wait()
}
