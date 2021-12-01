package main

import (
	"elvensubmarine/sonar"
	"elvensubmarine/utilities"
	"fmt"
)

func main() {
	measurements := utilities.ReadInputFile("input_1.txt")
	increases := sonar.CountDepthIncreases(measurements)
	fmt.Printf("Of %d measurements, %d were increases\n", len(measurements), increases)
}
