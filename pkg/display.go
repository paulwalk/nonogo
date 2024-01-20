package pkg

import (
	"github.com/fatih/color"
	"strconv"
	"strings"
)

func displayCell(cellValue int) string {
	var displayCharacter string
	switch cellValue {
	case SPACE:
		//displayCharacter = "☐"
		displayCharacter = "∙"
	case BLOCK:
		displayCharacter = "◼"
	default:
		//displayCharacter = "∙"
		displayCharacter = " "
	}
	//return fmt.Sprintf("%*v", Puzzle.DisplayPadding, displayCharacter)
	return color.BlueString("%*v", Puzzle.DisplayPadding, displayCharacter)
}

func displayLine(line *Line, displayClues bool) string {
	lineString := ""
	for _, cell := range line.Cells() {
		lineString += displayCell(*cell)
	}
	if displayClues {
		clueString := make([]string, 0)
		for _, clue := range line.Clue() {
			clueString = append(clueString, strconv.Itoa(clue))
		}
		lineString += " "
		lineString += strings.Join(clueString, ",")
	}
	lineString += color.YellowString(" %v", line.Index)
	return lineString
}
