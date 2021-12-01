package utilities

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInputFile(name string) (inputs []int) {
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
