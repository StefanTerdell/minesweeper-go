package main

import (
	"math"
	"math/rand"

	"github.com/buger/goterm"
)

type Cell struct {
	flagged bool
	mine    bool
	visible bool
	close   int
}

const (
	STATE_PLAYING = iota + 1
	STATE_WON
	STATE_LOST
)

type Grid = [][]Cell

func create_grid(size_x int, size_y int, incidence float32) Grid {
	grid := []([]Cell){}

	for y := 0; y < size_y; y++ {
		row := [](Cell){}

		for x := 0; x < size_x; x++ {
			row = append(row, Cell{
				flagged: false,
				mine:    rand.Float32() < incidence,
				visible: false,
				close:   0,
			})
		}

		grid = append(grid, row)
	}

	return count_mines(grid)
}

func print_grid(grid Grid) {

	goterm.MoveCursor(1, 1)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			cell := grid[y][x]

			if cell.visible == false {
				goterm.Print(" I")
			} else if cell.flagged {
				goterm.Print(" f")
			} else if cell.mine {
				goterm.Print(" *")
			} else if cell.close > 0 {
				goterm.Print(" ", cell.close)
			} else {
				goterm.Print("  ")
			}
		}
		goterm.Println("")
	}

	goterm.Flush()
}

func count_mines(grid Grid) Grid {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {

			if !grid[y][x].mine {
				for sy := int(math.Max(0, float64(y-1))); sy < len(grid) && sy <= y+1; sy++ {
					for sx := int(math.Max(0, float64(x-1))); sx < len(grid[y]) && sx <= x+1; sx++ {
						if sy == 0 && sx == 0 {
							continue
						}

						if grid[sy][sx].mine {
							grid[y][x].close += 1
						}
					}
				}

			}

		}
		println("")
	}

	return grid
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Returns true if you revealed a mine
func reveal(grid Grid, x int, y int) bool {
	if grid[y][x].visible {
		return false
	}

	grid[y][x].visible = true

	if grid[y][x].mine {
		return true
	}

	if grid[y][x].close == 0 {
		max_y := len(grid) - 1
		for sy := max(y-1, 0); sy <= min(y+1, max_y); sy++ {
			max_x := len(grid[y]) - 1
			for sx := max(x-1, 0); sx <= min(x+1, max_x); sx++ {
				if sx == 0 && sy == 0 {
					continue
				}

				reveal(grid, sx, sy)
			}
		}
	}

	return false
}
