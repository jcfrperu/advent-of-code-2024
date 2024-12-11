package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solutionPart01(lines []string) {
	m := BuildStrMatrix(lines)
	currentNode := findStartNode(m)
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
	m := BuildStrMatrix(lines)
	node := findStartNode(m)
	loopsCounter := 0
	// super force brute :P
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
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

func getPosKey(node Node[string]) string {
	return FormatInt(node.Row) + "|" + FormatInt(node.Col)
}

func getPosAndDirKey(node Node[string]) string {
	return FormatInt(node.Row) + "|" + FormatInt(node.Col) + "|" + node.Value
}

func turn90(node Node[string]) Node[string] {
	if node.Value == "^" {
		return node.Update(">")
	}
	if node.Value == ">" {
		return node.Update("v")
	}
	if node.Value == "v" {
		return node.Update("<")
	}
	if node.Value == "<" {
		return node.Update("^")
	}
	return node
}

func peek(node Node[string], m Matrix[string]) Node[string] {
	if node.Value == "^" {
		return m.GetUp(node)
	}
	if node.Value == ">" {
		return m.GetRight(node)

	}
	if node.Value == "v" {
		return m.GetDown(node)

	}
	if node.Value == "<" {
		return m.GetLeft(node)
	}
	return node
}

func move(node Node[string], m Matrix[string]) Node[string] {
	if node.Value == "^" {
		return m.GetUp(node).Update(node.Value)
	}
	if node.Value == ">" {
		return m.GetRight(node).Update(node.Value)

	}
	if node.Value == "v" {
		return m.GetDown(node).Update(node.Value)

	}
	if node.Value == "<" {
		return m.GetLeft(node).Update(node.Value)
	}
	fmt.Printf("invalid move")
	return node
}

func findStartNode(m Matrix[string]) Node[string] {
	rows := m.Rows()
	cols := m.Cols()

	startPoint := Node[string]{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m[i][j].Value == "^" {
				startPoint = Node[string]{
					Value: m[i][j].Value,
					Row:   i,
					Col:   j,
				}
			}
		}
	}
	return startPoint
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
