// TODO: add logging

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	fields          []int
	inputDelimeter  string
	outputDelimeter string
	separatedOnly   bool
}

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")

	var (
		opts      options
		fieldsStr string
	)

	flag.StringVar(&fieldsStr, "f", "", "specify the fields for printing")
	flag.StringVar(&opts.inputDelimeter, "d", "\t", "specify input fields delimeter")
	flag.StringVar(&opts.outputDelimeter, "output-delimeter", "\t", "specify output fields delimeter")

	flag.BoolVar(&opts.separatedOnly, "s", false, "do not print lines that do not contain the field separator character")

	flag.StringVar(&filepath, "file", "", "specify file")

	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	if len(fieldsStr) != 0 {
		fields, err := getFields(fieldsStr)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		opts.fields = fields
	}

	if debug {
		log.Printf("%s options: %#v\n", debugPrefix, opts)
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

	res := cut(lines, opts)

	for _, v := range res {
		fmt.Println(v)
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

func getFields(f string) ([]int, error) {
	fields := make([]int, 0)

	ranges := strings.Split(f, ",")

	for _, v := range ranges {
		values := strings.Split(v, "-")

		if len(values) == 1 {
			field, err := strconv.Atoi(values[0])
			if err != nil {
				return nil, fmt.Errorf("invalid field: %s", values[0])
			}

			field -= 1

			if field < 0 {
				return nil, fmt.Errorf("invalid field (<= 0): %s", values[0])
			}

			fields = append(fields, field)
		} else if len(values) == 2 {
			start, err := strconv.Atoi(values[0])
			if err != nil {
				return nil, fmt.Errorf("invalid field: %s", values[0])
			}

			end, err := strconv.Atoi(values[1])
			if err != nil {
				return nil, fmt.Errorf("invalid field: %s", values[1])
			}

			if start > end {
				start, end = end, start
			}

			start -= 1

			if start < 0 || end < 0 {
				return nil, fmt.Errorf("invalid field (<= 0)")
			}

			for i := start; i < end; i++ {
				fields = append(fields, i)
			}

		} else {
			return nil, fmt.Errorf("invalid range: %s", v)
		}
	}

	return fields, nil
}

func cut(lines []string, opts options) []string {
	res := make([]string, 0, len(lines))

	for _, v := range lines {
		fields := strings.Split(v, opts.inputDelimeter)

		if len(fields) == 1 {
			if !opts.separatedOnly {
				res = append(res, fields[0])
			}
		} else {
			targetFields := make([]string, 0, len(fields))

			if opts.fields == nil {
				targetFields = fields
			} else {
				for _, v := range opts.fields {
					if v >= len(fields) {
						continue
					} else {
						targetFields = append(targetFields, fields[v])
					}
				}
			}

			line := strings.Join(targetFields, opts.outputDelimeter)
			res = append(res, line)
		}
	}

	return res
}
