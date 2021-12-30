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

func Fold(points []*Point, cmd string) {
	splitCmd := strings.Split(strings.Fields(cmd)[2], "=")
	axis := splitCmd[0]
	amount, _ := strconv.Atoi(splitCmd[1])
	for i := 0; i < len(points); i++ {
		point := points[i]
		switch axis {
		case "x":
			if point.X > amount {
				point.X = amount - (point.X - amount)
			}
			break
		default:
			if point.Y > amount {
				point.Y = amount - (point.Y - amount)
			}
		}
	}
}

func CountUnique(points []*Point) int {
	countMap := make(map[string]bool)
	maxX := 0
	maxY := 0
	for _, point := range points {
		x := point.X
		y := point.Y
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		countMap[fmt.Sprintf("%d-%d", point.X, point.Y)] = true
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			c := " "
			if countMap[fmt.Sprintf("%d-%d", x, y)] {
				c = "o"
			}
			fmt.Print(c)
		}
		fmt.Print("\n")
	}
	return len(countMap)
}

func Solution() {
	file, _ := os.Open("day13/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	foldInstructions := false

	points := make([]*Point, 0)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			foldInstructions = true
			continue
		}
		if !foldInstructions {
			points = append(points, NewPoint(line))
			continue
		}
		Fold(points, line)
		count++

	}

	fmt.Println(count, ":", CountUnique(points))

}
