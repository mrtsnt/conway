package main

import (
    "time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	gridRows = 65
	gridCols = 120
)

var drawMode = true

func drawGrid(grid [][]bool, imd *imdraw.IMDraw) {
    imd.Clear()

	for y := len(grid) - 1; y >= 0; y-- {
		for x := range grid[y] {
			if grid[y][x] {
				xV := float64(x) * 10.0
				yV := float64(y) * 10.0
				imd.Push(pixel.V(xV+1.0, yV+1.0), pixel.V(xV+9.0, yV+9.0))
                imd.Color = colornames.White
				imd.Rectangle(0)
			}
		}
	}
}

func run() {
	win := newWindow()
	imd := imdraw.New(nil)
    grid := newGrid(gridRows, gridCols)

	for !win.Closed() {

        if drawMode {

            if win.JustPressed(pixelgl.KeySpace) {
                imd.Clear()
                win.Clear(colornames.Black)
                grid.reset()
            }

            if win.Pressed(pixelgl.MouseButtonLeft) {
                mousePos := win.MousePosition()
                msGridX := int(mousePos.X / 10)
                msGridY := int(mousePos.Y / 10)

                if msGridX >= 0 && msGridX < gridCols && msGridY >= 0 && msGridY < gridRows {

                    if !grid.gridA[msGridY][msGridX] {
                        xV := float64(msGridX) * 10.0
                        yV := float64(msGridY) * 10.0
                        imd.Push(pixel.V(xV+1.0, yV+1.0), pixel.V(xV+9.0, yV+9.0))
                        imd.Color = colornames.White
                        imd.Rectangle(0)
                        imd.Draw(win)
                        grid.gridA[msGridY][msGridX] = true
                    }
                }
            }

            drawMode = !win.JustPressed(pixelgl.KeyP)
        } else {
            imd.Clear()
            win.Clear(colornames.Black)
            grid.makeTurn()
            drawGrid(grid.activeGrid(), imd)
            imd.Draw(win)
            time.Sleep(time.Millisecond * 250)

            if win.JustPressed(pixelgl.KeySpace) {
                imd.Clear()
                win.Clear(colornames.Black)
                grid.reset()
                drawMode = true
            }
        }

		win.Update()
	}
}

func newWindow() *pixelgl.Window {
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
