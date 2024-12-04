package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"math"
)

func solutionPart01(lines []string) {
	var unsavedCounter = 0
	for _, line := range lines {
		var report = SplitInts(line, " ")

		var ascCounter = 0
		var descCounter = 0
		var equalsCounter = 0
		var isResolved = false
		for i := 0; i < len(report)-1; i++ {
			var diff = report[i+1] - report[i]

			if abs(diff) >= 4 {
				unsavedCounter++
				isResolved = true
				break
			} else {
				if diff > 0 {
					ascCounter++
				} else if diff < 0 {
					descCounter++
				} else {
					equalsCounter++
				}
			}
		}

		if !isResolved {
			if (ascCounter > 0 && descCounter > 0) || equalsCounter > 0 {
				unsavedCounter++
			}
		}
	}

	var saveCounter = len(lines) - unsavedCounter
	fmt.Printf("%d", saveCounter)
}

func isSafe(report []int) bool {
	var ascCounter = 0
	var descCounter = 0
	var equalsCounter = 0
	for i := 0; i < len(report)-1; i++ {
		var diff = report[i+1] - report[i]

		if abs(diff) >= 4 {
			return false
		}

		if diff > 0 {
			ascCounter++
		} else if diff < 0 {
			descCounter++
		} else {
			equalsCounter++
		}
	}
	if (ascCounter > 0 && descCounter > 0) || equalsCounter > 0 {
		return false
	}
	return true
}

func solutionPart02(lines []string) {
	var saveCounter = 0
	for _, line := range lines {
		var report = SplitInts(line, " ")
		var reportSize = len(report)

		if isSafe(report) {
			saveCounter++
		} else {
			for i := 0; i < reportSize; i++ {
				var newReport = make([]int, 0)
				for j := 0; j < reportSize; j++ {
					if j != i {
						newReport = append(newReport, report[j])
					}
				}

				if isSafe(newReport) {
					saveCounter++
					break
				}
			}
		}
	}
	fmt.Printf("%d", saveCounter)
}

func abs(value int) int {
	return int(math.Abs(float64(value)))
}

// https://adventofcode.com/2024/day/2
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9")
	//RunAdventOfCodeWithFile(solutionPart01, "day_02/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9")
	RunAdventOfCodeWithFile(solutionPart02, "day_02/testcases/input-part-01.txt")
}
