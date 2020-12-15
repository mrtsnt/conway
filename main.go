package main

import (
	"math/rand"
    t "time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	gridRows = 48
	gridCols = 64
)

func drawGrid(grid [][]bool, imd *imdraw.IMDraw) {
    imd.Color = colornames.White

	for y := len(grid) - 1; y >= 0; y-- {
		for x := range grid[y] {
			if grid[y][x] {
				xV := float64(x) * 10.0
				yV := float64(y) * 10.0
				imd.Push(pixel.V(xV+1.0, yV+1.0), pixel.V(xV+9.0, yV+9.0))
				imd.Rectangle(0)
			}
		}
	}
}

func testPop(grid [][]bool) {
	for y := range grid {
		for x := range grid[y] {
			r := rand.Intn(2)
			if r == 0 {
				grid[y][x] = true
			} else {
				grid[y][x] = false
			}
		}
	}
}

func run() {
	win := createWindow()
	imd := imdraw.New(nil)
    grid := createGrid()

	for !win.Closed() {
		imd.Clear()

        testPop(grid)
        drawGrid(grid, imd)

		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()

        t.Sleep(t.Millisecond * 500)
	}
}

func createGrid() [][]bool {
	grid := make([][]bool, gridRows)
	for r := range grid {
		grid[r] = make([]bool, gridCols)
	}

    return grid
}

func createWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "conway",
		Bounds: pixel.R(0, 0, gridCols*10, gridRows*10),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}

func main() {
	pixelgl.Run(run)
}
