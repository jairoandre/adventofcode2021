package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func QuickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := Partition(arr, lo, hi)
	QuickSort(arr, lo, p - 1)
	QuickSort(arr, p + 1, hi)
}

func Partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if arr[j] < pivot {
			i += 1
			Swap(arr, i, j)
		}
	}
	i += 1
	if pivot < arr[i] {
		Swap(arr, i, hi)
	}
	return i
}


func Swap(arr []int, i, j int) {
	aux := arr[j]
	arr[j] = arr[i]
	arr[i] = aux
}

func Test() {
	arr := []int{4, 2, 10, 2, 1, 3, 5, 6, 1, 3, 2}
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}

func Solution() {
	file, _ := os.Open("day7/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	scanner.Scan()

	sum := 0
	values := make([]int, 0)
	for _, str := range strings.Split(scanner.Text(),",") {
		val, _ := strconv.Atoi(str)
		sum += val
		values = append(values, val)
	}

	QuickSort(values, 0, len(values) - 1)

	avg := sum / len(values)

	fmt.Println(values)

	l := len(values)

	m := l / 2

	median := values[m]

	if l % 2 == 0 {
		median = (median + values[m - 1]) / 2
	}

	leastCost := 0

	for _, val := range values {
		steps := math.Abs(float64(val - avg))
		leastCost += CalcCost(int(steps))
	}

	fmt.Println(leastCost)
}


func CalcCost(steps int) int {
	cost := 0
	for i := 0; i < steps; i++ {
		cost += 1 + i
	}
	return cost
}