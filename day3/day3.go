package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution1() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	onesCounts := make([]int, 12)
	totalLines := 0

	for scanner.Scan() {
		line := scanner.Text()
		bits := strings.Split(line, "")
		for idx, bit := range bits {
			if bit == "1" {
				onesCounts[idx] += 1
			}
		}
		totalLines++
	}

	gammaRate := ""
	epsilonRate := ""

	halfTotal := totalLines / 2

	for _, count := range onesCounts {
		if count > halfTotal {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}
	fmt.Println("Gama: " + gammaRate)
	fmt.Println("Epsilon: " + epsilonRate)

	gamma, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Println(gamma * epsilon)

	test, _ := strconv.ParseInt("01001", 2, 64)
	fmt.Println(test)

}
