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
	Value int
	Visited bool
}

func Solution() {
	file, _ := os.Open("day11/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	cells := make([]Cell, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		for _, val := range line {
			value, _ := strconv.Atoi(string(val))
			cell := Cell {
				Value: value,
				Visited: false,
			}
			cells = append(cells, cell)
		}
	}
	fmt.Println(cells)
}