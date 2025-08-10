// TODO: add logging

package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const debugPrefix = "[DEBUG]"

var debug bool

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")

	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	sig := func(after time.Duration) <-chan any {
		c := make(chan any)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func or(channels ...<-chan any) <-chan any {
	if len(channels) == 0 {
		return nil
	}

	if len(channels) == 1 {
		return channels[0]
	}

	start := time.Now()

	out := make(chan any)
	sig := make(chan struct{})

	for i, v := range channels {
		go func(num int) {
			_, ok := <-v

			for ok {
				_, ok = <-v
			}

			sig <- struct{}{}
		}(i)
	}

	go func() {
		<-sig

		if debug {
			log.Printf("%s [%dms] one of the channels has been closed. Closing or channel...", debugPrefix, time.Since(start).Milliseconds())
		}

		close(out)
	}()

	return out
}
