package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
)

const (
	MIN = -90
	MAX = 60
)

var (
	debug      bool
	sizedata   int
	defaultarr bool
)

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.IntVar(&sizedata, "sizedata", 10, "amount of measurements")
	flag.BoolVar(&defaultarr, "default", false, "use default array of measurements (-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5)")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		if !defaultarr {
			log.Printf("Amount of measurements: %d\n", sizedata)
		}
	}

	m := make(map[int][]float64)

	var arr []float64
	if defaultarr {
		arr = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	} else {
		arr = make([]float64, 0, sizedata)
		for i := 0; i < sizedata; i++ {
			num := MIN + rand.Float64()*(MAX-MIN)
			num = math.Round(num*10) / 10
			arr = append(arr, num)
		}
	}

	if debug {
		log.Printf("[DEBUG] array of measurements: %v\n", arr)
	}

	for _, v := range arr {
		key := int(v) / 10 * 10
		m[key] = append(m[key], v)
	}

	fmt.Println(m)
}
