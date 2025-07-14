package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
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

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Amount of workers: %d\n", workers)
		log.Printf("Amount of messages to send: %d\n", sizedata)
		log.Printf("Interval beetween messages: %v\n", interval)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if debug {
			log.Println("\nReceived SIGINT. Shutting down...")
		}
		cancel()
	}()

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup, idx int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					if debug {
						log.Printf("[DEBUG] worker %d received shutdown signal\n", idx)
					}
					return
				case v, ok := <-ch:
					if ok {
						if debug {
							log.Printf("[DEBUG] worker %d received from chan: %d\n", idx, v)
						} else {
							fmt.Println(v)
						}
					} else {
						log.Printf("[DEBUG] worker %d discovered that the channel is closed\n", idx)
						return
					}
				}
			}
		}(ctx, &wg, i)
	}

	for i := 0; i < sizedata; i++ {
		select {
		case <-ctx.Done():
			if debug {
				log.Println("[DEBUG] Main goroutine received shutdown signal\n")
			}
			wg.Wait()
			return
		default:
			if debug {
				log.Printf("[DEBUG] Main goroutine sended to chan: %d\n", i)
			}
			ch <- i
			time.Sleep(interval)
		}
	}

	wg.Wait()
	close(ch)
}
