package day9

import (
	"adventofcode2021/day7"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkLevel(heightMap []string, visitedMap map[string]bool, curr uint8, x, y int) int {
	mapKey := fmt.Sprintf("%d,%d", x, y)
	if visitedMap[mapKey] {
		return 0
	}
	elem := heightMap[y][x]
	if elem < "9"[0] && elem > curr {
		visitedMap[mapKey] = true
		return calcBasinSize(heightMap, visitedMap, x, y)
	}
	return 0
}

func calcBasinSize(heightMap []string, visitedMap map[string]bool, x, y int) int {
	rows := len(heightMap)
	cols := len(heightMap[0])
	curr := heightMap[y][x]
	var top, bottom, left, right int
	if y > 0 {
		top = checkLevel(heightMap, visitedMap, curr, x, y - 1)
	}
	if y < rows - 1 {
		bottom = checkLevel(heightMap, visitedMap, curr, x, y + 1)
	}
	if x > 0 {
		left = checkLevel(heightMap, visitedMap, curr, x - 1, y)
	}
	if x < cols - 1 {
		right = checkLevel(heightMap, visitedMap, curr, x + 1, y)
	}
	return 1 + top + bottom + left + right
}


func Solution() {

	file, _ := os.Open("day9/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	heightMap := make([]string, 0)

	for scanner.Scan() {
		heightMap = append(heightMap, scanner.Text())
	}

	cols := len(heightMap[0])
	rows := len(heightMap)

	total := 0

	basins := make([]int, 0)

	for y := 0; y < rows ; y++ {
		for x := 0; x < cols; x++ {
			curr := heightMap[y][x]
			if y > 0 && heightMap[y-1][x] <= curr {
				continue
			}
			if y < rows - 1 && heightMap[y+1][x] <= curr {
				continue
			}
			if x > 0 && heightMap[y][x-1] <= curr {
				continue
			}
			if x < cols - 1 && heightMap[y][x+1] <= curr {
				continue
			}
			visitedMap := make(map[string]bool)
			basinSize := calcBasinSize(heightMap, visitedMap, x, y)
			basins = append(basins, basinSize)
			fmt.Printf("Low point (%d, %d); Basin size: %d\n", x, y, calcBasinSize(heightMap, visitedMap, x, y))
			c, _ := strconv.Atoi(string(curr))
			total += c + 1
		}
	}

	lb := len(basins)
	day7.QuickSort(basins, 0, lb - 1)

	fmt.Println(basins[lb-1]*basins[lb-2]*basins[lb-3])

	fmt.Println(basins)

	fmt.Println(total)

}
