package pkg

import "github.com/fatih/color"

type Cell int

const ROW = int(0)
const COL = int(1)
const UNKNOWN = Cell(0)
const SPACE = Cell(1)
const BLOCK = Cell(2)

func (c *Cell) displayString(displayPadding int) string {
	var displayCharacter string
	switch *c {
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
