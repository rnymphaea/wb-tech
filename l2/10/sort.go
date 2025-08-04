package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	_ "strings"
)

const debugPrefix = "[DEBUG]"

var (
	debug    bool
	check    bool
	unique   bool
	filepath string
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
	flag.StringVar(&filepath, "file", "", "sort lines from specified file")

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

	var (
		input    []string
		inputErr error
	)

	if len(filepath) > 0 {
		input, inputErr = readLinesFromFile(filepath)
		if inputErr != nil {
			fmt.Println("error:", inputErr)
			return
		}
	} else {
		var size int
		fmt.Printf("Enter the number of strings: ")
		fmt.Scan(&size)

		input, inputErr = readLinesFromStdin(size)
		if inputErr != nil {
			fmt.Println("error:", inputErr)
			return
		}
	}

	if debug {
		log.Printf("%s lines:\n", debugPrefix)
		for i := 0; i < len(input); i++ {
			log.Println(input[i])
		}
	}

	if err := validateOpts(&opts); err != nil {
		fmt.Println("error:", err)
		return
	}

	sort(input, &opts)
}

func readLinesFromFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func readLinesFromStdin(size int) ([]string, error) {
	lines := make([]string, size)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter %d lines that you want to sort: ", size)

	for i := 0; i < size; i++ {
		scanner.Scan()
		line := scanner.Text()

		if len(line) == 0 {
			break
		} else {
			lines[i] = line
		}
	}

	return lines, scanner.Err()
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

func sort(arr []string, opts *sortOptions) []string {
	return arr
}
