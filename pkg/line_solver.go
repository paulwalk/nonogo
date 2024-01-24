package pkg

func solveLine(line *Line) bool {
	progressFromAlgorithm1 := findCellsWhichAreSameInAllPotentialSolutions(line)
	progressFromAlgorithm2 := removeSolutionsWhichDoNotFitKnownCells(line)
	if progressFromAlgorithm1 || progressFromAlgorithm2 {
		return true
	} else {
		return false
	}
}

func findCellsWhichAreSameInAllPotentialSolutions(line *Line) bool {
	log.Debug("Finding cells which are blocks/spaces in all potential solutions....")
	progressWasMade := false
	for cellIndex, cell := range line.Cells() {
		if *cell == Unknown {
			foundNonBlockCell := false
			for _, solution := range line.PotentialSolutions {
				if solution[cellIndex] != Block {
					foundNonBlockCell = true
					break
				}
			}
			if foundNonBlockCell == false {
				line.setCell(cellIndex, Block)
				progressWasMade = true
			} else {
				foundNonSpaceCell := false
				for _, solution := range line.PotentialSolutions {
					if solution[cellIndex] != Space {
						foundNonSpaceCell = true
						break
					}
				}
				if foundNonSpaceCell == false {
					line.setCell(cellIndex, Space)
					progressWasMade = true
				}
			}
		}
	}
	return progressWasMade
}

func removeSolutionsWhichDoNotFitKnownCells(line *Line) bool {
	log.Debug("Finding solutions which do not fit known cells....")
	progressWasMade := false
	for cellIndex, cell := range line.Cells() {
		if *cell != Unknown {
			solutionIndex := 0 // output index
			for _, solution := range line.PotentialSolutions {
				if solution[cellIndex] == *cell {
					line.PotentialSolutions[solutionIndex] = solution
					solutionIndex++
				} else {
					progressWasMade = true
				}
			}
			for j := solutionIndex; j < len(line.PotentialSolutions); j++ {
				line.PotentialSolutions[j] = nil
			}
			line.PotentialSolutions = line.PotentialSolutions[:solutionIndex]
		}
	}
	return progressWasMade
}
