package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"regexp"
	"strings"
)

var EMPTY = ""

func solutionPart01(lines []string) {
	line := joinStrings(lines, EMPTY)
	product := findProduct(line)
	fmt.Printf("%d", product)
}

func solutionPart02(lines []string) {
	line := joinStrings(lines, EMPTY)
	cleanLine := removeBlocks(line, "don't()", "do()")
	product := findProduct(cleanLine)
	fmt.Printf("%d", product)
}

func findProduct(line string) int {
	sum := 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(line, -1)
	clean := EMPTY
	for _, match := range matches {
		clean = strings.ReplaceAll(match, "mul(", EMPTY)
		clean = strings.ReplaceAll(clean, ")", EMPTY)
		numbers := SplitInts(clean, ",")
		sum += numbers[0] * numbers[1]
	}
	return sum
}

func removeBlocks(input string, startPattern string, endPattern string) string {
	line := input
	startIndex := indexAt(line, startPattern, 0)
	for startIndex >= 0 {
		endIndex := indexAt(line, endPattern, startIndex)
		if endIndex >= 0 && endIndex >= startIndex {
			line = line[:startIndex] + line[endIndex:]
		} else {
			line = line[0:startIndex]
		}
		startIndex = indexAt(line, startPattern, 0)
	}
	return line
}

func joinStrings(lines []string, separator string) string {
	line := EMPTY
	for i, item := range lines {
		if i == 0 {
			line = item
		} else {
			line = line + separator + item
		}
	}
	return line
}

func indexAt(s, substr string, fromIndex int) int {
	if fromIndex >= 0 && fromIndex < len(s) {
		index := strings.Index(s[fromIndex:], substr)
		if index >= 0 {
			return index + fromIndex
		}
	}
	return -1
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
