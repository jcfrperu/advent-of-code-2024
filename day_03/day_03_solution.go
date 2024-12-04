package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"regexp"
	"strings"
)

func solutionPart01(lines []string) {
	var sum = 0
	for _, line := range lines {
		var re = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		var matches = re.FindAllString(line, -1)
		var clean string
		for _, match := range matches {
			clean = strings.ReplaceAll(match, "(", "")
			clean = strings.ReplaceAll(clean, ")", "")
			clean = strings.ReplaceAll(clean, ",", " ")
			clean = strings.ReplaceAll(clean, "mul", "")
			var numbers = SplitInts(clean, " ")
			sum += numbers[0] * numbers[1]
		}
	}
	fmt.Printf("%d", sum)
}

func findProduct(line string) int {
	// PRECONDITION: line must not contains `don't()`
	var sum = 0
	var re = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	var matches = re.FindAllString(line, -1)
	var clean string
	for _, match := range matches {
		clean = strings.ReplaceAll(match, "(", "")
		clean = strings.ReplaceAll(clean, ")", "")
		clean = strings.ReplaceAll(clean, ",", " ")
		clean = strings.ReplaceAll(clean, "mul", "")
		var numbers = SplitInts(clean, " ")
		sum += numbers[0] * numbers[1]
	}
	return sum
}

func solutionPart02(lines []string) {
	var sum = 0
	var on = true
	for _, line := range lines {
		var re = regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
		var matches = re.FindAllString(line, -1)
		var clean string

		for _, match := range matches {
			if match == "do()" {
				on = true
			} else if match == "don't()" {
				on = false
			} else {
				if on {
					clean = strings.ReplaceAll(match, "(", "")
					clean = strings.ReplaceAll(clean, ")", "")
					clean = strings.ReplaceAll(clean, ",", " ")
					clean = strings.ReplaceAll(clean, "mul", "")
					var numbers = SplitInts(clean, " ")
					sum += numbers[0] * numbers[1]
				}
			}
		}
	}
	fmt.Printf("%d", sum)
}

// https://adventofcode.com/2024/day/3
func main() {
	// part 01: using string or input fileÑ/
	//RunAdventOfCodeWithString(solutionPart01, "mul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	//RunAdventOfCodeWithFile(solutionPart01, "day_03/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "mul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	RunAdventOfCodeWithFile(solutionPart02, "day_03/testcases/input-part-02.txt")
}
