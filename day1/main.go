package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solution1() {
	file, err := os.Open("day1/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	previous := int(^uint(0) >> 1)
	count := 0

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if curr > previous {
			count++
		}
		previous = curr
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Count: %d", count))
}

func Solution2() {
	file, err := os.Open("day1/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buffer := make([]int, 0)
	count := 0

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		buffer = append(buffer, curr)
	}

	for i := 3; i < len(buffer); i++ {
		if biggerThanPreviousGroup(buffer, uint(i)) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Count: %d", count))
}

func biggerThanPreviousGroup(buffer []int, idx uint) bool {
	if idx < 3 {
		return false
	}
	return buffer[idx] > buffer[idx - 3]
}

