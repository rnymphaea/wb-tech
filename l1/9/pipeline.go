package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	debug    bool
	sizedata int
	interval time.Duration
)

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.IntVar(&sizedata, "sizedata", 10, "amount of messages to send")
	flag.DurationVar(&interval, "interval", 200*time.Millisecond, "interval between messages (seconds)")
	flag.Parse()

	var start time.Time

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Amount of messages to send: %d\n", sizedata)
		log.Printf("Interval beetween messages: %v\n", interval)
		start = time.Now()
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < sizedata; i++ {
			v := rand.Intn(100)
			if debug {
				log.Printf("[DEBUG] %dms sender: sending %d to ch1\n", time.Since(start).Milliseconds(), v)
			}
			ch1 <- v
			time.Sleep(interval)
		}
		if debug {
			log.Printf("[DEBUG] %dms sender: closing ch1\n", time.Since(start).Milliseconds())
		}
		close(ch1)
	}()

	go func() {
		for v := range ch1 {
			if debug {
				log.Printf("[DEBUG] %dms processor: received %d from ch1\n", time.Since(start).Milliseconds(), v)
				log.Printf("[DEBUG] %dms processor: sending %d to ch2\n", time.Since(start).Milliseconds(), v*2)
			}
			ch2 <- v * 2
		}
		if debug {
			log.Printf("[DEBUG] %dms processor: ch1 is closed. Closing ch2\n", time.Since(start).Milliseconds())
		}
		close(ch2)
	}()

	for v := range ch2 {
		if debug {
			log.Printf("[DEBUG] %dms receiver: received %d from ch2\n", time.Since(start).Milliseconds(), v)
		} else {
			fmt.Println(v)
		}
	}

	if debug {
		log.Printf("[DEBUG] %dms receiver: ch2 is closed. Quitting...\n", time.Since(start).Milliseconds())
	}
}
