package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	debug    bool
	timeout  int
	interval time.Duration
)

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.IntVar(&timeout, "timeout", 5, "number of seconds to finish the program")
	flag.DurationVar(&interval, "interval", 1*time.Second, "interval between messages (seconds)")
	flag.Parse()

	var start time.Time

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Timeout: %v\n", time.Duration(timeout)*time.Second)
		log.Printf("Interval beetween messages: %v\n", interval)
		start = time.Now()
	}

	quit := time.After(time.Duration(timeout) * time.Second)
	stopReceiver := make(chan struct{})

	ch := make(chan int)
	defer close(ch)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-stopReceiver:
				if debug {
					log.Printf("[DEBUG] %dms Receiver: time exceeded\n", time.Since(start).Milliseconds())
				}
				return
			case v := <-ch:
				if debug {
					log.Printf("[DEBUG] %dms Receiver got: %d\n", time.Since(start).Milliseconds(), v)
				} else {
					fmt.Println(v)
				}
			}
		}
	}()

	for {
		select {
		case <-quit:
			if debug {
				log.Printf("[DEBUG] %dms Sender: time exceeded\n", time.Since(start).Milliseconds())
			}

			stopReceiver <- struct{}{}
			wg.Wait()
			return
		default:
			v := rand.Intn(100)

			if debug {
				log.Printf("[DEBUG] %dms Sender send: %d", time.Since(start).Milliseconds(), v)
			}

			ch <- v
			time.Sleep(interval)
		}
	}

}
