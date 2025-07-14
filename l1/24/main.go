package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
)

const (
	MIN = -100
	MAX = 100
)

var (
	debug  bool
	random bool
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.BoolVar(&random, "random", false, "use the random array to check the correctness")
	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Random: %t\n", random)
	}

	var (
		x1, x2, y1, y2 float64
	)

	if random {
		x1 = GetRandomFloat()
		x2 = GetRandomFloat()
		y1 = GetRandomFloat()
		y2 = GetRandomFloat()
	} else {
		fmt.Printf("Enter (x y) of point 1: ")
		fmt.Scan(&x1, &y1)
		fmt.Printf("Enter (x y) of point 2: ")
		fmt.Scan(&x2, &y2)
	}

	p1 := NewPoint(x1, y1)
	p2 := NewPoint(x2, y2)

	if debug {
		log.Printf("[DEBUG] point 1: (%f, %f)\n", p1.x, p1.y)
		log.Printf("[DEBUG] point 2: (%f, %f)\n", p2.x, p2.y)
	} else if random {
		fmt.Printf("Point 1: (%f, %f)\n", p1.x, p1.y)
		fmt.Printf("Point 2: (%f, %f)\n", p2.x, p2.y)
	}

	dist := p1.Distance(p2)
	fmt.Println("Distance between points: ", dist)

}

func GetRandomFloat() float64 {
	return MIN + rand.Float64()*(MAX-MIN)
}
