package life

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) *World {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (w *World) Neighbors(x, y int) int {
	var counter int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if x+j >= len(w.Cells[y]) || y+i >= len(w.Cells) || x+j < 0 || y+i < 0 {
				continue
			}
			if i == 0 && j == 0 {
				continue
			}
			if w.Cells[y+i][x+j] {
				counter++
			}
		}
	}
	return counter
}
