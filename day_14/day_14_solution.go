package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strings"
)

type Robot struct {
	PosRow int
	PosCol int
	VelRow int
	VelCol int
}

func updateRobotPosition(r *Robot, rows int, cols int) {
	row := (r.PosRow + r.VelRow) % rows
	col := (r.PosCol + r.VelCol) % cols
	if row < 0 {
		row = rows + row
	}
	if col < 0 {
		col = cols + col
	}
	r.PosRow = row
	r.PosCol = col
}

func solutionPart01(lines []string) {
	robots, m := readInput(lines, 103, 101)

	rows := m.Rows()
	cols := m.Cols()

	iterNumber := 100
	for iter := 1; iter <= iterNumber; iter++ {
		for i := range robots {
			updateRobotPosition(&robots[i], rows, cols)
			m = buildMatrixFromRobots(robots, rows, cols)
		}
	}

	removeCrossItems(&m)
	factor := calculateSafetyFactor(&m)

	fmt.Printf("%d", factor)
}

func calculateSafetyFactor(m *Matrix[string]) int64 {
	rows := m.Rows()
	cols := m.Cols()
	halfRows := m.Rows() / 2
	halfCols := m.Cols() / 2

	factor01 := 0
	for i := 0; i < halfRows; i++ {
		for j := 0; j < halfCols; j++ {
			if m.GetAt(i, j) != "." {
				factor01 += ParseInt(m.GetAt(i, j))
			}
		}
	}

	factor02 := 0
	for i := 0; i < halfRows; i++ {
		for j := halfCols + 1; j < cols; j++ {
			if m.GetAt(i, j) != "." {
				factor02 += ParseInt(m.GetAt(i, j))
			}
		}
	}

	factor03 := 0
	for i := halfRows + 1; i < rows; i++ {
		for j := 0; j < halfCols; j++ {
			if IsInt(m.GetAt(i, j)) {
				factor03 += ParseInt(m.GetAt(i, j))
			}
		}
	}

	factor04 := 0
	for i := halfRows + 1; i < rows; i++ {
		for j := halfCols + 1; j < cols; j++ {
			if m.GetAt(i, j) != "." {
				factor04 += ParseInt(m.GetAt(i, j))
			}
		}
	}

	return int64(factor01) * int64(factor02) * int64(factor03) * int64(factor04)
}

func closenessMetric(robots []Robot) int {
	posMap := make(map[string]int)
	for i := range robots {
		posMap[buildKey(robots[i])]++
	}

	inMap := func(key string) bool {
		_, ok := posMap[key]
		return ok
	}

	closeness := 0
	for i := range robots {
		neighbours := getRobotNeighbours(robots[i])
		for _, neighbour := range neighbours {
			neighbourKey := buildKey(neighbour)
			if inMap(neighbourKey) {
				closeness += posMap[buildKey(neighbour)]
			}
		}
	}
	return closeness
}

func getRobotNeighbours(robot Robot) []Robot {
	neighbours := make([]Robot, 8)
	neighbours[0] = Robot{PosRow: robot.PosRow, PosCol: robot.PosCol - 1, VelRow: 0, VelCol: 0}
	neighbours[1] = Robot{PosRow: robot.PosRow, PosCol: robot.PosCol + 1, VelRow: 0, VelCol: 0}

	neighbours[2] = Robot{PosRow: robot.PosRow - 1, PosCol: robot.PosCol, VelRow: 0, VelCol: 0}
	neighbours[3] = Robot{PosRow: robot.PosRow + 1, PosCol: robot.PosCol, VelRow: 0, VelCol: 0}

	neighbours[4] = Robot{PosRow: robot.PosRow - 1, PosCol: robot.PosCol - 1, VelRow: 0, VelCol: 0}
	neighbours[5] = Robot{PosRow: robot.PosRow - 1, PosCol: robot.PosCol + 1, VelRow: 0, VelCol: 0}

	neighbours[6] = Robot{PosRow: robot.PosRow + 1, PosCol: robot.PosCol - 1, VelRow: 0, VelCol: 0}
	neighbours[7] = Robot{PosRow: robot.PosRow + 1, PosCol: robot.PosCol + 1, VelRow: 0, VelCol: 0}
	return neighbours
}

func removeCrossItems(m *Matrix[string]) {
	rows := m.Rows()
	cols := m.Cols()
	halfRows := rows / 2
	halfCols := cols / 2
	for r := 0; r < rows; r++ {
		m.SetAt(r, halfCols, " ")
	}
	for c := 0; c < cols; c++ {
		m.SetAt(halfRows, c, " ")
	}
}

func readInput(lines []string, rows int, cols int) ([]Robot, Matrix[string]) {
	robots := make([]Robot, len(lines))
	for i, line := range lines {
		leftRight := Split(line, " ")

		positions := Trim(strings.ReplaceAll(leftRight[0], "p=", ""))
		velocities := Trim(strings.ReplaceAll(leftRight[1], "v=", ""))

		robots[i] = Robot{
			PosRow: SplitGetIntAt(positions, ",", 1),
			PosCol: SplitGetIntAt(positions, ",", 0),
			VelRow: SplitGetIntAt(velocities, ",", 1),
			VelCol: SplitGetIntAt(velocities, ",", 0),
		}
	}
	m := buildMatrixFromRobots(robots, rows, cols)
	return robots, m
}

func buildMatrixFromRobots(robots []Robot, rows int, cols int) Matrix[string] {
	posMap := make(map[string]int)
	for i := range robots {
		posMap[buildKey(robots[i])]++
	}
	m := NewEmptyMatrix(".", rows, cols)
	for key, times := range posMap {
		r := SplitGetIntAt(key, ",", 0)
		c := SplitGetIntAt(key, ",", 1)
		m[r][c].Value = FormatInt(times)
	}
	return m
}

func buildKey(robot Robot) string {
	return fmt.Sprintf("%d,%d", robot.PosRow, robot.PosCol)
}

func NewEmptyMatrix[T any](value T, rows int, cols int) Matrix[T] {
	matrix := make([][]Cell[T], rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]Cell[T], cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = Cell[T]{
				Value: value,
				Row:   i,
				Col:   j,
			}
		}
	}
	return matrix
}

func solutionPart02(lines []string) {
	robots, m := readInput(lines, 103, 101)

	rows := m.Rows()
	cols := m.Cols()

	iterNumber := 10000
	maxCloseness := 0

	for iter := 1; iter <= iterNumber; iter++ {
		for i := range robots {
			updateRobotPosition(&robots[i], rows, cols)
		}
		closeness := closenessMetric(robots)
		if closeness > maxCloseness {
			fmt.Printf("iter: %d, closeness: %d, max-closeness: %d\n", iter, closeness, maxCloseness)
			m = buildMatrixFromRobots(robots, rows, cols)
			PrintStrMatrix(m)
			maxCloseness = closeness
		}
	}

	fmt.Printf("%d", iterNumber)
}

// https://adventofcode.com/2024/day/14
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "p=0,4 v=3,-3\np=6,3 v=-1,-3\np=10,3 v=-1,2\np=2,0 v=2,-1\np=0,0 v=1,3\np=3,0 v=-2,-2\np=7,6 v=-1,-3\np=3,0 v=-1,-2\np=9,3 v=2,3\np=7,3 v=-1,2\np=2,4 v=2,-3\np=9,5 v=-3,-3")
	//RunAdventOfCodeWithFile(solutionPart01, "day_14/testcases/input-part-01.txt")

	// part 02: using string or input file
	RunAdventOfCodeWithFile(solutionPart02, "day_14/testcases/input-part-02.txt")
}
