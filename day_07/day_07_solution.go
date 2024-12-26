package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

type Data struct {
	total   int64
	numbers []int64
}

func readInput(lines []string) []Data {
	nodes := make([]Data, 0)
	for _, line := range lines {
		total := ParseLong(SplitGetAt(line, ":", 0))
		numbers := SplitLongs(SplitGetAt(line, ":", 1), " ")
		nodes = append(nodes, Data{total, numbers})
	}
	return nodes
}

func solutionPart01(lines []string) {
	total := int64(0)
	nodes := readInput(lines)

	for _, node := range nodes {
		if isSolution(node.total, node.numbers) {
			total += node.total
		}
	}

	fmt.Printf("%v", total)
}

func isSolution(total int64, nums []int64) bool {
	groups := len(nums) - 1
	operators := []string{"+", "*"}
	permutations := Permute(operators, groups, true)

	for _, permutation := range permutations {
		sum := int64(0)

		if permutation[0] == "*" {
			sum += nums[0] * nums[1]
		} else if permutation[0] == "+" {
			sum += nums[0] + nums[1]
		}

		for i := 2; i < len(nums); i++ {
			operation := permutation[i-1]
			if operation == "*" {
				sum *= nums[i]
			} else if operation == "+" {
				sum += nums[i]
			}
		}

		if sum == total {
			return true
		}
	}
	return false
}

func solutionPart02(lines []string) {
	total := int64(0)
	nodes := readInput(lines)

	for _, node := range nodes {
		if isSolution2(node.total, node.numbers) {
			total += node.total
		}
	}

	fmt.Printf("%v", total)
}

func isSolution2(total int64, nums []int64) bool {
	groups := len(nums) - 1
	operators := []string{"+", "*", "||"}
	permutations := Permute(operators, groups, true)

	for _, permutation := range permutations {
		sum := int64(0)

		if permutation[0] == "*" {
			sum += nums[0] * nums[1]
		} else if permutation[0] == "+" {
			sum += nums[0] + nums[1]
		} else if permutation[0] == "||" {
			sum += ParseLong(FormatInt(nums[0]) + FormatInt(nums[1]))
		}

		for i := 2; i < len(nums); i++ {
			operation := permutation[i-1]
			if operation == "*" {
				sum = sum * nums[i]
			} else if operation == "+" {
				sum = sum + nums[i]
			} else if operation == "||" {
				sum = ParseLong(FormatInt(sum) + FormatInt(nums[i]))
			}
		}

		if sum == total {
			return true
		}
	}
	return false
}

// https://adventofcode.com/2024/day/7
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20")
	//RunAdventOfCodeWithFile(solutionPart01, "day_07/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20")
	RunAdventOfCodeWithFile(solutionPart02, "day_07/testcases/input-part-02.txt")
}
