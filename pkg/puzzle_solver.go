package pkg

import (
	"fmt"
	"go.uber.org/zap"
)

type PuzzleSolver struct {
	Heap   LineHeap
	Puzzle *Puzzle
}

var log *zap.SugaredLogger

func NewPuzzleSolver(puzzleFilePath string, logger *zap.SugaredLogger) (PuzzleSolver, error) {
	var err error
	log = logger
	puzzleSolver := PuzzleSolver{}
	puzzle, err := NewPuzzle(puzzleFilePath)
	if err != nil {
		log.Error(err.Error())
		return puzzleSolver, err
	}
	puzzleSolver.Puzzle = &puzzle
	puzzleSolver.Heap = NewLineHeap(&puzzle)
	return puzzleSolver, err
}

func (puzzleSolver *PuzzleSolver) Solve() int {
	iteration := 0
	progressWasMade := true
	for puzzleSolver.Heap.hasLinesToSolve() == true && progressWasMade == true {
		progressWasMade = false
		iteration += 1
		log.Debug(fmt.Sprintf("Beginning solution iteration %v...", iteration))
		for _, lineLabel := range puzzleSolver.Heap.LineLabels {
			if line, exists := puzzleSolver.Heap.Lines[lineLabel]; exists {
				log.Debug(fmt.Sprintf("Attempting to solve '%s'...", line.Label()))
				if solveLine(line) {
					progressWasMade = true
				}
				if line.processIfSolved() {
					puzzleSolver.Heap.removeLine(line)
				}
			}
		}
	}
	return iteration
}
