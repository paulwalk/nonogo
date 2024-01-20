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
	log.Info(fmt.Sprintf("Unsolved lines = %v out of %v", len(Solver.Heap.Lines), Puzzle.ColCount+Puzzle.RowCount))
	log.Info(fmt.Sprintf("Iterations = %v", iterations))
	log.Info(fmt.Sprintf("Execution took %s", elapsedTime))
	log.Info("Solve completed")
}
