package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

var (
	debug      bool
	goroutines int
)

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.IntVar(&goroutines, "goroutines", 10, "amount of goroutines")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Amount of goroutines: %d\n", goroutines)
	}

	var count int
	fmt.Printf("Amount of incrementing operations for each goroutine: ")
	fmt.Scan(&count)

	if debug {
		log.Printf("[DEBUG] expected value: %d\n", goroutines*count)
	}

	c := Counter{}

	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < count; j++ {
				if debug {
					log.Printf("[DEBUG] goroutine #%d: incrementing counter (current = %d)\n", i, c.value)
				}

				c.Inc()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Result: ", c.value)
}
