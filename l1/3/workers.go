package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	debug    bool
	workers  int
	sizedata int
	interval time.Duration
)

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.IntVar(&workers, "workers", 10, "amount of workers")
	flag.IntVar(&sizedata, "sizedata", 10, "amount of messages to send to chan")
	flag.DurationVar(&interval, "interval", 1*time.Second, "interval between messages (seconds)")
	flag.Parse()

	var start time.Time
	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Amount of workers: %d\n", workers)
		log.Printf("Amount of messages to send: %d\n", sizedata)
		log.Printf("Interval beetween messages: %v\n", interval)
		start = time.Now()
	}

	ch := make(chan int)

	for i := 0; i < workers; i++ {
		go func(idx int) {
			for {
				v, ok := <-ch
				if ok {
					if debug {
						log.Printf("[DEBUG] %dms worker %d received from chan: %d\n", time.Since(start).Milliseconds(), idx, v)
					} else {
						fmt.Println(v)
					}
				} else {
					log.Printf("[DEBUG] %dms worker %d discovered that the channel is closed\n", time.Since(start).Milliseconds(), idx)
					break
				}
			}
		}(i)
	}

	for i := 0; i < sizedata; i++ {
		if debug {
			log.Printf("[DEBUG] %dms main goroutine sended to chan: %d\n", time.Since(start).Milliseconds(), i)
		}
		ch <- i
		time.Sleep(interval)
	}

	close(ch)
}
