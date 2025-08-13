package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type command struct {
	args   []string
	input  string
	output string
	prev   *command
}

type sequence struct {
	commands []command
	pipe     bool
}

func main() {
	sigIntChan := make(chan os.Signal)
	signal.Notify(sigIntChan, syscall.SIGINT)

	reader := bufio.NewReader(os.Stdin)

	for {
		if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
			dir, _ := os.Getwd()
			fmt.Printf("\033[32m%s\033[0m$ ", dir)
		}

		go func() {
			<-sigIntChan
			fmt.Printf("\nInterrupt signal received")
		}()

		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				break
			}
			fmt.Fprintln(os.Stderr, "read error:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		seq := parse(input)
		if seq.pipe {
			seq.executePipeLine(sigIntChan)
		}
	}
}

func parse(input string) sequence {
	fields := strings.Fields(input)
	var seq sequence

	cmds := make([]command, 0)
	cmd := command{}

	ind := 0
	for ind < len(fields) {
		switch fields[ind] {
		case "<":
			if ind+1 < len(fields) {
				cmd.input = fields[ind+1]
			}
			ind++
			break
		case ">":
			if ind+1 < len(fields) {
				cmd.output = fields[ind+1]
			}
			ind++
			break
		case "|":
			cmds = append(cmds, cmd)
			cmd = command{}
			if len(cmds) > 0 {
				seq.pipe = true
			}
			ind++
			break
		default:
			cmd.args = append(cmd.args, fields[ind])
			ind++
		}
	}

	cmds = append(cmds, cmd)

	seq.commands = cmds
	return seq
}

func (s *sequence) executePipeLine(stop <-chan os.Signal) int {
	var status int

	return status
}
