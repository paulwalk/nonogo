package pkg

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
)

const UNKNOWN = int(0)
const SPACE = int(1)
const BLOCK = int(2)

type Puzzle struct {
	Title          string  `yaml:"title"`
	Author         string  `yaml:"author"`
	RowClueData    [][]int `yaml:"rows"`
	ColClueData    [][]int `yaml:"columns"`
	Grid           [][]int `yaml:"-"`
	ColCount       int     `yaml:"-"`
	RowCount       int     `yaml:"-"`
	DisplayPadding int
}

func NewPuzzle(filePath string) (Puzzle, error) {
	var err error
	puzzle := Puzzle{}
	puzzleData, err := os.ReadFile(filePath)
	if err != nil {
		return puzzle, err
	}
	err = yaml.Unmarshal(puzzleData, &puzzle)
	if err != nil {
		return puzzle, err
	}
	puzzle.RowCount = len(puzzle.RowClueData)
	puzzle.ColCount = len(puzzle.ColClueData)
	puzzle.Grid = make([][]int, puzzle.RowCount)
	for x := 0; x < puzzle.RowCount; x++ {
		puzzle.Grid[x] = make([]int, puzzle.ColCount)
		for y := 0; y < puzzle.ColCount; y++ {
			puzzle.Grid[x][y] = UNKNOWN
		}
	}
	largestColClueNum := 0
	for _, clue := range puzzle.ColClueData {
		for _, clueNum := range clue {
			if clueNum > largestColClueNum {
				largestColClueNum = clueNum
			}
		}
	}
	numberOfDigits := len(strconv.Itoa(largestColClueNum))
	puzzle.DisplayPadding = -(numberOfDigits + 1)
	return puzzle, err
}

func (puzzle *Puzzle) Dump() string {
	output := color.YellowString(puzzle.Title)
	output += "\n"
	for j := 0; j < puzzle.ColCount; j++ {
		output += color.YellowString("%*v", puzzle.DisplayPadding, strconv.Itoa(j))
	}
	output += "\n"
	for i := 0; i < puzzle.maxColClueLength(); i++ {
		for _, clue := range puzzle.ColClueData {
			if len(clue) >= (i + 1) {
				output += fmt.Sprintf("%*v", puzzle.DisplayPadding, strconv.Itoa(clue[i]))
			} else {
				output += fmt.Sprintf("%*v", puzzle.DisplayPadding, " ")
			}
		}
		output += "\n"
	}
	for i := 0; i < puzzle.RowCount; i++ {
		output += displayLine(NewLine(puzzle, ROW, i), true)
		output += "\n"
	}
	return output
}

func (puzzle *Puzzle) maxColClueLength() int {
	maxLength := 0
	for _, clue := range puzzle.ColClueData {
		if len(clue) > maxLength {
			maxLength = len(clue)
		}
	}
	return maxLength
}
