package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

var (
	debug      bool
	goroutines int
)

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.IntVar(&goroutines, "goroutines", 10, "amount of goroutines to write in map")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Amount of goroutines: %d\n", goroutines)
	}

	var sm sync.Map

	var wg sync.WaitGroup

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(c int) { // c stands for coefficient
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key, value := j, j*c
				if debug {
					log.Printf("[DEBUG] goroutine %d: set: key - %d, value - %d\n", c, key, value)
				}
				sm.Store(key, value)

				v, _ := sm.Load(key)
				value = v.(int)
				if debug {
					log.Printf("[DEBUG] goroutine %d: get: key - %d, value - %d\n", c, key, value)
				}
			}
		}(i)
	}

	wg.Wait()

	PrintMap(sm)
}

func PrintMap(sm sync.Map) {
	fmt.Printf("map[")
	sm.Range(func(k, v interface{}) bool {
		fmt.Printf("%v:%v ", k, v)
		return true
	})
	fmt.Printf("\b]\n")
}
