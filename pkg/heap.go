package pkg

import "sort"

type LineHeap struct {
	Lines      map[string]*Line
	LineLabels []string
	Puzzle     *Puzzle
}

func NewLineHeap(puzzle *Puzzle) LineHeap {
	heap := LineHeap{}
	heap.Puzzle = puzzle
	heap.Lines = make(map[string]*Line)
	log.Debugf("Puzzle rows - %v", puzzle.RowCount)
	for i := 0; i < puzzle.RowCount; i++ {
		line := NewLine(puzzle, Row, i)
		heap.Lines[line.Label()] = line
	}
	for i := 0; i < puzzle.ColCount; i++ {
		line := NewLine(puzzle, Col, i)
		heap.Lines[line.Label()] = line
	}
	heap.LineLabels = make([]string, 0, len(heap.Lines))
	for k := range heap.Lines {
		heap.LineLabels = append(heap.LineLabels, k)
	}
	sort.Slice(
		heap.LineLabels,
		func(i, j int) bool {
			return sortLineLabel(heap.LineLabels[i]) < sortLineLabel(heap.LineLabels[j])
		},
	)
	return heap
}

func (heap *LineHeap) hasLinesToSolve() bool {
	return len(heap.Lines) > 0
}

func (heap *LineHeap) removeLine(line *Line) {
	delete(heap.Lines, line.Label())
}
