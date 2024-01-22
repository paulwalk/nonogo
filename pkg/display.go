package pkg

import (
	"github.com/fatih/color"
	"strconv"
	"strings"
)

func displayCell(cellValue, displayPadding int) string {
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
	return color.BlueString("%*v", displayPadding, displayCharacter)
}

func displayLine(line *Line, displayClues bool) string {
	lineString := ""
	for _, cell := range line.Cells() {
		lineString += displayCell(*cell, line.Puzzle.DisplayPadding)
	}
	if displayClues {
		clueString := make([]string, 0)
		for _, clue := range line.Clue() {
			clueString = append(clueString, strconv.Itoa(clue))
		}
		lineString += " "
		lineString += strings.Join(clueString, ",")
	}
	lineString += color.YellowString(" %v", line.Index+1)
	return lineString
}
