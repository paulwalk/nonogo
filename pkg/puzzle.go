package pkg

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
)

type Puzzle struct {
	Title                   string   `yaml:"title"`
	Author                  string   `yaml:"author"`
	RowClueData             [][]int  `yaml:"rows"`
	ColClueData             [][]int  `yaml:"columns"`
	Grid                    [][]Cell `yaml:"-"`
	ColCount                int      `yaml:"-"`
	RowCount                int      `yaml:"-"`
	PaddingForDisplay       int
	DisplayClues            bool
	DisplayRowAndColNumbers bool
}

func NewPuzzle(filePath string, displayClues, displayRowAndColNumbers bool) (Puzzle, error) {
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
	puzzle.Grid = make([][]Cell, puzzle.RowCount)
	puzzle.DisplayClues = displayClues
	puzzle.DisplayRowAndColNumbers = displayRowAndColNumbers
	for x := 0; x < puzzle.RowCount; x++ {
		puzzle.Grid[x] = make([]Cell, puzzle.ColCount)
		for y := 0; y < puzzle.ColCount; y++ {
			puzzle.Grid[x][y] = Unknown
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
	puzzle.PaddingForDisplay = -(numberOfDigits + 1)
	return puzzle, err
}

func (puzzle *Puzzle) Dump() string {
	output := color.YellowString(puzzle.Title)
	output += "\n"
	if puzzle.DisplayRowAndColNumbers {
		for j := 1; j <= puzzle.ColCount; j++ {
			output += color.YellowString("%*v", puzzle.PaddingForDisplay, strconv.Itoa(j))
		}
		output += "\n"
	}
	if puzzle.DisplayClues {
		for i := 0; i < puzzle.maxColClueLength(); i++ {
			for _, clue := range puzzle.ColClueData {
				if len(clue) > i {
					output += fmt.Sprintf("%*v", puzzle.PaddingForDisplay, strconv.Itoa(clue[i]))
				} else {
					output += fmt.Sprintf("%*v", puzzle.PaddingForDisplay, " ")
				}
			}
			output += "\n"
		}
	}
	for i := 0; i < puzzle.RowCount; i++ {
		line := NewLine(puzzle, Row, i)
		output += line.displayString(puzzle.DisplayClues, puzzle.DisplayRowAndColNumbers)
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
