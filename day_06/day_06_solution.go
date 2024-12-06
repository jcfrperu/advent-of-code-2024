package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strconv"
)

func solutionPart01(lines []string) {
	m, currentNode := buildStringMatrix(lines)
	routes := make(map[string]int)

	routes[getPosKey(currentNode)]++

	// possible status: ^ > v <
	maxIter := 10000
	for currentNode.IsValid() && maxIter >= 0 {
		nextNode := peek(currentNode, m)
		if nextNode.Value == "#" || nextNode.Value == "0" {
			currentNode = turn90(currentNode)
		} else {
			currentNode = move(currentNode, m)
			routes[getPosKey(currentNode)]++
		}
		maxIter--
	}

	if maxIter < 0 {
		fmt.Println("ERROR: max iterations reach")
	}

	// removing last invalid node
	routesNumber := len(routes) - 1
	fmt.Printf("%d", routesNumber)
}

func solutionPart02(lines []string) {
	m, node := buildStringMatrix(lines)
	loopsCounter := 0
	// super force brute :P
	for i := 0; i < m.GetRowSize(); i++ {
		for j := 0; j < m.GetColSize(); j++ {
			if m[i][j].Value != "." {
				continue
			}
			m[i][j].Value = "0" // making walls in every single position xD

			if isLoop(m, node) {
				loopsCounter++
			}
			m[i][j].Value = "."
		}
	}
	fmt.Printf("%d", loopsCounter)
}

func isLoop(m Matrix[string], node Node[string]) bool {
	routes := make(map[string]int)
	routes[getPosKey(node)]++

	maxIter := 10000
	for node.IsValid() && maxIter >= 0 {
		nextNode := peek(node, m)
		if nextNode.Value == "#" || nextNode.Value == "0" {
			node = turn90(node)
		} else {
			node = move(node, m)
			routes[getPosAndDirKey(node)]++
			if routes[getPosAndDirKey(node)] > 1 {
				return true // return a known point in grid
			}
		}
		maxIter--
	}
	if maxIter < 0 {
		fmt.Println("ERROR: max iterations reach")
	}
	return false
}

func printMatrix(m Matrix[string]) {
	for i := 0; i < m.GetRowSize(); i++ {
		for j := 0; j < m.GetColSize(); j++ {
			fmt.Printf("%s", m[i][j].Value)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func getPosKey(node Node[string]) string {
	return strconv.Itoa(node.Row) + "|" + strconv.Itoa(node.Col)
}

func getPosAndDirKey(node Node[string]) string {
	return strconv.Itoa(node.Row) + "|" + strconv.Itoa(node.Col) + "|" + node.Value
}

func turn90(node Node[string]) Node[string] {
	if node.Value == "^" {
		node.Value = ">"
		return node
	}
	if node.Value == ">" {
		node.Value = "v"
		return node
	}
	if node.Value == "v" {
		node.Value = "<"
		return node
	}
	if node.Value == "<" {
		node.Value = "^"
		return node
	}
	return node
}

func peek(node Node[string], m Matrix[string]) Node[string] {
	if node.Value == "^" {
		return m.GetUp(node.Row, node.Col)
	}
	if node.Value == ">" {
		return m.GetRight(node.Row, node.Col)

	}
	if node.Value == "v" {
		return m.GetDown(node.Row, node.Col)

	}
	if node.Value == "<" {
		return m.GetLeft(node.Row, node.Col)
	}
	return node
}

func move(node Node[string], m Matrix[string]) Node[string] {
	if node.Value == "^" {
		n := m.GetUp(node.Row, node.Col)
		n.Value = node.Value
		return n
	}
	if node.Value == ">" {
		n := m.GetRight(node.Row, node.Col)
		n.Value = node.Value
		return n

	}
	if node.Value == "v" {
		n := m.GetDown(node.Row, node.Col)
		n.Value = node.Value
		return n

	}
	if node.Value == "<" {
		n := m.GetLeft(node.Row, node.Col)
		n.Value = node.Value
		return n
	}
	fmt.Printf("invalid move")
	return node
}

func buildStringMatrix(lines []string) (Matrix[string], Node[string]) {
	rowSize := len(lines)
	colSize := len(lines[0])

	startPoint := Node[string]{}

	matrix := make([][]Node[string], rowSize)
	for i := 0; i < rowSize; i++ {
		matrix[i] = make([]Node[string], colSize)
		for j := 0; j < colSize; j++ {
			value := string(lines[i][j])
			if value == "^" {
				startPoint = Node[string]{
					Value: value,
					Row:   i,
					Col:   j,
				}
			}
			matrix[i][j] = Node[string]{
				Value: value,
				Row:   i,
				Col:   j,
			}
		}
	}
	return matrix, startPoint
}

// https://adventofcode.com/2024/day/6
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...")
	//RunAdventOfCodeWithFile(solutionPart01, "day_06/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...")
	RunAdventOfCodeWithFile(solutionPart02, "day_06/testcases/input-part-02.txt")
}
