package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

type AntennaSet struct {
	List        []Cell[string]
	AntennaType string
}

func solutionPart01(lines []string) {
	antennas, m := readInput(lines)
	antinodes := make(map[string]Cell[string])
	for _, antenna := range antennas {
		for i := 0; i < len(antenna.List)-1; i++ {
			for j := i + 1; j < len(antenna.List); j++ {
				node01 := antenna.List[i]
				node02 := antenna.List[j]

				var antinode01 Cell[string]
				var antinode02 Cell[string]

				if float64(node02.Row-node01.Row)/float64(node02.Col-node01.Col) > 0 {
					antinode01 = Cell[string]{
						Value: "#",
						Row:   node01.Row - Abs(node02.Row, node01.Row),
						Col:   node01.Col - Abs(node02.Col, node01.Col),
					}
					antinode02 = Cell[string]{
						Value: "#",
						Row:   node02.Row + Abs(node02.Row, node01.Row),
						Col:   node02.Col + Abs(node02.Col, node01.Col),
					}
				} else {
					antinode01 = Cell[string]{
						Value: "#",
						Row:   node01.Row - Abs(node02.Row, node01.Row),
						Col:   node01.Col + Abs(node02.Col, node01.Col),
					}
					antinode02 = Cell[string]{
						Value: "#",
						Row:   node02.Row + Abs(node02.Row, node01.Row),
						Col:   node02.Col - Abs(node02.Col, node01.Col),
					}
				}

				if m.IsValid(antinode01) {
					antinodes[buildKey(antinode01)] = antinode01
				}

				if m.IsValid(antinode02) {
					antinodes[buildKey(antinode02)] = antinode02
				}
			}
		}
	}

	PrintStrMatrix(m)
	for key, antinode := range antinodes {
		fmt.Printf("%v->%v\n", key, antinode)
		m[antinode.Row][antinode.Col].Value = antinode.Value
	}
	PrintStrMatrix(m)

	fmt.Printf("%d\n", len(antinodes))
}

func buildKey(node Cell[string]) string {
	return FormatInt(node.Row) + "|" + FormatInt(node.Col)
}

func readInput(lines []string) (map[string]AntennaSet, Matrix[string]) {
	antennas := make(map[string]AntennaSet)
	antennasType := make(map[string]string)

	m := BuildStrMatrix(lines)

	rows := m.Rows()
	cols := m.Cols()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if m[r][c].Value != "." {
				antennasType[m[r][c].Value] = m[r][c].Value
			}
		}
	}

	for key := range antennasType {
		list := make([]Cell[string], 0)
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

// NOT WORKING
func solutionPart02(lines []string) {
	antennas, m := readInput(lines)
	antinodes := make(map[string]Cell[string])

	for _, antenna := range antennas {
		for i := 0; i < len(antenna.List)-1; i++ {
			for j := i + 1; j < len(antenna.List); j++ {
				node01 := antenna.List[i]
				node02 := antenna.List[j]

				var antinode01 Cell[string]
				var antinode02 Cell[string]

				if float64(node02.Row-node01.Row)/float64(node02.Col-node01.Col) > 0 {
					// first node
					antinode01 = Cell[string]{
						Value: "#",
						Row:   node01.Row - Abs(node02.Row, node01.Row),
						Col:   node01.Col - Abs(node02.Col, node01.Col),
					}
					if m.IsValid(antinode01) {
						antinodes[buildKey(antinode01)] = antinode01
						// add # as many at top+left
						top, bottom := clone(antinode01), clone(node01)
						potential := Cell[string]{
							Value: "#",
							Row:   top.Row - Abs(bottom.Row, top.Row),
							Col:   top.Col - Abs(bottom.Col, top.Col),
						}
						for m.IsValid(potential) {
							antinodes[buildKey(potential)] = potential
							top, bottom = clone(potential), clone(top)
							potential = Cell[string]{
								Value: "#",
								Row:   top.Row - Abs(bottom.Row, top.Row),
								Col:   top.Col - Abs(bottom.Col, top.Col),
							}
						}
					}

					// second node
					antinode02 = Cell[string]{
						Value: "#",
						Row:   node02.Row + Abs(node02.Row, node01.Row),
						Col:   node02.Col + Abs(node02.Col, node01.Col),
					}
					if m.IsValid(antinode02) {
						antinodes[buildKey(antinode02)] = antinode02
						// add # as many at bottom+right
						top, bottom := clone(node02), clone(antinode02)
						potential := Cell[string]{
							Value: "#",
							Row:   bottom.Row + Abs(bottom.Row, top.Row),
							Col:   bottom.Col + Abs(bottom.Col, top.Col),
						}
						for m.IsValid(potential) {
							antinodes[buildKey(potential)] = potential
							top, bottom = clone(bottom), clone(potential)
							potential = Cell[string]{
								Value: "#",
								Row:   bottom.Row + Abs(bottom.Row, top.Row),
								Col:   bottom.Col + Abs(bottom.Col, top.Col),
							}
						}
					}

				} else {
					// first node
					antinode01 = Cell[string]{
						Value: "#",
						Row:   node01.Row - Abs(node02.Row, node01.Row),
						Col:   node01.Col + Abs(node02.Col, node01.Col),
					}
					if m.IsValid(antinode01) {
						antinodes[buildKey(antinode01)] = antinode01
						// at many at right+top
						top, bottom := clone(antinode01), clone(node01)
						potential := Cell[string]{
							Value: "#",
							Row:   top.Row - Abs(bottom.Row, top.Row),
							Col:   top.Col + Abs(bottom.Col, top.Col),
						}
						for m.IsValid(potential) {
							antinodes[buildKey(potential)] = potential
							top, bottom = clone(potential), clone(top)
							potential = Cell[string]{
								Value: "#",
								Row:   top.Row - Abs(bottom.Row, top.Row),
								Col:   top.Col + Abs(bottom.Col, top.Col),
							}
						}
					}
					// second node
					antinode02 = Cell[string]{
						Value: "#",
						Row:   node02.Row + Abs(node02.Row, node01.Row),
						Col:   node02.Col - Abs(node02.Col, node01.Col),
					}
					if m.IsValid(antinode02) {
						antinodes[buildKey(antinode02)] = antinode02
						// at many at left+bottom
						top, bottom := clone(node02), clone(antinode02)
						potential := Cell[string]{
							Value: "#",
							Row:   bottom.Row + Abs(bottom.Row, top.Row),
							Col:   bottom.Col - Abs(bottom.Col, top.Col),
						}
						for m.IsValid(potential) {
							antinodes[buildKey(potential)] = potential
							top, bottom = clone(bottom), clone(potential)
							potential = Cell[string]{
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

	// counting antennas as well as antinodes
	for _, antenna := range antennas {
		for i := range antenna.List {
			//antennasCount += len(antenna.List)
			potential := Cell[string]{
				Value: antenna.List[i].Value,
				Row:   antenna.List[i].Row,
				Col:   antenna.List[i].Col,
			}
			antinodes[buildKey(potential)] = potential
		}
	}
	fmt.Printf("%d\n", len(antinodes))
}

func clone(node Cell[string]) Cell[string] {
	return Cell[string]{
		Value: node.Value,
		Row:   node.Row,
		Col:   node.Col,
	}
}

// https://adventofcode.com/2024/day/8
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "............\n............\n............\n............\n......A.....\n............\n............\n............\n............\n............\n............")
	//RunAdventOfCodeWithFile(solutionPart01, "day_08/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............")
	RunAdventOfCodeWithFile(solutionPart02, "day_08/testcases/input-part-02.txt")
}
