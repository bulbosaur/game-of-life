package life

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
)

var (
	ErrorFieldSize        = errors.New("dimensions in the file: different number of elements in lines")
	ErrorInvalidCharacter = errors.New("invalid character in LoadState file")
)

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

func (w *World) Next(x, y int) bool {
	n := w.Neighbors(x, y)
	alive := w.Cells[y][x]
	if n < 4 && n > 1 && alive {
		return true
	}
	if n == 3 && !alive {
		return true
	}
	return false
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

func (w *World) Seed() {
	for _, row := range w.Cells {
		for i := range row {
			if rand.Intn(10) == 1 {
				row[i] = true
			}
		}
	}
}

func (w *World) String() string {
	var (
		str         string
		brownSquare = "\xF0\x9F\x9F\xAB"
		greenSquare = "\xF0\x9F\x9F\xA9"
	)

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if w.Cells[y][x] {
				str += greenSquare
			} else {
				str += brownSquare
			}
		}
		str += "\n"
	}
	return str
}

func (w *World) SaveState(filename string) error {
	var (
		str string
	)

	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if w.Cells[y][x] {
				str += "1"
			} else {
				str += "0"
			}
		}
		if y+1 != w.Height {
			str += "\n"
		}
	}

	_, err = f.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}

func hasManyUnique(slice []int) bool {
	unique := make(map[int]bool)
	for _, v := range slice {
		unique[v] = true
	}
	return len(unique) > 1
}

func (w *World) LoadState(filename string) error {
	var (
		x         int
		y         int
		allWidths []int
	)

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		y++
		x += len(scanner.Text())
		allWidths = append(allWidths, x)
	}

	if hasManyUnique(allWidths) {
		return ErrorFieldSize
	}

	w.Height = y
	w.Width = x

	w.Cells = make([][]bool, w.Height)
	for i := range w.Cells {
		w.Cells[i] = make([]bool, w.Width)
	}

	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)

	y = 0
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			if string(line[x]) == "1" {
				w.Cells[y][x] = true
			} else if string(line[x]) == "0" {
				w.Cells[y][x] = false
			} else {
				return ErrorInvalidCharacter
			}
		}
		y++
	}

	return nil
}
