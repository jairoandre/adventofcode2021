package day10

import (
	"adventofcode2021/day7"
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	PARENTHESIS = []uint8{"("[0], ")"[0]}
	TAGS        = []uint8{"<"[0], ">"[0]}
	BRACKETS    = []uint8{"["[0], "]"[0]}
	CURLS       = []uint8{"{"[0], "}"[0]}
)

type Buffer struct {
	Items []uint8
}

func Close(str string, closeBuffer *Buffer, toClose uint8) (int, string, error) {
	length := len(str)
	if length == 0 {
		return 0, "", errors.New("incomplete")
	}
	nextChar := str[0]
	// Verify is Closing?
	if toClose == nextChar {
		closeBuffer.Items = closeBuffer.Items[:len(closeBuffer.Items)-1]
		return 0, str[1:], nil
	}
	points, tail, err := Open(str, closeBuffer)
	if err != nil {
		return points, "", err
	}
	return Close(tail, closeBuffer, toClose)
}

func Open(str string, closeBuffer *Buffer) (int, string, error) {
	if len(str) == 0 {
		return 0, "", nil
	}
	symbol := str[0]
	switch symbol {
	case PARENTHESIS[0]:
		closeBuffer.Items = append(closeBuffer.Items, PARENTHESIS[1])
		return Close(str[1:], closeBuffer, PARENTHESIS[1])
	case BRACKETS[0]:
		closeBuffer.Items = append(closeBuffer.Items, BRACKETS[1])
		return Close(str[1:], closeBuffer, BRACKETS[1])
	case CURLS[0]:
		closeBuffer.Items = append(closeBuffer.Items, CURLS[1])
		return Close(str[1:], closeBuffer, CURLS[1])
	case TAGS[0]:
		closeBuffer.Items = append(closeBuffer.Items, TAGS[1])
		return Close(str[1:], closeBuffer, TAGS[1])
	default:
		return Point(symbol), "", errors.New("corrupted")
	}
}

func Point(symbol uint8) int {
	switch symbol {
	case PARENTHESIS[1]:
		return 3
	case BRACKETS[1]:
		return 57
	case CURLS[1]:
		return 1197
	case TAGS[1]:
		return 25137
	default:
		return 0
	}
}

func Point2(symbol uint8) int {
	switch symbol {
	case PARENTHESIS[1]:
		return 1
	case BRACKETS[1]:
		return 2
	case CURLS[1]:
		return 3
	case TAGS[1]:
		return 4
	default:
		return 0
	}
}

func Solution() {
	file, _ := os.Open("day10/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	total := 0
	scores := make([]int, 0)
	for scanner.Scan() {
		closeBuffer := Buffer{Items: make([]uint8, 0)}
		line := scanner.Text()
		points, _, err := Open(line, &closeBuffer)
		if err != nil && err.Error() == "corrupted" {
			continue
		}
		score := 0
		for idx := len(closeBuffer.Items) - 1; idx >= 0; idx-- {
			symbol := closeBuffer.Items[idx]
			score *= 5
			score += Point2(symbol)

		}
		scores = append(scores, score)
		total += points
	}

	day7.QuickSort(scores, 0, len(scores)-1)

	if len(scores) > 0 {
		fmt.Println(scores[len(scores)/2])
	}
}
