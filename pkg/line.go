package pkg

import (
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type Line struct {
	Axis, Index, Length int
	PotentialSolutions  [][]Cell
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

func (line *Line) Cells() []*Cell {
	var cells []*Cell
	if line.Axis == ROW {
		cells = make([]*Cell, 0)
		for i := 0; i < line.Puzzle.ColCount; i++ {
			cells = append(cells, &line.Puzzle.Grid[line.Index][i])
		}
	} else {
		cells = make([]*Cell, 0)
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

func (line *Line) setCells(newCells []Cell) {
	for i, cell := range newCells {
		line.setCell(i, cell)
	}
}

func (line *Line) setCell(pos int, value Cell) {
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

func populatePotentialSolutions(clues []int, length int) [][]Cell {
	if len(clues) == 0 {
		solutions := make([][]Cell, 0)
		return append(solutions, getCellSlice(SPACE, length))
	} else {
		starts := length - clues[0]
		if len(clues) == 1 {
			solutions := make([][]Cell, 0)
			for i := 0; i <= starts; i++ {
				solution := make([]Cell, 0)
				solution = append(solution, getCellSlice(SPACE, i)...)
				solution = append(solution, getCellSlice(BLOCK, clues[0])...)
				solution = append(solution, getCellSlice(SPACE, starts-i)...)
				solutions = append(solutions, solution)
			}
			return solutions
		} else {
			solutions := make([][]Cell, 0)
			for i := 0; i < length-clues[0]; i++ {
				for _, j := range populatePotentialSolutions(clues[1:], starts-i-1) {
					solution := make([]Cell, 0)
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
			outputLine += cell.displayString(line.Puzzle.PaddingForDisplay)
		}
		outputLine += "\n"
	}
	return outputLine
}

func (line *Line) removePotentialSolution(solutionIndex int) {
	line.PotentialSolutions = append(line.PotentialSolutions[:solutionIndex], line.PotentialSolutions[solutionIndex+1:]...)
}

func (line *Line) displayString(displayClues, displayRowAndColNumbers bool) string {
	lineString := ""
	for _, cell := range line.Cells() {
		lineString += cell.displayString(line.Puzzle.PaddingForDisplay)
	}
	if displayClues {
		clueString := make([]string, 0)
		for _, clue := range line.Clue() {
			clueString = append(clueString, strconv.Itoa(clue))
		}
		lineString += " "
		lineString += strings.Join(clueString, ",")
	}
	if displayRowAndColNumbers {
		lineString += color.YellowString(" %v", line.Index+1)
	}
	return lineString
}
