package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
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
	ClosedChannelExit()
	PanicExit()
	TimeAfterExit()
	OsExit()
}

// ConditionalExit демонстрирует остановку горутины по флагу stop.
// Горутина проверяет условие !stop на каждой итерации и завершается, когда флаг становится true.
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
		fmt.Printf("%s: the condition stops being fulfilled. Quitting...\n", funcName)
	}

	stop = true
	wg.Wait()
}

// ChannelExit демонстрирует остановку горутины через сигнальный канал.
// Использует select для проверки канала stop. При получении сигнала горутина завершает работу.
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
		fmt.Printf("%s: sending stop signal to channel. Quitting...\n", funcName)
	}

	stop <- struct{}{}
	wg.Wait()
}

// ContextWithTimeoutExit демонстрирует остановку горутины по истечении таймаута контекста.
// Использует context.WithTimeout для автоматической отмены через указанное время.
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

// ContextWithCancelExit демонстрирует остановку горутины через отмену контекста.
// Использует context.WithCancel для ручного управления завершением работы.
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

// RuntimeGoExit демонстрирует остановку текущей горутины через runtime.Goexit().
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
			log.Printf("[DEBUG] %s: [%dms] goroutine: quitting after sleep\n", funcName, time.Since(start).Milliseconds())
		} else {
			fmt.Printf("%s: quitting goroutine after sleep\n", funcName)
		}

		runtime.Goexit()
	}()

	if debug {
		log.Printf("%s: [%dms] waiting for goroutine", funcName, time.Since(start).Milliseconds())
	}
	wg.Wait()
}

// ClosedChannelExit демонстрирует остановку горутины при закрытии канала.
// Горутина завершает работу при выходе из range по закрытому каналу.
func ClosedChannelExit() {
	var (
		wg       sync.WaitGroup
		start    time.Time
		interval time.Duration = 200 * time.Millisecond
		funcName string        = "func ClosedChannelExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine by closing a channel")
		log.Printf("Interval between messages sending to channel: %v\n", interval)
		start = time.Now()
	}

	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range ch {
			if debug {
				log.Printf("[DEBUG] %s: [%dms] receiver got: %d\n", funcName, time.Since(start).Milliseconds(), v)
			} else {
				fmt.Printf("%s: receiver: channel is not closed\n", funcName)
			}
		}

		if debug {
			log.Printf("[DEBUG] %s: [%dms] receiver: channel is closed\n", funcName, time.Since(start).Milliseconds())
		} else {
			fmt.Printf("%s: receiver discovered that channel is closed. Quitting...\n", funcName)
		}
	}()

	for i := 0; i < 5; i++ {
		if debug {
			log.Printf("[DEBUG] %s: [%dms] sender: sending %d\n", funcName, time.Since(start).Milliseconds(), i)
		}
		ch <- i
		time.Sleep(interval)
	}

	if debug {
		log.Printf("[DEBUG] %s: [%dms] sender: closing the channel\n", funcName, time.Since(start).Milliseconds())
	}
	close(ch)
	wg.Wait()
}

// PanicExit демонстрирует аварийное завершение горутины через panic с последующим восстановлением через recover.
func PanicExit() {
	var (
		wg       sync.WaitGroup
		start    time.Time
		timeout  time.Duration = 1 * time.Second
		funcName string        = "func RuntimeGoExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine by panic")
		log.Printf("Timeout: %v\n", timeout)
		start = time.Now()
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s: recovered from panic:", funcName, r)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		defer func() {
			if r := recover(); r != nil {
				if debug {
					log.Printf("[DEBUG] %s: [%dms] goroutine recovered from panic: %v\n", funcName, time.Since(start).Milliseconds(), r)
				} else {
					fmt.Printf("%s: goroutine recovered from panic: %v\n", funcName, r)
				}
			}
		}()

		if debug {
			log.Printf("[DEBUG] %s: [%dms] entering to goroutine\n", funcName, time.Since(start).Milliseconds())
		} else {
			fmt.Printf("%s: entering to goroutine\n", funcName)
		}

		time.Sleep(timeout)
		if debug {
			log.Printf("[DEBUG] %s: [%dms] goroutine: calling panic...\n", funcName, time.Since(start).Milliseconds())
		}

		panic("stop goroutine")
	}()

	wg.Wait()
}

// TimeAfterExit демонстрирует остановку горутины по истечении времени через time.After.
func TimeAfterExit() {
	var (
		wg       sync.WaitGroup
		start    time.Time
		interval time.Duration = 200 * time.Millisecond
		timeout  time.Duration = 1 * time.Second
		funcName string        = "func TimeAfterExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine with time.After")
		log.Printf("Interval between messages of goroutine: %v\n", interval)
		log.Printf("Timeout: %v\n", timeout)
		start = time.Now()
	}

	quit := time.After(timeout)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-quit:
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
	}()

	if debug {
		log.Printf("[DEBUG] %s: [%dms] sleeping for %v...\n", funcName, time.Since(start).Milliseconds(), timeout.Milliseconds())
	}

	time.Sleep(timeout)

	wg.Wait()
}

// OsExit демонстрирует завершение всей программы из горутины через os.Exit.
func OsExit() {
	var (
		wg       sync.WaitGroup
		start    time.Time
		timeout  time.Duration = 1 * time.Second
		funcName string        = "func OsExit"
	)

	if debug {
		log.Println("\nDemonstrating of exiting a goroutine with os.Exit")
		log.Printf("Timeout: %v\n", timeout)
		start = time.Now()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if debug {
			log.Printf("[DEBUG] %s: [%dms] entering to goroutine and sleeping for %v\n", funcName, time.Since(start).Milliseconds(), timeout)
		} else {
			fmt.Printf("%s: entering to goroutine\n", funcName)
		}

		time.Sleep(timeout)

		if debug {
			log.Printf("[DEBUG] %s: [%dms] goroutine: calling os.Exit...\n", funcName, time.Since(start).Milliseconds())
		} else {
			fmt.Printf("%s: goroutine is calling os.Exit\n", funcName)
		}

		os.Exit(0)
	}()

	wg.Wait()
}
