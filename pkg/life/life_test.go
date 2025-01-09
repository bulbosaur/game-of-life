package life

import (
	"testing"
)

func TestNeighbors(t *testing.T) {
	testCases := []struct {
		world    *World
		x, y     int
		expected int
	}{
		{
			world: func() *World {
				w := NewWorld(3, 3)
				return w
			}(),
			x:        1,
			y:        1,
			expected: 0,
		},
		{
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
			world: func() *World {
				w := NewWorld(3, 3)
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
		got := tc.world.Neighbors(tc.x, tc.y)
		if got != tc.expected {
			t.Errorf("Neighbors(%d, %d) = %d; want %d",
				tc.x, tc.y, got, tc.expected)
		}
	}
}

func TestNext(t *testing.T) {
	testCases := []struct {
		world    *World
		x, y     int
		expected bool
	}{
		{
			world: func() *World {
				w := NewWorld(3, 3)
				w.Cells[0][0] = true
				w.Cells[0][1] = true
				w.Cells[1][0] = true
				return w
			}(),
			x:        1,
			y:        1,
			expected: true,
		},
		{
			world: func() *World {
				w := NewWorld(3, 3)
				w.Cells[0][0] = true
				w.Cells[0][1] = true
				w.Cells[1][1] = true
				return w
			}(),
			x:        1,
			y:        1,
			expected: true,
		},
		{
			world: func() *World {
				w := NewWorld(3, 3)
				w.Cells[0][0] = true
				w.Cells[0][1] = true
				w.Cells[1][0] = true
				w.Cells[1][1] = true
				return w
			}(),
			x:        1,
			y:        1,
			expected: true,
		},
		{
			world: func() *World {
				w := NewWorld(3, 3)
				w.Cells[0][0] = true
				w.Cells[0][1] = true
				w.Cells[0][2] = true
				w.Cells[1][0] = true
				w.Cells[1][1] = true
				return w
			}(),
			x:        1,
			y:        1,
			expected: false,
		},
		{
			world: func() *World {
				w := NewWorld(3, 3)
				w.Cells[0][0] = true
				w.Cells[0][1] = true
				return w
			}(),
			x:        1,
			y:        1,
			expected: false,
		},
	}

	for _, tc := range testCases {
		got := tc.world.Next(tc.x, tc.y)
		if got != tc.expected {
			t.Errorf("Next(%d, %d) = %v; want %v",
				tc.x, tc.y, got, tc.expected)
		}
	}
}
