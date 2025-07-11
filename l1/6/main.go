package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"runtime"
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
	ContextWithTimeoutExit()
	ContextWithCancelExit()
	RuntimeGoExit()
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
				fmt.Printf("%s: goroutine is working\n", funcName)
			}
			time.Sleep(interval)
		}

		if debug {
			log.Printf("[DEBUG] %s: [%vms] goroutine: condition is no longer fulfilled\n", funcName, time.Since(start).Milliseconds())
		} else {
			fmt.Printf("%s: goroutine stops working\n", funcName)
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
					log.Printf("[DEBUG] %s: [%vms] goroutine: received stop signal from channel\n", funcName, time.Since(start).Milliseconds())
				} else {
					fmt.Printf("%s: goroutine stops working\n", funcName)
				}
				return
			default:
				if debug {
					log.Printf("[DEBUG] %s: [%vms] goroutine is working\n", funcName, time.Since(start).Milliseconds())
				} else {
					fmt.Printf("%s: goroutine is working\n", funcName)
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

func ContextWithTimeoutExit() {
	var (
		wg       sync.WaitGroup
		start    time.Time
		interval time.Duration = 200 * time.Millisecond
		timeout  time.Duration = 1 * time.Second
		funcName string        = "func ContextWithTimeoutExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine with context.WithTimeout")
		log.Printf("Interval between messages of goroutine: %v\n", interval)
		log.Printf("Timeout: %v\n", timeout)
		start = time.Now()
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				if debug {
					log.Printf("[DEBUG] %s: [%dms] goroutine: time exceeded\n", funcName, time.Since(start).Milliseconds())
				} else {
					fmt.Printf("%s: goroutine stops working\n", funcName)
				}
				return
			default:
				if debug {
					log.Printf("[DEBUG] %s: [%dms] goroutine is working\n", funcName, time.Since(start).Milliseconds())
				} else {
					fmt.Printf("%s: goroutine is working\n", funcName)

				}
				time.Sleep(interval)
			}
		}
	}(ctx)

	if debug {
		log.Printf("[DEBUG] %s: [%dms] sleeping for %v...\n", funcName, time.Since(start).Milliseconds(), timeout.Milliseconds())
	}

	time.Sleep(timeout)

	wg.Wait()
}

func ContextWithCancelExit() {
	var (
		wg       sync.WaitGroup
		start    time.Time
		interval time.Duration = 200 * time.Millisecond
		timeout  time.Duration = 1 * time.Second
		funcName string        = "func ContextWithCancelExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine with context.WithCancel")
		log.Printf("Interval between messages of goroutine: %v\n", interval)
		log.Printf("Timeout: %v\n", timeout)
		start = time.Now()
	}

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				if debug {
					log.Printf("[DEBUG] %s: [%dms] goroutine: received cancel signal\n", funcName, time.Since(start).Milliseconds())
				} else {
					fmt.Printf("%s: goroutine stops working...\n", funcName)
				}
				return
			default:
				if debug {
					log.Printf("[DEBUG] %s: [%dms] goroutine is working\n", funcName, time.Since(start).Milliseconds())
				} else {
					fmt.Printf("%s: goroutine is working\n", funcName)

				}
				time.Sleep(interval)
			}
		}
	}(ctx)

	time.Sleep(timeout)
	if debug {
		log.Printf("[DEBUG] %s: [%dms] calling cancel()\n", funcName, time.Since(start).Milliseconds())
	} else {
		fmt.Printf("%s: calling cancel(). Quitting...\n", funcName)
	}

	cancel()
	wg.Wait()
}

func RuntimeGoExit() {
	var (
		wg       sync.WaitGroup
		start    time.Time
		timeout  time.Duration = 1 * time.Second
		funcName string        = "func RuntimeGoExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine with runtime.Goexit")
		log.Printf("Timeout: %v\n", timeout)
		start = time.Now()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if debug {
			log.Printf("[DEBUG] %s: [%dms] entering to goroutine and sleeping for %v...\n", funcName, time.Since(start).Milliseconds(), timeout)
		} else {
			fmt.Printf("%s: entering to goroutine\n", funcName)
		}

		time.Sleep(timeout)

		if debug {
			log.Printf("[DEBUG] %s: [%dms] goroutine: quitting after sleep...\n", funcName, time.Since(start).Milliseconds())
		} else {
			fmt.Printf("%s: quitting goroutine after sleep...\n", funcName)
		}

		runtime.Goexit()
	}()

	if debug {
		log.Printf("%s: [%dms] waiting for goroutine", funcName, time.Since(start).Milliseconds())
	}
	wg.Wait()
}
