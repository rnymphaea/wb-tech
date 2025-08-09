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

const (
	debugPrefix = "[DEBUG]"
	maxStrings  = 50
)

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

	if opts.after < 0 {
		fmt.Println("invalid num lines. Using default (0)")
		opts.after = 0
	}

	if opts.before < 0 {
		fmt.Println("invalid num lines. Using default (0)")
		opts.before = 0
	}

	if opts.context < 0 {
		fmt.Println("invalid num lines. Using default (0)")
		opts.context = 0
	}

	if opts.context > 0 {
		opts.after = opts.context
		opts.before = opts.context
	}

	if opts.ignoreCase {
		pattern = strings.ToLower(pattern)
		if !opts.strict {
			pattern = "(?i)" + pattern
		}
	}

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
		log.Printf("Options: %#v\n", opts)
		log.Printf("Pattern: \"%s\"\n", pattern)
	}

	var (
		re         *regexp.Regexp
		reParseErr error
	)

	if !opts.strict {
		re, reParseErr = regexp.Compile(pattern)
		if reParseErr != nil {
			fmt.Println("error:", reParseErr)
			return
		}
	}

	var input io.Reader

	if len(filepath) > 0 {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		defer file.Close()

		input = file
	} else {
		fmt.Printf("Enter the text (Ctrl + D to stop): ")
		input = os.Stdin
	}

	lines, err := readlines(input)
	if err != nil {
		fmt.Println("error:", err)
	}

	if debug {
		log.Printf("\n%s lines: ", debugPrefix)
		for i, v := range lines {
			if i == maxStrings && maxStrings != len(lines)-1 {
				log.Printf("... and %d line(-s) more\n", len(lines)-maxStrings-1)
				break
			}

			log.Println(v)
		}

		log.Println()
	}

	var res []bool

	if !opts.strict {
		res = grepRE(lines, re, opts.invert)
	} else {
		res = grepStrict(lines, pattern, opts.ignoreCase, opts.invert)
	}

	var count int
	if opts.count {
		for _, v := range res {
			if v {
				count++
			}
		}

		fmt.Println("Number of occurrences:", count)
		return
	}

	fmt.Println("Result: ")
	for i, v := range res {
		if v {
			printline(lines, i, opts.before, opts.after, opts.printIndex)
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

func grepRE(lines []string, re *regexp.Regexp, invert bool) []bool {
	matches := make([]bool, len(lines))

	for i, v := range lines {
		matches[i] = re.MatchString(v)

		if invert {
			matches[i] = !matches[i]
		}
	}

	return matches
}

func grepStrict(lines []string, pattern string, ignoreCase, invert bool) []bool {
	matches := make([]bool, len(lines))

	for i, v := range lines {
		str := v
		if ignoreCase {
			str = strings.ToLower(str)
		}

		matches[i] = strings.Contains(str, pattern)

		if invert {
			matches[i] = !matches[i]
		}
	}

	return matches
}

func printline(lines []string, target, before, after int, printIndex bool) {
	start := target - before
	if start < 0 {
		start = 0
	}

	end := target + after + 1
	if end > len(lines) {
		end = len(lines)
	}

	for i := start; i < end; i++ {
		if (start != target || end != target+1) && i == target {
			fmt.Printf("---> ")
		}
		if printIndex {
			fmt.Printf("%d. ", i+1)
		}
		fmt.Println(lines[i])
	}

}
