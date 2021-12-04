package main

import (
	"elvensubmarine/sonar"
	"elvensubmarine/utilities"
	"fmt"
)

func main() {
	measurements := utilities.ReadSonarInputFile("inputs/day_1.txt")
	increases := sonar.CountDepthIncreases(measurements)
	fmt.Printf("Of %d sonar measurements, %d were increases\n", len(measurements), increases)

	windows := sonar.SlidingWindow(measurements, 3)
	windowed_increases := sonar.CountDepthIncreases(windows)
	fmt.Printf("Of %d windows, %d were increases", len(windows), windowed_increases)
}
