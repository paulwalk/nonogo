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
	puzzleSolver, err := NewPuzzleSolver(puzzleFilePath, log)
	if err != nil {
		log.Fatal(err.Error())
	}
	iterations := puzzleSolver.Solve()
	elapsedTime := time.Since(start)
	fmt.Print(puzzleSolver.Puzzle.Dump())
	unsolvedLines := len(puzzleSolver.Heap.Lines)
	log.Info(fmt.Sprintf("Unsolved lines = %v out of %v", unsolvedLines, puzzleSolver.Puzzle.ColCount+puzzleSolver.Puzzle.RowCount))
	log.Info(fmt.Sprintf("Iterations = %v", iterations))
	log.Info(fmt.Sprintf("Execution took %s", elapsedTime))
	if unsolvedLines > 0 {
		log.Info("Failed to solve puzzle")
	} else {
		log.Info("Puzzle solved!")
	}
}
