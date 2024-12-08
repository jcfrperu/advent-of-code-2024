package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strconv"
)

type AntennaSet struct {
	List        []Node[string]
	AntennaType string
}

func solutionPart01(lines []string) {
	antennas, m := readInput(lines)
	antinodes := make(map[string]Node[string])
	for _, antenna := range antennas {
		for i := 0; i < len(antenna.List)-1; i++ {
			for j := i + 1; j < len(antenna.List); j++ {
				node01 := antenna.List[i]
				node02 := antenna.List[j]

				if Abs(node02.Row, node01.Row) == 0 {
					continue
				}

				if Abs(node02.Col, node01.Col) == 0 {
					continue
				}

				var antinode01 Node[string]
				var antinode02 Node[string]

				if float64(node02.Row-node01.Row)/float64(node02.Col-node01.Col) > 0 {
					antinode01 = Node[string]{
						Value: "#",
						Row:   node01.Row - Abs(node02.Row, node01.Row),
						Col:   node01.Col - Abs(node02.Col, node01.Col),
					}
					antinode02 = Node[string]{
						Value: "#",
						Row:   node02.Row + Abs(node02.Row, node01.Row),
						Col:   node02.Col + Abs(node02.Col, node01.Col),
					}
				} else {
					antinode01 = Node[string]{
						Value: "#",
						Row:   node01.Row - Abs(node02.Row, node01.Row),
						Col:   node01.Col + Abs(node02.Col, node01.Col),
					}
					antinode02 = Node[string]{
						Value: "#",
						Row:   node02.Row + Abs(node02.Row, node01.Row),
						Col:   node02.Col - Abs(node02.Col, node01.Col),
					}
				}

				if m.IsValid(antinode01.Row, antinode01.Col) {
					antinodes[buildKey(antinode01)] = antinode01
				}
				if m.IsValid(antinode02.Row, antinode02.Col) {
					antinodes[buildKey(antinode02)] = antinode02
				}
			}
		}
	}

	//printMatrix(m)
	//for _, antinode := range antinodes {
	//	fmt.Printf("%v\n", antinode)
	//	m[antinode.Row][antinode.Col].Value = antinode.Value
	//}
	//printMatrix(m)

	fmt.Printf("%d\n", len(antinodes))
}

func buildKey(node Node[string]) string {
	return strconv.Itoa(node.Row) + "|" + strconv.Itoa(node.Col)
}

func buildKeyWithAntenna(node Node[string], antennaType string) string {
	return strconv.Itoa(node.Row) + "|" + strconv.Itoa(node.Col) + "|" + antennaType
}

func printMatrix(m Matrix[string]) {
	fmt.Printf("\n")
	for i := 0; i < m.GetRowSize(); i++ {
		for j := 0; j < m.GetColSize(); j++ {
			fmt.Printf(m[i][j].Value)
		}
		fmt.Printf("\n")
	}
}

func readInput(lines []string) (map[string]AntennaSet, Matrix[string]) {
	antennas := make(map[string]AntennaSet)
	antennasType := make(map[string]string)

	m := BuildStringMatrix(lines)

	rows := m.GetRowSize()
	cols := m.GetColSize()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if m[r][c].Value != "." {
				antennasType[m[r][c].Value] = m[r][c].Value
			}
		}
	}

	for key := range antennasType {
		list := make([]Node[string], 0)
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				antennaType := m[r][c].Value
				if antennaType == key {
					list = append(list, m[r][c])
				}
			}
		}
		antennas[key] = AntennaSet{list, key}
	}
	return antennas, m
}

func solutionPart02(lines []string) {
	antennas, m := readInput(lines)
	antinodes := make(map[string]Node[string])
	for _, antenna := range antennas {
		for i := 0; i < len(antenna.List)-1; i++ {
			for j := i + 1; j < len(antenna.List); j++ {
				node01 := antenna.List[i]
				node02 := antenna.List[j]

				//if Abs(node02.Row, node01.Row) == 0 {
				//	continue
				//}
				//
				//if Abs(node02.Col, node01.Col) == 0 {
				//	continue
				//}

				var antinode01 Node[string]
				var antinode02 Node[string]

				if float64(node02.Row-node01.Row)/float64(node02.Col-node01.Col) > 0 {
					// first node
					antinode01 = Node[string]{
						Value: "#",
						Row:   node01.Row - Abs(node02.Row, node01.Row),
						Col:   node01.Col - Abs(node02.Col, node01.Col),
					}
					if m.IsValid(antinode01.Row, antinode01.Col) {
						antinodes[buildKeyWithAntenna(antinode01, antenna.AntennaType)] = antinode01
						// add # as many at top+left
						top, bottom := clone(antinode01), clone(node01)
						potential := Node[string]{
							Value: "#",
							Row:   top.Row - Abs(bottom.Row, top.Row),
							Col:   top.Col - Abs(bottom.Col, top.Col),
						}
						for m.IsValid(potential.Row, potential.Col) {
							antinodes[buildKeyWithAntenna(potential, antenna.AntennaType)] = potential
							top, bottom = clone(potential), clone(top)
							potential = Node[string]{
								Value: "#",
								Row:   top.Row - Abs(bottom.Row, top.Row),
								Col:   top.Col - Abs(bottom.Col, top.Col),
							}
						}
					}

					// second node
					antinode02 = Node[string]{
						Value: "#",
						Row:   node02.Row + Abs(node02.Row, node01.Row),
						Col:   node02.Col + Abs(node02.Col, node01.Col),
					}
					if m.IsValid(antinode02.Row, antinode02.Col) {
						antinodes[buildKeyWithAntenna(antinode02, antenna.AntennaType)] = antinode02
						// add # as many at bottom+right
						top, bottom := clone(node02), clone(antinode02)
						potential := Node[string]{
							Value: "#",
							Row:   bottom.Row + Abs(bottom.Row, top.Row),
							Col:   bottom.Col + Abs(bottom.Col, top.Col),
						}
						for m.IsValid(potential.Row, potential.Col) {
							antinodes[buildKeyWithAntenna(potential, antenna.AntennaType)] = potential
							top, bottom = clone(bottom), clone(potential)
							potential = Node[string]{
								Value: "#",
								Row:   bottom.Row + Abs(bottom.Row, top.Row),
								Col:   bottom.Col + Abs(bottom.Col, top.Col),
							}
						}
					}

				} else {
					// first node
					antinode01 = Node[string]{
						Value: "#",
						Row:   node01.Row - Abs(node02.Row, node01.Row),
						Col:   node01.Col + Abs(node02.Col, node01.Col),
					}
					if m.IsValid(antinode01.Row, antinode01.Col) {
						antinodes[buildKeyWithAntenna(antinode01, antenna.AntennaType)] = antinode01
						// at many at right+top
						top, bottom := clone(antinode01), clone(node01)
						potential := Node[string]{
							Value: "#",
							Row:   top.Row - Abs(bottom.Row, top.Row),
							Col:   top.Col + Abs(bottom.Col, top.Col),
						}
						for m.IsValid(potential.Row, potential.Col) {
							antinodes[buildKeyWithAntenna(potential, antenna.AntennaType)] = potential
							top, bottom = clone(potential), clone(top)
							potential = Node[string]{
								Value: "#",
								Row:   top.Row - Abs(bottom.Row, top.Row),
								Col:   top.Col + Abs(bottom.Col, top.Col),
							}
						}
					}
					// second node
					antinode02 = Node[string]{
						Value: "#",
						Row:   node02.Row + Abs(node02.Row, node01.Row),
						Col:   node02.Col - Abs(node02.Col, node01.Col),
					}
					if m.IsValid(antinode02.Row, antinode02.Col) {
						antinodes[buildKeyWithAntenna(antinode02, antenna.AntennaType)] = antinode02
						// at many at left+bottom
						top, bottom := clone(node02), clone(antinode02)
						potential := Node[string]{
							Value: "#",
							Row:   bottom.Row + Abs(bottom.Row, top.Row),
							Col:   bottom.Col - Abs(bottom.Col, top.Col),
						}
						for m.IsValid(potential.Row, potential.Col) {
							antinodes[buildKeyWithAntenna(potential, antenna.AntennaType)] = potential
							top, bottom = clone(bottom), clone(potential)
							potential = Node[string]{
								Value: "#",
								Row:   bottom.Row + Abs(bottom.Row, top.Row),
								Col:   bottom.Col - Abs(bottom.Col, top.Col),
							}
						}
					}

				}
			}
		}
	}

	printMatrix(m)
	for _, antinode := range antinodes {
		//fmt.Printf("%v\n", antinode)
		m[antinode.Row][antinode.Col].Value = antinode.Value
	}
	printMatrix(m)

	fmt.Printf("%d\n", len(antinodes))
}

func clone(node Node[string]) Node[string] {
	return Node[string]{
		Value: node.Value,
		Row:   node.Row,
		Col:   node.Col,
	}
}

// https://adventofcode.com/2024/day/8
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............")
	//RunAdventOfCodeWithFile(solutionPart01, "day_08/TESTCASES/input-part-01.txt")

	// part 02: using string or input file
	RunAdventOfCodeWithString(solutionPart02, "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............")
	//RunAdventOfCodeWithFile(solutionPart02, "day_08/testcases/input-part-02.txt")
}
