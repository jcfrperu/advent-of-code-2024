package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

type DataBlock struct {
	size   int
	index  int
	starts int
	ends   int
}

type FreeBlock struct {
	size   int
	starts int
	ends   int
}

func solutionPart01(lines []string) {
	result := readInput(lines)

	fmt.Printf("%s\n", result)

	left := 0
	right := len(result) - 1

	for left < right {
		for result[left] != "." {
			left++
		}
		for result[right] == "." {
			right--
		}

		// swap
		result[left] = result[right]
		result[right] = "."

		left++
		right--
	}

	sum := 0
	for i, value := range result {
		if value == "." {
			continue
		}
		sum += i * ParseInt(value)
	}

	fmt.Printf("%d\n", sum)
}

func readInput(lines []string) []string {
	line := lines[0]

	result := make([]string, 0)
	index := 0
	for i := 0; i < len(line); i++ {
		size := ParseInt(string(line[i]))
		if i%2 == 0 {
			count := size
			for count > 0 {
				result = append(result, FormatInt(index))
				count--
			}
			index++
		} else {
			count := size
			for count > 0 {
				result = append(result, ".")
				count--
			}
		}
	}
	return result
}

func solutionPart02(lines []string) {
	count := 0
	fmt.Printf("%d", count)
}

// https://adventofcode.com/2024/day/09
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "2333133121414131402")
	RunAdventOfCodeWithFile(solutionPart01, "day_09/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "")
	//RunAdventOfCodeWithFile(solutionPart02, "day_09/testcases/input-part-02.txt")
}
