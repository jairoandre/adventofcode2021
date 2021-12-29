package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(line string) *Point {
	values := strings.Split(line, ",")
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])
	return &Point{
		X: x,
		Y: y,
	}
}

func Solution() {
	file, _ := os.Open("day13/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	foldInstructions := false

	points := make([]*Point, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			foldInstructions = true
			continue
		}
		if !foldInstructions {
			points = append(points, NewPoint(line))
		}
	}

	fmt.Println(points)
}
