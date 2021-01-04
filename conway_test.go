package main

import (
    "testing"
    "fmt"
)

func TestGetNeighbours(t *testing.T) {
    grid := make([][]bool, 2)
    grid[0] = make([]bool, 2)
    grid[1] = make([]bool, 2)
    grid[0][0], grid[0][1], grid[1][0] = true, true, true

    var params = []struct{
        x, y int
        want int
    }{
        {0, 0, 2},
        {0, 1, 2},
        {1, 0, 2},
        {1, 1, 3},
    }

    for _, param := range params {
        name := fmt.Sprintf("x:%d, y:%d, want:%d", param.x, param.y, param.want)
        t.Run(name, func (t *testing.T) {
            got := getNeighbours(param.x, param.y, grid)
            if got != param.want {
                t.Errorf("got %d, want %d", got, param.want)
            }
        })
    }
}
