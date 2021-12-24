package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LanternFishArray [9]int

func (arr LanternFishArray) NextState() LanternFishArray {
	var nextState LanternFishArray
	for i := 0; i < 8; i++ {
		nextState[i] = arr[i+1]
	}
	nextState[8] = arr[0]
	nextState[6] += arr[0]
	return nextState
}

func (arr LanternFishArray) Total() int {
	sum := 0
	for _, curr := range arr {
		sum += curr
	}
	return sum
}

func Solution() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	state := LanternFishArray{}

	for _, str := range strings.Split(scanner.Text(), ",") {
		idx, _ := strconv.Atoi(str)
		state[idx] += 1
	}

	fmt.Println(state)

	for i := 0; i < 256; i++ {
		state = state.NextState()
		fmt.Println(state)
	}

	fmt.Println(state.Total())

}
