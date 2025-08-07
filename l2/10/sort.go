package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
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
	sep                  string
	numeric              bool
	reverse              bool
	month                bool
	human                bool
	ignoreTrailingBlanks bool
}

type line struct {
	text string
	key  int
	sep  string
}

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")
	flag.BoolVar(&check, "c", false, "check whether the given text is already sorted")
	flag.BoolVar(&unique, "u", false, "print unique strings only")
	flag.StringVar(&filepath, "file", "", "sort lines from specified file")

	var opts sortOptions
	flag.IntVar(&opts.key, "k", 0, "specify a sort field")
	flag.StringVar(&opts.sep, "t", "\t", "use character separator as the field separator")
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

	res := sort(input, &opts)
	fmt.Println("Result:")

	for _, v := range res {
		fmt.Println(v)
	}
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

		if len(line) != 0 {
			lines[i] = line
		} else {
			break
		}
	}

	return lines, scanner.Err()
}

func validateOpts(opts *sortOptions) error {
	if opts.key < 1 {
		fmt.Println("invalid key, using default (key=1)")
		opts.key = 0
	} else {
		opts.key -= 1
	}

	if !(opts.numeric && opts.month) && !(opts.numeric && opts.human) && !(opts.month && opts.human) {
		return nil
	} else {
		return fmt.Errorf("mutually exclusive flags")
	}
}

func sort(arr []string, opts *sortOptions) []string {
	res := make([]string, len(arr))
	lines := make([]line, len(arr))

	key := opts.key
	sep := opts.sep

	for i, str := range arr {
		lines[i] = line{
			text: str,
			key:  key,
			sep:  sep,
		}
	}

	var cmp func(a, b line) int
	if opts.numeric {
		cmp = cmpNumeric
	}

	slices.SortStableFunc(lines, cmp)

	for i, v := range lines {
		res[i] = v.text
	}

	return res
}

func cmpNumeric(a, b line) int {
	const funcName = "cmpNumeric"

	txt1 := strings.Split(a.text, a.sep)
	txt2 := strings.Split(b.text, b.sep)

	if debug {
		log.Printf("%s %s: after splitting got: %q, %q\n", debugPrefix, funcName, txt1, txt2)
	}

	if a.key >= len(txt1) {
		if b.key >= len(txt2) {
			return strings.Compare(a.text, b.text)
		} else {
			return -1
		}
	} else if b.key >= len(txt2) {
		return 1
	}

	num1, err1 := strconv.ParseFloat(txt1[a.key], 64)
	num2, err2 := strconv.ParseFloat(txt2[b.key], 64)

	if debug {
		log.Printf("%s %s: num1: %f, err1: %v, num2: %f, err2: %v\n", debugPrefix, funcName, num1, err1, num2, err2)
	}

	if err1 != nil {
		if err2 != nil {
			if debug {
				log.Printf("%s %s: both strings [%s] and [%s] don't have numbers at col %d\n", debugPrefix, funcName, a.text, b.text, a.key)
			}

			return strings.Compare(a.text, b.text)
		} else {
			return -1
		}
	} else if err2 != nil {
		return 1
	}

	if num1 < num2 {
		return -1
	} else if num1 == num2 {
		return 0
	}

	return 1
}
