package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	_ "os"
	"strconv"
	"strings"
)

const debugPrefix = "[DEBUG]"

var (
	debug    bool
	filepath string
)

type options struct {
	fields        []int
	all           bool
	delimeter     string
	onlySeparated bool
}

func main() {
	flag.BoolVar(&debug, "debug", false, "show logs to check the correctness of execution")

	var (
		opts      options
		fieldsStr string
	)

	flag.StringVar(&fieldsStr, "f", "", "specify the fields for printing")
	flag.StringVar(&opts.delimeter, "d", "\t", "specify fields delimeter")
	flag.BoolVar(&opts.onlySeparated, "s", false, "do not print lines that do not contain the field separator character")

	flag.StringVar(&filepath, "file", "", "specify file")

	flag.Parse()

	if debug {
		log.SetFlags(0)
		log.Println("Debug mode started")
	}

	if len(fieldsStr) == 0 {
		opts.all = true
	} else {
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
