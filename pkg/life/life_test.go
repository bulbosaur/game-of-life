package life

import (
	"os"
	"path/filepath"
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

func TestSaveState(t *testing.T) {
	tmpDir := "test_tmp"
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		world    *World
		filename string
		want     string
		wantErr  bool
	}{
		{
			world: &World{
				Height: 3,
				Width:  3,
				Cells: [][]bool{
					{true, false, true},
					{false, true, false},
					{true, false, true},
				},
			},
			filename: filepath.Join(tmpDir, "test1.txt"),
			want:     "101\n010\n101",
			wantErr:  false,
		},
		{
			world: &World{
				Height: 2,
				Width:  2,
				Cells: [][]bool{
					{true, true},
					{false, false},
				},
			},
			filename: filepath.Join(tmpDir, "test2.txt"),
			want:     "11\n00",
			wantErr:  false,
		},
		{
			world: &World{
				Height: 1,
				Width:  1,
				Cells: [][]bool{
					{true},
				},
			},
			filename: filepath.Join(tmpDir, "test3.txt"),
			want:     "1",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		err := tt.world.SaveState(tt.filename)

		if (err != nil) != tt.wantErr {
			t.Errorf("SaveState() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		content, err := os.ReadFile(tt.filename)
		if err != nil {
			t.Errorf("Failed to read test file: %v", err)
			return
		}

		got := string(content)
		if got != tt.want {
			t.Errorf("SaveState() content = %v, want %v", got, tt.want)
		}
	}
}
