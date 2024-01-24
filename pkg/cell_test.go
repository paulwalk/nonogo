package pkg

import (
	"fmt"
	"testing"
)

func TestCell_displayString(t *testing.T) {
	type args struct {
		displayPadding int
	}
	tests := []struct {
		name string
		c    Cell
		args args
		want string
	}{
		{"test displayString() for Unknown", Cell(Unknown), args{3}, fmt.Sprintf("%*v", 3, UnknownDisplayString)},
		{"test displayString() for Space", Cell(Space), args{3}, fmt.Sprintf("%*v", 3, SpaceDisplayString)},
		{"test displayString() for Block", Cell(Block), args{3}, fmt.Sprintf("%*v", 3, BlockDisplayString)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.displayString(tt.args.displayPadding); got != tt.want {
				t.Errorf("displayString() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
