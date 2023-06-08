package core_test

import (
    "fmt"
    "testing"
    "time"

    "github.com/gapisani/arrows/core"
)

func rotArr(arr rune, dir core.Direction) string {
    switch dir {
    case core.WEST:
        return string(arr)
    case core.NORTH:
        return string(arr+1)
    case core.EAST:
        return string(arr+2)
    case core.SOUTH:
        return string(arr+3)
    }
    return string(arr)
}

func render(g core.Grid) {
    w, h := g.Dimensions()
    for y := uint(0); y < h; y++ {
        for x := uint(0); x < w; x++ {
            cell := g.GetCell(x, y)
            dir := (*cell).Dir()
            lit := (*cell).Check()
            switch (*cell).(type) {
            case *core.Wire:
                if(lit) {
                    fmt.Print(rotArr('⬅', dir))
                } else {
                    fmt.Print(rotArr('←', dir))
                }
            case *core.Source:
                fmt.Print("@")
            case *core.None:
                fmt.Print(" ")
            case *core.MemCell:
                if(lit) {
                    fmt.Print("#")
                } else {
                    fmt.Print("O")
                }
            default:
                fmt.Print("x")
            }
        }
        fmt.Println()
    }
}

func TestMain(t *testing.T) {
    g := core.Grid{}
    g.Init(50, 50)
    w, h := g.Dimensions()
    for i := uint(h)-1; i > 0; i-- {
        *g.GetCell(1, i) = core.Cell(&core.MemCell{})
        *g.GetCell(2, i) = core.Cell(&core.Get{})
        (*g.GetCell(2, i)).SetDir(core.EAST)
        for j := uint(3); j < w; j++ {
            *g.GetCell(j, i) = core.Cell(&core.Wire{})
            (*g.GetCell(j, i)).SetDir(core.EAST)
        }
    }
    *g.GetCell(1, h-1) = core.Cell(&core.Source{})
    g.FAST = true
    g.RecountUpdate()
    for t := 0; t <= 100; t++ {
        g.Update()
        render(g)
        time.Sleep(time.Millisecond * 1000)
    }
}
