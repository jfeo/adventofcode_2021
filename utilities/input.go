package utilities

import (
	"bufio"
	"elvensubmarine/submarine"
	"os"
	"strconv"
	"strings"
)

func ReadSonarInputFile(name string) (inputs []int) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	inputs = make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		inputs = append(inputs, int(input))
	}
	return
}

func ReadCommandsInputFile(name string) (commands []submarine.Command) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	commands = make([]submarine.Command, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")

		if len(row) != 2 {
			panic("Each row must contain 2 items")
		}

		x, err := strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			panic(err)
		}

		cmd := submarine.Command{
			Type: row[0],
			X:    int(x),
		}

		commands = append(commands, cmd)
	}
	return
}
