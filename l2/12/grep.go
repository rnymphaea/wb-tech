package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

const debugPrefix = "[DEBUG]"

var (
	debug    bool
	filepath string
)

type options struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	strict     bool
	printIndex bool
}

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")

	var opts options

	flag.IntVar(&opts.after, "A", 0, "print num lines of trailing context after matching lines")
	flag.IntVar(&opts.before, "B", 0, "print num lines of leading context before matching lines")
	flag.IntVar(&opts.context, "C", 0, "print num lines of leading and trailing output context")
	flag.BoolVar(&opts.count, "c", false, "suppress normal output; instead print a count of matching lines for each input file")
	flag.BoolVar(&opts.ignoreCase, "i", false, "ignore case distinctions in patterns and input data, so that characters that differ only in case match each other")
	flag.BoolVar(&opts.invert, "v", false, "invert the sense of matching, to select non-matching lines")
	flag.BoolVar(&opts.strict, "F", false, "interpret patterns as fixed strings, not regular expressions")
	flag.BoolVar(&opts.printIndex, "n", false, "prefix each line of output with the 1-based line number within its input file")
	flag.StringVar(&filepath, "file", "", "specify file")

	flag.Parse()

	var pattern string

	if len(flag.Args()) == 0 {
		fmt.Println("no pattern provided")
		return
	} else {
		pattern = flag.Args()[0]
	}

	if opts.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Options: %#v\n", opts)
		log.Printf("Pattern: \"%s\"\n", pattern)
	}

	if !opts.strict {
		expr, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
	}

	var input io.Reader

	if len(filepath) > 0 {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		defer file.Close()

		input = file
	} else {
		fmt.Printf("Enter the text: ")
		input = os.Stdin
	}

	lines, err := readlines(input)
	if err != nil {
		fmt.Println("error: ", err)
	}

	if debug {
		log.Printf("%s lines: ", debugPrefix)
		for _, v := range lines {
			log.Println(v)
		}
	}

}

func readlines(r io.Reader) ([]string, error) {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
