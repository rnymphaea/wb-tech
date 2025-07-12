package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

type safemap struct {
	mu sync.RWMutex
	m  map[int]int
}

func (sm *safemap) Set(key, value int) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

func (sm *safemap) Get(key int) (int, bool) {
	sm.mu.RLock()
	v, ok := sm.m[key]
	sm.mu.RUnlock()
	return v, ok
}

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

	sm := safemap{
		m: make(map[int]int),
	}

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
				sm.Set(key, value)

				value, _ = sm.Get(key)
				if debug {
					log.Printf("[DEBUG] goroutine %d: get: key - %d, value - %d\n", c, key, value)
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(sm.m)
}
