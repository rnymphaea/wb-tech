package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var debug bool

func sleep(duration time.Duration) {
	done := make(chan struct{})
	var start time.Time

	go func() {
		start = time.Now()
		for {
			if time.Since(start) >= duration {
				close(done)
				return
			}
		}
	}()

	<-done

	if debug {
		log.Printf("[DEBUG] %dms sleep: time exceeded\n", time.Since(start).Milliseconds())
	}
}

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	var ms int
	fmt.Printf("Enter number of milliseconds to sleep: ")
	fmt.Scan(&ms)

	duration := time.Duration(ms) * time.Millisecond

	sleep(duration)

	fmt.Println("Time exceeded")

}
