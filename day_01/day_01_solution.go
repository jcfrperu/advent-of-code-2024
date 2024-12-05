package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"math"
	"slices"
)

func solutionPart01(lines []string) {
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		parts := Split(line, " ")
		left = append(left, ParseInt(parts[0]))
		right = append(right, ParseInt(parts[3]))
	}

	slices.Sort(left)
	slices.Sort(right)

	dif := float64(0)
	for i := 0; i < len(left); i++ {
		dif += math.Abs(float64(left[i] - right[i]))
	}

	fmt.Printf("%d", int(dif))
}

func solutionPart02(lines []string) {
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		parts := Split(line, " ")
		left = append(left, ParseInt(parts[0]))
		right = append(right, ParseInt(parts[3]))
	}

	rightFreq, _, _, _ := Freq(right, true)

	dif := 0
	for i := 0; i < len(left); i++ {
		dif += left[i] * rightFreq[left[i]]
	}

	fmt.Printf("%d", dif)
}

// https://adventofcode.com/2024/day/1
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "3   4\n4   3\n2   5\n1   3\n3   9\n3   3")
	//RunAdventOfCodeWithFile(solutionPart01, "day_01/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "3   4\n4   3\n2   5\n1   3\n3   9\n3   3")
	RunAdventOfCodeWithFile(solutionPart02, "day_01/testcases/input-part-02.txt")
}
