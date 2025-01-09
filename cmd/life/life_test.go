package life

import (
	"testing"
)

func TestNeighbors(t *testing.T) {

	type test struct {
		cells     [][]bool
		X         int
		Y         int
		neighbors int
	}

	tests := []test{
		{
			cells: [][]bool{
				{true, false, false},
				{true, true, false},
				{true, true, false},
			},
			X:         1,
			Y:         1,
			neighbors: 4,
		},
		{
			cells: [][]bool{
				{true, false, false},
				{true, true, false},
				{true, true, false},
			},
			X:         0,
			Y:         0,
			neighbors: 2,
		},
	}

	for _, tt := range tests {
		w := World{
			Height: len(tt.cells),
			Width:  len(tt.cells[0]),
			Cells:  tt.cells,
		}
		n := w.Neighbors(tt.X, tt.Y)

		if n != tt.neighbors {
			t.Fatalf("expected: %v, got: %v", tt.neighbors, n)
		}
	}
}
