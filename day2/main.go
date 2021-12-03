package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution1() {
	file, err := os.Open("day2/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fSum := 0
	dSum := 0
	aSum := 0

	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")

		command := arr[0]
		amount, _ := strconv.Atoi(arr[1])

		switch command {
		case "forward":
			fSum += amount
			dSum += amount * aSum
		case "down":
			aSum += amount
		case "up":
			aSum -= amount
		default:
			// nothing
		}


	}

	fmt.Println(fmt.Sprintf("Count: %d", fSum * dSum))

}
