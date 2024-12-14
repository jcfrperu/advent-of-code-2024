package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solutionPart01(lines []string) {
	m := BuildStrMatrix(lines)

	pattern := "XMAS"

	count := 0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			if m[r][c].Value == "X" {
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}
						isMatch := true
						for i := 0; i < 4; i++ {
							r2 := r + dr*i
							c2 := c + dc*i
							matches := m.IsValidAt(r2, c2) && m[r2][c2].Value == string(pattern[i])
							if !matches {
								isMatch = false
								break
							}
						}
						if isMatch {
							count++
						}
					}
				}
			}
		}
	}

	fmt.Printf("%d", count)
}

func solutionPart02(lines []string) {
	m := BuildStrMatrix(lines)

	patterns := [4]string{"MSSM", "SSMM", "MMSS", "SMMS"}

	count := 0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			if m[r][c].IsValid() && m[r][c].Value == "A" {
				diag := m.GetDiagonalNeighborsAt(r, c)
				for _, pattern := range patterns {
					isMatch := true
					for i := 0; i < 4; i++ {
						matches := diag[i].IsValid() && diag[i].Value == string(pattern[i])
						if !matches {
							isMatch = false
						}
					}
					if isMatch {
						count++
					}
				}
			}
		}
	}

	fmt.Printf("%d", count)
}

// https://adventofcode.com/2024/day/4
func main() {
	// part 01: using string or input fileÑ/
	//RunAdventOfCodeWithString(solutionPart01, "....XXMAS.\n.SAMXMS...\n...S..A...\n..A.A.MS.X\nXMASAMX.MM\nX.....XA.A\nS.S.S.S.SS\n.A.A.A.A.A\n..M.M.M.MM\n.X.X.XMASX")
	//RunAdventOfCodeWithFile(solutionPart01, "day_04/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, ".M.S......\n..A..MSMS.\n.M.S.MAA..\n..A.ASMSM.\n.M.S.M....\n..........\nS.S.S.S.S.\n.A.A.A.A..\nM.M.M.M.M.\n..........")
	RunAdventOfCodeWithFile(solutionPart02, "day_04/testcases/input-part-02.txt")
}
