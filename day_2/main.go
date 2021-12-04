package main

import (
	"elvensubmarine/submarine"
	"elvensubmarine/utilities"
	"fmt"
)

func main() {
	commands := utilities.ReadCommandsInputFile("inputs/day_2.txt")

	sub := submarine.Submarine{Depth: 0, Horizontal: 0}
	sub.SimplePilotMany(commands)
	fmt.Printf("Simple piloting: Reached depth %d, horizontal position %d, giving the product %d\n",
		sub.Depth, sub.Horizontal, sub.Depth*sub.Horizontal)

	sub = submarine.Submarine{Depth: 0, Horizontal: 0}
	sub.PilotMany(commands)
	fmt.Printf("Improved piloting: Reached depth %d, horizontal position %d, giving the product %d\n",
		sub.Depth, sub.Horizontal, sub.Depth*sub.Horizontal)
}
