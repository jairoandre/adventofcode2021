package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	value int
	checked bool
}

func (c *Cell) Check(val int) {
	if c.value == val {
		c.checked = true
	}
}

func NewCell(val int) *Cell {
	cell := Cell{}
	cell.checked = false
	cell.value = val
	return &cell
}

type Row struct {
	cells []Cell
}

func NewRow() *Row {
	row := Row{}
	row.cells = make([]Cell, 0)
	return &row
}

func (r *Row) AddCell(val int) {
	r.cells = append(r.cells, *NewCell(val))
}

func (r *Row) Check(val int) {
	for _, cell := range r.cells {
		cell.Check(val)
	}
}

func (r *Row) IsAllChecked() bool {
	for _, cell := range r.cells {
		if !cell.checked {
			return false
		}
	}
	return true
}

type Sheet struct {
	rows []Row
}

func (s *Sheet) Check() bool {
	// check rows
	for _, row := range s.rows {
		if row.IsAllChecked() {
			return true
		}
	}
	return false
}

func (s *Sheet) AddRow(row *Row) {
	s.rows = append(s.rows, *row)
}

func NewSheet() *Sheet {
	sheet := Sheet{}
	sheet.rows = make([]Row, 0)
	return &sheet
}

func Solution1() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calls := make([]int, 0)
	scanner.Scan()
	for _, strCall := range strings.Split(scanner.Text(), ",") {
		intCall, _ := strconv.Atoi(strCall)
		calls = append(calls, intCall)
	}

	sheets := make([]Sheet, 0)

	for scanner.Scan() {
		scanner.Text() // blank line
		sheet := NewSheet()
		for i := 0; i < 5; i++ {
			row := NewRow()
			for _, strCell := range strings.Split(scanner.Text(), ",") {
				intCell, _ := strconv.Atoi(strCell)
				row.AddCell(intCell)
			}
			sheet.AddRow(row)
		}
		sheets = append(sheets, *sheet)
	}

}
