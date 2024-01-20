package pkg

import "sort"

type LineHeap struct {
	Lines      map[string]*Line
	LineLabels []string
}

func (heap *LineHeap) hasLinesToSolve() bool {
	return len(heap.Lines) > 0
}

func (heap *LineHeap) Initialise() {
	heap.Lines = make(map[string]*Line)
	log.Debugf("Puzzle rows - %v", Puzzle.RowCount)
	for i := 0; i < Puzzle.RowCount; i++ {
		line := NewLine(ROW, i)
		heap.Lines[line.Label()] = line
	}
	for i := 0; i < Puzzle.ColCount; i++ {
		line := NewLine(COL, i)
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
}

func (heap *LineHeap) removeLine(line *Line) {
	delete(heap.Lines, line.Label())
}
