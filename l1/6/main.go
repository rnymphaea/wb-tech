package main

import (
	_ "context"
	"flag"
	"fmt"
	"log"
	_ "runtime"
	"sync"
	"time"
)

var debug bool

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	ConditionalExit()
	ChannelExit()
}

func ConditionalExit() {
	var (
		stop     bool = false
		wg       sync.WaitGroup
		start    time.Time
		interval time.Duration = 200 * time.Millisecond
		timeout  time.Duration = 1 * time.Second
		funcName string        = "func ConditionalExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine on condition")
		log.Printf("Interval between messages of goroutine: %v\n", interval)
		log.Printf("Timeout: %v\n", timeout)
		start = time.Now()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for !stop {
			if debug {
				log.Printf("[DEBUG] %s: [%vms] goroutine is working\n", funcName, time.Since(start).Milliseconds())
			} else {
				fmt.Printf("%s: goroutine is working...\n", funcName)
			}
			time.Sleep(interval)
		}

		if debug {
			log.Printf("[DEBUG] %s: [%vms] goroutine: condition is no longer fulfilled\n", funcName, time.Since(start).Milliseconds())
		} else {
			fmt.Printf("%s: goroutine stops working...\n", funcName)
		}
	}()

	time.Sleep(timeout)

	if debug {
		log.Printf("[DEBUG] %s: [%vms] the condition stops being fulfilled\n", funcName, time.Since(start).Milliseconds())
	} else {
		fmt.Println("The condition stops being fulfilled. Quitting...")
	}

	stop = true
	wg.Wait()
}

func ChannelExit() {
	var (
		wg       sync.WaitGroup
		start    time.Time
		interval time.Duration = 200 * time.Millisecond
		timeout  time.Duration = 1 * time.Second
		funcName string        = "func ChannelExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine on channel notification")
		log.Printf("Interval between messages of goroutine: %v\n", interval)
		log.Printf("Timeout: %v\n", timeout)
		start = time.Now()
	}

	stop := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-stop:
				if debug {
					log.Printf("[DEBUG] %s: [%vms] goroutine received stop signal from channel\n", funcName, time.Since(start).Milliseconds())
				} else {
					fmt.Printf("%s: goroutine stops working...\n", funcName)
				}
				return
			default:
				if debug {
					log.Printf("[DEBUG] %s: [%vms] goroutine is working\n", funcName, time.Since(start).Milliseconds())
				} else {
					fmt.Printf("%s: goroutine is working...\n", funcName)
				}
				time.Sleep(interval)
			}
		}
	}()

	time.Sleep(timeout)

	if debug {
		log.Printf("[DEBUG] %s: [%vms] sending stop signal to channel\n", funcName, time.Since(start).Milliseconds())
	} else {
		fmt.Println("Sending stop signal to channel. Quitting...")
	}

	stop <- struct{}{}
	wg.Wait()
}
