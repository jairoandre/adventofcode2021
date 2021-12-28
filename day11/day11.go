package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
https://adventofcode.com/2021/day/11
*/

type Cell struct {
	Value   int
	Flashed bool
	Ticked  bool
}

func (c *Cell) Tick() bool {
	// already flashed
	if c.Flashed {
		c.Value = 0
		return false
	}
	c.Value++
	if c.Value > 9 {
		c.Value = 0
		c.Flashed = true
		return true
	}
	c.Flashed = false
	return false
}

type Grid struct {
	Cols  int
	Rows  int
	Cells []Cell
}

func (g *Grid) NextState() {
	for x := 0; x < g.Cols; x++ {
		for y := 0; y < g.Rows; y++ {
			g.TickCell(x, y)
		}
	}
}

func (g *Grid) Print() {
	for y := 0; y < g.Rows; y++ {
		for x := 0; x < g.Cols; x++ {
			fmt.Print(g.Cells[y*g.Cols+x].Value)
		}
		fmt.Print("\n")
	}
	fmt.Println("-------")
}

func (g *Grid) Get(x, y int) *Cell {
	idx := g.Cols*y + x
	if idx < 0 || idx > len(g.Cells)-1 {
		return nil
	}
	return &g.Cells[idx]
}

func (g *Grid) TickCell(x, y int) {
	if x > g.Cols-1 || y > g.Rows-1 || x < 0 || y < 0 {
		return
	}
	cell := g.Get(x, y)
	if cell == nil {
		return
	}
	if !cell.Tick() {
		return
	}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			g.TickCell(x+i, y+j)
		}
	}
}

func (g *Grid) CountFlashed() int {
	result := 0
	for idx := 0; idx < len(g.Cells); idx++ {
		if g.Cells[idx].Flashed {
			result++
		}
	}
	return result
}

func (g *Grid) ResetAllFlashed() {
	for idx := 0; idx < len(g.Cells); idx++ {
		g.Cells[idx].Flashed = false
	}
}

func Solution() {
	file, _ := os.Open("day11/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	cells := make([]Cell, 0)
	cols := 0
	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)
		for _, val := range line {
			value, _ := strconv.Atoi(string(val))
			cell := Cell{
				Value:   value,
				Flashed: false,
			}
			cells = append(cells, cell)
		}
	}
	grid := Grid{
		Cols:  cols,
		Rows:  len(cells) / cols,
		Cells: cells,
	}

	result := 0

	fmt.Println("before any step ")
	grid.Print()
	//for i := 0; i < 100; i++ {
	step := 1
	cellsCount := len(grid.Cells)
	for {
		grid.NextState()
		grid.Print()
		result = grid.CountFlashed()
		if result == cellsCount {
			break
		}
		grid.ResetAllFlashed()
		step++
	}

	fmt.Println(step)

}
