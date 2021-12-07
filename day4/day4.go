package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	value int
	checked bool
}

func (c *Cell) Check(val int) bool {
	if c.value == val {
		c.checked = true
		return true
	}
	return false
}

func NewCell(val int) *Cell {
	cell := Cell{}
	cell.checked = false
	cell.value = val
	return &cell
}

type Cells struct {
	cells []*Cell
}

func NewRow() *Cells {
	row := Cells{}
	row.cells = make([]*Cell, 0)
	return &row
}

func (c *Cells) AddCell(cell *Cell) {
	c.cells = append(c.cells, cell)
}

func (c *Cells) Check(val int) bool {
	for idx, _ := range c.cells {
		if c.cells[idx].Check(val) {
			return true
		}
	}
	return false
}

func (c *Cells) IsAllChecked() bool {
	for _, cell := range c.cells {
		if !cell.checked {
			return false
		}
	}
	return true
}

type Sheet struct {
	rows []Cells
	cols []Cells
}

func (s *Sheet) Check(val int) bool {
	for rowIdx, _ := range s.rows {
		for cellIdx, _ := range s.rows[rowIdx].cells {
			if s.rows[rowIdx].cells[cellIdx].Check(val) {
				return true
			}
		}
	}
	return false
}

func (s *Sheet) DidBingo() bool {
	// check rows
	for _, row := range s.rows {
		if row.IsAllChecked() {
			return true
		}
	}
	for _, col := range s.cols {
		if col.IsAllChecked() {
			return true
		}
	}
	return false
}

func (s *Sheet) Calc(call int) int {
	sumUnchecked := 0
	for _, row := range s.rows {
		for _, cell := range row.cells {
			if !cell.checked {
				sumUnchecked += cell.value
			}
		}
	}
	return sumUnchecked * call
}

func (s *Sheet) AddRow(row *Cells) {
	s.rows = append(s.rows, *row)
}

func (s *Sheet) AddCol(row *Cells) {
	s.cols = append(s.cols, *row)
}

func NewSheet() *Sheet {
	sheet := Sheet{}
	sheet.rows = make([]Cells, 0)
	sheet.cols = make([]Cells, 0)
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
			sheet.AddCol(NewRow())
		}


		for i := 0; i < 5; i++ {
			row := NewRow()
			scanner.Scan()
			raw := scanner.Text()
			strArr := strings.Fields(raw)
			for idx, strCell := range strArr {
				intCell, _ := strconv.Atoi(strCell)
				cell := NewCell(intCell)
				sheet.cols[idx].AddCell(cell)
				row.AddCell(cell)
			}
			sheet.AddRow(row)
		}
		sheets = append(sheets, *sheet)
	}

	countBingos := 0

	winners := make(map[int]bool, 0)

	Exit:
	for callIdx, call := range calls {
		for idx, sheet := range sheets {
			if sheet.Check(call) {
				if sheet.DidBingo() && !winners[idx] {
					fmt.Println(fmt.Sprintf("Bingo on sheet %d, call %d at %d. Position %d:", idx, call, callIdx, len(winners)))
					winners[idx] = true
					if countBingos == len(sheets) - 1 {
						fmt.Println(sheet.Calc(call))
						break Exit
					}
					countBingos += 1
				}
			}
		}
	}

	fmt.Println("Done!")

}
