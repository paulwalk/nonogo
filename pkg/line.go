package pkg

import (
	"strconv"
)

const ROW = int(0)
const COL = int(1)

type Line struct {
	Axis, Index, Length int
	PotentialSolutions  [][]int
	Puzzle              *Puzzle
}

func NewLine(puzzle *Puzzle, axis, index int) *Line {
	line := Line{}
	line.Axis = axis
	line.Index = index
	line.Puzzle = puzzle
	if axis == ROW {
		line.Length = puzzle.ColCount
	} else {
		line.Length = puzzle.RowCount
	}
	line.PotentialSolutions = populatePotentialSolutions(line.Clue(), line.Length)
	return &line
}

func (line *Line) Label() string {
	if line.Axis == ROW {
		return "ROW " + strconv.Itoa((line.Index)+1)
	} else {
		return "COL " + strconv.Itoa((line.Index)+1)
	}
}

func (line *Line) Cells() []*int {
	var cells []*int
	if line.Axis == ROW {
		cells = make([]*int, 0)
		for i := 0; i < line.Puzzle.ColCount; i++ {
			cells = append(cells, &line.Puzzle.Grid[line.Index][i])
		}
	} else {
		cells = make([]*int, 0)
		for i := 0; i < line.Puzzle.RowCount; i++ {
			cells = append(cells, &line.Puzzle.Grid[i][line.Index])
		}
	}
	return cells
}

func (line *Line) Clue() []int {
	if line.Axis == ROW {
		return line.Puzzle.RowClueData[line.Index]
	} else {
		return line.Puzzle.ColClueData[line.Index]
	}
}

func (line *Line) setCells(newCells []int) {
	for i, cell := range newCells {
		line.setCell(i, cell)
	}
}

func (line *Line) setCell(pos, value int) {
	if *line.Cells()[pos] != value {
		*line.Cells()[pos] = value
	}
}

func (line *Line) processIfSolved() bool {
	if len(line.PotentialSolutions) == 1 {
		line.setCells(line.PotentialSolutions[0])
		return true
	} else {
		return false
	}

}

func populatePotentialSolutions(clues []int, length int) [][]int {
	if len(clues) == 0 {
		solutions := make([][]int, 0)
		return append(solutions, getCellSlice(SPACE, length))
	} else {
		starts := length - clues[0]
		if len(clues) == 1 {
			solutions := make([][]int, 0)
			for i := 0; i <= starts; i++ {
				solution := make([]int, 0)
				solution = append(solution, getCellSlice(SPACE, i)...)
				solution = append(solution, getCellSlice(BLOCK, clues[0])...)
				solution = append(solution, getCellSlice(SPACE, starts-i)...)
				solutions = append(solutions, solution)
			}
			return solutions
		} else {
			solutions := make([][]int, 0)
			for i := 0; i < length-clues[0]; i++ {
				for _, j := range populatePotentialSolutions(clues[1:], starts-i-1) {
					solution := make([]int, 0)
					solution = append(solution, getCellSlice(SPACE, i)...)
					solution = append(solution, getCellSlice(BLOCK, clues[0])...)
					solution = append(solution, SPACE)
					solution = append(solution, j...)
					solutions = append(solutions, solution)
				}
			}
			return solutions
		}
	}
}

func (line *Line) dumpPotentialSolutions() string {
	outputLine := ""
	for _, solution := range line.PotentialSolutions {
		for _, cell := range solution {
			outputLine += displayCell(cell, line.Puzzle.DisplayPadding)
		}
		outputLine += "\n"
	}
	return outputLine
}

func (line *Line) removePotentialSolution(solutionIndex int) {
	line.PotentialSolutions = append(line.PotentialSolutions[:solutionIndex], line.PotentialSolutions[solutionIndex+1:]...)
}
