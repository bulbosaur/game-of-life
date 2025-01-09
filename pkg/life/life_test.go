package life

import (
	"testing"
)

func TestNeighbors(t *testing.T) {
	testCases := []struct {
		name     string
		world    *World
		x, y     int
		expected int
	}{
		{
			name: "Пустой мир 3x3",
			world: func() *World {
				w := NewWorld(3, 3)
				return w
			}(),
			x:        1,
			y:        1,
			expected: 0,
		},
		{
			name: "Одна живая клетка сверху",
			world: func() *World {
				w := NewWorld(3, 3)
				w.Cells[0][1] = true
				return w
			}(),
			x:        1,
			y:        1,
			expected: 1,
		},
		{
			name: "Все соседи живые",
			world: func() *World {
				w := NewWorld(3, 3)
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						w.Cells[i][j] = true
					}
				}
				return w
			}(),
			x:        1,
			y:        1,
			expected: 8,
		},
		{
			name: "Проверка угловой клетки",
			world: func() *World {
				w := life.NewWorld(3, 3)
				w.Cells[0][1] = true
				w.Cells[1][0] = true
				w.Cells[1][1] = true
				return w
			}(),
			x:        0,
			y:        0,
			expected: 3,
		},
		{
			name: "Проверка граничной клетки",
			world: func() *World {
				w := NewWorld(3, 3)
				w.Cells[0][0] = true
				w.Cells[0][2] = true
				w.Cells[1][1] = true
				return w
			}(),
			x:        1,
			y:        0,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.world.Neighbors(tc.x, tc.y)
			if got != tc.expected {
				t.Errorf("Neighbors(%d, %d) = %d; want %d",
					tc.x, tc.y, got, tc.expected)
			}
		})
	}
}
