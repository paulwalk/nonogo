package cmd

import (
	"fmt"
	. "github.com/paulwalk/nonogo/pkg"
	"github.com/spf13/cobra"
	"time"
)

var solveCmd = &cobra.Command{
	Use: "solve",
	Run: func(cmd *cobra.Command, args []string) {
		initialiseApplication()
		solve()
	},
}

func solve() {
	start := time.Now()
	log.Info("Solve starting...")
	var Solver = PuzzleSolver{}
	err := Solver.Initialise(puzzleFilePath, log)
	if err != nil {
		log.Fatal(err.Error())
	}
	iterations := Solver.Solve()
	elapsedTime := time.Since(start)
	fmt.Print(Puzzle.Dump())
	unsolvedLines := len(Solver.Heap.Lines)
	log.Info(fmt.Sprintf("Unsolved lines = %v out of %v", unsolvedLines, Puzzle.ColCount+Puzzle.RowCount))
	log.Info(fmt.Sprintf("Iterations = %v", iterations))
	log.Info(fmt.Sprintf("Execution took %s", elapsedTime))
	if unsolvedLines > 0 {
		log.Info("Failed to solve puzzle")
	} else {
		log.Info("Puzzle solved!")
	}
}
