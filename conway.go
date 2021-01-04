package main

import (
    "fmt"
)

type grid int
const (
    A grid = iota
    B
)

type gameGrid struct {
    gridA [][]bool
    gridB [][]bool
    currentGrid grid
}

func newGrid(r int, c int) *gameGrid {
    grid := new(gameGrid)

    initGrid := func () [][]bool {
        g := make([][]bool, r)
        for row := range g {
            g[row] = make([]bool, c)
        }
        return g
    }

    grid.gridA = initGrid()
    grid.gridB = initGrid()
    grid.currentGrid = A

    return grid
}

func (g *gameGrid) makeTurn() {
    fmt.Println("making turn")
    var src, tgt [][]bool
    if g.currentGrid == A {
        src, tgt, g.currentGrid = g.gridA, g.gridB, B
    } else {
        src, tgt, g.currentGrid = g.gridB, g.gridA, A
    }

    for y := range src {
        for x := range src[y] {
            neighbours := getNeighbours(x, y, src)
            if src[y][x] {
                tgt[y][x] = neighbours == 2 || neighbours == 3
            } else {
                tgt[y][x] = neighbours == 3
            }
        }
    }
}

func getNeighbours(x int, y int, grid [][]bool) int {
    neigbours := 0

    for i := y - 1; i <= y + 1; i++ {
        for j := x - 1; j <= x + 1; j++ {
            if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] {
                fmt.Println("checking {X} {Y}", x, y)
                neigbours++
            }
        }
    }

    return neigbours
}

func (g *gameGrid) activeGrid() [][]bool {
    if g.currentGrid == A {
        return g.gridA
    } else {
        return g.gridB
    }
}

func (g *gameGrid) reset() {
    for y := range g.gridA {
        for x := range g.gridA[y] {
            g.gridA[y][x] = false
            g.gridB[y][x] = false
        }
    }
    g.currentGrid = A
}
