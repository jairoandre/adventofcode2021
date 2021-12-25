package day11

import (
	"bufio"
	"fmt"
	"os"
)

/*
https://adventofcode.com/2021/day/11
 */


func Solution() {
	file, _ := os.Open("day11/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}