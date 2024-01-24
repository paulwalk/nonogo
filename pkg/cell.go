package pkg

import "github.com/fatih/color"

type Cell int

const Row = int(0)
const Col = int(1)
const Unknown = Cell(0)
const Space = Cell(1)
const Block = Cell(2)
const UnknownDisplayString = " "
const SpaceDisplayString = "∙"
const BlockDisplayString = "◼"

func (c *Cell) displayString(displayPadding int) string {
	var displayCharacter string
	switch *c {
	case Space:
		//displayCharacter = "☐"
		displayCharacter = SpaceDisplayString
	case Block:
		displayCharacter = BlockDisplayString
	default:
		//displayCharacter = "∙"
		displayCharacter = UnknownDisplayString
	}
	return color.BlueString("%*v", displayPadding, displayCharacter)
}
