package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

type Data struct {
	buttonARow int
	buttonACol int
	buttonBRow int
	buttonBCol int
	prizeRow   int
	prizeCol   int
}

func solutionPart01(lines []string) {
	nodes := readInput(lines)
	for _, node := range nodes {
		fmt.Printf("%v\n", node)
	}

	count := 0
	for _, node := range nodes {
		count += getTokens(node)
	}

	fmt.Printf("%d", count)
}

func getTokens(node Data) int {

	possibleValues := createPermutations(1, 100)
	permutations := Permute(possibleValues, 2, false) // groups of 2

	//TODO: consider divisibility criteria
	for _, permutation := range permutations {
		times01 := permutation[0]
		times02 := permutation[1]

		// conditions
		if node.prizeRow == times01*node.buttonARow+times02*node.buttonBRow &&
			node.prizeCol == times01*node.buttonACol+times02*node.buttonBCol {
			return times01*3 + times02 // toke
		}
	}

	// is not solution
	return 0
}

func createPermutations(start int, end int) []int {
	permutations := make([]int, 0)
	for i := start; i <= end; i++ {
		permutations = append(permutations, i)
	}
	return permutations
}

func readInput(lines []string) []Data {
	nodes := make([]Data, 0)

	node := Data{}
	for _, line := range lines {
		if StartsWith(line, "Button A") {
			_, err := fmt.Sscanf(line, "Button A: X+%d, Y+%d", &node.buttonARow, &node.buttonACol)
			CheckError(err)
		}
		if StartsWith(line, "Button B") {
			_, err := fmt.Sscanf(line, "Button B: X+%d, Y+%d", &node.buttonBRow, &node.buttonBCol)
			CheckError(err)
		}
		if StartsWith(line, "Prize") {
			_, err := fmt.Sscanf(line, "Prize: X=%d, Y=%d", &node.prizeRow, &node.prizeCol)
			CheckError(err)
		}
		if IsBlank(line) {
			nodes = append(nodes, node)
			node = Data{}
		}
	}

	return nodes
}

func solutionPart02(lines []string) {
	count := 0
	fmt.Printf("%d", count)
}

// https://adventofcode.com/2024/day/1
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400\n\nButton A: X+26, Y+66\nButton B: X+67, Y+21\nPrize: X=12748, Y=12176\n\nButton A: X+17, Y+86\nButton B: X+84, Y+37\nPrize: X=7870, Y=6450\n\nButton A: X+69, Y+23\nButton B: X+27, Y+71\nPrize: X=18641, Y=10279")
	RunAdventOfCodeWithFile(solutionPart01, "day_13/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400\n\nButton A: X+26, Y+66\nButton B: X+67, Y+21\nPrize: X=12748, Y=12176\n\nButton A: X+17, Y+86\nButton B: X+84, Y+37\nPrize: X=7870, Y=6450\n\nButton A: X+69, Y+23\nButton B: X+27, Y+71\nPrize: X=18641, Y=10279")
	//RunAdventOfCodeWithFile(solutionPart02, "day_13/testcases/input-part-02.txt")
}
