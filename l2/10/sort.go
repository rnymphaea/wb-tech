package main

import (
	"flag"
	"fmt"
	"log"
	_ "strings"
)

var (
	debug  bool
	check  bool
	unique bool
)

type sortOptions struct {
	key                  int
	numeric              bool
	reverse              bool
	month                bool
	human                bool
	ignoreTrailingBlanks bool
}

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.BoolVar(&check, "c", false, "check whether the given text is already sorted")
	flag.BoolVar(&unique, "u", false, "print unique strings only")

	var opts sortOptions
	flag.IntVar(&opts.key, "k", 1, "specify a sort field")
	flag.BoolVar(&opts.numeric, "n", false, "sort numerically")
	flag.BoolVar(&opts.reverse, "r", false, "reverse the result of comparison")
	flag.BoolVar(&opts.month, "M", false, "sort by month")
	flag.BoolVar(&opts.human, "h", false, "sort numerically, first by numeric sign (negative, zero, or positive); then by SI suffix; and finally by numeric value")
	flag.BoolVar(&opts.ignoreTrailingBlanks, "b", false, "ignore trailing blanks")

	flag.Parse()
	fmt.Println(opts)

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	if err := validateOpts(&opts); err != nil {
		fmt.Println("error:", err)
		return
	}
}

func validateOpts(opts *sortOptions) error {
	if opts.key < 1 {
		fmt.Println("invalid key, using default (key=1)")
		opts.key = 1
	}

	if !(opts.numeric && opts.month) && !(opts.numeric && opts.human) && !(opts.month && opts.human) {
		return nil
	} else {
		return fmt.Errorf("mutually exclusive flags")
	}
}
