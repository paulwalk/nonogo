package pkg

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
)

const UNKNOWN = int(0)
const SPACE = int(1)
const BLOCK = int(2)

type PuzzleStruct struct {
	Title          string  `yaml:"title"`
	Author         string  `yaml:"author"`
	RowClueData    [][]int `yaml:"rows"`
	ColClueData    [][]int `yaml:"columns"`
	Grid           [][]int `yaml:"-"`
	ColCount       int     `yaml:"-"`
	RowCount       int     `yaml:"-"`
	DisplayPadding int
}

func (puzzle *PuzzleStruct) Initialise(filePath string) error {
	var err error
	puzzleData, err := ioutil.ReadFile(filePath)
	if err == nil {
		err = yaml.Unmarshal([]byte(puzzleData), &puzzle)
		if err == nil {
			puzzle.RowCount = len(puzzle.RowClueData)
			puzzle.ColCount = len(puzzle.ColClueData)
			puzzle.Grid = make([][]int, puzzle.RowCount)
			for x := 0; x < puzzle.RowCount; x++ {
				puzzle.Grid[x] = make([]int, puzzle.ColCount)
				for y := 0; y < puzzle.ColCount; y++ {
					puzzle.Grid[x][y] = UNKNOWN
				}
			}
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
	return err
}

func (puzzle *PuzzleStruct) Dump() string {
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
		output += displayLine(NewLine(ROW, i), true)
		output += "\n"
	}
	return output
}

func (puzzle *PuzzleStruct) maxColClueLength() int {
	maxLength := 0
	for _, clue := range puzzle.ColClueData {
		if len(clue) > maxLength {
			maxLength = len(clue)
		}
	}
	return maxLength
}
