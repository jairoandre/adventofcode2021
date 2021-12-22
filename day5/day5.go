package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


type PointMap map[string]int

func ReadPoint(str []string) (int, int) {
	x1, _ := strconv.Atoi(str[0])
	y1, _ := strconv.Atoi(str[1])
	return x1, y1
}

func Solution() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pointMap := make(PointMap, 0)

	count := 0

	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		x1, y1 := ReadPoint(strings.Split(values[0], ","))
		x2, y2 := ReadPoint(strings.Split(values[2], ","))
		diagonal := x1 != x2 && y1 != y2
		stepX := 1
		if x1 > x2 {
			stepX = -1
		}
		stepY := 1
		if y1 > y2 {
			stepY = -1
		}
		yB := y1
		LoopX:
		for x := x1; (stepX > 0 && x <= x2) || (stepX < 0 && x >= x2); x += stepX {
			for y := yB; (stepY > 0 && y <= y2) || (stepY < 0 && y >= y2); y += stepY {
				key := fmt.Sprintf("%d,%d", x, y)
				pointMap[key] += 1
				if curr := pointMap[key]; curr == 2 {
					count++
				}
				if diagonal {
					yB += stepY
					continue LoopX
				}
			}
		}

	}
	fmt.Println(count)
}

